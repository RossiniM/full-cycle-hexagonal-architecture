package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestErrorJson(t *testing.T) {
	msg := "Hello json"

	result := jsonError(msg)

	require.Equal(t, []byte(`{"message":"Hello json"}`), result)
}
