package services

import "github.com/dgrijalva/jwt-go"

type jwtService struct {
}

var JwtService = new(jwtService)

// 所有需要颁发 token 的用户模型必须实现这个接口

type JwtUser interface {
	GetUid() struct{}
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	jwt.StandardClaims
}
