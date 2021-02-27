package main

type onCommand struct {
	device
}

func (o *onCommand) execute() {
	o.device.on()
}
