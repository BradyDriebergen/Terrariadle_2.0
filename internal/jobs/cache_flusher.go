package jobs

import (
	"context"
	"log"
	"terrariadle-backend/internal/store"
	"time"
)

func StartFlushJob(ctx context.Context, s *store.CachedUserStore) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := s.FlushDirty(ctx); err != nil {
				log.Printf("flush error: %v", err)
			}
			s.EvictStale()
		case <-ctx.Done():
			if err := s.FlushDirty(context.Background()); err != nil {
				log.Printf("final flush error: %v", err)
			}
			return
		}
	}
}
