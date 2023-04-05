package a

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var name, expectName, actualName string

func TestValid(t *testing.T) {
	require.Equal(t, expectName, "")
	require.Equal(t, "", actualName)
	require.Equal(t, expectName, name)
	require.Equal(t, name, actualName)
	require.Equal(t, expectName, actualName)
}

func TestInValid(t *testing.T) {
	require.Equal(t, actualName, "")         // want "the actual-like variable name must comes after the expect-like variable name"
	require.Equal(t, "", expectName)         // want "the expected-like variable name must comes before the actual-like variable name"
	require.Equal(t, actualName, name)       // want "the actual-like variable name must comes after the expect-like variable name"
	require.Equal(t, name, expectName)       // want "the expected-like variable name must comes before the actual-like variable name"
	require.Equal(t, actualName, expectName) // want "the actual-like variable name must comes after the expect-like variable name"
}
