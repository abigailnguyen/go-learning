package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/iterator"
)

func implicit() {
	ctx := context.Background()

	// For API packages whose import path is starting with "cloud.google.com/go",
	// such as cloud.google.com/go/storage in this case, if there are no credentials provided, the client library
	// will look for credentials in the environment.
	storageClient, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}

	it := storageClient.Buckets(ctx, "pc-universal-print-connector")

	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bucketAttrs.Name)
	}

	// For packages whose import path is starting with "google.golang.org/api",
	// such as google.golang.org/api/cloudkms/v1, use NewService to create the client.
	kmsService, err := cloudkms.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_ = kmsService
}

func makeGetRequest(w io.Writer, targetURL string) error {
	ctx := context.Background()

	// client is a http.Client that automatically adds an "Authorization" header
	// to any requests made.
	client, err := idtoken.NewClient(ctx, targetURL)
	if err != nil {
		return fmt.Errorf("idtoken.NewClient: %v", err)
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		return fmt.Errorf("client.Get: %v", err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	return nil
}

func main() {
	functionURL := "https://us-central1-pc-universal-print-connector.cloudfunctions.net/update-stats"
	var buf bytes.Buffer
	makeGetRequest(&buf, functionURL)
	s := buf.String()
	fmt.Print(s)
}
