package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"

	"github.com/TahjibNil75/inventory-management/models"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenerateQRCode(data models.QRCodeParams) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding json")
		return nil
	}
	qrCode, err := qr.Encode(string(jsonData), qr.M, qr.Auto)
	if err != nil {
		fmt.Println("Error Generating QR Code")
		return nil
	}

	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		fmt.Println("Error Scaling QR Code")
		return nil
	}

	pngImage := new(bytes.Buffer)
	err = png.Encode(pngImage, qrCode)
	if err != nil {
		fmt.Println("failed to convert in PNG")
		return nil
	}
	return pngImage.Bytes()
}
