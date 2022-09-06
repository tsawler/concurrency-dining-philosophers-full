package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Philosopher is a simple struct used to hold a philosopher's name, and the position of their left and
// right fork in the map that holds the mutexes for those forks.
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers is our list of philosophers. We define their name, assign a left and right fork using ints that
// match the map of forks. Note that each philosopher shares one fork with the person next to them (five philosophers,
// and five forks).
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// Define a few variables.
var hunger = 3                  // how many times a philosopher eats
var eat = 1 * time.Second       // how long it takes to eat
var think = 3 * time.Second     // how long a philosopher thinks
var sleepTime = 1 * time.Second // how long to wait when printing things out
var orderFinished []string      // the order in which philosophers finish dining and leave
var orderMutex sync.Mutex       // a mutex for the slice orderFinished

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")
	time.Sleep(sleepTime)

	dine()

	fmt.Printf("Order finished: %s.\n", strings.Join(orderFinished, ", "))
	fmt.Println("The table is empty.")
}

func dine() {
	// Uncomment the next three lines to set delays to zero while developing to speed things up.
	//eat = 0 * time.Second
	//sleepTime = 0 * time.Second
	//think = 0 * time.Second

	// wg is the WaitGroup that keeps track of how many philosophers are still at the table. When
	// it reaches zero, everyone is finished eating and has left. We add 5 (the number of philosophers) to this
	// wait group.
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// we want everyone to be seated before they start eating, so create a WaitGroup for that, and set it to 5.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks. Forks are assigned using the fields leftFork and rightFork in the Philosopher type.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < 5; i++ {
		forks[i] = &sync.Mutex{}
	}

	// Start the meal by iterating through our slice of Philosophers.
	for i := 0; i < len(philosophers); i++ {
		// fire off each philosopher's goroutine
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	// Wait for the philosophers to finish. This blocks until the wait group is 0.
	wg.Wait()
}

// diningProblem is the function fired off as a goroutine for each of our philosophers. It takes one
// philosopher, our WaitGroup to determine when everyone is done, a map containing the mutexes for every
// fork on the table, and a WaitGroup used to pause execution of every instance of this goroutine
// until everyone is seated at the table.
func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	// Decrement our WaitGroup by one when this goroutine exits.
	defer wg.Done()
	fmt.Println(philosopher.name, "is seated at the table.")

	// Decrement the seated WaitGroup by one.
	seated.Done()

	// Wait until everyone is seated.
	seated.Wait()

	// We'll also define a WaitGroup for this goroutine specifically, so that the "leaving table"
	// message does not get printed until after the message that the forks are dropped.
	eating := &sync.WaitGroup{}
	eating.Add(hunger)

	// Have this philosopher eat and think "hunger" times (3).
	for i := hunger; i > 0; i-- {
		// Get a lock on the left and right forks. We have to choose the lower numbered fork first in order
		// to avoid a logical race condition, which is not detected by the -race flag in tests; if we don't do this,
		// we have the potential for a deadlock, since two philosophers will wait endlessly for the same fork.
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			forks[philosopher.leftFork].Lock()
		} else {
			forks[philosopher.leftFork].Lock()
			forks[philosopher.rightFork].Lock()
		}

		fmt.Printf("\t%s has both forks, and is eating.\n", philosopher.name)
		time.Sleep(eat)

		// The philosopher starts to think, but does not drop the forks yet.
		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(think)

		// Unlock the mutexes for both forks.
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)

		// Decrement the eating WaitGroup by 1.
		eating.Done()
	}

	// Wait until all messages have been printed.
	eating.Wait()

	fmt.Println(philosopher.name, "is satisfied.")
	time.Sleep(sleepTime)
	fmt.Println(philosopher.name, "left the table.")

	// Update the list of finished diners.
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}
