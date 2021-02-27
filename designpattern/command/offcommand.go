package main

type offCommand struct {
	device
}

func (o *offCommand) execute() {
	o.device.off()
}
