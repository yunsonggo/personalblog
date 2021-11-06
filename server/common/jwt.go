package common

import (
	"2021/yunsongcailu/yunsong_server/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jc = config.Conf.Jwt
var jwtKey = []byte(jc.JwtKey)
var jwtIssuer = jc.Issuer

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(userId int64) (string,error) {
	// 有效期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: userId,
		StandardClaims:jwt.StandardClaims{
			// 有效期
			ExpiresAt: expirationTime.Unix(),
			// 开始时间
			IssuedAt: time.Now().Unix(),
			// 发放者
			Issuer: jwtIssuer,
			// 主题
			Subject: "user token",
		},
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString ,err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token,*Claims,error) {
	claims := &Claims{}
	token,err := jwt.ParseWithClaims(tokenString,claims,func(token *jwt.Token) (i interface{},err error) {
		return jwtKey,nil
	})
	return token,claims,err
}