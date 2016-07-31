package timeutil

import "time"

func CurrentMillis() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetMillis(t time.Time) int64 {
    return t.UnixNano() / int64(time.Millisecond)
}