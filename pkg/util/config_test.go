package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {

	_, err := LoadConfig("../..")
	require.NoError(t, err)

	//fmt.Printf("%v", config.Token)
}
