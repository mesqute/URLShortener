package utilites

import "math/rand"

// GenerateToken создает токен заданной длины
func GenerateToken(length int) string {
	const dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	var str []byte

	for i := 0; i < length; i++ {
		rnd := rand.Intn(63)
		str = append(str, dictionary[rnd])
	}
	return string(str)
}
