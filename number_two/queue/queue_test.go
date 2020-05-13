package queue

import(	
	"testing"
	
	"reflect"
	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T){
	queue := New("jhon", "doe")
	queue.Enqueue("alex")

	assert.Equal(t, queue.Size(), 3)	
	assert.True(t, reflect.DeepEqual(queue.GetCustomers(), []interface{}{"jhon", "doe", "alex"}))
}