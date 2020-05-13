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

func TestPeek(t *testing.T){
	queue := New()

	queue.Enqueue("jhon")
	customer, err := queue.Peek()
	assert.Nil(t, err)
	assert.Equal(t, customer, "jhon")

	queue.Dequeue()	
	_, err = queue.Peek()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "queue is empty")
	
	queue.Enqueue("doe")
	customer, err = queue.Peek()
	assert.Nil(t, err)
	assert.Equal(t, customer, "doe")
}
func TestGetCustomers (t *testing.T){
	queue := New()

	queue.Enqueue("jhon", "doe", "alex")
	assert.True(t, reflect.DeepEqual(queue.GetCustomers(), []interface{}{"jhon", "doe", "alex"}))
}
