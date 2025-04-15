package storage_test

import (
	"os"
	"testing"
	"fmt"
	"mongodb_clone/internal/storage"

)

func TestPager(t *testing.T) {
	path := "test.db"
	defer os.Remove(path)

	pager, err := storage.NewPager(path)
	if err != nil {
		t.Fatalf("failed to create pager: %v", err)
	}
	defer pager.Close()

	pageID := int64(1)
	page, err := pager.GetPage(pageID)
	if err != nil {
		t.Fatalf("failed to get page: %v", err)
	}

	copy(page.Data, []byte("Hello World"))
	page.Dirty = true

	if err := pager.FlushPage(page); err != nil {
		t.Fatalf("failed to flush page: %v", err)
	}

	pager2, _ := storage.NewPager(path)
	defer pager2.Close()

	readPage, err := pager2.GetPage(pageID)
	if err != nil {
		t.Fatalf("failed to get page after flush: %v", err)
	}

	if string(readPage.Data[:11]) != "Hello World" {
		t.Fatalf("expected 'Hello World' got '%s'", string(readPage.Data[:11]))
	}
}



func TestPagerOperations(t *testing.T) {
	// Initialize the Pager by opening/creating the database file "data.db"
	pager, err := storage.NewPager("data.db")
	if err != nil {
		t.Fatalf("failed to initialize pager: %v", err)
	}
	defer pager.Close()

	// Get the first page (page ID = 0)
	page, err := pager.GetPage(0)
	if err != nil {
		t.Fatalf("failed to get page: %v", err)
	}

	// Modify the page data (writing a string into the byte array)
	copy(page.Data, []byte("Hello Document Storage!"))
	page.Dirty = true 

	// Write the dirty page back to the file
	if err := pager.FlushPage(page); err != nil {
		t.Fatalf("failed to flush page: %v", err)
	}

	// Retrieve the page again to ensure it was saved
	readPage, err := pager.GetPage(0)
	if err != nil {
		t.Fatalf("failed to retrieve page after flush: %v", err)
	}

	// Verify that the content is as expected
	expected := "Hello Document Storage!"
	actual := string(readPage.Data[:len(expected)])
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	// Output the retrieved data for confirmation
	fmt.Println(actual) 
	// This will print: "Hello Document Storage!"

	// if err := os.Remove("data.db"); err != nil {
	// 	t.Fatalf("failed to delete file: %v", err)
	// }
}
