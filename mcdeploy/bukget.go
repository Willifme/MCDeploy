package mcdeploy

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func perror(err error) {

	if err != nil {

		panic(err.Error())

	}

}

func GetJSON(url string) []uint8 {

	response, err := http.Get(url)

	perror(err)

	body, err := ioutil.ReadAll(response.Body)

	perror(err)

	return body

}

func GetDownload(url string, file_name string) {

	output, err := os.Create(file_name)

	perror(err)

	defer output.Close()

	response, err := http.Get(url)

	perror(err)

	_, err = io.Copy(output, response.Body)

	perror(err)

}
