package fixapp

import (
	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp/printer"
	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp/reader"
	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp/types"
)

func Fixapp(path string) ([]types.Employee, error) {
	//	fmt.Printf("Enter data file path: ")
	//	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "fixapp/data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		return nil, err
	}

	err = printer.PrintStaff(staff)
	if err != nil {
		return nil, err
	}
	return staff, nil
}
