package util

import (
	"daidai-server/pkg/setting"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"log"
)

var (
	client  *oss.Client
	bucket  *oss.Bucket
	fileUrl string
)

func init() {
	sec, err := setting.Cfg.GetSection("aliyunoss")
	if err != nil {
		log.Fatal(2, "Fail to get section 'aliyunoss': %v", err)
	}
	endoint := sec.Key("endpoint").String()
	accessKeyId := sec.Key("keyid").String()
	accessKeySecret := sec.Key("keysecret").String()
	bucketName := sec.Key("bucketname").String()
	fileUrl = sec.Key("fileUrl").String()
	client, err := oss.New(endoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Fatal(2, "Fail to get section 'oss': %v", err)
	}
	// 创建存储空间。
	bucket, err = client.Bucket(bucketName)
	if err != nil {
		log.Fatal(2, "Fail to get section 'oss': %v", err)
	}
}

func UpLoadFile(path string, file io.Reader) (url string, err error) {
	return fileUrl + "/" + path, bucket.PutObject(path, file)
}
