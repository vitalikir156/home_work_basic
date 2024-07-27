package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Client(u string) (string, error) {
	result, err := http.Get(u+"/get") //nolint
	if err != nil {
		return "", err
	}
	defer result.Body.Close()
	resultpost, err := http.Post(u+"/save", "", bytes.NewReader([]byte("Test send"))) //nolint
	if err != nil {
		return "", err
	}
	defer resultpost.Body.Close()

	bodypost, err := io.ReadAll(resultpost.Body)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return "", err
	}
	var out strings.Builder
	out.WriteString(string(bodypost))
	out.WriteString(fmt.Sprintf("\n%v\n", resultpost.StatusCode))
	out.WriteString(string(body))
	out.WriteString(fmt.Sprintf("\n%v", result.StatusCode))
	return out.String(), err
}
