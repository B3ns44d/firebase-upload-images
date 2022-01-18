package controllers

import (
	"io"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

func Upload(c *gin.Context) {

	var err error

	ctx := appengine.NewContext(c.Request)
	opt := option.WithCredentialsFile("env/serviceAccountKey.json")

	storageClient, err = storage.NewClient(ctx, opt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	defer file.Close()

	bucket := os.Getenv("BUCKET_NAME")

	object := storageClient.Bucket(bucket).Object(handler.Filename)
	wc := object.NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	if err := wc.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	url := "https://storage.cloud.google.com/" + bucket + "/" + wc.Attrs().Name

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"url":     url,
		"error":   false,
	})
}
