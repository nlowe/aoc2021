package tilemap

import (
	"fmt"

	"github.com/nlowe/aoc2021/challenge"

	"github.com/beefsack/go-astar"
	"github.com/nlowe/aoc2021/util"
)

type TileContainer[T any] struct {
	tileMap *TileMap[T]

	Value T
	x     int
	y     int
}

func (t TileContainer[T]) PathNeighbors() (results []astar.Pather) {
	if t.tileMap.NeighborFunc != nil {
		neighbors := t.tileMap.NeighborFunc(t)

		// TODO: Is there a way to cast this so we don't have to copy?
		results = make([]astar.Pather, len(neighbors))
		for i, v := range t.tileMap.NeighborFunc(t) {
			results[i] = v
		}

		return
	}

	for _, delta := range []struct {
		x int
		y int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		c, ok := t.tileMap.ContainerAt(t.x+delta.x, t.y+delta.y)
		if !ok {
			continue
		}

		results = append(results, c)
	}

	return
}

func (t TileContainer[T]) PathNeighborCost(to astar.Pather) float64 {
	if t.tileMap.CostFunc != nil {
		return t.tileMap.CostFunc(t, to.(TileContainer[T]))
	}

	// Assume all paths are equal cost
	return 1
}

func (t TileContainer[T]) PathEstimatedCost(to astar.Pather) float64 {
	toSpot := to.(TileContainer[T])

	if t.tileMap.EstimateFunc != nil {
		return t.tileMap.EstimateFunc(t, toSpot)
	}

	return float64(util.ManhattanDistance(t.x, t.y, toSpot.x, toSpot.y))
}

// TileMap represents a fixed size grid of runes. The top-left tile is [0,0]
type TileMap[T any] struct {
	tiles []TileContainer[T]
	w     int
	h     int

	NeighborFunc func(container TileContainer[T]) []TileContainer[T]
	CostFunc     func(a, b TileContainer[T]) float64
	EstimateFunc func(a, b TileContainer[T]) float64
}

func FromInput(input *challenge.Input) *TileMap[rune] {
	return FromInputOf[rune](input, func(v rune) rune { return v })
}

func FromInputOf[T any](input *challenge.Input, convert func(rune) T) *TileMap[T] {
	lines := input.LineSlice()

	m := Of[T](len(lines[0]), len(lines))

	for row, line := range lines {
		for column, tile := range line {
			m.SetTile(column, row, convert(tile))
		}
	}

	return m
}

func New(w, h int) *TileMap[rune] {
	return Of[rune](w, h)
}

func Of[T any](w, h int) *TileMap[T] {
	return &TileMap[T]{
		tiles: make([]TileContainer[T], w*h),
		w:     w,
		h:     h,
	}
}

func (t *TileMap[T]) Size() (int, int) {
	return t.w, t.h
}

func (t *TileMap[T]) outOfBounds(x, y int) bool {
	return x < 0 || y < 0 || x >= t.w || y >= t.h
}

func (t *TileMap[T]) indexOf(x, y int) (int, bool) {
	return x + (t.w * y), !t.outOfBounds(x, y)
}

func (t *TileMap[T]) SetTile(x, y int, tile T) {
	idx, ok := t.indexOf(x, y)
	if !ok {
		panic(fmt.Errorf("out of bounds tile access: [%d, %d] is not within the %dx%d map", x, y, t.w, t.h))
	}

	t.tiles[idx] = TileContainer[T]{tileMap: t, Value: tile, x: x, y: y}
}

func (t *TileMap[T]) ContainerAt(x, y int) (TileContainer[T], bool) {
	idx, ok := t.indexOf(x, y)
	if !ok {
		return TileContainer[T]{}, false
	}

	return t.tiles[idx], true
}

func (t *TileMap[T]) TileAt(x, y int) (T, bool) {
	c, ok := t.ContainerAt(x, y)
	return c.Value, ok
}
func (t *TileMap[T]) PathBetween(startX, startY, endX, endY int) ([]astar.Pather, float64, bool) {
	start, ok := t.ContainerAt(startX, startY)
	if !ok {
		return nil, 0, false
	}

	end, ok := t.ContainerAt(endX, endY)
	if !ok {
		return nil, 0, false
	}

	return astar.Path(start, end)
}
