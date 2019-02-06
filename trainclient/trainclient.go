package trainclient

type TrainClient struct {
	ApiKey string
}

type TrainGetter interface {
	GetAllDepartures(crs string) []Train
	GetAllArrivals(crs string) []Train
}

func (tc TrainClient) GetAllDepartures(crs string) []Train {
	trains := []Train{Train{crs, "STP", "10:55", "10:57"}}
	return trains
}

func (tc TrainClient) GetAllArrivals(crs string) []Train {
	trains := []Train{Train{"STP", crs, "10:55", "10:57"}}
	return trains
}
