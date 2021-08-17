package server

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	s := Create(":8080")
	go Start(s)
	defer Stop(s)

	response, err := http.Get("http://localhost:8080/ping")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := readString(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, "pong\n", body)
}

func TestPong(t *testing.T) {
	s := Create(":8080")
	go Start(s)
	defer Stop(s)

	response, err := http.Get("http://localhost:8080/pong")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := readString(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, "ping\n", body)
}

func readString(r io.Reader) (string, error) {
	buf := new(bytes.Buffer)

	_, err := buf.ReadFrom(r)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
