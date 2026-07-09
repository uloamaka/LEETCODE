/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, h := 0, n

    for l <= h {
        mid := l + (h-l)/2

        switch guess(mid) {
        case 0:
            return mid
        case 1:
            l = mid + 1
        default: // -1
            h = mid
        }
    }

    return 0
}