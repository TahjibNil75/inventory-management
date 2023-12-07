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

func UpdateUploadQRCode(ctx *gin.Context) {
	var qrCodeParams models.QRCodeParams
	if err := ctx.ShouldBindJSON(&qrCodeParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update qrcode",
		})
		return
	}
	// Check if the asset with the given tag is already uploaded
	if !CheckIfAssetExistsinS3(qrCodeParams.AssetTag) {
		err := DeleteFromS3("qrcode/" + qrCodeParams.AssetTag + ".png")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to delete previos qrcode",
			})
			return
		}
	}
	// Generate the updated qrcode
	qrCode := GenerateQRCode(qrCodeParams)

	// Upload to s3
	err := uploadToS3(qrCode, "qrcode/"+qrCodeParams.AssetTag+".png")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to upload qrcode",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "qrcode updated successfully",
	})
}

// func DeleteQRCode(ctx *gin.Context) {
// 	assetTag := ctx.Param("assetTag")

// 	if !CheckIfAssetExistsinS3(assetTag) {
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"message": "Asset not found",
// 		})
// 		return
// 	}
// 	err := DeleteFromS3("qrcode/" + assetTag + ".png")
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Failed to delete qrcode",
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "qrcode deleted successfully",
// 	})
// }
