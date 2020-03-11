package goodreads

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type getAllSeriesForWorkResponse struct {
	Results []Series `xml:"series_works>series_work"`
}

type getOneSeriesResponse struct {
	SeriesWithWorks
}

// GetAllSeriesForWork See all series a work is in
func (c client) GetAllSeriesForWork(ctx context.Context, workID int) ([]Series, error) {
	var response = getAllSeriesForWorkResponse{
		Results: []Series{},
	}

	err := c.Get(ctx, fmt.Sprintf("/series/work/%d", workID), url.Values{}, &response)

	if err != nil {
		return []Series{}, fmt.Errorf("failed to get the series for the work #%d: %w", workID, err)
	}

	return response.Results, nil
}

// GetOneSeries Get info on a given series, includes the works in the series
// For pagination 0 or 1 seems to be the same thing.
// The pagination will paginate the works if there's more than 100~
func (c client) GetOneSeries(ctx context.Context, serieID int, page int) (SeriesWithWorks, error) {
	var response = getOneSeriesResponse{
		SeriesWithWorks{
			Works: []Work{},
		},
	}

	q := url.Values{}
	q.Set("page", strconv.Itoa(page))

	err := c.Get(ctx, fmt.Sprintf("/series/show/%d", serieID), q, &response)

	if err != nil {
		return SeriesWithWorks{}, fmt.Errorf("failed to get the work for the series #%d in page #%d: %w", serieID, page, err)
	}

	return response.SeriesWithWorks, nil
}
