package utils

import (
	"dmp/db"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io/ioutil"
	"mime/multipart"

	"log"
)

// UploadFile :  upload file to gridFS
func UploadFile(file *multipart.FileHeader, filename string) error {
	bucket, err := gridfs.NewBucket(db.InitiateMongoClient().Database("dmp-files"))
	if err != nil {
		log.Fatal(err)
	}
	fileContent, _ := file.Open()
	data, err := ioutil.ReadAll(fileContent) // why the long names though?
	if err != nil {
		log.Fatal(err)
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer uploadStream.Close()
	_, err = uploadStream.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
