package mem

import (
	"github.com/mitchellh/copystructure"
)

// Clone Deep copy any object.
func Clone[T any](value T) T {
	b, _ := copystructure.Copy(value)
	cloned, _ := b.(T)
	return cloned
}
