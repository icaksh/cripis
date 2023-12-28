package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/icaksh/cripis/app/models"
	"github.com/jung-kurt/gofpdf"
	"mime/multipart"
	"time"
)

const (
	accountId       = "aa"
	accessKeyId     = "bb"
	accessKeySecret = "cc"
	bucketName      = "dd"
)

func UploadFiletoS3(user string, file *multipart.FileHeader) (bool, string, error) {
	var err error
	returnSuccess := false
	returnMessage := "NULL"

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
			HostnameImmutable: true,
			Source:            aws.EndpointSourceCustom,
		}, nil
	})
	fmt.Println(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)

	client := s3.NewFromConfig(cfg)

	currentDate := time.Now()

	dateString := currentDate.Format("2006/01")
	dir := dateString + "/" + user
	fileName := dir + "/" + file.Filename

	bucket := aws.String(bucketName)
	key := aws.String(fileName)

	srcFile, err := file.Open()
	if err != nil {
		returnSuccess = false
		returnMessage = "Failed to open file"
	}
	defer srcFile.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: bucket,
		Key:    key,
		Body:   srcFile,
	})

	if err != nil {
		returnSuccess = false
		returnMessage = "Failed to upload object" + fileName
		err = err
	} else {
		returnSuccess = true
		returnMessage = "https://cdn.icaksh.my.id/" + fileName
		err = nil
	}
	return returnSuccess, returnMessage, err
}

func UploadFiletoS3PDF(user string, file *gofpdf.Fpdf, v *models.Trademark) (bool, string, error) {
	var err error
	returnSuccess := false
	returnMessage := "NULL"

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
			HostnameImmutable: true,
			Source:            aws.EndpointSourceCustom,
		}, nil
	})
	fmt.Println(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)

	client := s3.NewFromConfig(cfg)

	currentDate := time.Now()

	dateString := currentDate.Format("2006/01")
	dir := dateString + "/" + user
	fileName := dir + "/" + v.RegisterNumber + ".pdf"

	bucket := aws.String(bucketName)
	key := aws.String(fileName)

	var pdfBuffer bytes.Buffer
	if err := file.Output(&pdfBuffer); err != nil {
		returnSuccess = false
		returnMessage = "Failed to open file"
	}

	pdfReader := bytes.NewReader(pdfBuffer.Bytes())

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: bucket,
		Key:    key,
		Body:   pdfReader,
	})

	if err != nil {
		returnSuccess = false
		returnMessage = "Failed to upload object" + fileName
		err = err
	} else {
		returnSuccess = true
		returnMessage = "https://cdn.icaksh.my.id/" + fileName
		err = nil
	}
	return returnSuccess, returnMessage, err
}
