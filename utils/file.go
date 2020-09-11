package utils

import (
	bytes "bytes"
	"context"
	"dmp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io/ioutil"
	"mime/multipart"
	"time"

	"log"
)

// UploadFile :  upload file to gridFS
func UploadFile(file *multipart.FileHeader, filename string) (string, error) {
	bucket, err := gridfs.NewBucket(db.InitiateMongoClient().Database("dmp-files"))
	if err != nil {
		log.Fatal(err)
	}
	fileContent, _ := file.Open()
	data, err := ioutil.ReadAll(fileContent)
	if err != nil {
		log.Fatal(err)
	}
	fileID := GenerateRandomNumber()
	uploadStream, err := bucket.OpenUploadStreamWithID(fileID, filename)
	if err != nil {
		log.Fatal(err)
	}
	defer uploadStream.Close()
	_, err = uploadStream.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	return fileID, err
}

// DownloadFile :  download file from gridFS
func DownloadFile(fileID string) []byte {
	connection := db.InitiateMongoClient().Database("dmp-files")

	fsFiles := connection.Collection("fs.files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var results bson.M
	log.Println("FILEID : ", fileID)
	err := fsFiles.FindOne(ctx, bson.M{"_id": fileID}).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(results)

	bucket, _ := gridfs.NewBucket(connection)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStream(fileID, &buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(dStream)
	return buf.Bytes()
}
