package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define a few variables.
var hunger = 3                  // how many times a philosopher eats
var eat = 1 * time.Second       // how long it takes to eat
var think = 3 * time.Second     // how long a philosopher thinks
var sleepTime = 1 * time.Second // how long to wait when printing things out
var orderFinished []string      // the order in which philosophers finish dining and leave
var orderMutex sync.Mutex       // a mutex for the slice orderFinished

// define a wait group

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(philosopher.name, "is seated at the table.")
	seated.Done()

	seated.Wait()

	// we'll define a waitgroup for this goroutine specifically, so that the "leaving table"
	// message does not get printed until after the message that the forks are dropped
	eating := &sync.WaitGroup{}
	eating.Add(hunger)

	for i := hunger; i > 0; i-- {
		// get a lock on the left and right forks
		forks[philosopher.leftFork].Lock()
		forks[philosopher.rightFork].Lock()

		fmt.Println(philosopher.name, "has both forks, and is eating.")
		time.Sleep(eat)

		// the philosopher starts to think, but does not drop the forks yet
		fmt.Printf("%s is thinking.\n", philosopher.name)
		time.Sleep(think)

		// unlock the mutexes for both forks
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)

		// decrement the eating waitgroup by 1
		eating.Done()
	}

	// wait until all messages have been printed
	eating.Wait()

	fmt.Println(philosopher.name, "is satisfied.")
	time.Sleep(sleepTime)
	fmt.Println(philosopher.name, "left the table.")

	// update the list of finished eaters
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")
	time.Sleep(sleepTime)

	// set delays to zero while developing
	eat = 0 * time.Second
	sleepTime = 0 * time.Second
	think = 0 * time.Second

	// wg is the waitgroup that keeps track of how many philosophers are still at the table. When
	// it reaches zero, everyone is finished eating and has left. We add 5 (the number of philosophers) to this
	// wait group.
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// we want everyone to be seated before they start eating, so create a waitgroup for that, and set it to 5.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks. Forks are assigned using the fields leftFork and rightFork in the Philosopher type
	forks := map[int]*sync.Mutex{
		0: &sync.Mutex{},
		1: &sync.Mutex{},
		2: &sync.Mutex{},
		3: &sync.Mutex{},
		4: &sync.Mutex{},
	}

	// start the meal by iterating through our slice of Philosophers
	for i := 0; i < len(philosophers); i++ {
		// fire off each philosopher's goroutine
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	// Wait for the philosophers to finish. This blocks until the wait group is 0.
	wg.Wait()
	fmt.Printf("Order finished: %s\n", strings.Join(orderFinished, ", "))
	fmt.Println("The table is empty.")
}
