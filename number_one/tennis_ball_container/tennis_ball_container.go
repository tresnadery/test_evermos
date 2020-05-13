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