package main

import (
	"bufio"
	"fmt"
	"github.com/greengeko/kiwi-go"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Node struct {
	data     kiwi.Flight
	parent   *Node
	children []*Node
}

func print(node Node) {
	for i := 0; i < len(node.data.Route); i++ {
		fmt.Print("Flight from " + node.data.Route[i].FlyFrom + " to " + node.data.Route[i].FlyTo + " and flight number: ")
		fmt.Print(node.data.Route[i].OperatingFlightNo)

	}
	unixTimeUTC := time.Unix(int64(node.data.DepartureTimeUTC), 0) //gives unix time stamp in utc
	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339)
	fmt.Print("departure on: " + unitTimeInRFC3339)
	fmt.Print("for the price: ")
	fmt.Print(node.data.Price)

}

func printFile(node Node) {
	var price int = 0
	//var filename = RandStringBytes(10)
	var filename string = node.data.FlyFrom + node.data.FlyTo
	//var filenameNew string = filename
	nuovofile, errore := os.Create(filename + ".txt")
	if errore == nil {
		fmt.Println("File creato: ", nuovofile.Name())
	} else {
		fmt.Println("Errore: ", errore)
		panic(errore)
	}
	myWriter := bufio.NewWriter(nuovofile)
	myWriter.WriteString("\nciao a tutti")
	for node.parent != nil {

		for i := 0; i < len(node.data.Route); i++ {
			myWriter.WriteString("Flight from " + node.data.Route[i].FlyFrom + "to " + node.data.Route[i].FlyTo + " and flight number: ")
			myWriter.WriteString(node.data.Route[i].OperatingFlightNo)
			//filenameNew = filenameNew + node.data.FlyTo

		}
		unixTimeUTC := time.Unix(int64(node.data.DepartureTimeUTC), 0) //gives unix time stamp in utc
		unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339)
		myWriter.WriteString("departure on: " + unitTimeInRFC3339)
		myWriter.WriteString("for the price: ")
		myWriter.WriteString(strconv.Itoa(int(node.data.Price)))
		price = price + int(node.data.Price)
		node.data = (*node.parent).data
		node.parent = (*node.parent).parent

	}
	myWriter.WriteString("ARRIVO a " + node.data.FlyTo + "with price: " + strconv.Itoa(price))
	/*e := os.Rename(filename+".txt", filenameNew+".txt")
	if e != nil {
		log.Fatal(e)
	}*/
	myWriter.Flush()
	nuovofile.Close()
	fmt.Println("File popolato")
	//node = *node.parent

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
