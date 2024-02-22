package deep

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMax(t *testing.T) {
	i, err := Max([]int{2, 3, 1})
	require.NoError(t, err)
	require.Equal(t, 3, i)

	_, err = Max[float64](nil)
	require.Error(t, err)

	s, err := Max([]string{"B", "C", "A"})
	require.NoError(t, err)
	require.Equal(t, s, "C")
}

func TestRelu(t *testing.T) {
	out := Relu(7)
	require.Equal(t, 7, out)
	out = Relu(-2)
	require.Equal(t, 0, out)

	fout := Relu(7.0)
	require.Equal(t, 7.0, fout)

	mout := Relu(time.April)
	require.Equal(t, time.April, mout)
}

/*
func TestReluInt(t *testing.T) {
	out := ReluInt(7)
	require.Equal(t, 7, out)
	out = ReluInt(-2)
	require.Equal(t, 0, out)
}

func TestReluFloat64(t *testing.T) {
	out := ReluInt(7)
	require.Equal(t, 7, out)
	out = ReluInt(-2)
	require.Equal(t, 0, out)
}

*/

func f() {
	time.Sleep(100 * time.Millisecond)
	const sleepTime = 100
	time.Sleep(sleepTime * time.Millisecond)

	n := 100
	// time.Sleep(n * time.Millisecond)
	time.Sleep(time.Duration(n) * time.Millisecond)
}
