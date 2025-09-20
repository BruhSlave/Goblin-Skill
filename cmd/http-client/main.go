package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	urls := []string{"https://ya.ru", "https://ya.ru/?nr=1", "https://ya.ru/showcaptcha?cc=1&..."}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
	}

	for _, url := range urls {
		response, err := client.Get(url)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Printf("\n\n\n\nResponse from %v\nnstatus: %v\n", url, response.Status)
		fmt.Printf("Response body \n\n %s", string(body))
	}
}
