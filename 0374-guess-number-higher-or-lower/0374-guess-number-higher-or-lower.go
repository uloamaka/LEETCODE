/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left, right := 0, n

    for left < right {
        midNum := left + (right-left) / 2

        guessRes := guess(midNum)
        if guessRes == 0 {
            return midNum
        } else if guessRes == -1 {
            right = midNum
        } else {
            left = midNum+1
        }
    }
    return n
}