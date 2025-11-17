package generation

import (
	"bytes"
	"fmt"
	"net/http"
)

func Submit(year, day, part int, answer string) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	body := fmt.Sprintf("level=%d&answer=%s", part, answer)
	res, e := prepareRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)), true)
	if e != nil {
		return "", fmt.Errorf("error creating/sending request: %w", e)
	}
	defer res.Body.Close()
	msg, e := ArticleParagraphText(res.Body)
	if e != nil {
		return "", fmt.Errorf("error processing response body: %w", e)
	}
	return msg, nil
}
