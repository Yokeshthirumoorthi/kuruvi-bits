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
	// "os"
	"time"
	// "encoding/json"
	"google.golang.org/grpc"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "github.com/kuruvi-bits/transform/pb"
)

const (
	address     = "192.168.1.100:8003"
	// defaultName = "world"
	CADDT_SERVER_ENDPOINT = "192.168.1.100:2015"
	UPLOADS_VOL = "album-uploads"
	RESIZED_VOL = "album-resized"	
	FACES_VOL = "album-faces"
)

type Message struct {
	AlbumName string
	PhotoName string
}

type URL struct {
	Upload string
	Resized string
	Faces string
}

func GetURL(message Message) URL {
	albumName := message.AlbumName
	photoName := message.PhotoName
	server := "http://" + CADDT_SERVER_ENDPOINT + "/" 
	uploadURL := server + UPLOADS_VOL + "/" + albumName + "/uploads/" + photoName
	resizedURL := server + RESIZED_VOL + "/" + albumName + "/" + photoName
	facesURL := server + FACES_VOL + "/" + albumName + "/" + photoName
	fmt.Println(uploadURL)
	url := URL{
		Upload: uploadURL,
		Resized: resizedURL,
		Faces: facesURL,
	}
	return url
}

func Exif(message Message) *pb.ExifData {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewExifCoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	Url := GetURL(message).Upload

	exifData, err := c.ExtractExif(ctx, &pb.PhotoURL{Url: Url})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return exifData
}


