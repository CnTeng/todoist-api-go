package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type request[T any] struct {
	endpoint string
	token    string
	params   any
	body     any

	client *http.Client
}

func newRequest[T any](client *http.Client, endpoint, token string) *request[T] {
	return &request[T]{
		client:   client,
		endpoint: endpoint,
		token:    token,
	}
}

func (r *request[T]) withParams(params any) *request[T] {
	r.params = params
	return r
}

func (r *request[T]) withBody(body any) *request[T] {
	r.body = body
	return r
}

func (r *request[T]) do(ctx context.Context, method string) (*T, error) {
	u, err := url.Parse(r.endpoint)
	if err != nil {
		return nil, err
	}

	if r.params != nil {
		params, err := query.Values(r.params)
		if err != nil {
			return nil, err
		}
		u.RawQuery = params.Encode()
	}

	var body io.Reader
	if r.body != nil {
		b, err := json.Marshal(r.body)
		if err != nil {
			return nil, err
		}
		fmt.Println("Request URL:", string(b))
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.token)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
