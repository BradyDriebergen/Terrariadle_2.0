package services

import "math/rand/v2"

func shuffleChunks(chunks []string) {
	rand.Shuffle(len(chunks), func(i, j int) {
		chunks[i], chunks[j] = chunks[j], chunks[i]
	})
}
