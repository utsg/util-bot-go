package util

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func getBodyFromResponse(response http.Response) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func IsListContains(list []string, value string) bool {
	for _, val := range list {
		if value == val {
			return true
		}
	}
	return false
}

func IsUserAllowed(userName string) bool {
	accessList := strings.Split(os.Getenv("ACCESS_LIST"), ",")
	return IsListContains(accessList, userName)
}

func GetIp() string {
	response, err := http.Get("https://api.ipify.org")
	if err != nil {
		return err.Error()
	}
	return getBodyFromResponse(*response)
}
