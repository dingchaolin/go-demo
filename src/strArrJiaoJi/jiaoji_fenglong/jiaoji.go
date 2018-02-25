package intersection

func StrArr(args ...[]string) []string {
	sumNum := len(args)
	if sumNum == 0 {
		return []string{}
	}
	if sumNum == 1 {
		return args[0]
	}
	if sumNum == 2 {
		if len(args[0]) > len(args[1]) {
			return intersection(args[0], args[1])
		} else {
			return intersection(args[0], args[1])
		}
	}
	// now >= 3
	NewArr := intersection(args[0], args[1])
	// 拿 NewArr(来自0，1) 分别和 第 2，3，4，5，6 ... 的做比较
	for i := 2; sumNum > i; i++ {
		if len(NewArr) > len(args[i]) {
			NewArr = intersection(NewArr, args[i])
		} else {
			NewArr = intersection(args[i], NewArr)
		}
	}
	return NewArr[:]
}

func intersection(args1, args2 []string) []string {
	// default len(args1) > len(args2)
	var NewArr []string;
	j := len(args2)
	k := len(args1)
	for i := 0; i < j; i++ {
		for m := 0; m < k; m++ {
			if args2[i] == args1[m] {
				NewArr = append(NewArr, args1[m])
			}

		}
	}
	return NewArr[:]
}