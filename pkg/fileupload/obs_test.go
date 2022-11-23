package main

import (
	"bluebell/pkg/fileupload/obs"
	"fmt"
	"testing"
)

const (
	endpoint           = "obs.cn-south-1.myhuaweicloud.com"
	ak                 = "OMHUUOBKCEF1YUMVFQOG"
	sk                 = "BGbRqlBsjeFvNm3cP6Yr1dH0SfyyNFT6Xdzgvf1F"
	bucketName         = "bluebell-oss"
	PUBLIC_FILEFOLDER  = "public/"
	PRIVATE_FILEFOLDER = "private/"
)

var obsClient *obs.ObsClient

func getObsClient() *obs.ObsClient {
	var err error
	if obsClient == nil {
		obsClient, err = obs.New(ak, sk, endpoint)
		if err != nil {
			panic(err)
		}
	}
	return obsClient
}

func putFile(folder, file, obsfilename string) {
	input := &obs.PutFileInput{}
	input.Bucket = bucketName
	input.Key = folder + obsfilename
	input.SourceFile = file
	output, err := getObsClient().PutFile(input)
	if err == nil {
		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
		fmt.Printf("ETag:%s, StorageClass:%s\n", output.ETag, output.StorageClass)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.StatusCode)
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}

func TestUpload(t *testing.T) {
	// 创建ObsClient结构体
	var obsClient, err = obs.New(ak, sk, endpoint)
	if err == nil {
		// 使用访问OBS
		fmt.Println("连接成功")
		putFile(PUBLIC_FILEFOLDER, "text.txt", "123")
		// 关闭obsClient
		obsClient.Close()
	}
	t.Fatal("can not New Client,#err: ", err.Error())
}
