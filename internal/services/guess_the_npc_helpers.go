package services

import "terrariadle-backend/internal/domain"

func toSearchableNpcs(n []domain.SearchNpcResult) []SearchNpcData {
	npcs := make([]SearchNpcData, len(n))
	for _, sn := range n {
		npcs = append(npcs, SearchNpcData{
			NpcID: sn.NpcID,
			Name:  sn.Name,
			Path:  sn.Path,
		})
	}

	return npcs
}
