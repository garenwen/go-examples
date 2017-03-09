package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/png"
	"log"
	"os"
	"fmt"
	"image/jpeg"
	"image/draw"
	"github.com/nfnt/resize"
	"strings"
	"net/http"
	"bytes"
)

func writePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	// err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println(file.Name())
}

func QRCode() {
	base64 := "http://baidu.com"
	log.Println("Original data:", base64)
	code, err := qr.Encode(base64, qr.H, qr.Unicode)
	// code, err := code39.Encode(base64)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data: ", code.Content())

	if base64 != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, 300, 300)
	if err != nil {
		log.Fatal(err)
	}

	writePng("test.png", code)
}

func Merge() {
	file, err := os.Create("dst.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Open("logo.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	img, _ := jpeg.Decode(file1)

	file2, err := os.Open("test.png")
	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()
	img2, _ := png.Decode(file2)

	jpg := image.NewRGBA(image.Rect(0, 0, 300, 300))

	img = resize.Resize(33*3, 18*3, img, resize.Lanczos3)

	x := jpg.Bounds().Size().X / 2 - img.Bounds().Size().X / 2
	y := jpg.Bounds().Size().Y / 2 - img.Bounds().Size().Y / 2

	rect := image.Rect(x + 22, y,x + 18*3+ 22,y + 18*3)
	draw.Draw(jpg, jpg.Bounds(), img2, img2.Bounds().Min, draw.Over)                   //首先将一个图片信息存入jpg
	draw.Draw(jpg, rect, img, image.Point{22,0}, draw.Over)   //将另外一张图片信息存入jpg

	// draw.DrawMask(jpg, jpg.Bounds(), img, img.Bounds().Min, img2, img2.Bounds().Min, draw.Src) // 利用这种方法不能够将两个图片直接合成？目前尚不知道原因。

	jpeg.Encode(file, jpg, nil)
}

// 生成商铺二维码并上传到七牛云
func MakeQRCode(linkUrl, logoUrl string, logoImg *image.Image) (res []byte) {
	codeImg, err := qr.Encode(linkUrl, qr.H, qr.Unicode)
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	if linkUrl != codeImg.Content() {
		return
	}

	codeImg, err = barcode.Scale(codeImg, 300, 300)
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	jpg := image.NewRGBA(image.Rect(0, 0, 300, 300))

	draw.Draw(jpg, jpg.Bounds(), codeImg, codeImg.Bounds().Min, draw.Over)

	var img image.Image
	if logoUrl != "" {
		respUrl, err := http.Get(logoUrl)
		if err != nil {
			log.Println("error:", err.Error())
			return
		}
		defer respUrl.Body.Close()

		if respUrl.StatusCode != http.StatusOK {
			log.Println("response with code :", respUrl.StatusCode)
			return
		}

		if strings.Contains(logoUrl, ".png") {
			img, err = png.Decode(respUrl.Body)
			if err!= nil {
				log.Println(err.Error())
			}
		} else {
			img, err = jpeg.Decode(respUrl.Body)
			if err != nil {
				log.Println(err.Error())
			}
		}

		if err == nil {
			img = resize.Resize(33*3, 18*3, *logoImg, resize.Lanczos3)
			x := jpg.Bounds().Size().X / 2 - img.Bounds().Size().X / 2
			y := jpg.Bounds().Size().Y / 2 - img.Bounds().Size().Y / 2

			rect := image.Rect(x + 22, y,x + 18*3+ 22,y + 18*3)
			draw.Draw(jpg, rect, *logoImg, image.Point{22,0}, draw.Over)   //将另外一张图片信息存入jpg
		}
	}

	var resBytes []byte
	respBuffer := bytes.NewBuffer(resBytes)
	jpeg.Encode(respBuffer, jpg, nil)


	return resBytes

}