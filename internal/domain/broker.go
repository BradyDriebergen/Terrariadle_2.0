package domain

import "sync"

type GameMode string

const (
	GameModeDailySlash  GameMode = "daily_slash"
	GameModeConnections GameMode = "connections"
	GameModeGuessTheNpc GameMode = "guess_the_npc"
	GameModeHangman     GameMode = "hangman"
	GameModeTerraTrivia GameMode = "terratrivia"
)

type GuessCountBroker interface {
	Subscribe() chan GuessCountEvent
	Unsubscribe(ch chan GuessCountEvent)
	Publish(event GuessCountEvent)
}

type GuessCountEvent struct {
	GameMode GameMode `json:"game_mode"`
	Count    int      `json:"count"`
}

type Broker struct {
	subscribers map[chan GuessCountEvent]struct{}
	mu          sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[chan GuessCountEvent]struct{}),
	}
}

func (b *Broker) Subscribe() chan GuessCountEvent {
	ch := make(chan GuessCountEvent, 8) // buffered so slow clients don't block
	b.mu.Lock()
	b.subscribers[ch] = struct{}{}
	b.mu.Unlock()
	return ch
}

func (b *Broker) Unsubscribe(ch chan GuessCountEvent) {
	b.mu.Lock()
	delete(b.subscribers, ch)
	close(ch)
	b.mu.Unlock()
}

func (b *Broker) Publish(event GuessCountEvent) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for ch := range b.subscribers {
		select {
		case ch <- event:
		default:
			// subscriber is too slow, drop the event rather than blocking
		}
	}
}
