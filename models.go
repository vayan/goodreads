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
// - OriginalPublicationDate (partial data sometimes)
// - Maybe some more, be careful :)
type Work struct {
	WorkID        int    `xml:"id"`
	BookID        int    `xml:"best_book>id"`
	BestBookID    int    `xml:"best_book_id"`
	OriginalTitle string `xml:"original_title"`
	Title         string `xml:"best_book>title"`
	ImageURL      string `xml:"best_book>image_url"`
	SmallImageURL string `xml:"best_book>small_image_url"`
	Author        Author `xml:"best_book>author"`
	OriginalPublicationDate
}

// OriginalPublicationDate holds the year / month / day
// 0 value means the API don't have data for the field
type OriginalPublicationDate struct {
	Year  int `xml:"original_publication_year"`
	Month int `xml:"original_publication_month"`
	Day   int `xml:"original_publication_day"`
}

// PublicationDate same as OriginalPublicationDate but different xml attr :(
type PublicationDate struct {
	Year  int `xml:"publication_year"`
	Month int `xml:"publication_month"`
	Day   int `xml:"publication_day"`
}

// AuthorWithBooks include a partial author and his books
type AuthorWithBooks struct {
	Author
	Books []Book `xml:"books>book"`
}

// Author the guy who wrote the thing
type Author struct {
	ID            int    `xml:"id"`
	Name          string `xml:"name"`
	About         string `xml:"about"`
	ImageURL      string `xml:"image_url"`
	SmallImageURL string `xml:"small_image_url"`
	LargeImageURL string `xml:"large_image_url"`
	WorkCount     int    `xml:"works_count"`
	Gender        string `xml:"gender"`
	Hometown      string `xml:"hometown"`
	BornDate      string `xml:"born_at"`
	DiedAt        string `xml:"died_at"`
}

// Book the paper thing u know
type Book struct {
	ID                 int      `xml:"id"`
	Title              string   `xml:"title"`
	Description        string   `xml:"description"`
	ImageURL           string   `xml:"image_url"`
	SmallImageURL      string   `xml:"small_image_url"`
	NumPage            int      `xml:"num_pages"`
	Format             string   `xml:"format"`
	EditionInformation string   `xml:"edition_information"`
	Publisher          string   `xml:"publisher"`
	Work               Work     `xml:"work"`
	Authors            []Author `xml:"authors>author"`
	PublicationDate
}
