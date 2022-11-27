package main

func main() {
	for i, v := range solveA(50) {
		if v == 1 {
			println(i)
		}
	}
}

// func main() {
// 	const n = 50

// 	ansA := solveA(n)
// 	ansADash := make([]int, 0)
// 	for i, v := range ansA {
// 		if v == 1 {
// 			ansADash = append(ansADash, i)
// 		}
// 	}
// 	fmt.Println(ansADash)
// 	sort.Ints(ansADash)

// 	ansB := solveB(n)
// 	ansBDash := make([]int, 0)
// 	for k, v := range ansB {
// 		if v {
// 			ansBDash = append(ansBDash, k)
// 		}
// 	}
// 	sort.Ints(ansBDash)

// 	fmt.Println(reflect.DeepEqual(ansADash, ansBDash)) // true
// }

func solveA(n int) []int {
	// ans[素数]=1にする
	// 開始時は全要素を1にしておき、除外する数は0にする
	ans := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		ans[i] = 1
	}

	for i := 2; i <= n; i++ {
		// 除外済みかどうか
		if ans[i] == 0 {
			continue
		}

		for j := i + 1; j <= n; j++ {
			// 除外済みかどうか
			if ans[j] == 0 {
				continue
			}

			// 除外対象かどうか
			if j%i == 0 {
				ans[j] = 0
			}
		}
	}
	return ans
}

func solveB(n int) map[int]bool {
	ans := make(map[int]bool, n)

	for i := 2; i <= n; i++ {
		// 確認済みかどうか
		if _, ok := ans[i]; ok {
			continue
		}
		ans[i] = true

		for j := i + 1; j <= n; j++ {
			// 確認済みかどうか
			if _, ok := ans[j]; ok {
				continue
			}

			// 除外対象かどうか
			if j%i == 0 {
				ans[j] = false
			}
		}
	}
	return ans
}
