package feeder

type Item struct {
	// RSS and Shared fields
	Title       string
	Links       []*Link
	Description string
	Author      Author
	Categories  []*Category
	Comments    string
	Enclosures  []*Enclosure
	Guid        *string
	PubDate     string
	Source      *Source

	// Atom specific fields
	Id           string
	Generator    *Generator
	Contributors []string
	Content      *Content
}
