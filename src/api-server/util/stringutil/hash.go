package stringutil

import (
    "crypto/md5"
    "fmt"
)

func MD5String(in string) string {
    data := []byte(in)
    return fmt.Sprintf("%x", md5.Sum(data))
}
