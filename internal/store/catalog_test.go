package store

import (
	"context"
	"terrariadle/internal/domain"
	"terrariadle/internal/testutils"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Tests that GetSearchableWeapons maps weapon data into the search result shape
func TestGetSearchableWeapons(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeCatalogRepo()
	weapons := fakeRepo.Weapons

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	want := make([]domain.SearchWeaponResult, len(weapons))
	for i, w := range weapons {
		want[i] = domain.SearchWeaponResult{
			WeaponID: w.ID,
			Name:     w.Name,
			Path:     w.Info.ImagePath,
		}
	}

	got := store.GetSearchableWeapons()

	// This helper sort function sorts the slice, used in cmpopts.SortSlice
	less := func(a, b domain.SearchWeaponResult) bool { return a.WeaponID < b.WeaponID }

	if diff := cmp.Diff(want, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests that GetSearchableNpcs maps npc data into the search result shape
func TestGetSearchableNpcs(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeCatalogRepo()
	npcs := fakeRepo.Npcs

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	want := make([]domain.SearchNpcResult, len(npcs))
	for i, n := range npcs {
		want[i] = domain.SearchNpcResult{
			NpcID: n.ID,
			Name:  n.NPC,
			Path:  n.NpcPath,
		}
	}

	got := store.GetSearchableNpcs()

	// This helper sort function sorts the slice, used in cmpopts.SortSlice
	less := func(a, b domain.SearchNpcResult) bool { return a.NpcID < b.NpcID }

	if diff := cmp.Diff(want, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}
