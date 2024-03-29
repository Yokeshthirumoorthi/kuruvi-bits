/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package services

import (
	"fmt"
	"context"
	"log"
	"os"
	"time"
    "io"
    "net/http"
	// "encoding/json"
	"google.golang.org/grpc"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "github.com/kuruvi-bits/transform/pb"
	utils "github.com/kuruvi-bits/transform/utils"
)

const (
	FACE_DETECT_ENDPOINT = "192.168.1.100:8006"
	FACE_DESCRIBE_ENDPOINT = "192.168.1.100:8009"
)

func DescribeFace(message utils.Message) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(FACE_DESCRIBE_ENDPOINT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFaceDescribeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Url := utils.GetResizedVolPath(message)

	faceDescription, err := c.DescribeFaces(ctx, &pb.DescribeFaceRequest{
		AlbumName: message.AlbumName,
		PhotoName: message.PhotoName,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println("Face describe", faceDescription)
	// return faceBoxes	
}

func CropFace(message utils.Message, boundingBox *pb.BoundingBox, index int) {
	dirPath := fmt.Sprintf("%s/%s", utils.FACES_VOL, message.AlbumName)
	fileName  := fmt.Sprintf("%d_%s", index, message.PhotoName)
    filePath := fmt.Sprintf("%s/%s", dirPath, fileName)
    utils.CreateDirIfNotExist(dirPath)

    url := utils.GetFaceCropURL(message, boundingBox)

    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }
    defer response.Body.Close()

    //open a file for writing
    file, err := os.Create(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Use io.Copy to just dump the response body to the file. This supports huge files
    _, err = io.Copy(file, response.Body)
    if err != nil {
        log.Fatal(err)
	}
	
	faceMessage := utils.Message{
		AlbumName: message.AlbumName,
		PhotoName: fileName,
	}

	DescribeFace(faceMessage)

    fmt.Println("Success!")
}

func DetectFaces(message utils.Message) *pb.BoundingBoxes {
	// Set up a connection to the server.
	conn, err := grpc.Dial(FACE_DETECT_ENDPOINT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFaceCoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	Url := utils.GetResizedVolPath(message)

	faceBoxes, err := c.DetectFaces(ctx, &pb.PhotoURL{Url: Url})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println("Face detect", faceBoxes)
	return faceBoxes
}


