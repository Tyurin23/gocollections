package collections

type Vector struct {
	size    int
	storage []interface{}
}

func (v *Vector) Size() int {
	return v.size
}

// Appends the specified element in this vector
func (v *Vector) Push(val interface{}) {
	if v.storage == nil {
		v.storage = make([]interface{}, 1)
	}
	v.resizeIfNeed(v.size + 1)
	v.storage[v.size] = val
	v.size++
}

// Returns last element and remove it from vector
// Return nil if vector is empty
func (v *Vector) Pop() interface{} {
	if v.size == 0 {
		return nil
	}
	val := v.storage[v.size-1]
	v.storage[v.size-1] = nil
	v.size--
	v.resizeIfNeed(v.size)
	return val
}

// Inserts the element at the specified position in this vector.
// Shifts the element to the right.
func (v *Vector) Add(val interface{}, index int) {
	if !v.checkIndex(index) {
		panic("Array index out of bounds")
	}
	v.resizeIfNeed(v.size + 1)
	for i := v.size - 1; i >= index; i-- {
		v.storage[i+1] = v.storage[i]
	}
	v.storage[index] = val
	v.size++
}

// Replaces the element at the specified position in this vector with
// the specified element.
func (v *Vector) Set(val interface{}, index int) {
	if !v.checkIndex(index) {
		panic("Array index out of bounds")
	}
	v.storage[index] = val
}

// Returns the element at the specified position in this vector.
func (v *Vector) Get(index int) interface{} {
	if !v.checkIndex(index) {
		panic("Array index out of bounds")
	}
	return v.storage[index]
}

// Returns the index of the first occurrence  of the specified
// element in this vector, or -1 if this vector does not contain
// the element.
func (v *Vector) IndexOf(val interface{}) int {
	for i := 0; i < v.size; i++ {
		if val == v.storage[i] {
			return i
		}
	}
	return -1
}

// Removes the element at the specified position in this vector.
// Shifts any subsequent elements to the left.
func (v *Vector) Remove(index int) interface{} {
	if !v.checkIndex(index) {
		panic("Array index out of bounds")
	}
	val := v.storage[index]
	v.resizeIfNeed(v.size - 1)
	for i := index + 1; i < v.size; i++ {
		v.storage[i-1] = v.storage[i]
	}
	v.storage[v.size-1] = nil
	v.size--
	return val
}

func (v *Vector) checkIndex(index int) bool {
	if index < 0 && v.size >= index {
		return false
	}
	return true
}

func (v *Vector) resizeIfNeed(size int) {
	c := cap(v.storage)
	if size == c {
		v.resize(c * 2)
	} else if size <= c/4 {
		v.resize(c / 2)
	}
}

func (v *Vector) resize(newSize int) {
	n := make([]interface{}, newSize)
	copy(n, v.storage)
	v.storage = n
}
