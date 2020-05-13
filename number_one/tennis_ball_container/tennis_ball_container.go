package tennis_ball_container

import(
	"errors"
	"math/rand"
	"fmt"
	"time"
)

// Containers structure
type Containers struct{
	array [][]int
}
// New factory to generate containers
func New(containerSize ...int) *Containers{
	container := &Containers{make([][]int, 0, len(containerSize))}	
	for _, val := range containerSize{
		container.AddContainer(int(val))
	}
	return container
}
// RandomizeArray factory to generate random array int
func RandomizeArray()[]int{
	var array []int
	rand.Seed(time.Now().UnixNano())
	totalData := rand.Intn(100)
	for count := 1; count <= totalData; count++{
		array = append(array, rand.Intn(100))
	}
	return array
}
// RandomizeNumber factory to generate random int
func RandomizeNumber(max int)int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}