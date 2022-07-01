package azureStorage

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func randomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%s.jpg", strconv.Itoa(r.Int()))
}

func GetAccountInfo() (string, string, string, string) {
	azrKey := os.Getenv("AZR_KEY")
	azrBlobAccountName := os.Getenv("AZR_BLOB_ACC_NAME")
	azrPrimaryBlobServiceEndpoint := fmt.Sprintf("https://%s.blob.core.windows.net/", azrBlobAccountName)
	azrBlobContainer := os.Getenv("AZR_BLOB_CONTAINER")

	return azrKey, azrBlobAccountName, azrPrimaryBlobServiceEndpoint, azrBlobContainer
}

func UploadBytesToBlob(b []byte) (string, error) {
	azrKey, accountName, endPoint, container := GetAccountInfo()            // This is our account info method
	u, _ := url.Parse(fmt.Sprint(endPoint, container, "/", randomString())) // This uses our Blob Name Generator to create individual blob urls
	credential, errC := azblob.NewSharedKeyCredential(accountName, azrKey)  // Finally we create the credentials object required by the uploader
	if errC != nil {
		return "", errC
	}

	// Another Azure Specific object, which combines our generated URL and credentials
	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background() // We create an empty context (https://golang.org/pkg/context/#Background)

	// Provide any needed options to UploadToBlockBlobOptions (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadToBlockBlobOptions)
	o := azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: "image/jpg", //  Add any needed headers here
		},
	}

	// Combine all the pieces and perform the upload using UploadBufferToBlockBlob (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadBufferToBlockBlob)
	_, errU := azblob.UploadBufferToBlockBlob(ctx, b, blockBlobUrl, o)
	return blockBlobUrl.String(), errU
}
