package main

func main() {
	manager := NewStationManager()
	pass := passengerTrain{name: "东方之珠", mediator: manager}
	freight := freightTrain{name: "长江三号", mediator: manager}
	pass.arrive()
	freight.arrive()
}
