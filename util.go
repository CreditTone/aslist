package aslist

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	log "github.com/CreditTone/colorfulog"
	"reflect"
	"strings"
)

func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func SmartGanerateUniqueId(data interface{}) string {
	bs, err := Encode(data)
	if err != nil {
		log.Warn("SmartGanerateUniqueId", err)
		return ""
	}
	typeStr := reflect.TypeOf(data).String()
	typeBs := []byte(typeStr)
	dataBytes := append(typeBs, bs...)
	Md5Inst := md5.New()
	Md5Inst.Write(dataBytes)
	result := Md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}

func SmartGanerateUniqueIdWithIgnorePoint(data interface{}) string {
	bs, err := Encode(data)
	if err != nil {
		log.Warn("SmartGanerateUniqueIdWithIgnorePoint", err)
		return ""
	}
	typeStr := reflect.TypeOf(data).String()
	if strings.HasPrefix(typeStr, "*") {
		typeStr = typeStr[1:]
	}
	typeBs := []byte(typeStr)
	dataBytes := append(typeBs, bs...)
	Md5Inst := md5.New()
	Md5Inst.Write(dataBytes)
	result := Md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}
