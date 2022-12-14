package Odometer

import (
	"fmt"
	"log"
)

func ExampleFor() {
	o := make([]int, 5)
	end := []int{0, 1, 2, 3, 4} // 里程跑到1234的时候终止
	count := 0
	err := For(o, func() (isCarry, isEnd bool, err error) {
		// 九进一
		for _, v := range o {
			if v > 9 {
				isCarry = true
				break
			} else {
				isCarry = false
			}
		}
		if !isCarry {
			fmt.Println(o)
		}
		// 打上终止标志
		isEnd = true
		for i, v := range end {
			if o[i] != v {
				isEnd = false
				break
			}
		}
		if isEnd {
			return
		}

		if !isCarry && !isEnd {
			count++
		}

		if len(o) > 5 {
			err = fmt.Errorf("this is error")
			return
		}

		return
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("里程表结果:%v  循环次数:%v\n", o, count)
}
