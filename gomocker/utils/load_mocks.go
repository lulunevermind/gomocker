package utils

import (
	"io/ioutil"
	"strings"
)

func LoadReqResps(dir string) map[string]string {

	mapping := map[string]string{}

	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {

		if strings.HasSuffix(f.Name(), ".req") || strings.HasSuffix(f.Name(), ".resp") {

			resp, err := ioutil.ReadFile(dir + f.Name())
			if err != nil {
				panic(err)
			}
			as_string := string(resp)
			mapping[f.Name()] = as_string
		}
	}
	return mapping
}
