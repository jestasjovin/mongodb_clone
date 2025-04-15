package storage

import (
	"os" // For file handling
)

type Pager struct {
	file  *os.File         
	cache map[int64]*Page   
}

// Opens (or creates) a file where the pages will be stored and prepares a cache
func NewPager(path string) (*Pager, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err // Return error if file cannot be opened
	}
	return &Pager{
		file:  f,                           
		cache: make(map[int64]*Page),        // Initialize empty cache for pages
	}, nil
}

// GetPage retrieves a page by its ID. 
// If the page is in cache, it returns from cache
// Otherwise, it loads the page from disk and stores it in the cache
func (p *Pager) GetPage(pageID int64) (*Page, error) {
	if page, ok := p.cache[pageID]; ok {
		return page, nil
	}

	page := NewPage(pageID)
	offset := pageID * PageSize

	_, err := p.file.Seek(offset, 0)
	if err != nil {
		return nil, err
	}

	n, err := p.file.Read(page.Data)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}
	if n == 0 { // fresh page, untouched
		p.cache[pageID] = page
		return page, nil
	}

	p.cache[pageID] = page
	return page, nil
}

// Writes the contents of a dirty page (modified page) back to the disk
func (p *Pager) FlushPage(page *Page) error {
	// If page is not dirty, skip the flush
	if !page.Dirty {
		return nil
	}

	offset := page.ID * PageSize 

	// Seek to the position in the file to write the page
	_, err := p.file.Seek(offset, 0)
	if err != nil {
		return err 
	}

	// Write the page data back to the file
	_, err = p.file.Write(page.Data)
	if err != nil {
		return err 
	}

	// Mark the page as not dirty after successful write
	page.Dirty = false
	return nil
}

// Close flushes any remaining dirty pages and closes the file
func (p *Pager) Close() error {
	// Loop through the cache and flush any dirty pages
	for _, page := range p.cache {
		if page.Dirty {
			if err := p.FlushPage(page); err != nil {
				return err 
				// Return error if flushing a dirty page fails
			}
		}
	}

	// Close the file after flushing all pages
	return p.file.Close()
}
