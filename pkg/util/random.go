package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomInt32(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomProductTypeID() int64 {
	productTypes := []int64{PRODUCT, SERVICE}
	n := len(productTypes)
	return productTypes[rand.Intn(n)]
}
