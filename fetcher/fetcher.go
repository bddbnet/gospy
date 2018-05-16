package fetcher

import (
	"fmt"
	"math/rand"
	"net/http"

	"bufio"

	"log"

	"io/ioutil"

	"crypto/sha1"
	"time"

	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	var t = time.Tick(time.Millisecond * time.Duration(RandInt(3000, 5000)))
	<-t
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 获取编码
	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)

	// 转码
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func JsonFetch(url string) ([]byte, error) {
	var t = time.Tick(time.Millisecond * time.Duration(RandInt(3000, 5000)))
	<-t
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error found code is %s\n", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

// 延迟

func JsonPostFetch(id string) ([]byte, error) {
	var t = time.Tick(time.Second * time.Duration(RandInt(8, 15)))
	<-t

	// this api not need csrf
	apiUrl := "https://space.bilibili.com/ajax/member/GetInfo"

	// only result by this way
	s := `
------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="mid"

%s
------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="csrf"

%s
------WebKitFormBoundary7MA4YWxkTrZu0gW--`
	payload := strings.NewReader(fmt.Sprintf(s, string(id), FakeCsrf()))
	req, _ := http.NewRequest("POST", apiUrl, payload)

	req.Header.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Set("Referer", fmt.Sprintf("https://space.bilibili.com/%s/", string(id)))
	req.Header.Set("Host", "space.bilibili.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error found code is %s\n", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	return bytes, err

	//return ioutil.ReadAll(resp.Body)
}

// 获取编码
func DetermineEncoding(reader *bufio.Reader) encoding.Encoding {
	b, err := reader.Peek(1024)
	if err != nil {
		log.Printf("DetermineEncoding error ,%s", err)
		return unicode.UTF8
	}

	encoding, _, _ := charset.DetermineEncoding(b, "")
	return encoding
}

func NilFetch(url string) ([]byte, error) {
	return []byte{}, nil
}

func FakeCsrf() string {
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprint(time.Now().Minute())))

	str := fmt.Sprintf("%x", hash.Sum(nil))
	return string([]byte(str)[:33])
}

func RandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}
