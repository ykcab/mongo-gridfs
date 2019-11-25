package main
​
import (
	"context"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"time"
	"path"
​
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)
​
func check(e error) {
	if e != nil {
		panic(e)
	}
}
​
func main() {
​
	
	data, err := ioutil.ReadFile("test.deb")
	
	fmt.Println(data)
	check(err)
​
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
​
	bucket, bucketErr := gridfs.NewBucket(
		client.Database("articles"),
	)
​
	if bucketErr != nil {
		log.Fatal("bucketErr")
		os.Exit(1)
	}
​
	uploadStream, uploadStreamErr := bucket.OpenUploadStream(
		"fs",
	)
	if uploadStreamErr != nil {
		fmt.Println(uploadStreamErr)
		os.Exit(1)
	}
	defer uploadStream.Close()
​
	fileSize, writeErr := uploadStream.Write(data)
	fmt.Println(fileSize)
​
	if writeErr != nil {
			log.Fatal("writeErr")
			os.Exit(1)
	}
	log.Printf("Write file to DB was succesful, File size: %d", fileSize)
}