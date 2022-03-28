package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMax(nums []int64, maxs int64, i *int, j *int) int64 {
	if nums[*i] >= nums[*j] {
		maxs += nums[*i]
		*i++
	} else {
		maxs += nums[*j]
		*j--
	}
	return maxs
}

func main() {
	var nums []int64
	input := bufio.NewReader(os.Stdin)
	str, _ := input.ReadString('\n')
	words := strings.Fields(str)
	for _, v := range words {
		// w, _ := strconv.Atoi(v)
		w, _ := strconv.ParseInt(v, 10, 64)
		nums = append(nums, w)
	}

	var first, second int64
	first = 0
	second = 0
	i := 0
	j := len(nums) - 1
	choose := 0
	for {
		if i > j {
			break
		}
		if choose == 0 {
			first = getMax(nums, first, &i, &j)
			choose = 1
		} else {
			second = getMax(nums, second, &i, &j)
			choose = 0
		}
	}
	fmt.Println(first, second)
}
