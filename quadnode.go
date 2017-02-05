package rquad

import (
	"fmt"
	"image"
)

// Color is the set of colors that can take a Node.
type Color byte

const (
	// Black is the color of leaf nodes
	// that are considered as obstructed.
	Black Color = 0 + iota

	// White is the color of leaf nodes
	// that are considered as free.
	White

	// Gray is the color of non-leaf nodes
	// that contain both black and white children.
	Gray
)

const _Color_name = "BlackWhiteGray"

var _Color_index = [...]uint8{0, 5, 10, 14}

func (i Color) String() string {
	if i >= Color(len(_Color_index)-1) {
		return fmt.Sprintf("Color(%d)", i)
	}
	return _Color_name[_Color_index[i]:_Color_index[i+1]]
}

// Node defines the interface for a quadtree node.
type Node interface {

	// Parent returns the quadtree node that is the parent of current one.
	Parent() Node

	// Child returns current node child at specified quadrant.
	Child(Quadrant) Node

	// Bounds returns the bounds of the rectangular area represented by this
	// quadtree node.
	Bounds() image.Rectangle

	// Color returns the node Color.
	Color() Color

	// Location returns the node inside its parent quadrant
	Location() Quadrant
}

// NodeList is a slice of Node instances.
type NodeList []Node
