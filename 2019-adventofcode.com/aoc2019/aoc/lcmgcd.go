package aoc

func Gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func LcmArray(arr []int64) int64 {
	lcm := arr[0]
	for i := 1; i < len(arr); i++ {
		lcm = Lcm(lcm, arr[i])
	}
	return lcm
}

func Lcm(x, y int64) int64 {
	var gcd int64
	if x > y {
		gcd = Gcd(x, y)
	} else {
		gcd = Gcd(y, x)
	}
	return x * y / gcd
}
