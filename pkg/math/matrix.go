// В pkg/math/matrix.go или pkg/matrix/rotate.go
package math // или package matrix

// Rotate90 принимает квадратную матрицу mat и возвращает новую, повернутую по часовой стрелке на 90°.
func Rotate90(mat [][]int) [][]int {
	n := len(mat)
	if n == 0 || len(mat[0]) != n {
		return nil // или выбрасывайте ошибку, как вам удобнее
	}
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[j][n-1-i] = mat[i][j]
		}
	}
	return res
}
