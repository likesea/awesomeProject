package main

func minWindow(s string, t string) string {
	need := make(map[string]int)
	var sub string = s
	found := false
	for _, b := range t {
		need[string(b)] = 0
	}
	left, right := 0, 0
	for ; right < len(s); {
		k := string(s[right])
		right++
		if _, ok := need[k]; ok {
			need[k] += 1
		}
		for ; shrink(need); {
			found = true
			k := string(s[left])
			//左边滑动后，满足
			if _, ok := need[k]; ok {
				need[k] -= 1
				//仍满足
				if need[k]>0{
					temp := s[left+1 : right]
					if len(temp) <= len(sub) {
						sub = temp
					}
					left++
				}else{
					temp := s[left : right]
					if len(temp) <= len(sub) {
						sub = temp
					}
					//不满足,则跳出
					left++
					break
				}
			}else{
				left++
			}
		}
	}
	if found{
		return sub
	}else{
		return ""
	}
}
func shrink(need map[string]int) bool {
	for _, v := range need {
		if v == 0 {
			return false
		}
	}
	return true
}
