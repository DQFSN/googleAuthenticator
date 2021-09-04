/*******************************************************
	File Name: main.go
	Author: An
	Mail:lijian@cmcm.com
	Created Time: 14/11/25 - 10:24:49
	Modify Time: 14/11/25 - 10:24:49
 *******************************************************/
package main

import (
	"fmt"
	"github.com/DQFSN/googleAuthenticator"
)

func createSecret(ga *googleAuthenticator.GAuth) string {
	secret, err := ga.CreateSecret(16)
	if err != nil {
		return ""
	}
	return secret
}

func getCode(ga *googleAuthenticator.GAuth, secret string) string {
	code, err := ga.GetCode(secret)
	if err != nil {
		return "*"
	}
	return code
}

func verifyCode(ga *googleAuthenticator.GAuth, secret, code string) bool {
	// 1:30sec
	ret, err := ga.VerifyCode(secret, code, 1)
	if err != nil {
		return false
	}
	return ret
}

func main() {
	for i := 0; i < 100000; i++ {
		ga := googleAuthenticator.NewGAuth()
		//secret := "O3KAB4R2BPDJXJZR"
		secret, _ := ga.CreateSecret(16)
		fmt.Println(secret)
		code := getCode(ga, secret)
		//fmt.Println(code)
		if !verifyCode(ga, secret, code) {
			fmt.Println(secret)
		}

	}
}
