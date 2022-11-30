package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const BaseUrl = "https://adventofcode.com"
const UserAgent = "github.com/hkennyv/aocfetch"

func FetchDay(year, day int) ([]byte, error) {
	url := makeAocUrl(year, day)

	body, err := fetchAocUrl(url)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Full request (from chrome browser, redacted session cookie)
//
//	curl 'https://adventofcode.com/2015/day/6/input' \
//	  -H 'authority: adventofcode.com' \
//	  -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
//	  -H 'accept-language: en-US,en;q=0.9' \
//	  -H 'cache-control: max-age=0' \
//	  -H 'cookie: session=<REDACTED>' \
//	  -H 'referer: https://adventofcode.com/2015/day/6' \
//	  -H 'sec-ch-ua: "Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"' \
//	  -H 'sec-ch-ua-mobile: ?0' \
//	  -H 'sec-ch-ua-platform: "macOS"' \
//	  -H 'sec-fetch-dest: document' \
//	  -H 'sec-fetch-mode: navigate' \
//	  -H 'sec-fetch-site: same-origin' \
//	  -H 'sec-fetch-user: ?1' \
//	  -H 'upgrade-insecure-requests: 1' \
//	  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36' \
//	  --compressed
func fetchAocUrl(url string) ([]byte, error) {
	session := getSession()
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{}
	cookie.Name = "session"
	cookie.Value = session

	req.AddCookie(cookie)
	req.Header.Add("user-agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// check status code
	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	return body, nil
}

func makeAocUrl(year, day int) string {
	return fmt.Sprintf("%s/%d/day/%d/input", BaseUrl, year, day)
}

// TODO: refactor this to read from a config file
func getSession() string {
	data, err := os.ReadFile("./session.txt")
	if err != nil {
		fmt.Println("No 'session.txt' found, please create this file with your session cookie.")
		os.Exit(1)
	}

	return strings.TrimSpace(string(data))
}
