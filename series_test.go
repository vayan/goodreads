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

func TestClient_GetAllSeriesForWork(t *testing.T) {
	var ctx = context.TODO()

	t.Run("it decodes the series response into the given struct", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/get_all_series_for_work_with_result.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		series, err := client.GetAllSeriesForWork(ctx, 111)

		assert.NoError(t, err)
		assert.Equal(t, []Series{{
			ID:               57883,
			Title:            "\n    Naruto\n",
			Description:      "\n    Naruto is a manga series about  Uzumaki Naruto, a young ninja.\n",
			Note:             "\n    Please only add the primary manga volumes to this list. Anything else, including omnibus editions and box-sets, WILL be removed.\n",
			SeriesWorksCount: 72,
			PrimaryWorkCount: 72,
			Numbered:         true,
		}}, series)
	})

	t.Run("it decodes the series response into an empty slice if no result", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/get_all_series_for_work_with_no_result.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		series, err := client.GetAllSeriesForWork(ctx, 111)

		assert.NoError(t, err)
		assert.Equal(t, []Series{}, series)
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

		series, err := client.GetAllSeriesForWork(ctx, 111)

		assert.EqualError(t, err, "failed to get the series for the work #111: request failed for '//series/work/111': 500 Internal Server Error")
		assert.Equal(t, []Series{}, series)
	})
}

func TestClient_GetOneSeries(t *testing.T) {
	var ctx = context.TODO()

	t.Run("it decodes the series response into the given struct", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/get_one_series_with_some_works.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		works, err := client.GetOneSeries(ctx, 111, 0)

		assert.NoError(t, err)
		assert.Equal(t, SeriesWithWorks{
			Series{
				ID:               193556,
				Title:            "\n    The Band\n",
				Description:      "\n        blabla\n        blabla\n",
				Note:             "\n",
				SeriesWorksCount: 3,
				PrimaryWorkCount: 3,
				Numbered:         true,
			},
			[]Work{
				{
					WorkID:        51246585,
					BookID:        30841984,
					Title:         "Kings of the Wyld (The Band, #1)",
					OriginalTitle: "Kings of the Wyld",
					ImageURL:      "https://image.jpg",
					SmallImageURL: "",
					Author: Author{
						ID:   15388346,
						Name: "Nicholas Eames",
					},
					OriginalPublicationDate: OriginalPublicationDate{
						Year:  2017,
						Month: 2,
						Day:   21,
					},
				},
				{
					WorkID:        56340013,
					BookID:        35052265,
					Title:         "Bloody Rose (The Band, #2)",
					OriginalTitle: "Bloody Rose",
					ImageURL:      "https://image2.jpg",
					SmallImageURL: "",
					Author: Author{
						ID:   15388346,
						Name: "Nicholas Eames",
					},
					OriginalPublicationDate: OriginalPublicationDate{
						Year:  2018,
						Month: 8,
						Day:   28,
					},
				},
				{
					WorkID:        52587935,
					BookID:        31932963,
					Title:         "Outlaw Empire (The Band, #3)",
					OriginalTitle: "Outlaw Empire",
					ImageURL:      "https://image3.png",
					SmallImageURL: "",
					Author: Author{
						ID:   15388346,
						Name: "Nicholas Eames",
					},
					OriginalPublicationDate: OriginalPublicationDate{
						Year:  0,
						Month: 0,
						Day:   0,
					},
				},
			},
		}, works)
	})

	t.Run("it decodes the series response into the given struct with empty work array if no works", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadFile("fixtures/get_one_series_with_no_works.xml")
			_, _ = fmt.Fprintln(w, string(content))
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}

		works, err := client.GetOneSeries(ctx, 111, 0)

		assert.NoError(t, err)
		assert.Equal(t, SeriesWithWorks{
			Series{
				ID:               193556,
				Title:            "\n    The Band\n",
				Description:      "\n    blabla\n    blabla\n",
				Note:             "\n",
				SeriesWorksCount: 3,
				PrimaryWorkCount: 3,
				Numbered:         true,
			},
			[]Work{},
		}, works)
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

		works, err := client.GetOneSeries(ctx, 111, 0)

		assert.EqualError(t, err, "failed to get the work for the series #111 in page #0: request failed for '//series/show/111': 500 Internal Server Error")
		assert.Equal(t, SeriesWithWorks{}, works)
	})
}
