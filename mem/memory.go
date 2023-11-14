package mem

import (
	"github.com/mitchellh/copystructure"
)

func Clone[T any](value T) T {
	b, _ := copystructure.Copy(value)
	cloned, _ := b.(T)
	return cloned
}
