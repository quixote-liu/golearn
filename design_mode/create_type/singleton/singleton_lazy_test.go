package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, GetLazyInstance(), singleton)
}

func BenchmarkGetLazyInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if GetLazyInstance() != singleton {
			b.Fail()
		}
	}
}
