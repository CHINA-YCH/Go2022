package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: bcrypt
 * @Version: ...
 * @Date: 2022-11-01 15:42:22
 */

func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		return false
	}
	return true
}
