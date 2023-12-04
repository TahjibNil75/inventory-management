package src

import (
	"net/http"

	"github.com/TahjibNil75/inventory-management/models"
	"github.com/gin-gonic/gin"
)

func GenerateAndUploadQRCode(ctx *gin.Context) {
	var qrcodeParams models.QRCodeParams
	if err := ctx.ShouldBindJSON(&qrcodeParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	qrCode := GenerateQRCode(qrcodeParams)
	err := uploadToS3(qrCode, "qrcode/"+qrcodeParams.AssetTag+".png")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to upload qrcode in s3",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "qrcode was successfully uploaded to s3",
	})

}
