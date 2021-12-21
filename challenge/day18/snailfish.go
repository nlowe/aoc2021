package day18

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2021/challenge"

	"github.com/nlowe/aoc2021/util"
)

// snailfish represents a snailfish number, which is essentially a tree with 2 children per node
type snailfish struct {
	parent *snailfish

	left  *snailfish
	right *snailfish

	v int
}

func Parse(line string) (*snailfish, string) {
	// Nested snailfish?
	if line[0] == '[' {
		line = line[1:]
		result := &snailfish{}

		// If so, parse the left and right snailfish
		result.left, line = Parse(line)
		result.left.parent = result

		if line[0] == ',' {
			// Coming up from a nested snailfish, consume the ','
			line = line[1:]
		}

		result.right, line = Parse(line)
		result.right.parent = result

		return result, line[1:]
	}

	idx := strings.IndexAny(line, "],")

	// Otherwise, this should be a literal, separated by a ']' or ','
	parts := strings.SplitN(line, string(line[idx]), 2)

	if line[idx] == ']' {
		// Don't consume closing brace, that's the parent snailfish's job
		line = "]" + parts[1]
	} else {
		line = parts[1]
	}

	return &snailfish{v: util.MustAtoI(parts[0])}, line
}

func (s *snailfish) String() string {
	if s == nil {
		return ""
	}

	if s.left == nil && s.right == nil {
		return strconv.Itoa(s.v)
	}

	if s.left == nil || s.right == nil {
		panic(fmt.Errorf("bad snailfish %p: left=%p, right=%p, v=%d", s, s.left, s.right, s.v))
	}

	line := strings.Builder{}
	line.WriteRune('[')
	line.WriteString(s.left.String())
	line.WriteRune(',')
	line.WriteString(s.right.String())
	line.WriteRune(']')

	return line.String()
}

func AddAll(challenge *challenge.Input) *snailfish {
	var s *snailfish
	for line := range challenge.Lines() {
		parsed, left := Parse(line)
		if left != "" {
			panic(fmt.Errorf("failed to parse '%s': leftover: %s", line, left))
		}

		if s == nil {
			s = parsed
		} else {
			s = s.add(parsed)
		}
	}

	return s
}

func (s *snailfish) add(other *snailfish) *snailfish {
	result := &snailfish{
		left:  s,
		right: other,
	}

	// Don't forget to re-parent the child nodes
	s.parent = result
	other.parent = result

	result.reduce()
	return result
}

// reduce simplifies a snailfish number by repeatedly calling explode then split until both perform no operations
func (s *snailfish) reduce() {
	for {
		// First, try to explode a number
		if s.explode() {
			// If we did, start the reduction over
			continue
		}

		// Then, try splitting a number
		if s.split() {
			// If we did, start the reduction over
			continue
		}

		// If neither simplification could be performed, we're done reducing
		return
	}
}

// explode "balances" the snailfish number by exploding the pair nested four levels deep. It does this by:
//
// * The left number of the pair is added to the rightmost child of the pair's parent's left child, if any
// * The right number of the pair is added to the leftmost child of the pair's parent's right child, if any
// * The entire pair is replaced with a literal 0
//
// Return true iff a pair explodes. At most one explosion per call.
func (s *snailfish) explode() bool {
	target := findExplodeTarget(0, s)
	if target == nil {
		return false
	}

	// The left target is the rightmost child of the parent's left child
	leftTarget := target
	// Go "up" until we can go left
	for leftTarget.parent != nil {
		old := leftTarget
		leftTarget = leftTarget.parent

		if leftTarget.left != old {
			// Now that we can finally go left, do so
			leftTarget = leftTarget.left
			break
		}
	}

	if leftTarget != s {
		// Go "right" until we hit a leaf
		for leftTarget.right != nil {
			leftTarget = leftTarget.right
		}
	}

	// The right target is the leftmost child of the parent's right child
	rightTarget := target
	// Go "up" until we can go right
	for rightTarget.parent != nil {
		old := rightTarget
		rightTarget = rightTarget.parent

		if rightTarget.right != old {
			// Now that we can finally go right, do so
			rightTarget = rightTarget.right
			break
		}
	}

	if rightTarget != s {
		// Go "left" until we hit a leaf
		for rightTarget.left != nil {
			rightTarget = rightTarget.left
		}
	}

	if leftTarget != nil {
		// If we found a left target, add the target's left leaf to it
		leftTarget.v += target.left.v
	}

	if rightTarget != nil {
		// If we found a right target, add the target's right leaf to it
		rightTarget.v += target.right.v
	}

	// Replace the target with a literal 0
	target.left = nil
	target.right = nil
	target.v = 0

	return true
}

// findExplodeTarget finds the first node whose descendents are literals and is nested 4 levels deep
func findExplodeTarget(n int, s *snailfish) *snailfish {
	if s.left == nil && s.right == nil {
		// Hit a leaf, go back up
		return nil
	}

	if n == 4 && s.left.left == nil && s.right.right == nil {
		return s
	}

	if target := findExplodeTarget(n+1, s.left); target != nil {
		return target
	}

	return findExplodeTarget(n+1, s.right)
}

// split finds the first snailfish literal >= 10 and turns it into a pair where the left element is the floor of the
// literal divided by 2, and the right element is the ceil of the literal divided by 2.
//
// Returns true iff a split was performed. At most one split per call.
func (s *snailfish) split() bool {
	target := findSplitTarget(s)
	if target == nil {
		return false
	}

	left := int(math.Floor(float64(target.v) / 2.0))
	right := int(math.Ceil(float64(target.v) / 2.0))

	target.v = 0
	target.left = &snailfish{parent: target, v: left}
	target.right = &snailfish{parent: target, v: right}

	return true
}

func findSplitTarget(s *snailfish) *snailfish {
	if s.left == nil && s.right == nil {
		if s.v >= 10 {
			return s
		}

		return nil
	}

	if target := findSplitTarget(s.left); target != nil {
		return target
	}

	return findSplitTarget(s.right)
}

// magnitude returns the magnitude of the number, which is 3 times its left child plus 2 times its right child.
//
// The magnitude of a literal is the literal value itself.
func (s *snailfish) magnitude() int {
	if s.left == nil && s.right == nil {
		return s.v
	}

	return 3*s.left.magnitude() + 2*s.right.magnitude()
}
