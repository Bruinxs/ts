package ts

func CmpStr(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, sv1 := range s1 {
		equal := false
		for _, sv2 := range s2 {
			if sv1 == sv2 {
				equal = true
				break
			}
		}
		if !equal {
			return false
		}
	}

	for _, sv1 := range s2 {
		equal := false
		for _, sv2 := range s1 {
			if sv1 == sv2 {
				equal = true
				break
			}
		}
		if !equal {
			return false
		}
	}

	return true
}

func CmpStr_Strict(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, l := 0, len(s1); i < l; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
