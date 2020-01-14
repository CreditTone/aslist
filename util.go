package aslist

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	log "github.com/CreditTone/colorfulog"
	"reflect"
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
	typeBs := []byte(reflect.TypeOf(data).String())
	dataBytes := append(typeBs, bs...)
	Md5Inst := md5.New()
	Md5Inst.Write(dataBytes)
	result := Md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}
