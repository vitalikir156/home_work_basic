package printer

import (
	"errors"
	"fmt"

	"github.com/vitalikir156/home_work_basic/hw06_testing/fixapp/types"
)

func PrintStaff(staff []types.Employee) error {
	var err error
	for i := 0; i < len(staff) && err == nil; i++ {
		if staff[i].UserID == 0 || staff[i].Name == "" || staff[i].DepartmentID == 0 || staff[i].Age == 0 {
			err = errors.New("bad JSON data")
		} else {
			fmt.Println(staff[i])
		}
	}
	return err
}
