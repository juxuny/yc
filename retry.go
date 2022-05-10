package yc

func Retry(f func() (isEnd bool, err error), times ...int) error {
	count := 3
	if len(times) > 0 {
		count = times[0]
	}
	var lastError error
	var isEnd bool
	for i := 0; i < count; i++ {
		isEnd, lastError = f()
		if isEnd || lastError == nil {
			return lastError
		}
	}
	return lastError
}
