package main

import "math/rand"

func RandomString(n int) string {
	var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
	b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}