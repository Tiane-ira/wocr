package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

func GetFileList(dirPath string, filterSuffix []string, recurcive bool) (files []string, err error) {
	if recurcive {
		err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && Contains(filterSuffix, filepath.Ext(path)) {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return
		}
	} else {
		files, err = walkThisDir(dirPath, filterSuffix)
	}
	return
}

func walkThisDir(dirPath string, filterSuffix []string) (files []string, err error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return
	}
	defer dir.Close()
	fileList, err := dir.ReadDir(-1)
	if err != nil {
		return
	}
	for _, info := range fileList {
		if !info.IsDir() && Contains(filterSuffix, filepath.Ext(info.Name())) {
			path := filepath.Join(dirPath, info.Name())
			files = append(files, path)
		}
	}
	return
}

func Contains(arr []string, ele string) bool {
	for _, key := range arr {
		if key == ele {
			return true
		}
	}
	return false
}

func Base64(filePath string, urlEncode bool) (res string) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}
	baseData := base64.StdEncoding.EncodeToString(data)
	if urlEncode {
		baseData = url.QueryEscape(baseData)
	}
	return baseData
}

func Sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func Hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

func GetSavePath(dir string) string {
	date := time.Now().Unix()
	return filepath.Join(dir, fmt.Sprintf("发票_%d.xlsx", date))
}

func GetFieldNames(obj interface{}) []string {
	var fieldNames []string

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() != reflect.Struct {
		return fieldNames
	}

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames
}

func GetFilePtr(filename string) (f *os.File, err error) {
	f, err = os.Open(filename)
	if err != nil {
		return
	}
	return
}
