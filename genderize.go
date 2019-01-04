package genderize

import (
	"fmt"
	"net/url"

	"github.com/savaki/httpctx"
	"golang.org/x/net/context"
)

const (
	endpoint = "https://api.genderize.io"
)

type Results []Result

type Result struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
	Count       int     `json:"count"`
}

type Client struct {
	Client httpctx.HttpClient
}

func New() *Client {
	return &Client{
		Client: httpctx.NewClient(),
	}
}

func (c *Client) Query(names ...string) (Results, error) {
	ctx := context.Background()
	return c.QueryWithContext(ctx, names...)
}

func (c *Client) QueryWithContext(ctx context.Context, names ...string) (Results, error) {
	params := url.Values{}
	for index, name := range names {
		key := fmt.Sprintf("name[%d]", index)
		params.Add(key, name)
	}

	results := Results{}
	err := c.Client.Get(ctx, endpoint, &params, &results)
	return results, err
}
