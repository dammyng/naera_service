package helpers

import (
	"math/rand"
	"net/http"
	"time"
)

func SetupCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method == "OPTIONS" {
		(*w).Header().Set("Access-Control-Max-Age", "1728000")
		(*w).Header().Set("Response-Code", "204")
	}

	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}



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