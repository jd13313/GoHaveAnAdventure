package utility

import "math/rand"

func RandomResponse(responses []string) string {
	return responses[rand.Intn(len(responses))]
}
