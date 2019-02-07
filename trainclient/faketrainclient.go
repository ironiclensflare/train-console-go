package trainclient

type FakeTrainClient struct{}

func (c FakeTrainClient) GetAllDepartures(crs string) []Train {
	trains := getFakeTrains()
	return trains
}

func getFakeTrains() []Train {
	trains := []Train{Train{"NOT", "STP", "10:55", "10:56"}, Train{"NOT", "STP", "12:04", "On Time"}, Train{"NOT", "STP", "13:17", "On Time"}, Train{"NOT", "STP", "14:56", "Delayed"}}
	return trains
}
