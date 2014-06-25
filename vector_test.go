package collections

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNewVector(t *testing.T) {
	vec := new(Vector)
	assert.NotNil(t, vec)
	assert.Equal(t, 0, vec.Size())
}

func TestPush(t *testing.T) {
	vec := new(Vector)
	val := 10
	vec.Push(val)

	assert.Equal(t, 1, vec.Size())
}

func TestManyPushes(t *testing.T) {

}

func TestGet(t *testing.T) {
	vec := new(Vector)
	val := 10
	vec.Push(val)

	assert.Equal(t, 1, vec.Size())
	assert.Equal(t, val, vec.Get(0))
	// assert.Panics(t, func() { vec.Get(0) })
}

func TestSet(t *testing.T) {
	vec := zeroVector(10)
	val := 10
	ind := 3

	old := vec.Get(ind)
	assert.Equal(t, 10, vec.Size())
	assert.Equal(t, 0, old)
	vec.Set(val, ind)
	assert.Equal(t, 10, vec.Size())
	assert.Equal(t, val, vec.Get(ind))

}

func TestAdd(t *testing.T) {
	vec := zeroVector(10)
	val := 10
	ind := 3

	old := vec.Get(ind)
	assert.Equal(t, 10, vec.Size())
	assert.Equal(t, 0, old)
	vec.Add(val, ind)
	assert.Equal(t, 11, vec.Size())
	assert.Equal(t, val, vec.Get(ind))
}

func TestIndexOf(t *testing.T) {
	vec := zeroVector(10)
	val := 10
	ind := 3
	vec.Set(val, ind)

	assert.Equal(t, ind, vec.IndexOf(val))
}

func TestManyPushPop(t *testing.T) {
	vec := new(Vector)
	size := 10000

	sz := 0
	for i := 0; i < size; i++ {
		if r := rand.Int(); r%2 == 0 {
			sz++
			vec.Push(i)
			assert.Equal(t, sz, vec.Size())
		} else {

			val := vec.Pop()
			if val != nil {
				sz--
			}
			assert.Equal(t, sz, vec.Size())
		}

	}
}

func BenchmarkPush(b *testing.B) {
	vec := new(Vector)
	val := 1000
	for i := 0; i < b.N; i++ {
		vec.Push(val)
	}
}

func zeroVector(size int) *Vector {
	vec := new(Vector)
	for i := 0; i < size; i++ {
		vec.Push(0)
	}
	return vec
}
