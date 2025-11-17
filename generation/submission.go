package generation

import (
	"advent-of-go/utils"
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func Submit(year, day, part int, answer string) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	body := fmt.Sprintf("level=%d&answer=%s", part, answer)
	req, e := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))
	if e != nil {
		return "", fmt.Errorf("error creating request: %w", e)
	}
	cookieContents, e := utils.GetFileContents("private/cookie.txt")
	if e != nil {
		return "", fmt.Errorf("error reading cookie file: %w", e)
	}
	c := http.Cookie{Value: cookieContents, Name: "session"}
	req.AddCookie(&c)
	req.Header.Add("User-Agent", "github.com/bxdn/advent-of-go by @bxdn")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, e := (&http.Client{Timeout: 5 * time.Second}).Do(req)
	if e != nil {
		return "", fmt.Errorf("error sending request: %w", e)
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("bad status code: %d", res.StatusCode)
	}
	defer res.Body.Close()
	msg, e := ArticleParagraphText(res.Body)
	if e != nil {
		return "", fmt.Errorf("error processing response body: %w", e)
	}
	return msg, nil
}
