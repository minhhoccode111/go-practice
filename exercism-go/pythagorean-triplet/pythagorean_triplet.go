package pythagorean

import "sort"

func Sum(p int) (res []Triplet) {
	for m := 2; m < p/2; m++ {
		for n := 1; n < m; n++ {
			if (m-n)%2 == 0 || gcd(m, n) != 1 {
				continue
			}
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			sum := a + b + c
			if p%sum != 0 {
				continue
			}
			k := p / sum
			ka, kb, kc := k*a, k*b, k*c
			triplet := Triplet{ka, kb, kc}
			sort.Ints(triplet[:])
			res = append(res, triplet)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		for x := 0; x < 3; x++ {
			if res[i][x] != res[j][x] {
				return res[i][x] < res[j][x]
			}
		}
		return false
	})
	return
}

func Range(min, max int) (res []Triplet) {
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			c2 := a*a + b*b
			c := isqrt(c2)
			if c > max || c < b || c*c != c2 {
				continue
			}
			res = append(res, Triplet{a, b, c})
		}
	}
	return
}

type Triplet [3]int

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isqrt(n int) int {
	x := 0
	for x*x <= n {
		x++
	}
	return x - 1
}
