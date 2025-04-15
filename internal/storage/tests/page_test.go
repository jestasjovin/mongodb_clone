package storage_test

import (
	"testing"
	"mongodb_clone/internal/storage"

)

func TestNewPage(t *testing.T) {
	pageID := int64(42)
	page := storage.NewPage(pageID)

	if page.ID != pageID {
		t.Fatalf("expected ID %d, got %d", pageID, page.ID)
	}
	if len(page.Data) != storage.PageSize {
		t.Fatalf("expected page size %d, got %d", storage.PageSize, len(page.Data))
	}
	if page.Dirty {
		t.Fatalf("new page should not be dirty")
	}
}

func TestPageModification(t *testing.T) {
	page := storage.NewPage(1)
	testData := []byte("test content")

	copy(page.Data, testData)
	page.Dirty = true

	if !page.Dirty {
		t.Fatalf("expected Dirty to be true after modification")
	}
	if string(page.Data[:len(testData)]) != string(testData) {
		t.Fatalf("expected data '%s', got '%s'", string(testData), string(page.Data[:len(testData)]))
	}

	
}
