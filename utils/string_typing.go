package utils

type sortableStringSlice []string

func (t sortableStringSlice) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t sortableStringSlice) Len() int {
	return len(t)
}

func (t sortableStringSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
