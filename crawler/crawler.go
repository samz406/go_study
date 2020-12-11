package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	download("http://wwww.baidu.com")

}

func download(url string) {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	response, err := client.Do(req)

	if err != nil {
		fmt.Println("read error", err)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("body read error", err)
		return
	}

	fmt.Println(string(body))

}
