import(
	"testing"
	"fmt"
	"reflect"	

	"github.com/stretchr/testify/assert"
)


func TestVerifiedContainer(t *testing.T){	
	containers := New()	
	for i := 0; i < 2; i++{
		containers.AddContainer(2)
	}
	containers.Put(0)	
	containers.Put(1)
	containers.Put(1)

	assert.Equal(t, containers.IsFull(1), true)
	assert.Equal(t, containers.Size(1), containers.Capacity(1))
	assert.Equal(t, fmt.Sprintf("Container number %d is full", 2), containers.Verified(1))
	assert.True(t, reflect.DeepEqual(containers.array[1], []int{1, 1}))
}