package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileReaderGood(t *testing.T) {
	file, err := readfile("dmesg.log")
	leng := len(file)
	require.NoError(t, err)
	require.Equal(t, leng, 929)
}

func TestFileReaderBad(t *testing.T) {
	_, err := readfile("noreqfile")
	require.Error(t, err)
}

func TestFilter(t *testing.T) {
	file, err := readfile("dmesg.log")
	out := filterfile(file, "ACPI:")
	require.NoError(t, err)
	require.Equal(t, len(out), 66)
}
