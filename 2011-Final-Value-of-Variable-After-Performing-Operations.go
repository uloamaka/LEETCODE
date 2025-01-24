func finalValueAfterOperations(operations []string) int {
    x := 0
    for _, operation := range operations {
        switch operation {
            case "++X", "X++":
                x++
            case "--X", "X--":
                x--
        }
    }
    return x
}
    // intialise the value of x to be zero value of int
    // loop throught the slice of operations for each string perform rhe operation on x
    // return the final value of x 