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

func bytesToMd5Hex(dataBytes []byte) string {
	Md5Inst := md5.New()
	Md5Inst.Write(dataBytes)
	result := Md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}

func MultiFieldsGanerateUniqueId(ignorePoint bool, datas ...interface{}) string {
	finallyBytes := []byte{}
	for _, data := range datas {
		if data != nil {
			bs, err := Encode(data)
			if err != nil {
				log.Warn("MultiGanerateUniqueId", err)
				continue
			}
			typeStr := reflect.TypeOf(data).String()
			if ignorePoint && strings.HasPrefix(typeStr, "*") {
				typeStr = typeStr[1:]
			}
			typeBs := []byte(typeStr)
			finallyBytes = append(finallyBytes, bs...)
			finallyBytes = append(finallyBytes, typeBs...)
		}
	}
	return bytesToMd5Hex(finallyBytes)
}

func SmartGanerateUniqueId(data interface{}) string {
	return MultiFieldsGanerateUniqueId(false, data)
}

func SmartGanerateUniqueIdWithIgnorePoint(data interface{}) string {
	return MultiFieldsGanerateUniqueId(true, data)
}
