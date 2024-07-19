package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vitalikir156/home_work_basic/hw13_http/client"
	"github.com/vitalikir156/home_work_basic/hw13_http/server"
)

func TestHttpGood(t *testing.T) {
	go server.Server(":25565")

	out, err := client.Client("http://localhost:25565")

	require.NoError(t, err)
	require.Equal(t, "Test send\n200\nhello there!\n200", out)
}

func TestHttpBad(t *testing.T) {
	_, err := client.Client("http://localhost:25555")

	require.Error(t, err)
}
