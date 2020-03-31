package goodreads

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type getOneBook struct {
	Book Book `xml:"book"`
}

// GetOneBook retrieve a specific book
func (c client) GetOneBook(ctx context.Context, bookID int) (Book, error) {
	var response = getOneBook{}

	q := url.Values{}
	q.Set("id", strconv.Itoa(bookID))

	err := c.Get(ctx, fmt.Sprintf("/book/show"), q, &response)

	if err != nil {
		return Book{}, fmt.Errorf("failed to get the book #%d: %w", bookID, err)
	}

	return response.Book, nil
}
