package google

import (
	"context"
	"encoding/json"
	"learning_context/userip"
	"net/http"
	"os"
)

type Result struct {
	Title, Link string
}

type Results []Result

// The Google Web Search API request includes the search query and the user IP as query parameters
func Search(ctx context.Context, query string) (Results, error) {
	// Prepare the Google Search API request.
	// Get creds here - https://developers.google.com/custom-search/v1/using_rest
	// https://www.googleapis.com/customsearch/v1?key=<ENV_KEY>&cx=<ENV_CX>&q=<query>
	req, err := http.NewRequest("GET", "https://www.googleapis.com/customsearch/v1", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)
	q.Set("key", os.Getenv("KEY"))
	q.Set("cx", os.Getenv("CX"))

	// If ctx is carrying the user IP address, forward it to the server.
	// Google APIs use the user IP to distinguish server-initiated requests
	// from end-user requests.
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()

	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Parse the JSON search result.
		// https://developers.google.com/custom-search/v1/reference/rest/v1/Search#result
		var data struct {
			Items []struct {
				Title string
				Link  string
			}
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		for _, res := range data.Items {
			results = append(results, Result{Title: res.Title, Link: res.Link})
		}
		return nil
	})
	// httpDo waits for the closure we provided to return, so it's safe to
	// read results here.
	return results, err
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() {
		c <- f(http.DefaultClient.Do(req))
	}()
	select {
	case <-ctx.Done():
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
