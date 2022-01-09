package service

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var FQDN = map[int]string{
	1: "www.google.com",
	2: "www.amazon.co.jp",
}

func TestReadData(t *testing.T) {
	for _, fqdn := range FQDN {
		res, err := ReadDataAndRewiteURL(fqdn)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile("test/rewite_"+fqdn+".html", []byte(res), os.ModePerm)
	}
}
