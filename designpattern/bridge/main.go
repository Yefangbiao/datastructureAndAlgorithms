package main

func main() {
	hp := &hpPrint{}
	epson := &epsonPrint{}

	mac := NewMac(hp)
	mac.ComputerPrint()

	windows := NewWindows(epson)
	windows.ComputerPrint()
}
