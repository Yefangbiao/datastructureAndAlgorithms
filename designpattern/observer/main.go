package main

func main() {
	c1 := cSubscriber{}
	c2 := cSubscriber{}
	c3 := cSubscriber{}
	c4 := cSubscriber{}

	p := publisher{[]subscriber{}}
	i := iphone{p: &p}
	i.p.addSubscriber(&c1)
	i.p.addSubscriber(&c2)
	i.p.addSubscriber(&c3)
	i.p.addSubscriber(&c4)

	i.newMessage("iphone12 is coming\n")
}
