package services

import "math/rand/v2"

func shuffleOptions(options []string) {
	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})
}
