package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"regexp"
	"strings"
	"os"
)

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

func getAdbList(url string, adbList chan []string) {
	var adbListTmp []string
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	adbContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	adbLines := strings.Split(string(adbContents), "\n")
	re := regexp.MustCompile(`\|\|([a-z0-9\.\-]+[a-z]+[a-z0-9\.\-]+)\^$`)

	for line := range adbLines {
		match := re.FindStringSubmatch(adbLines[line])
		if len(match) == 2 {
			adbListTmp = append(adbListTmp, "address=/"+match[1]+"/0.0.0.0/")
		}
	}
	adbList <- adbListTmp
}

func main() {
	cmd := parseCmd()

	url := strings.Split(cmd.urls, "|")

	adbList := make(chan []string, len(url))

	for i := range url {
		go getAdbList(url[i], adbList)
	}

	var mode int

	if cmd.saveMode == "a" {
		mode = os.O_CREATE | os.O_RDWR | os.O_APPEND
	} else {
		mode = os.O_CREATE | os.O_RDWR | os.O_TRUNC
	}

	fp, err := os.OpenFile(cmd.savePath, mode, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	defer fp.Close()

	var adbContents []string

	for range url {
		adbContents = MergeSlice(adbContents, <-adbList)
	}

	fp.WriteString(strings.Join(adbContents, "\n"))
}
