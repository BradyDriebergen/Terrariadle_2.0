package repo

import (
	"context"
	"terrariadle/internal/db"
	"testing"
)

// func TestFindOne(t *testing.T) {
// 	ctx := context.Background()
// 	userRepo := mockUserRepo(t)

// 	user := generateUser()
// 	want :=

// 	tests := []struct {
// 		name    string
// 		seed    *userData
// 		wantErr error
// 	}{
// 		{name: "ExistingUser", seed: &user},
// 		{name: "NonExistingUser", seed: nil, wantErr: ErrNotFound},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.seed != nil {
// 				err := db.Upsert(ctx, testMongo, userRepo.userCollection, db.Filter{"_id": tt.seed.ID}, *tt.seed)
// 				if err != nil {
// 					t.Fatalf("upsert failed: %v", err)
// 				}
// 			}

// 			got, err := userRepo.GetUser(ctx, tt.seed.UserID)
// 			if tt.wantErr != nil {
// 				if err != tt.wantErr {
// 					t.Fatalf("got err %v, want %v", err, tt.wantErr)
// 				}
// 				return
// 			}
// 			if err != nil {
// 				t.Fatalf("findone failed: %v", err)
// 			}

// 			if diff := cmp.Diff(want, got); diff != "" {
// 				t.Errorf("answer mismatch (-want +got):\n%s", diff)
// 			}

// 			if got.Mode != tt.seed.Mode {
// 				t.Errorf("got mode %q, want %q", got.Mode, tt.seed.Mode)
// 			}
// 		})
// 	}
// }

func mockUserRepo(t *testing.T) *MongoUserRepo {
	t.Helper()

	name := "repo_" + t.Name()

	t.Cleanup(func() {
		_ = db.DeleteAll(context.Background(), testMongo, name)
	})

	return NewUserRepo(testMongo, name)
}

func generateUser() userData {
	return userData{}
}
