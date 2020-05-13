package tennis_ball_container

import(
	"errors"
	"math/rand"
	"fmt"
	"time"
)
var (
	errEmptyContainer = errors.New("container is empty")
	errIndexLessThanZero = errors.New("index is less than 0")
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
// AddContainer add container with size
func (c *Containers) AddContainer(size int){
	c.array = append(c.array, make([]int, 0, size))
}
// Put add ball to container
func (c *Containers) Put(index int){	
	c.array[index] = append(c.array[index], 1)	
}
// Remove remove ball from container
func (c *Containers) Remove(index int)(error){
	if c.IsEmpty(index){
		return errEmptyContainer
	}	
	c.array[index] = c.array[index][:len(c.array[index])- 1]
	return nil
}
// Size return size of the Container
func (c *Containers) Size(index int)int{
	return len(c.array[index])
}
// Capacity return capacity of the Container
func (c *Containers) Capacity(index int) int{
	return cap(c.array[index])
}
// IsFull checks if the container is full
func (c *Containers) IsFull(index int) bool{
	return c.Size(index) == c.Capacity(index)
}
// IsEmpty checks if the container is empty
func (c *Containers) IsEmpty(index int) bool{	
	return c.Size(index) == 0
}
// Verified print container verified
func (c *Containers) Verified(index int)string{
	return fmt.Sprintf("Container number %d is full", index+1)
}