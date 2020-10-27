package qr

import (
	"bytes"
	"encoding/base64"
	"github.com/skip2/go-qrcode"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	log "github.com/sirupsen/logrus"
)

// 生成的小边框二维码
func GenSmallPaddingQR(url string, size int, useBase64 bool) []byte {
	newUrl := url
	if useBase64 {
		base64Url, err := base64.StdEncoding.DecodeString(url)
		if err == nil {
			newUrl = string(base64Url)
		}
	}

	// Create the barcode
	qrCode, err := qr.Encode(newUrl, qr.M, qr.Auto)
	if err != nil {
		log.Errorf("Generate small padding qr failed, url: %s, size: %d, useBase64: %t, err: %s", url, size, useBase64, err)
		return []byte{}
	}

	// Scale the barcode to size * size pixels
	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		log.Errorf("Generate small padding qr failed, url: %s, size: %d, useBase64: %t, err: %s", url, size, useBase64, err)
		return []byte{}
	}

	// encode the barcode as png
	var b bytes.Buffer
	err = png.Encode(&b, qrCode)
	if err != nil {
		log.Errorf("Generate small padding qr failed, url: %s, size: %d, useBase64: %t, err: %s", url, size, useBase64, err)
		return []byte{}
	}

	return b.Bytes()
}

// 生成的大边框二维码
func GenBigPaddingQR(url string, size int, useBase64 bool) []byte {
	newUrl := url
	if useBase64 {
		base64Url, err := base64.StdEncoding.DecodeString(url)
		if err == nil {
			newUrl = string(base64Url)
		}
	}
	buf, err := qrcode.Encode(newUrl, qrcode.Medium, size)
	if err != nil {
		log.Errorf("Generate big padding qr failed, url: %s, size: %d, useBase64: %t, err: %s", url, size, useBase64, err)
		return []byte{}
	}
	return buf
}
