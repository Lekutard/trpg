package main

import (
	"app/usecase"
	"fmt"
	"strconv"
)

func main() {
	roll := usecase.GraspDices(6, 2)
	type sum int
	type details []int
	for i := 0; i < 10; i++ {
		sum, details := roll.ThrowDices()
		detailStr := "["
		for i := 0; i < len(details); i++ {
			if i > 0 {
				detailStr = detailStr + ","
			}
			detailStr = detailStr + strconv.Itoa(details[i])
		}
		detailStr = detailStr + "]"

		fmt.Println(sum)
		fmt.Println(detailStr)
	}
}
