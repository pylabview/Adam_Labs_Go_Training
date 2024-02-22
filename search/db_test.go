package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	db := NewDB()
	matches := db.Search("president")
	require.Equal(t, 46, len(matches))
}
