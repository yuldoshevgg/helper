package helper

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"paynet/pb_installment_service/config"
	"paynet/pb_installment_service/genproto/installment_service"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"github.com/xtgo/uuid"
)

func UploadImage(installmentId string, cfg config.Config) (*installment_service.CreateAttachmentRequest, error) {

	fName1 := uuid.NewRandom()

	dst, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "error while getting os directory")
	}

	minioClient, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyID, cfg.MinioSecretAccessKey, ""),
		Secure: true,
		Transport: &http.Transport{
			DisableCompression: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "error while connectiong minio")
	}

	file1, err := os.Open(dst + "/document/" + config.GenerateFile)
	if err != nil {
		return nil, errors.Wrap(err, "error while opening generated file")
	}
	defer file1.Close()

	_, err = minioClient.FPutObject(
		context.Background(),
		cfg.MinioMovieUploadBucketName,
		fName1.String()+config.GenerateFile,
		dst+"/document/"+config.GenerateFile,
		minio.PutObjectOptions{ContentType: ""},
	)

	if err != nil {
		return nil, errors.Wrap(err, "error while putting object in minio")
	}

	url1 := fmt.Sprintf("https://%v/%v/%v", cfg.MinioEndpoint, cfg.MinioMovieUploadBucketName, fName1.String()+config.GenerateFile)
	fmt.Println(url1)

	attachments := []*installment_service.Attachment{}
	attachments = append(attachments, &installment_service.Attachment{
		ObjectId:   installmentId,
		Type:       "installment_generation",
		FileName:   url1,
		BucketName: cfg.MinioMovieUploadBucketName,
		FileType:   "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	})

	return &installment_service.CreateAttachmentRequest{
		Attachment: attachments,
	}, nil
}

func UploadRefundDoc(installmentId string, cfg config.Config, filename string) (*installment_service.CreateAttachmentRequest, error) {

	fName := uuid.NewRandom()

	dst, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "error while getting os directory")
	}

	minioClient, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyID, cfg.MinioSecretAccessKey, ""),
		Secure: true,
		Transport: &http.Transport{
			DisableCompression: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "error while connectiong minio")
	}

	file1, err := os.Open(dst + "/document/" + filename)
	if err != nil {
		return nil, errors.Wrap(err, "error while opening file")
	}
	defer file1.Close()

	_, err = minioClient.FPutObject(
		context.Background(),
		cfg.MinioMovieUploadBucketName,
		fName.String()+filename,
		dst+"/document/"+filename,
		minio.PutObjectOptions{ContentType: ""},
	)

	if err != nil {
		return nil, errors.Wrap(err, "error while putting object in minio")
	}

	url1 := fmt.Sprintf("https://%v/%v/%v", cfg.MinioEndpoint, cfg.MinioMovieUploadBucketName, fName.String()+filename)

	attachments := []*installment_service.Attachment{}
	attachments = append(attachments, &installment_service.Attachment{
		ObjectId:   installmentId,
		Type:       "refund_generate_doc",
		FileName:   url1,
		BucketName: cfg.MinioMovieUploadBucketName,
		FileType:   "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	})

	return &installment_service.CreateAttachmentRequest{
		Attachment: attachments,
	}, nil
}
