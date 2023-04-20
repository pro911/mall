package tools

import (
	"math/rand"
	"time"
)

// RandomBool 随机返回一个bool值 可以传入概率:例如 RandomBool(0.8)
func RandomBool(p float64) bool {
	rand.NewSource(time.Now().UnixNano())
	return rand.Float64() < p
}
