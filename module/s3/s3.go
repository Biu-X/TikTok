package s3

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/log"
	"github.com/eleven26/goss/v3"
	"io"
)

var (
	s3  *goss.Goss
	err error
)

func Init() {
	cfg := &config.S3Config
	s3, err = goss.New(goss.WithConfig((*goss.Config)(cfg)))
	if err != nil {
		log.Logger.Fatalf("init goss faild: %v", err)
	}
}

// Put saves the content read from r to the key of oss.
func Put(key string, r io.Reader) error {
	return s3.Put(key, r)
}

// PutFromFile saves the file pointed to by the `localPath` to the oss key.
func PutFromFile(key string, localPath string) error {
	return s3.PutFromFile(key, localPath)
}

// Get gets the file pointed to by key.
func Get(key string) (io.ReadCloser, error) {
	return s3.Get(key)
}

// GetString gets the file pointed to by key and returns a string.
func GetString(key string) (string, error) {
	return s3.GetString(key)
}

// GetBytes gets the file pointed to by key and returns a byte array.
func GetBytes(key string) ([]byte, error) {
	return s3.GetBytes(key)
}

// GetToFile saves the file pointed to by key to the localPath.
func GetToFile(key string, localPath string) error {
	return s3.GetToFile(key, localPath)
}

// Delete the file pointed to by key.
func Delete(key string) error {
	return s3.Delete(key)
}

// Exists determines whether the file exists.
func Exists(key string) (bool, error) {
	return s3.Exists(key)
}

// Size fet the file size.
func Size(key string) (int64, error) {
	return s3.Size(key)
}
