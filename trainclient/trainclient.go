package trainclient

import (
  "fmt"
)

type TrainClient struct {
  ApiKey string
}

func (tc TrainClient) GetAllDepartures(crs string) []Train {
  trains := make([]Train, 0)
  trains = append(trains, Train{crs, "STP", "10:55", "10:56"})
  fmt.Println("Created a train")
  return trains
}

