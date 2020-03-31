package goodreads

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetOneBook(t *testing.T) {
	var ctx = context.TODO()

	t.Run("returns the book", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/get_one_book.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		book, err := client.GetOneBook(ctx, 1)

		assert.NoError(t, err)
		assert.Equal(t, Book{
			ID:                 862041,
			Title:              "Harry Potter Series Box Set (Harry Potter, #1-7)",
			Description:        "foo bar",
			ImageURL:           "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1534298934l/862041._SX98_.jpg",
			SmallImageURL:      "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1534298934l/862041._SX50_.jpg",
			NumPage:            4100,
			Format:             "",
			EditionInformation: "",
			Publisher:          "Arthur A. Levine Books",
			Work: Work{
				WorkID:        2962492,
				BookID:        0,
				BestBookID:    862041,
				Title:         "",
				ImageURL:      "",
				SmallImageURL: "",
				Author:        Author{},
				OriginalPublicationDate: OriginalPublicationDate{
					Year:  2007,
					Month: 10,
					Day:   1,
				},
			},
			Authors: []Author{
				{
					ID:            1077326,
					Name:          "J.K. Rowling",
					About:         "",
					ImageURL:      "https://images.gr-assets.com/authors/1510435123p5/1077326.jpg",
					SmallImageURL: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
					LargeImageURL: "",
					WorkCount:     0,
					Gender:        "",
					Hometown:      "",
					BornDate:      "",
					DiedAt:        "",
				},
			},
			PublicationDate: PublicationDate{
				Year:  2007,
				Month: 10,
				Day:   1,
			},
		}, book)
	})

	t.Run("returns an error if call failed", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "wtf", http.StatusInternalServerError)
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		book, err := client.GetOneBook(ctx, 1)

		assert.EqualError(t, err, "failed to get the book #1: request failed for '//book/show': 500 Internal Server Error")
		assert.Equal(t, Book{}, book)
	})
}
