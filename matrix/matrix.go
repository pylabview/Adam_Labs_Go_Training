package matrix

type Matrix [][]int

func New(rows, cols int) Matrix {
	m := make([][]int, rows)
	for r := range m {
		m[r] = make([]int, cols)
	}

	return m
}

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Cols() int {
	return len(m[0])
}

func (m Matrix) SumRows() int {
	total := 0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			total += m[r][c]
		}
	}
	return total
}

func (m Matrix) SumCols() int {
	total := 0
	for c := 0; c < m.Cols(); c++ {
		for r := 0; r < m.Rows(); r++ {
			total += m[r][c]
		}
	}
	return total
}
