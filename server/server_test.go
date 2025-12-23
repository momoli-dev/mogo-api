package server_test

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/momoli-dev/mogo/server"
)

func TestNewServer_OK(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/test", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	params := &server.Params{
		Addr:    ":8081",
		Handler: handler,
	}

	srv := server.NewServer(params)

	go func() {
		err := srv.Start()
		assert.NoError(t, err)
	}()
	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:8081/test")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	expectedBody := "Hello, World!"
	require.Equal(t, expectedBody, string(body))

	err = srv.Shutdown()
	require.NoError(t, err)

	_, err = http.Get("http://localhost:8081/test")
	require.Error(t, err)
}

func TestNewServer_InvalidAddr(t *testing.T) {
	handler := http.NewServeMux()
	params := &server.Params{
		Addr:    "invalid-addr",
		Handler: handler,
	}

	srv := server.NewServer(params)

	err := srv.Start()
	require.Error(t, err)
}

func TestNewServer_NilHandler(t *testing.T) {
	params := &server.Params{
		Addr:    ":8082",
		Handler: nil,
	}

	require.Panics(t, func() {
		server.NewServer(params)
	})
}

func TestServer_MustStart_Panic(t *testing.T) {
	handler := http.NewServeMux()
	params := &server.Params{
		Addr:    "invalid-addr",
		Handler: handler,
	}

	srv := server.NewServer(params)

	require.Panics(t, func() {
		srv.MustStart()
	})
}

func TestServer_Graceful(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/test", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Graceful!"))
	})

	params := &server.Params{
		Addr:    ":8083",
		Handler: handler,
	}

	srv := server.NewServer(params)
	shutdownChan := srv.StartGracefully()

	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:8083/test")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	shutdownChan <- os.Interrupt
	time.Sleep(100 * time.Millisecond)

	_, err = http.Get("http://localhost:8083/test")
	require.Error(t, err)
}
