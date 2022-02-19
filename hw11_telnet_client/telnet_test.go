package main

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTelnetClient_basic(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:")
	require.NoError(t, err)
	defer func() { require.NoError(t, l.Close()) }()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		in := &bytes.Buffer{}
		out := &bytes.Buffer{}

		timeout, err := time.ParseDuration("10s")
		require.NoError(t, err)

		client := NewTelnetClient(l.Addr().String(), timeout, ioutil.NopCloser(in), out)
		require.NoError(t, client.Connect())
		defer func() { require.NoError(t, client.Close()) }()

		in.WriteString("hello\n")
		err = client.Send()
		require.NoError(t, err)

		err = client.Receive()
		require.NoError(t, err)
		require.Equal(t, "world\n", out.String())
	}()

	go func() {
		defer wg.Done()

		conn, err := l.Accept()
		require.NoError(t, err)
		require.NotNil(t, conn)
		defer func() { require.NoError(t, conn.Close()) }()

		request := make([]byte, 1024)
		n, err := conn.Read(request)
		require.NoError(t, err)
		require.Equal(t, "hello\n", string(request)[:n])

		n, err = conn.Write([]byte("world\n"))
		require.NoError(t, err)
		require.NotEqual(t, 0, n)
	}()

	wg.Wait()
}

func TestTelnetClient_errorConnect(t *testing.T) {
	tests := []struct {
		telnetClient
	}{
		{
			telnetClient: telnetClient{
				address: "incorrect address",
				in:      os.Stdin,
				out:     os.Stdout,
			},
		},
	}

	for _, tc := range tests {
		err := tc.telnetClient.Connect()
		require.Error(t, err)
	}
}

func TestTelnetClient_(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:")
	require.NoError(t, err)
	defer func() { require.NoError(t, l.Close()) }()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		in := &bytes.Buffer{}
		out := &bytes.Buffer{}

		timeout := time.Second
		require.NoError(t, err)

		client := NewTelnetClient(l.Addr().String(), timeout, ioutil.NopCloser(in), out)
		require.NoError(t, client.Connect())
		defer func() { require.NoError(t, client.Close()) }()

		in.WriteString("hello\n")
		in.WriteString("world\n")
		in.WriteString("!")
		in.WriteString("!")
		in.WriteString("!")
		err = client.Send()
		require.NoError(t, err)

		err = client.Receive()
		require.NoError(t, err)
		require.Equal(t, "Some information bla-bla\n", out.String())
	}()

	go func() {
		defer wg.Done()

		conn, err := l.Accept()
		require.NoError(t, err)
		require.NotNil(t, conn)
		defer func() { require.NoError(t, conn.Close()) }()

		request := make([]byte, 1024)
		n, err := conn.Read(request)
		require.NoError(t, err)
		require.Equal(t, "hello\nworld\n!!!", string(request)[:n])

		n, err = conn.Write([]byte("Some information bla-bla\n"))
		require.NoError(t, err)
		require.NotEqual(t, 0, n)
	}()

	wg.Wait()
}
