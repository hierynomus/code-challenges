package aoc

// // This file is for stuff related to the chinese remainder theorem!

// // Solves x=a mod m; x=b mod n by using the chinese remainder theorem.
// func SolveCrt(a, m, b, n int) int {
// 	s, t, _ := ExtendedGcd(m, n)
// 	return Lpr(b*s*m+a*n*t, m*n)
// }

// // Represents an entry in the Extended Chinese Remainder Theorem
// type CrtEntry struct {
// 	A, M int
// }

// type AModN struct {
// 	A, N int
// }

// // Solves the solution to x=(a1 mod m1); x=(a2 mod m2); x=...
// //
// // If len(eqs) == 0, it panics.
// func SolveCrtMany(eqs []CrtEntry) int {
// 	if len(eqs) == 0 {
// 		panic("cannot have 0 entries to solve")
// 	}
// 	if len(eqs) == 1 {
// 		return Lpr(eqs[0].A, eqs[0].M)
// 	}
// 	eqs2 := make([]CrtEntry, len(eqs))
// 	copy(eqs2, eqs)

// 	for i := 1; i < len(eqs2); i++ {
// 		x := SolveCrt(eqs2[i-1].A, eqs2[i-1].M, eqs2[i].A, eqs2[i].M)
// 		eqs2[i] = CrtEntry{x, eqs2[i-1].M * eqs2[i].M}
// 	}
// 	return eqs2[len(eqs2)-1].A
// }

// def solve_crt(pair_modular_equations):
//     remainder_pair = (0, 1) # remainder, coefficient

//     for modular_equation in pair_modular_equations:
//         coefficient = remainder_pair[1]

//         for k in range(1, modular_equation[1]):
//             if (coefficient * k) % modular_equation[1] == 1:
//                 remainder_pair = ((((modular_equation[0] - remainder_pair[0]) * k) % modular_equation[1]) * remainder_pair[1] + remainder_pair[0], remainder_pair[1] * modular_equation[1])
//                 break

//     return remainder_pair

// func NewAModN(a, n) AModN {
// 	return AModN{A: -a % n, N: n}
// }

// 	def get_modular_equations(bus_ids):
//     k = 0

//     modular_equations = list()

//     for bus_id in bus_ids:
//         if bus_id == 'x':
//             k += 1
//             continue

//         modular_equations.append((-k % bus_id, bus_id))
//         k += 1

//     return modular_equations

// func SolveCrtManyIntern(eqs []CrtEntry) int {
// 	f := eqs[0]
// 	s := eqs[1]
// 	x := SolveCrt(f.A, f.M, s.A, s.M)
// 	if len(eqs) == 2 {
// 		return x
// 	}
// 	eqs[1] = CrtEntry{x, f.M * s.M}
// 	return SolveCrtManyIntern(eqs[1:])
// }

// // Finds x and y such that: Gcd(a, b) = ax + by. (By the extended euclidean algorithm)
// //
// // This implementation is based on
// // https://en.wikibooks.org/wiki/Algorithm_Implementation/Mathematics/Extended_Euclidean_algorithm#Iterative_algorithm_3
// func ExtendedGcd(a, b int) (x, y, gcd int) {
// 	x0, x1, y0, y1 := 1, 0, 0, 1

// 	for a != 0 {
// 		var q int
// 		q, b, a = b/a, a, b%a
// 		x0, x1 = x1, x0-q*x1
// 		y0, y1 = y1, y0-q*y1
// 	}
// 	return y0, x0, b
// }

// // Finds the least positive residue of a number
// // in a given modulus. Note that this is very slightly
// // different from the remainder (%) operator when working
// // with negative numbers.
// func Lpr(a, m int) int {
// 	return (a%m + m) % m
// }
