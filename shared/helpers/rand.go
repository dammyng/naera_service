package helpers

import (
	"math/rand"
	"time"

	"github.com/twinj/uuid"
)


const intset = "0123456789"

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charset2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

var seedRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandInt(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = intset[seedRand.Intn(len(intset))]
	}
	return string(b)
}

func RandUpperAlpha(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}
	return string(b)
}

func RandAlpha(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset2[seedRand.Intn(len(charset))]
	}
	return string(b)
}

func GuidId() string {
	return uuid.NewV4().String()
}


