package util

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// Get makes HTTP GET request and returns status code and response body in bytes.
func Get(ctx context.Context, t *testing.T, client *http.Client, url string) (int, []byte) {
	t.Logf("GET %s", url)

	response, err := client.Get(url)
	require.NoError(t, err)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	return response.StatusCode, body
}

// Post makes HTTP POST request with json request body and returns status code and response body in bytes.
func Post(ctx context.Context, t *testing.T, client *http.Client, url string, requestBody any) (int, []byte) {
	t.Logf("POST %s", url)

	requestBodyBytes, err := json.Marshal(requestBody)
	require.NoError(t, err)

	response, err := client.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
	require.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	return response.StatusCode, body
}

// Delete makes HTTP GET request and returns status code and response body in bytes
func Delete(ctx context.Context, t *testing.T, client *http.Client, url string) (int, []byte) {
	t.Logf("DELETE %s", url)

	req, err := http.NewRequest("DELETE", url, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp.StatusCode, body
}
