//go:build gendb

package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseRSS(t *testing.T) {
	const rssFile = "testdata/rss.xml"
	file, err := os.Open(rssFile)
	require.NoErrorf(t, err, "open %q", rssFile)
	defer file.Close()

	items, err := parseRSS(file)
	require.NoError(t, err, "parse")
	require.Equal(t, 72, len(items))
	require.Contains(t, items[0].Title, "presidential immunity")
	require.Contains(t, items[len(items)-1].Title, "budget is tight")
}
