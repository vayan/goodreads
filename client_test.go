package goodreads

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type fakeResponse struct {
	Bar string `xml:"bar"`
}

func TestNewClient(t *testing.T) {
	t.Run("it returns a new client with the given api key", func(t *testing.T) {
		gr := NewClient("awesomesuperapikey11")

		assert.Equal(t, client{
			APIKey: "awesomesuperapikey11",
			format: "xml",
			domain: "https://www.goodreads.com",
			http: &http.Client{
				Timeout: 10 * time.Second,
			},
		}, gr)
	})
}

func TestClient_Get(t *testing.T) {
	var ctx = context.TODO()

	t.Run("it decodes the response into the given struct", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintln(w, "<foo><bar>hello</bar></foo>")
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}
		resp := fakeResponse{}

		err := client.Get(ctx, "bar", url.Values{}, &resp)

		assert.NoError(t, err)
		assert.Equal(t, fakeResponse{Bar: "hello"}, resp)
	})

	t.Run("it returns error when we cannot build the url", func(t *testing.T) {
		client := client{
			APIKey: "123",
			domain: "://🥺",
			http:   nil,
		}
		err := client.Get(ctx, "bar", url.Values{}, nil)

		assert.EqualError(t, err, "could not build search url: parse \"://🥺/bar\": missing protocol scheme")
	})

	t.Run("it returns an error when we cannot build the request", func(t *testing.T) {
		client := client{
			APIKey: "123",
			domain: "http://foo",
			http:   nil,
		}
		err := client.Get(nil, "bar", url.Values{}, nil)

		assert.EqualError(t, err, "failed to build request for '/bar': net/http: nil Context")
	})

	t.Run("it returns an error when the server returns an HTTP error (>= 400)", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
		}
		resp := fakeResponse{}

		err := client.Get(ctx, "bar", url.Values{}, &resp)

		assert.EqualError(t, err, "request failed for '/bar': 404 Not Found")
	})

	t.Run("it returns an error when the requests fails on our side", func(t *testing.T) {
		var expiredCtx, _ = context.WithTimeout(ctx, 0*time.Second)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
			format: "xml",
		}
		resp := fakeResponse{}

		err := client.Get(expiredCtx, "bar", url.Values{}, &resp)

		assert.EqualError(t, err, fmt.Sprintf("request failed for '/bar': Get \"%s/bar?format=xml&key=123\": context deadline exceeded", ts.URL))
	})

	t.Run("it returns an error when we fail to decode the response from the server", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintln(w, "<foo><bar>hello</d></v>")
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
			format: "xml",
		}
		resp := fakeResponse{}

		err := client.Get(ctx, "bar", url.Values{}, &resp)

		assert.EqualError(t, err, "failed to decode response for '/bar': XML syntax error on line 1: element <bar> closed by </d>")
	})

	t.Run("it returns an error when server does not return xml at all", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintln(w, "")
		}))
		defer ts.Close()

		client := client{
			APIKey: "123",
			domain: ts.URL,
			http:   ts.Client(),
			format: "xml",
		}
		resp := fakeResponse{}

		err := client.Get(ctx, "bar", url.Values{}, &resp)

		assert.EqualError(t, err, "failed to decode response for '/bar': EOF")
	})
}
