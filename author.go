package goodreads

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type getOneAuthorResponse struct {
	Author Author `xml:"author"`
}

type getAuthorBooks struct {
	AuthorWithBooks AuthorWithBooks `xml:"author"`
}

// GetOneAuthor will returns the details of the given author ID
func (c client) GetOneAuthor(ctx context.Context, authorID int) (Author, error) {
	var response = getOneAuthorResponse{}

	q := url.Values{}
	q.Set("id", strconv.Itoa(authorID))

	err := c.Get(ctx, fmt.Sprintf("/author/show"), q, &response)

	if err != nil {
		return Author{}, fmt.Errorf("failed to get the author #%d: %w", authorID, err)
	}

	return response.Author, nil
}

// GetAuthorBooks returns some details of the author and a paginated list of his books
// For pagination 0 or 1 seems to be the same thing.
// The pagination will paginate the works if there's more than 100~
func (c client) GetAuthorBooks(ctx context.Context, authorID int, page int) (AuthorWithBooks, error) {
	var response = getAuthorBooks{
		AuthorWithBooks{
			Author: Author{},
			Books:  []Book{},
		},
	}

	q := url.Values{}
	q.Set("page", strconv.Itoa(page))
	q.Set("id", strconv.Itoa(authorID))

	err := c.Get(ctx, fmt.Sprintf("/author/list"), q, &response)

	if err != nil {
		return AuthorWithBooks{}, fmt.Errorf("failed to get the books for the author #%d in page #%d: %w", authorID, page, err)
	}

	return response.AuthorWithBooks, nil
}
