package main

import (
	"container/list"
	"fmt"
	"github.com/greengeko/kiwi-go"
	"time"
)

type Node struct {
	data     kiwi.Flight
	parent   *Node
	children list.List
}

func print(node Node) {
	for i := 0; i < len(node.data.Route); i++ {
		fmt.Print("Flight from " + node.data.Route[i].FlyFrom + "to " + node.data.Route[i].FlyTo + " and flight number: ")
		fmt.Print(node.data.Route[i].OperatingFlightNo)

	}
	unixTimeUTC := time.Unix(int64(node.data.DepartureTimeUTC), 0) //gives unix time stamp in utc
	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339)
	fmt.Print("departure on: " + unitTimeInRFC3339)
	fmt.Print("for the price: ")
	fmt.Print(node.data.Price)

}
