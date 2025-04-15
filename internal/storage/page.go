package storage

// Size of page defaulting 4KB
const PageSize = 4096 


// Represents a single page in the database (holds data)
type Page struct {
	ID    int64  
	Data  []byte // Data stored in the page
	Dirty bool   // Flag to track if the page has been modified
}

// neww Page instance with a specified ID
// Initializes an empty byte slice to store the page data
func NewPage(id int64) *Page {
	return &Page{
		ID:    id,
		Data:  make([]byte, PageSize), 
		Dirty: false, 
	}
}
