package main

func main() {
	p := &patient{name: "tom"}

	r := &register{}
	c := &checkout{}
	m := &medical{}
	pay := &payment{}

	r.setNext(c)
	c.setNext(m)
	m.setNext(pay)

	r.execute(p)
	r.execute(p)

}
