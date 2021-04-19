package timex

import "time"

const (
	nanoToMilliDivisor int64 = 1e6
	milliToSecDivisor  int64 = 1e3
)

var (
	now = time.Now
)

func CurrentUnixMillisecond() int64 {
	return ToUnixMillisecond(now())
}

func ToUnixMillisecond(t time.Time) int64 {
	return t.UnixNano() / nanoToMilliDivisor
}

func FromUnixMillisecond(milli int64) time.Time {
	sec := milli / milliToSecDivisor
	nano := (milli % milliToSecDivisor) * nanoToMilliDivisor
	return time.Unix(sec, nano)
}
