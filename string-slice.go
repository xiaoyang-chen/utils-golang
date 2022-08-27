package utils

func RmDuplicate(ss []string) (ssRet []string) {

	var lenSS = len(ss)
	ssRet = make([]string, lenSS)
	var ssRetIdx int
	var tmp = make(map[string]struct{}, lenSS)
	var exist bool
	for _, s := range ss {
		if _, exist = tmp[s]; !exist {
			ssRet[ssRetIdx] = s
			ssRetIdx++
			tmp[s] = struct{}{}
		}
	}
	ssRet = ssRet[:ssRetIdx]
	return
}
