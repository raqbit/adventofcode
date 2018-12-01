package shared

import "io/ioutil"

func LoadInputFile(loc string) (string, error) {
	data, err := ioutil.ReadFile(loc)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
