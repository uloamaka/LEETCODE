type SubrectangleQueries struct {
    rectangle_r [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
    rez := SubrectangleQueries{rectangle}
    return rez
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1, col1, row2, col2, newValue int)  {
    for i := row1; i <= row2; i++ {
        for j := col1; j <= col2; j++ {
            this.rectangle_r[i][j] = newValue
        }
    }
}


func (this *SubrectangleQueries) GetValue(row, col int) int {
    return this.rectangle_r[row][col]
}


/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */