package intersection

//import "errors"

func intersection(strArr1 []string, strArr2 []string) []string {
	var jiaojiArr []string

	var mapArr1 map[string]bool
	mapArr1 = make(map[string]bool)
	for _, item := range strArr1 {
		mapArr1[item] = true
	}

	for _, item := range strArr2 {
		if _, ok := mapArr1[item]; ok {
			jiaojiArr = append(jiaojiArr, item)
		}
	}
	return jiaojiArr
}

func Intersection(strArr1, strArr2 []string, args ...interface{}) []string {
	if len(args) == 0 {
		return intersection(strArr1, strArr2)
	} else {
		jiaojiArr := intersection(strArr1, strArr2)
		for _, item := range args {
			if value, ok := item.([]string); ok {
				jiaojiArr = intersection(jiaojiArr, value)
				if len(jiaojiArr) == 0 {
					return jiaojiArr
				}
			}

		}
		return jiaojiArr
	}
}



//func Intersection(args ...interface{}) ([]string, error) {
//	var strArr1, strArr2 []string
//	if len( args ) < 2 {
//		return nil, errors.New("args length must >= 2")
//	} else {
//		value0, ok0 := args[0].([]string)
//		value1, ok1 := args[1].([]string)
//		if len(args) == 2 {
//			if  ok0 && ok1 {
//				arr := intersection(value0, value1)
//				return arr, nil
//			}else{
//				return nil, errors.New("args item must be string array")
//			}
//		} else {
//			jiaojiArr := intersection(strArr1, strArr2)
//			for _, item := range args {
//				if value, ok := item.([]string); ok {
//					jiaojiArr = intersection(jiaojiArr, value)
//					if len(jiaojiArr) == 0 {
//						return jiaojiArr
//					}
//				}
//
//			}
//			return jiaojiArr
//		}
//	}
//}
