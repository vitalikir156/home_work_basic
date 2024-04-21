package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/vitalikir156/home_work_basic/hw02_fix_app/types"
)

func ReadJSON(filePath string, _ int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	_ = json.Unmarshal(bytes, &data)

	res := data

	return res, nil
}
