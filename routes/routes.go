package routes

import (
	"github.com/TahjibNil75/inventory-management/src"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.POST("/upload-to-s3", src.GenerateAndUploadQRCode)
	r.PUT("/update-qrcode", src.UpdateUploadQRCode)
	// r.DELETE("/delete/:assetTag", src.DeleteQRCode)

}
