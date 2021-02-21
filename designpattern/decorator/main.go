package main

func main() {
	tomato := &tomato{price: 10}
	potato := &potato{price: 20}
	cookWithtomatoandpotato := topChief{m: []material{
		tomato,
		potato,
	}}
	cookWithtomatoandpotato.cook()
}
