package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	assert.Equal(t, Instance(), singleton)
}

func BenchmarkInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if singleton != Instance() {
			b.Fail()
		}
	}
}
