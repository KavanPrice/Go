package utils

import (
	"math/rand"
	"time"
)

func RandomiseSlice(slice []interface{}) (s []interface{}) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range slice {
		newPosition := r.Intn(len(slice) - 1)
		slice[i], slice[newPosition] = slice[newPosition], slice[i]
	}
	s = slice
	return
}
