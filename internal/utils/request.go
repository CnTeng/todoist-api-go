package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type request[T any] struct {
	endpoint string
	token    string
	params   url.Values
	body     any

	client *http.Client
}

func NewRequest[T any](client *http.Client, endpoint, token string) *request[T] {
	return &request[T]{
		client:   client,
		endpoint: endpoint,
		token:    token,
		params:   url.Values{},
	}
}

func (r *request[T]) WithParameters(p url.Values) *request[T] {
	r.params = p
	return r
}

func (r *request[T]) WithBody(body any) *request[T] {
	r.body = body
	return r
}

func (r *request[T]) Get(ctx context.Context) (*T, error) {
	return r.do(ctx, http.MethodGet)
}

func (r *request[T]) Post(ctx context.Context) (*T, error) {
	return r.do(ctx, http.MethodPost)
}

func (r *request[T]) Delete(ctx context.Context) (*T, error) {
	return r.do(ctx, http.MethodDelete)
}

func (r *request[T]) do(ctx context.Context, method string) (*T, error) {
	u, err := url.Parse(r.endpoint)
	if err != nil {
		return nil, err
	}
	u.RawQuery = r.params.Encode()

	var body io.Reader
	if r.body != nil {
		b, err := json.Marshal(r.body)
		if err != nil {
			return nil, err
		}
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
