package hammingdistance

import "fmt"

func hammingDistance(a, b int) int {
	var ret int = 0
	if a&1 != b&1 {
		ret++
	}
	if a&2 != b&2 {
		ret++
	}
	if a&4 != b&4 {
		ret++
	}
	if a&8 != b&8 {
		ret++
	}
	if a&16 != b&16 {
		ret++
	}
	if a&32 != b&32 {
		ret++
	}
	if a&64 != b&64 {
		ret++
	}
	if a&128 != b&128 {
		ret++
	}
	fmt.Printf("HammingDistance: %08b and %08b = %d\n", byte(a), byte(b), ret)
	return ret
}

// func totalHammingDistance(nums []int) int {
// 	res := 0
// 	for i, len := uint(0), len(nums); i < 32; i++ {
// 		ones := 0
// 		for _, val := range nums {
// 			if val>>i&1 == 1 {
// 				ones++
// 			}
// 		}
// 		if ones > 0 {
// 			res += ones * (len - ones)
// 		}
// 	}
// 	return res
// }
