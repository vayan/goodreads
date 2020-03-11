package goodreads

// SeriesWithWorks include the series and its works
type SeriesWithWorks struct {
	Series
	Works []Work `xml:"series>series_works>series_work>work"`
}

// Series describe a series
type Series struct {
	ID               int    `xml:"series>id"`
	Title            string `xml:"series>title"`
	Description      string `xml:"series>description"`
	Note             string `xml:"series>note"`
	SeriesWorksCount int    `xml:"series>series_works_count"`
	PrimaryWorkCount int    `xml:"series>primary_work_count"`
	Numbered         bool   `xml:"series>numbered"`
}

// Work describes a work
// Optional fields:
// - SmallImageURL
// - PublicationDate (partial data sometimes)
// - Maybe some more, be careful :)
type Work struct {
	WorkID        int    `xml:"id"`
	BookID        int    `xml:"best_book>id"`
	Title         string `xml:"best_book>title"`
	ImageURL      string `xml:"best_book>image_url"`
	SmallImageURL string `xml:"best_book>small_image_url"`
	Author        Author `xml:"best_book>author"`
	PublicationDate
}

// PublicationDate holds the year / month / day
// 0 value means the API don't have data for the field
type PublicationDate struct {
	Year  int `xml:"original_publication_year"`
	Month int `xml:"original_publication_month"`
	Day   int `xml:"original_publication_day"`
}

// Author the guy who wrote the thing
type Author struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}
