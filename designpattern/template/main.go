package main

func main() {
	concrete1 := &concrete1{}
	ab1 := abstract{
		iAbstract: concrete1,
	}
	ab1.Start()
	ab1.step1()
	ab1.step2()
	ab1.step3()
	ab1.step4()

	concrete2 := &concrete2{}
	ab2 := abstract{
		iAbstract: concrete2,
	}
	ab2.Start()
	ab2.step1()
	ab2.step2()
	ab2.step3()
	ab2.step4()
}
