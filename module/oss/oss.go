package oss

import (
	"io"

	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/eleven26/goss/v3"
)

var (
	oss *goss.Goss
	err error
)

func Init() {
	cfg := &config.OSSConfig
	oss, err = goss.New(goss.WithConfig((*goss.Config)(cfg)))
	if err != nil {
		log.Logger.Errorf("init goss faild: %v", err)
	}
}

// Put saves the content read from r to the key of oss.
func Put(key string, r io.Reader) error {
	return oss.Put(key, r)
}

// PutFromFile saves the file pointed to by the `localPath` to the oss key.
func PutFromFile(key string, localPath string) error {
	return oss.PutFromFile(key, localPath)
}

// Get gets the file pointed to by key.
func Get(key string) (io.ReadCloser, error) {
	return oss.Get(key)
}

// GetString gets the file pointed to by key and returns a string.
func GetString(key string) (string, error) {
	return oss.GetString(key)
}

// GetBytes gets the file pointed to by key and returns a byte array.
func GetBytes(key string) ([]byte, error) {
	return oss.GetBytes(key)
}

// GetToFile saves the file pointed to by key to the localPath.
func GetToFile(key string, localPath string) error {
	return oss.GetToFile(key, localPath)
}

// Delete the file pointed to by key.
func Delete(key string) error {
	return oss.Delete(key)
}

// Exists determines whether the file exists.
func Exists(key string) (bool, error) {
	return oss.Exists(key)
}

// Size fet the file size.
func Size(key string) (int64, error) {
	return oss.Size(key)
}
