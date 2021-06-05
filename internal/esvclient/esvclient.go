package esvclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	ESVApiKey = "ESV_API_KEY"
	ESVApiUrl = "https://api.esv.org/v3/passage/text/?include-headings=false&include-footnotes=false&include-short-copyright=false&q="
)

type ESVClient struct {
	BaseURL *url.URL
	*http.Client
}

func NewClient(timeout time.Duration) (c *ESVClient, err error) {
	c = new(ESVClient)

	c.Client = &http.Client{
		Timeout: timeout,
	}

	return c, nil
}

func validParam(v string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.Replace(space.ReplaceAllString(v, " "), " ", "+", 1)
}

func (c *ESVClient) GetVerses(param string) (string, error) {
	passage := ""
	key := os.Getenv(ESVApiKey)
	fmt.Println(key)
	if key == "" {
		return passage, fmt.Errorf("no environment variable: %s", ESVApiKey)
	}

	path := fmt.Sprintf("%s%s", ESVApiUrl, validParam(param))
	fmt.Println(path)
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", key))
	response, err := c.Do(req)
	if err != nil {
		return passage, err
	}

	type result struct {
		Query     string   `json:"query"`
		Canonical string   `json:"canonical"`
		Passages  []string `json:"passages"`
	}

	var res result
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &res)
	if err != nil {
		return passage, err
	}

	if len(res.Passages) > 0 {
		passage = res.Passages[0]
	}

	return passage, nil
}
