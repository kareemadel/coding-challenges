func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	var layers int
	if m <= n {
		layers = (m + 1) / 2
	} else {
		layers = (n + 1) / 2
	}
	var result []int
	counter := m * n
	for layer := 0; layer < layers; layer++ {
		first := layer
		lastRow, lastColumn := m-layer-1, n-layer-1
		i, j := layer, layer
		for ok := true; ok; ok = (counter > 0 && (i != layer || j != layer)) {
			result = append(result, matrix[i][j])
			i, j = move(i, j, first, lastRow, lastColumn)
			counter--
		}
	}
	return result
}

func move(i, j, first, lastRow, lastColumn int) (int, int) {
	if i == first && j < lastColumn {
		j++
	} else if j == lastColumn && i < lastRow {
		i++
	} else if i == lastRow && j > first {
		j--
	} else if j == first && i > first {
		i--
	}
	return i, j
}