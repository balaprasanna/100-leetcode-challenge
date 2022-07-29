func isPalindrome(x int) bool {
	var y int = x
	var rev int = 0

	for y > 0 {
		rem := y % 10
		y = y / 10
		rev = (rev * 10) + rem
	}
	return x == rev
}