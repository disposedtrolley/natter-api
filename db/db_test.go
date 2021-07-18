package db_test

import (
	"testing"

	"github.com/disposedtrolley/natter-api/db"
	"github.com/stretchr/testify/require"
)

func TestNewInMemory(t *testing.T) {
	_, err := db.NewInMemory()
	require.Nil(t, err)
}
