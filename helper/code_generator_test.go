package helper

import (
	"crypto/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateCode(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		var r = require.New(t)

		code, err := GenerateCode(1)
		r.NoError(err)

		r.True(len(code) > 0)
	})

	t.Run("Error Rand", func(t *testing.T) {
		var r = require.New(t)

		var randBC = rand.Reader
		rand.Reader = strings.NewReader("")

		_, err := GenerateCode(1)
		r.Error(err)

		rand.Reader = randBC
	})
}
