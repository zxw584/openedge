package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	assert.Equal(t, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, U64ToB(1))
	assert.Equal(t, uint64(1), U64([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}))
}
