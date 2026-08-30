package services

import "terrariadle/internal/domain"

func toSearchableNpcs(n []domain.SearchNpcResult) []SearchNpcData {
	npcs := make([]SearchNpcData, len(n))
	for i, sn := range n {
		npcs[i] = SearchNpcData{
			NpcID: sn.NpcID,
			Name:  sn.Name,
			Path:  sn.Path,
		}
	}

	return npcs
}
