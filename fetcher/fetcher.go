package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Printf("cannot create request with url: %s", url)
	}

	request.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyReader := bufio.NewReader(resp.Body)
		e := determinEncoding(bodyReader)
		utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
		return ioutil.ReadAll(utf8Reader)
	} else {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
}

func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher warning: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
