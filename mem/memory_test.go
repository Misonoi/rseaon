package mem

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClone(t *testing.T) {
	type TestS struct {
		A int32
		B int64
		C *int64
	}

	c := new(int64)
	*c = 5

	A := &TestS{
		A: 1,
		B: 0,
		C: c,
	}

	B := Clone(A)

	println(&A)
	println(&B)
	println(A.C)
	println(B.C)

	fmt.Printf("%T", B)

	assert.Equal(t, A.B, B.B)
	assert.Equal(t, A.A, B.A)
}
