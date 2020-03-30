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

func TestClient_GetOneAuthor(t *testing.T) {
	var ctx = context.TODO()

	t.Run("returns the parsed author", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/get_one_author.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		author, err := client.GetOneAuthor(ctx, 15388346)

		assert.NoError(t, err)
		assert.Equal(t, Author{
			ID:            15388346,
			Name:          "Nicholas Eames",
			About:         "",
			ImageURL:      "https://images.gr-assets.com/authors/1468878466p5/15388346.jpg",
			SmallImageURL: "https://images.gr-assets.com/authors/1468878466p2/15388346.jpg",
			LargeImageURL: "https://images.gr-assets.com/authors/1468878466p7/15388346.jpg",
			WorkCount:     8,
			Gender:        "male",
			Hometown:      "",
			BornDate:      "",
			DiedAt:        "",
		}, author)
	})

	t.Run("returns error if the called failed", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "wtf", http.StatusInternalServerError)
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		author, err := client.GetOneAuthor(ctx, 15388346)

		assert.EqualError(t, err, "failed to get the author #15388346: request failed for '//author/show': 500 Internal Server Error")
		assert.Equal(t, Author{}, author)
	})
}

func TestClient_GetAuthorBooks(t *testing.T) {
	var ctx = context.TODO()

	t.Run("returns the author and his books where there's results", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/get_author_books_with_results.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		author, err := client.GetAuthorBooks(ctx, 15388346, 1)

		assert.NoError(t, err)
		assert.Equal(t, AuthorWithBooks{
			Author: Author{
				ID:            15388346,
				Name:          "Nicholas Eames",
				About:         "",
				ImageURL:      "",
				SmallImageURL: "",
				LargeImageURL: "",
				WorkCount:     0,
				Gender:        "",
				Hometown:      "",
				BornDate:      "",
				DiedAt:        "",
			},
			Books: []Book{
				{
					ID:                 30841984,
					Title:              "Kings of the Wyld (The Band, #1)",
					Description:        "good description",
					ImageURL:           "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1477027207l/30841984._SX98_.jpg",
					SmallImageURL:      "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1477027207l/30841984._SY75_.jpg",
					NumPage:            502,
					Format:             "Paperback",
					EditionInformation: "",
					Publisher:          "Orbit",
					Work: Work{
						WorkID: 51246585,
					},
					Authors: []Author{
						{
							ID:            15388346,
							Name:          "Nicholas Eames",
							About:         "",
							ImageURL:      "https://images.gr-assets.com/authors/1468878466p5/15388346.jpg",
							SmallImageURL: "https://images.gr-assets.com/authors/1468878466p2/15388346.jpg",
							LargeImageURL: "",
							WorkCount:     0,
							Gender:        "",
							Hometown:      "",
							BornDate:      "",
							DiedAt:        "",
						},
					},
					PublicationDate: PublicationDate{
						Year:  2017,
						Month: 2,
						Day:   21,
					},
				},
				{
					ID:                 35052265,
					Title:              "Bloody Rose (The Band, #2)",
					Description:        "good description 2",
					ImageURL:           "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1509483649l/35052265._SX98_.jpg",
					SmallImageURL:      "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1509483649l/35052265._SY75_.jpg",
					NumPage:            544,
					Format:             "Paperback",
					EditionInformation: "",
					Publisher:          "Orbit",
					Work: Work{
						WorkID: 56340013,
					},
					Authors: []Author{
						{
							ID:            15388346,
							Name:          "Nicholas Eames",
							About:         "",
							ImageURL:      "https://images.gr-assets.com/authors/1468878466p5/15388346.jpg",
							SmallImageURL: "https://images.gr-assets.com/authors/1468878466p2/15388346.jpg",
							LargeImageURL: "",
							WorkCount:     0,
							Gender:        "",
							Hometown:      "",
							BornDate:      "",
							DiedAt:        "",
						},
					},
					PublicationDate: PublicationDate{
						Year:  2018,
						Month: 8,
						Day:   30,
					},
				},
			},
		}, author)
	})

	t.Run("returns the author and his empty books if no results", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/get_author_books_with_no_results.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		author, err := client.GetAuthorBooks(ctx, 15388346, 1)

		assert.NoError(t, err)
		assert.Equal(t, AuthorWithBooks{
			Author: Author{
				ID:            15388346,
				Name:          "Nicholas Eames",
				About:         "",
				ImageURL:      "",
				SmallImageURL: "",
				LargeImageURL: "",
				WorkCount:     0,
				Gender:        "",
				Hometown:      "",
				BornDate:      "",
				DiedAt:        "",
			},
			Books: []Book{},
		}, author)
	})

	t.Run("returns an error if called failed", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "wtf", http.StatusInternalServerError)
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		author, err := client.GetAuthorBooks(ctx, 15388346, 1)

		assert.EqualError(t, err, "failed to get the books for the author #15388346 in page #1: request failed for '//author/list': 500 Internal Server Error")
		assert.Equal(t, AuthorWithBooks{}, author)
	})

}
