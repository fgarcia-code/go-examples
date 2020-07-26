package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func fetchHTML(URL string) (io.ReadCloser, error) {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("error fetching URL: %v\n", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("response status code was %d\n", resp.StatusCode)
		return nil, err
	}
	ctype := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ctype, "text/html") {
		log.Fatalf("response content type was %s not text/html\n", ctype)
		return nil, err
	}
	return resp.Body, nil
}

func extractTitle(body io.ReadCloser) (string, error) {
	defer body.Close()
	tokenizer := html.NewTokenizer(body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
			return "", tokenizer.Err()
		}
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if "title" == token.Data {
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					return tokenizer.Token().Data, nil
				}
			}
		}
	}
	return "", nil
}

func fetchTitle(URL string) (string, error) {
	body, err := fetchHTML(URL)
	if err != nil {
		log.Fatalf("response can't be read %s not text/html\n", err)
		return "", err
	}
	title, err := extractTitle(body)
	if err != nil {
		log.Fatalf("while reading the title %s not text/html\n", err)
		return "", err
	}
	return title, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:\\n  pagetitle <url>\\n")
	}
	title, err := fetchTitle(os.Args[1])
	if err != nil {
		log.Fatalf("error fetching page title: %v\n", err)
	}

	fmt.Println("Page title:", title)
}