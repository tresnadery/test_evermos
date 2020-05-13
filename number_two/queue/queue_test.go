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

func TestDequeue(t *testing.T){
	queue := New("jhon", "doe")

	customer, err := queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)
	assert.True(t, reflect.DeepEqual(queue.GetCustomers(), []interface{}{"doe"}))
	assert.Equal(t, customer, "jhon")

	customer, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, queue.IsEmpty(), true)
	assert.Equal(t, queue.Size(), 0)
	assert.Equal(t, customer, "doe")

	_, err = queue.Dequeue()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "queue is empty")

}