package main

import (
	"os"
	"io/ioutil"
	"io"
	"bytes"
	"net/http"
)

func DownLoad() {
	url := "https://o9kwm2s9v.qnssl.com/mall/1000058/test_img/nRurLJwMFLGboDlmuCAf20160825041559.png"
	out, _ := os.Create("pic.png")
	defer out.Close()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	pix, _ := ioutil.ReadAll(resp.Body)
	io.Copy(out, bytes.NewReader(pix))
}
