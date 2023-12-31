package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"terhandle/internal/features/request-teknisi/dto"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func R2Config() aws.Config {
	godotenv.Load()
	accountId := os.Getenv("ACCOUNT_ID")
	accessKeyId := os.Getenv("ACCESSKEY_ID")
	accessKeySecret := os.Getenv("SECRET_ACCES_KEY")

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func UploadFile(fileForm *multipart.FileHeader) (string, error) {

	client := s3.NewFromConfig(R2Config())
	godotenv.Load()
	bucketName := os.Getenv("BUCKET_NAME")

	file, err := fileForm.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	randomKey := uuid.New().String()
	contentType := "image/jpeg"

	if ext := filepath.Ext(fileForm.Filename); ext != "" {
		ext = strings.TrimPrefix(ext, ".")
		switch ext {
		case "jpg", "jpeg":
			contentType = "image/jpeg"
		case "png":
			contentType = "image/png"
		default:
			return "",errors.New("file harus berupa gambar png/jpeg")
		}
	}

	input := &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &randomKey,
		Body:        file,
		ContentType: aws.String(contentType),
	}
	_, err = client.PutObject(context.TODO(), input)
	if err != nil {
		return "", err
	}
	// publicURL := fmt.Sprintf("https://pub-d82858f6b028480cb736d00b2cfdbe32.r2.dev/%s", randomKey)
	publicURL := fmt.Sprintf("https://cloud.apicode.my.id/%s", randomKey)

	return publicURL, nil
}

func ProcessUploadFiles(files []*multipart.FileHeader) ([]dto.RequestFoto, error) {
	var fotos []dto.RequestFoto

	for _, file := range files {
		url, err := UploadFile(file)
		if err != nil {
			return nil, errors.New("file harus berupa gambar png/jpeg")
		}

		fotos = append(fotos, dto.RequestFoto{Foto: url})
	}

	return fotos, nil
}
