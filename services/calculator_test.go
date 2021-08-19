package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/viniciusmazon/go-test/services"
)

// func TestSum(t *testing.T) {
// 	if services.Sum(2, 2) != 4 {
// 		t.Error("it should be 4")
// 	}
// }

// func TestMultiply(t *testing.T) {
// 	if services.Multiply(2, 2) != 4 {
// 		t.Error("it should be 4")
// 	}
// }

func TestSum(t *testing.T) {
	require.Equal(t, services.Sum(2, 2), 4)
}

func TestMultiply(t *testing.T) {
	require.Equal(t, services.Multiply(2, 2), 4)
}
