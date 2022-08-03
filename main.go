package main

import (
	"fmt"
	kiwi "github.com/greengeko/kiwi-go"
	"strings"
	"time"
)

var k = kiwi.New()
var flyTo = "ATH"

// var flyTo = "45.46-9.18-250km"
var mother Node
var options []kiwi.Flight
var avoid0 []string

func main() {
	fmt.Println("Hello, World!")
	parametersTest := kiwi.Parameters{
		//FlyFrom: "1.35-103.81-250km",
		FlyFrom: "SIN",
		FlyTo:   "ATH",
		Partner: "gekostopovers",
		//PriceFrom: 100,
		//OneForCity:    true,
		Limit:         1000,
		DepartureFrom: "2022-09-05T00:00",
		DepartureTo:   "2022-09-15T00:00",
		//ArrivalFrom
		//ArrivalTo
	}
	resp, _ := k.GetFlights(&parametersTest)

	flight2 := kiwi.Flight{}
	for i := 0; i < len(resp.Flights); i++ {
		fmt.Print(flight2.FlyFrom)
		fmt.Print(flight2.FlyTo)
		fmt.Println(flight2.Price)
	}
	//k := kiwi.New()

	params := kiwi.Parameters{
		//FlyFrom: "1.35-103.81-250km",
		FlyFrom: "SIN",
		Partner: "gekostopovers",
		//PriceFrom: 100,
		OneForCity:    true,
		Limit:         1000,
		DepartureFrom: "2022-09-05T00:00",
		DepartureTo:   "2022-09-15T00:00",
		//ArrivalFrom
		//ArrivalTo
	}
	avoid0 = append(avoid0, params.FlyFrom)
	var tail int = searchFlights(mother, params, 0, avoid0)
	//fmt.Println(options)
	fmt.Println(tail)
	for y := 0; y < len(options); y++ {
		/*fmt.Println("prezzo:")
		fmt.Println(options[y].Price)*/
		/*for i := 0; i < len(options[y].Route); i++ {

			fmt.Println("step")
			fmt.Println(i)
			fmt.Println(options[y].Route[i].FlyFrom)
			fmt.Println(options[y].Route[i].FlyTo)
		}*/
	}

}

func searchFlights(n Node, parameters kiwi.Parameters, price float64, avoid []string) int {
	var num = 0
	var ok = true
	price = price + n.data.Price
	if price > 400 && n.parent != &mother && &n != &mother && price-n.data.Price != 0 && price-n.data.Price < 400 {
		createFlight(*n.parent, price-n.data.Price)
	} else {
		for a := 0; a < len(n.data.Route); a++ {
			avoid = append(avoid, n.data.Route[a].FlyTo)
		}
		for j := 0; j < len(avoid); j++ {
			if avoid[j] == n.data.FlyTo && n.data.FlyFrom != "SIN" {
				ok = false
				if price < 400 {
					createFlight(n, price)
				}
			}
		}
		//fmt.Println("nuovo da" + params.FlyFrom)
		if n.data.FlyTo == flyTo && price < 400 {
			createFlight(n, price) //da nodi risale per formare una lista e metterla in options
		} else if ok {
			//createFlight(n)
			resp, _ := k.GetFlights(&parameters)
			flight := kiwi.Flight{}
			for i := 0; i < len(resp.Flights); i++ {
				flight = resp.Flights[i]
				//fmt.Println(flight.FlyFrom + "to" + flight.FlyTo + "for" + strconv.Itoa(int(flight.Price)))
				var flightNode = Node{data: flight, parent: &n}
				n.children = append(n.children, &flightNode)
				//printFile(flightNode)

				unixTimeUTCTO := time.Unix(int64(flight.ArrivalTimeUTC), 0).AddDate(0, 0, 4)
				unixTimeUTCFROM := time.Unix(int64(flight.ArrivalTimeUTC), 0).AddDate(0, 0, 1) //gives unix time stamp in utc
				unitTimeInRFC3339TO := unixTimeUTCTO.Format(time.RFC3339)
				unitTimeInRFC3339FROM := unixTimeUTCFROM.Format(time.RFC3339)
				var departureTo string = strings.Split(unitTimeInRFC3339TO, "T")[0]
				var departureFrom string = strings.Split(unitTimeInRFC3339FROM, "T")[0]
				departureFrom = departureFrom + "T00:00"
				departureTo = departureTo + "T00:00"
				flightParams := kiwi.Parameters{FlyFrom: flight.FlyTo, Partner: "gekostopovers", Limit: 1000,
					OneForCity: true, DepartureFrom: departureFrom, DepartureTo: departureTo} //aggiungere param temporali
				//fmt.Println("da" + flight.FlyTo)
				num = 1 + searchFlights(flightNode, flightParams, price, avoid)

			}

		}
	}
	return num
}

func createFlight(n Node, price float64) {
	/*var flight = kiwi.Flight{}
	flight.Price = price*/
	printFile(n)
	for n.parent != nil {
		print(n)

		//createFlight(*n.parent, price)

		/*for i := 0; i < len(n.data.Route); i++ {
			flight.Route = append(flight.Route, n.data.Route[i])

		}
		var route = kiwi.Route{FlyTo: "---"}
		flight.Route = append(flight.Route, route)*/
		//var parent Node = n.Parent()
		n.data = (*n.parent).data
		n.parent = (*n.parent).parent
		//n.children = (*n.parent).children

	}
	fmt.Println("flight added ")

	/*fmt.Println(flight.Price)
	for i := 0; i < len(flight.Route); i++ {
		fmt.Print(flight.Route[i].FlyTo)
	}

	options = append(options, flight)

	/*fmt.Println("prezzo:")
	fmt.Println(options[0].Price)
	for i := 0; i < len(options[0].Route); i++ {
		fmt.Println("step")
		fmt.Println(i)
		fmt.Println(options[0].Route[i].FlyFrom)
		fmt.Println(options[0].Route[i].FlyTo)
	}*/

}
