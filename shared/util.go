package shared

import (
	"io/ioutil"
	"regexp"
)

type Result func()

var NoopResult Result = func() {
	println("no results")
}

func LoadInputFile(loc string) (string, error) {
	data, err := ioutil.ReadFile(loc)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func RegexMatch(matcher *regexp.Regexp, input string) map[string]string {
	match := matcher.FindStringSubmatch(input)
	result := make(map[string]string)
	for i, name := range matcher.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}
