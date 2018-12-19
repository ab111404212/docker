package util

import (
	"deploy/logger"
	"deploy/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"unsafe"
)

func ReadConfig(file string) (v interface{}) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Println(err.Error())
		return nil
	}
	v = Decoding(b)
	return
}

func Encoding(v interface{}) (b []byte) {
	b, err := json.Marshal(v)
	if err != nil {
		logger.Println(err.Error())
		return
	}
	return
}

func Decoding(b []byte) interface{} {
	v := new(interface{})
	err := json.Unmarshal(b, v)
	if err != nil {
		logger.Println(err.Error())
		return err
	}

	// val := reflect.Indirect(reflect.ValueOf(v))
	p := (*model.Value)(unsafe.Pointer(v))
	p1 := *(*model.Value)(unsafe.Pointer(v))
	p2 := (*(*interface{})(unsafe.Pointer(v)))
	log.Println(p, p1.A, p2)
	return v
}
