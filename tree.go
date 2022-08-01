package main

import (
	"container/list"
	"github.com/greengeko/kiwi-go"
)

type Node struct {
	data     kiwi.Flight
	parent   *Node
	children list.List
}
