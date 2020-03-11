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

func TestClient_Search(t *testing.T) {
	var ctx = context.TODO()

	t.Run("it decodes the search response into the given struct", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/search_with_result.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		books, err := client.Search(ctx, "hairy pooter", 0)

		assert.NoError(t, err)
		assert.Equal(t, []Work{{
			WorkID:        1111,
			BookID:        35052265,
			Title:         "Harry Pooter and The Funny Guy",
			ImageURL:      "https://big.jpg",
			SmallImageURL: "https://small.jpg",
			Author: Author{
				ID:   15388346,
				Name: "John",
			},
			PublicationDate: PublicationDate{
				Year:  2018,
				Month: 8,
				Day:   28,
			},
		}}, books)
	})

	t.Run("it decodes the search response into an empty slice if no result", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/search_with_no_result.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		books, err := client.Search(ctx, "hairy pooter", 0)

		assert.NoError(t, err)
		assert.Equal(t, []Work{}, books)
	})

	t.Run("it returns an error if something went wrong", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "wtf", http.StatusInternalServerError)
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		books, err := client.Search(ctx, "hairy pooter", 0)

		assert.EqualError(t, err, "'hairy pooter' search at page 0 failed: request failed for '//search/index': 500 Internal Server Error")
		assert.Equal(t, []Work{}, books)
	})
}
