package main

import (
    "fmt"
    rtctokenbuilder "github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder2"
)

func main() {
    appID := "970CA35de60c44645bbae8a215061b33"
    appCertificate := "5CFd2fd1755d40ecb72977518be15d3b"
    channelName := "7d72365eb983485397e3e3f9d460bdda"
    uid := uint32(2882341273)
    uidStr := "2882341273"
    expire := uint32(600)

    token, _ := rtctokenbuilder.BuildTokenWithUid(appID, appCertificate, channelName, uid, rtctokenbuilder.RolePublisher, expire)
    fmt.Println("BuildTokenWithUid: " + token)

    token, _ = rtctokenbuilder.BuildTokenWithAccount(appID, appCertificate, channelName, uidStr, rtctokenbuilder.RolePublisher, expire)
    fmt.Println("BuildTokenWithAccount: " + token)
}
