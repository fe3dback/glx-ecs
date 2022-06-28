package component

import "github.com/fe3dback/glx-ecs/ecs"

const Node2DTypeID = "Node2D-7c40b8e315a5"

type Node2D struct {
	X float64
	Y float64
}

func NewNode2D(x, y float64) *Node2D {
	return &Node2D{
		X: x,
		Y: y,
	}
}

func (c Node2D) TypeID() ecs.ComponentTypeID {
	return Node2DTypeID
}
