package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leeexing/go-social/pkg/setting"
)

var jwtSecret = []byte(setting.JWTSecret)

// Claims 结构体。自己根据需求设置需要的参数
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// NtClaims Nuctech相关服务定义的token数据
// IssuedAt 外层结构体中覆盖 'StandardClaims' 中字段的默认数据类型(int64)
type NtClaims struct {
	NtUID          string `json:"nt_uid"`
	NtEcode        string `json:"nt_ecode"`
	NtName         string `json:"nt_name"`
	NtRole         string `json:"nt_role"`
	NtOid          string `json:"nt_oid"`
	NtVerification string `json:"nt_verification"`

	IssuedAt string `json:"iat"`
	jwt.StandardClaims
}

// GenerateToken token生成方法
func GenerateToken(username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "token.anjianba.cn?v=",
		},
	}
	// 使用指定的签名方法创建签名对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secrect签名并获取完整的编码后的字符串token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken token解密
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// ParseNtToken 解密 nuctech 服务发过来的票
func ParseNtToken(token string) (*NtClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &NtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*NtClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
