package s3

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"time"
	"urbathon-backend-2023/pkg/config"
	"urbathon-backend-2023/pkg/errorHandler"
)

func UploadPhotos(photos *[]multipart.FileHeader) (*[]string, *errorHandler.HttpErr) {
	var urls []string
	for _, photo := range *photos {
		filename := fmt.Sprintf("%s_%s", time.Now().Format(config.DateTimeLayout), photo.Filename)
		openedFile, _ := photo.Open()
		url, err := BucketBase.UploadFile("urbathon", filename, openedFile)
		if err != nil {
			return nil, errorHandler.New("Yandex S3 not available", http.StatusBadRequest)
		}
		urls = append(urls, url)
	}

	return &urls, nil
}
