package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cover_PartiallyCovered(t *testing.T) {
	assert.Equal(t, "test****", Cover("testtest", 4))
}

func Test_Cover_AllCovered(t *testing.T) {
	assert.Equal(t, "********", Cover("testtest", 0))
}

func Test_Cover_PartiallyCovered_UTF8(t *testing.T) {
	assert.Equal(t, "Hello,***", Cover("Hello, 世界", 6))
}

func Test_Cover_PartiallyCovered_UTF8_Reverse(t *testing.T) {
	assert.Equal(t, "世界, He***", Cover("世界, Hello", 6))
}
