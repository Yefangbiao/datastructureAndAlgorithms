package main

func main() {
	tv := &tvDevice{}
	oncommand := &onCommand{tv}
	offcommand := &offCommand{tv}

	b1 := button{command: oncommand}
	b2 := button{command: offcommand}

	b1.press()
	b2.press()
}
