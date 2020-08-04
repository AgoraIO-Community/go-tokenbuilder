package main

import (
	"fmt"

	rtmtokenbuilder "github.com/AgoraIO-Community/go-tokenbuilder/rtmtokenbuilder"
)

func main() {
	appID := "970CA35de60c44645bbae8a215061b33"
	appCertificate := "5CFd2fd1755d40ecb72977518be15d3b"
	userId := "test_user"
	expire := uint32(600)

	token, _ := rtmtokenbuilder.BuildToken(appID, appCertificate, userId, expire)
	fmt.Println("BuildToken: " + token)
}
