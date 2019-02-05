package main

import (
  "github.com/ironiclensflare/train-console/trainclient"
  "fmt"
)

func main() {
  client := trainclient.TrainClient{"12345"}
  trains := client.GetAllDepartures("NOT")
  fmt.Println(trains)
}

