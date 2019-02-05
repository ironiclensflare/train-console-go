package trainclient

import (
	"fmt"
)

type TrainClient struct {
	ApiKey string
}

func (tc TrainClient) GetAllDepartures(crs string) []Train {
	trains := []Train{Train{crs, "STP", "10:55", "10:57"}}
	fmt.Println("Created a train")
	return trains
}
