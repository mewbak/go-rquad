package quadtree

import "image"

// Color is the set of colors that can take a QNode.
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

// QNode defines the interface for a quadtree node.
type QNode interface {
	Parent() QNode

	Bounds() image.Rectangle
	Color() Color

	// ForEachNeighbour calls the given function for each neighbour of current
	// node.
	ForEachNeighbour(func(QNode))
}
