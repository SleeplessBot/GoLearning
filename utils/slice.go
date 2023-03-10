package utils

func FindInt64(i64s []int64, v int64) int {
	for index, value := range i64s {
		if value == v {
			return index
		}
	}
	return -1
}

func FindStr(i64s []string, v string) int {
	for index, value := range i64s {
		if value == v {
			return index
		}
	}
	return -1
}

func CompareStrSlice(s1, s2 []string) (s1Only []string, s1AndS2 []string, s2Only []string) {
	s1Set := make(map[string]bool)
	for _, s := range s1 {
		s1Set[s] = true
	}
	s2Set := make(map[string]bool)
	for _, s := range s2 {
		s2Set[s] = true
		if s1Set[s] {
			s1AndS2 = append(s1AndS2, s)
		} else {
			s2Only = append(s2Only, s)
		}
	}

	for _, s := range s1 {
		if !s2Set[s] {
			s1Only = append(s1Only, s)
		}
	}
	return
}

func Strs2Map(ss []string) map[string]bool {
	m := make(map[string]bool)
	for _, s := range ss {
		m[s] = true
	}
	return m
}

func Map2Strs(m map[string]bool) []string {
	ss := make([]string, 0, len(m))
	for k := range m {
		ss = append(ss, k)
	}
	return ss
}

func AppendStrIfNoEmpty(ss []string, s string) []string {
	if s != "" {
		ss = append(ss, s)
	}
	return ss
}

func DeduplicateUint(uis []uint) (res []uint) {
	m := make(map[uint]bool)
	for _, ui := range uis {
		m[ui] = true
	}
	for k := range m {
		res = append(res, k)
	}
	return
}

func IntersectUint(uis1, uis2 []uint) (res []uint) {
	m1 := make(map[uint]bool)
	for _, ui := range uis1 {
		m1[ui] = true
	}

	for _, ui := range uis2 {
		if m1[ui] {
			res = append(res, ui)
		}
	}

	return res
}
