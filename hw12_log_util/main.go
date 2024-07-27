package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

func main() {
	var path string
	var out strings.Builder

	filter := pflag.StringP("level", "l", "", "Select level for filter")
	watch := pflag.BoolP("watch", "w", false, "displays all found lines after statistics")
	pflag.StringVarP(&path, "path", "p", "", "Path to the file for analysis ")
	logoutflie := pflag.StringP("output", "o", "", "Select output path for log")
	pflag.Parse()
	env, ok := os.LookupEnv("CONFIG_ENV")
	if ok {
		godotenv.Load(env)
	}
	filterenv, ok := os.LookupEnv("LOG_ANALYZER_LEVEL")

	switch {
	case len(*filter) > 0:
	case ok:
		{
			*filter = filterenv
		}
	default:
		{
			*filter = "info"
		}
	}
	pathenv, ok := os.LookupEnv("LOG_ANALYZER_FILE")
	switch {
	case len(path) > 0:

	case ok:
		{
			path = pathenv
		}
	default:
		{
			fmt.Println("No path to log detected! Exiting...")
			return
		}
	}
	logoutenv, ok := os.LookupEnv("LOG_ANALYZER_OUTPUT")

	switch {
	case len(*logoutflie) > 0:
	case ok:
		{
			*logoutflie = logoutenv
		}
	}
	fmt.Printf("options selected:\nfilter: %v\npath to original log file: %v\ndisplay found lines:", *filter, path)
	fmt.Printf(" %v\npath to script output (if path empty all goes to STDOUT): %v\n\n", *watch, *logoutflie)

	stringsliceinput, err := readfile(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	sorted := filterfile(stringsliceinput, *filter)
	sortedstr := strings.Join(sorted, "\n")

	out.WriteString(fmt.Sprintf("number of input lines:%v\n", len(stringsliceinput)))
	out.WriteString(fmt.Sprintf("number of filtered lines:%v\n", len(sorted)))
	if *watch {
		out.WriteString(sortedstr)
	}

	if len(*logoutflie) > 0 {
		file, ero := os.Create(*logoutflie)
		if ero != nil {
			fmt.Printf("Error: %v\n", ero)
			return
		}
		defer file.Close()
		file.WriteString(out.String())
	} else {
		fmt.Println(out.String())
	}
}

func readfile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
			return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	s := strings.Split(string(bytes), "\n")

	return s, nil
}

func filterfile(input []string, sort string) []string {
	output := []string{}
	for _, v := range input {
		if strings.Contains(v, sort) {
			output = append(output, v)
		}
	}
	return output
}
