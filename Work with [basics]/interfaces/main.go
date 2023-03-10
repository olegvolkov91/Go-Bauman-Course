package main

import "fmt"

type Animal interface {
	Runner
	Flyer
}

type Flyer interface {
	Fly()
}

type Runner interface {
	Run()
}

type Eagle struct{}

func (e *Eagle) Fly() {}

type Kiwi struct{}

func (k *Kiwi) Run() {}

type Pigeon struct{}

func (p *Pigeon) Fly() {}
func (p *Pigeon) Run() {}

func main() {
	pigeon := &Pigeon{}
	kiwi := &Kiwi{}
	eagle := &Eagle{}

	Fly(pigeon)
	//Fly(kiwi) // <- error because Kiwi does not have Fly method
	Fly(eagle)

	Run(pigeon)
	Run(kiwi)
	//Run(eagle) // <- error because Eagle does not have Run method

	var animal Animal = &Pigeon{}
	fmt.Printf("Animal %T\n", animal)
	TypeFinder(animal)
}

func Run(r Runner) {
	r.Run()
}

func Fly(f Flyer) {
	f.Fly()
}

func TypeFinder(i interface{}) {
	switch v := i.(type) {
	case Runner:
		fmt.Printf("This is a runner %v\n", v)
	case Flyer:
		fmt.Printf("This is a flyer %v\n", v)
	default:
		fmt.Println("Unknown type")
	}
}
