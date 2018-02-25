package intersection

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

