package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return randomString(6)
}

func RandomBalance() int64 {
	return randomInt(0, 10000)
}

func RandomAccountID() int64 {
	return randomInt(1, 100000)
}

func RandomCurrency() string {
	currencies := []string{"INR", "USD", "CAD", "EUR"}
	return currencies[rand.Intn(len(currencies))]
}
