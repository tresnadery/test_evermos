package main

import (
	"fmt"
	"github.com/tresnadery/test_evermos/number_two/queue"
)

func main() {
	totalProduct := 5
	queue := queue.New("jhon", "doe", "Audrey", "Clarissa")
	queue.Enqueue("Maisey")
	queue.Enqueue("Kirstin")

	for i := totalProduct; i > 0; i-- {
		queue.Dequeue()
		totalProduct--
		if totalProduct <= 0 {
			fmt.Println("Sold Out")
			break
		}
	}
	if queue.IsEmpty() {
		fmt.Println("queue is clear")
		return
	}
	fmt.Printf("Person on queue : %+v \n", queue.GetCustomers())
}
