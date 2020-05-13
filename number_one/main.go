package main

import(
	"fmt"
	tennisBallContainer "github.com/tresnadery/test_evermos/number_one/tennis_ball_container"	
)
func main(){
	index := 0
	containers := tennisBallContainer.New()
	totalContainer := tennisBallContainer.RandomizeArray()
	for _, val := range totalContainer{
		containers.AddContainer(val)
	}
	for{
		index = tennisBallContainer.RandomizeNumber(len(totalContainer))
		containers.Put(index)
		if containers.IsFull(index){
			break
		}
	}
	fmt.Printf("Container number %d is full \n", index + 1)
}