package stringutil

import "strconv"

func ValueOfFloat64(value float64) string {
    return strconv.FormatFloat(float64(value), 'f', 6, 64)
}

func ValueOfInt(value int64) string {
    return strconv.FormatInt(value, 10)
}