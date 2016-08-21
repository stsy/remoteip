package remoteip

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func downloadString(url string) (html string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	html = string(bytes)
	return
}

// Get remote ip from url.
func Get(url string) (ip string, err error) {
	source, err := downloadString(url)
	if err != nil {
		return
	}

	re, err := regexp.Compile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	if err != nil {
		return
	}

	if re.MatchString(source) {
		ip = re.FindString(source)
		return
	}

	return
}
