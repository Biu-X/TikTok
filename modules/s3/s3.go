package s3

import (
	"biu-x.org/TikTok/modules/config"
	"github.com/eleven26/goss/v3"
)

var (
	s3  *goss.Goss
	err error
)

func Init() {
	s3, err = goss.New(goss.WithConfig((*goss.Config)(&config.S3Config)))
	if err != nil {
		return
	}
}
