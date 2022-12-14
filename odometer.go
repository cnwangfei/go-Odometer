package Odometer

var carryCount int // 记录连续进位次数
// For
//
//	！！小心死循环
//	@描述: 里程表循环函数  当连续进位次数==里程表切片长度时，自动终止循环
//	@参数 odometer []int 里程表切片
//	@参数 fun func() (isCarry ,isEnd bool,err error) isCarry里程进位标志，ture即进位。isEnd 主动终止循环标志。err 错误
//	@返回 error
func For(odometer []int, fun func() (isCarry, isEnd bool, err error)) error {
	var isCarry bool
	// 像里程表一样的循环：末位数走完后归0，左一位+1
	length := len(odometer)
	for 1 == 1 {
		// 判断末位是否走到头了，即应该进位，如果不进位，就进入下一个循环继续+1
		// 具体判定是否进位的代码
		{
			carry, end, err := fun()
			if err != nil {
				return err
			}
			if end {
				return nil
			}
			isCarry = carry
		}
		if !isCarry { // 不进位时
			odometer[length-1]++ // 末位+1
			carryCount = 0       // 进位次数归0
			continue
		}
		carryCount++ // 进位时，进位次数+1
		// 如果连续进位次数==里程表位数，自动终止循环
		if carryCount == length {
			return nil
		}

		// 进位时，倒序检查，从倒数第二位开始向左检查
		for i := length - 2; i >= 0; i-- {
			// 如果当前右边一位！=0，那么代表当前需要进一位，它的右边通通归0,然后继续下一个循环
			if odometer[i+1] != 0 /*&& odometer[i] == 0*/ {
				odometer[i]++
				for i2 := range odometer {
					if i2 > i {
						odometer[i2] = 0
					}
				}
				break
			}

			// 如果末位！=0，而当前位!=0 就检查下一位,直到i==0代表里程表用完，返回err
			if i == 0 { //全部检查完
				//odometer is end
				return nil
			}
		}
	}
	return nil
}
