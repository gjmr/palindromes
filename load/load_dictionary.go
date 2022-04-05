package load

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func Load(filename string) []string {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	loaded_txt := regexp.MustCompile("\r\n|\n").Split(string(text), -1)
	for i, v := range loaded_txt {
		loaded_txt[i] = strings.ToLower(v)
	}

	return loaded_txt
}
