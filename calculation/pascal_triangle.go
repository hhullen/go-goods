package calculation

func GeneratePascalTriangle(numRows int) (res [][]int) {
	if numRows <= 0 {
		return
	}

    res = make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        res[i] = make([]int, i+1)
    }
    res[0][0] = 1

    var recF func(int,int) int
    recF = func(row, col int) int {
        if row < 0 || col < 0 || col >= len(res[row]) {
            return 0
        }

        if res[row][col] == 0 {
            res[row][col] = recF(row-1, col-1) + recF(row-1, col)
        }

        return res[row][col]
    }

    lastArr := res[len(res)-1]
    for i := 0; i < len(lastArr); i++ {
        lastArr[i] = recF(numRows-1, i)
    }

    return
}

func GenerateRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }

    arr1 := make([]int, 1, rowIndex+1)
    arr2 := make([]int, 2, rowIndex+1)
    arr1[0] = 1

    for i := 0; i < rowIndex; i++ {
        for j := 0; j < len(arr2); j++ {
            summ := 0
            if j-1 >= 0 {
                summ += arr1[j-1]
            }
            if j < len(arr1) {
                summ += arr1[j]
            }

            arr2[j] = summ
        }

        arr1 = append(arr1, 0, 0)
        arr1, arr2 = arr2, arr1
    }

    return arr1
}
