package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	bufSize = 1024 * 8
)

func main() {
	var url = flag.String("u", "", "url сайта")
	var fileName = flag.String("fn", "", "название файла")

	log.SetFlags(0)
	flag.Parse()

	if *url == "" {
		panic("Пустой url")
	}

	resp := getResponse(*url)
	if *fileName == "" {
		urlSplit := strings.Split(*url, "/")
		fileName = &urlSplit[len(urlSplit)-1]
	}
	*fileName = "download\\" + *fileName + ".html"
	writeToFile(*fileName, resp)
}

func getResponse(url string) *http.Response {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}

func writeToFile(fileName string, resp *http.Response) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufferedWriter := bufio.NewWriterSize(file, bufSize)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(bufferedWriter, resp.Body)
	if err != nil {
		panic(err)
	}
}
