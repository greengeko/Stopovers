package main

import (
	"fmt"
	kiwi "github.com/greengeko/kiwi-go"
)

var k = kiwi.New()
var flyTo = "ATH"

// var flyTo = "45.46-9.18-250km"
var mother Node
var options []kiwi.Flight

func main() {
	fmt.Println("Hello, World!")
	//k := kiwi.New()

	params := kiwi.Parameters{
		//FlyFrom: "1.35-103.81-250km",
		FlyFrom: "SIN",
		Partner: "gekostopovers",
		//DepartureFrom: "",
		//DepartureTo:   "",
		//ArrivalFrom
		//ArrivalTo
	}

	searchFlights(mother, params, 0)
	//fmt.Println(options)

}

func searchFlights(n Node, params kiwi.Parameters, price float64) {
	price = price + n.data.Price
	//fmt.Println("nuovo da" + params.FlyFrom)
	if n.data.FlyTo == flyTo || price > 1000 {
		createFlight(n) //da nodi risale per formare una lista e metterla in options
	} else {
		//createFlight(n)
		resp, _ := k.GetFlights(&params)
		flight := kiwi.Flight{}
		for i := 0; i < len(resp.Flights); i++ {
			flight = resp.Flights[i]
			var flightNode = Node{data: flight, parent: &n}
			n.children.PushFront(flightNode)
			flightParams := kiwi.Parameters{FlyFrom: flight.FlyTo} //aggiungere param temporali
			//fmt.Println("da" + flight.FlyTo)
			searchFlights(flightNode, flightParams, price)
		}
	}
}

func createFlight(n Node) {
	var flight = kiwi.Flight{}
	for n.parent != &mother && n.parent != nil {
		for i := 0; i < len(n.data.Route); i++ {
			flight.Route = append(flight.Route, n.data.Route[i])
			flight.Price = flight.Price + n.data.Price
		}
		n = *n.parent
	}
	options = append(options, flight)
	fmt.Println("prezzo:")
	fmt.Println(options[0].Price)
	for i := 0; i < len(options[0].Route); i++ {
		fmt.Println("step")
		fmt.Println(i)
		fmt.Println(options[0].Route[i].FlyFrom)
		fmt.Println(options[0].Route[i].FlyTo)

	}

}
