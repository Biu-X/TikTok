package routers

import (
	v1 "biu-x.org/TikTok/routers/api/v1"
	"errors"
	"log"
)

var StartGinServeErr = errors.New("start Gin filed")

func Init() {
	err := NewWeb()
	if err != nil {
		return
	}
}

func NewWeb() error {
	e := v1.NewAPI()
	err := e.Run()
	if err != nil {
		log.Fatalf("%v: %v\n", StartGinServeErr, err)
		return err
	}
	return nil
}
