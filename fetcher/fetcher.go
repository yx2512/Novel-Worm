package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var RateLimiter = time.NewTicker(100 * time.Millisecond).C

func Fetch(url string) ([]byte, error) {
	<-RateLimiter
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Printf("cannot create request with url: %s", url)
	}

	request.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyReader := bufio.NewReader(resp.Body)
		e := determineEncoding(bodyReader)
		utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
		return ioutil.ReadAll(utf8Reader)
	} else {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher warning: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
