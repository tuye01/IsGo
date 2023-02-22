package utils

import (
	"math/rand"
	"time"
)

// 随机生成
func RandomeString(n int) string {
	var letters = []byte("asdfghjklqfjkhyuugilwertyuiop")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
