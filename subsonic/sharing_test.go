package subsonic

import (
	"strconv"
	"testing"
	"time"
)

func runSharingTests(client Client, t *testing.T) {
	sampleAlbum := getSampleAlbum(client)

	t.Run("GetShares", func(t *testing.T) {
		_, err := client.GetShares()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("CreateShare", func(t *testing.T) {
		description := "Test share"
		expires := time.Now().UTC().Round(time.Millisecond).Add(time.Hour)
		expiresStr := strconv.FormatInt(expires.UnixMilli(), 10)
		share, err := client.CreateShare(sampleAlbum.ID, map[string]string{"description": description, "expires": expiresStr})
		if err != nil {
			t.Fatal(err)
		}
		if share == nil {
			t.Fatal("No share returned")
		}
		if len(share.Entry) != 1 {
			t.Fatalf("Expected 1 entry, got %d", len(share.Entry))
		}
		if share.Entry[0].ID != sampleAlbum.ID {
			t.Errorf("Expected ID %q, got %q", sampleAlbum.ID, share.Entry[0].ID)
		}
		if share.Description != description {
			t.Errorf("Expected description %q, got %q", description, share.Description)
		}
		if share.Expires != expires {
			t.Errorf("Expected expires %q, got %q", expires, share.Expires)
		}
	})

	t.Run("UpdateShare", func(t *testing.T) {
		description := "Test share"
		expires := time.Now().UTC().Round(time.Millisecond).Add(time.Hour)
		expiresStr := strconv.FormatInt(expires.UnixMilli(), 10)
		share, err := client.CreateShare(sampleAlbum.ID, map[string]string{"description": description, "expires": expiresStr})
		if err != nil {
			t.Fatal(err)
		}

		newDescription := "Updated share"
		newExpires := time.Now().UTC().Round(time.Millisecond).Add(3 * time.Hour)
		newExpiresStr := strconv.FormatInt(newExpires.UnixMilli(), 10)
		err = client.UpdateShare(share.ID, map[string]string{"description": newDescription, "expires": newExpiresStr})
		if err != nil {
			t.Fatal(err)
		}

		shares, err := client.GetShares()
		if err != nil {
			t.Fatal(err)
		}
		for _, s := range shares {
			if s.ID == share.ID {
				if s.Description != newDescription {
					t.Errorf("Expected description %q, got %q", newDescription, s.Description)
				}
				if s.Expires != newExpires {
					t.Errorf("Expected expires %q, got %q", newExpires, s.Expires)
				}
				return
			}
		}
		t.Fatalf("Share %q not found", share.ID)
	})

	t.Run("DeleteShare", func(t *testing.T) {
		share, err := client.CreateShare(sampleAlbum.ID, nil)
		if err != nil {
			t.Fatal(err)
		}

		err = client.DeleteShare(share.ID)
		if err != nil {
			t.Fatal(err)
		}

		shares, err := client.GetShares()
		if err != nil {
			t.Fatal(err)
		}
		for _, s := range shares {
			if s.ID == share.ID {
				t.Fatalf("Share %q not deleted. Share: %+v", share.ID, s)
			}
		}
	})
}
