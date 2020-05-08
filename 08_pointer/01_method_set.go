package main

import "fmt"

type transport interface {
	// This is the method set of the interface
	start()
	stop()
}

type car struct {
	brand string
	doors int
	speed int
	passenger int
}

// Car's method set, it implements transport interface
func (c car) start() {
	fmt.Printf("Car %v with %v doors can have %v passengers, it is travelling at %v km/h\n",c.brand,c.doors,c.passenger,c.speed)
}

func (c car) stop() {
	fmt.Printf("%v is stopped\n",c.brand)
}

type airplane struct {
	brand string
	passenger int
	speed int
}

// Airplane's method set, it doesn't implements transport interface because it have pointer's receiver(r *receiver)
func (a *airplane) start() {
	fmt.Printf("Airplane %v can have %v passengers, it is travelling at %v km/h\n",a.brand,a.passenger,a.speed)
}

func (a *airplane) stop() {
	fmt.Printf("%v is stopped\n",a.brand)
}


// Called methods with interface as parameter
func run(t transport)  {
	t.start()
}

func stop(t transport)  {
	t.stop()
}

func main() {
	c := car{
		brand: "Versa",
		doors: 4,
		speed: 80,
		passenger: 5,
	}

	a := airplane{
		brand: "Boeing 737",
		speed: 840,
		passenger: 132,
	}

	// You can pass either with type 'car' or '*car' because the methods have the receiver of type (r receiver)
	fmt.Println("\nCalling with type 'car'")
	run(c)
	stop(c)

	fmt.Println("\nCalling with type '*car'")
	run(&c)
	stop(&c)

	// You can't pass 'airplane' because it doesn't implements the transport interface
	fmt.Println("\nNot calling with type 'airplane'")
	// run(a)
	// stop(a)

	// You can pass '*airplane' because the pointer implements the transport interface
	fmt.Println("\nCalling with type '*airplane'")
	run(&a)
	stop(&a)
	fmt.Println()
}
