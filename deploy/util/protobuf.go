package util

import (
	"deploy/logger"
	"deploy/protogo"

	"github.com/golang/protobuf/proto"
)

func MarshalTest() {
	pr := person.Person{}
	pr.Person = &person.Person_Person{"wyj", "Man"}
	pr.Email = []*person.Person_Email{&person.Person_Email{"1110", "18217007736@163.com"}}
	b, err := proto.Marshal(&pr)
	if err != nil {
		logger.Println(err.Error())
		return
	}
	logger.Println(string(b))
}
