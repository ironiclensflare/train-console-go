package main

import (
	"fmt"

	"github.com/ironiclensflare/train-console/trainclient"
)

func main() {
	client := trainclient.TrainClient{ApiKey: "12345"}
	departures := client.GetAllDepartures("NOT")
	arrivals := client.GetAllArrivals("STP")
	fmt.Println(departures, arrivals)
}
