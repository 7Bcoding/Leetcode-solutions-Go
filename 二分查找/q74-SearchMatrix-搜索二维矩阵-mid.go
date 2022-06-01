package 二分查找

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i <= len(matrix)-1; i++ {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] > target {
				break
			}
		}
	}
	return false
}
