/*
 * @FilePath: /workflow-server/pkg/jwt/jwt.go
 * @Author: maggot-code
 * @Date: 2023-08-15 20:41:15
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-16 01:48:09
 * @Description:
 */
package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

type MetadataClaims struct {
	jwt.RegisteredClaims
	Metadata interface{} `json:"metadata"`
}

func New(uer string, exp time.Time, metadata interface{}) *MetadataClaims {
	return &MetadataClaims{
		Metadata: metadata,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    uer,
			ID:        uuid.NewV4().String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
}

func Issue(key string, mc *MetadataClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)

	str, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return str, nil
}

func Parse(key, str string) (*MetadataClaims, error) {
	token, err := jwt.ParseWithClaims(str, &MetadataClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	mc, ok := token.Claims.(*MetadataClaims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	return mc, nil
}

func Refresh(key string, exp time.Time, mc *MetadataClaims) (string, error) {
	mc.ID = uuid.NewV4().String()
	mc.NotBefore = jwt.NewNumericDate(time.Now())
	mc.IssuedAt = jwt.NewNumericDate(time.Now())
	mc.ExpiresAt = jwt.NewNumericDate(exp)

	return Issue(key, mc)
}

// 根据term时间判断是否需要刷新token
func NeedRefresh(term time.Duration, mc *MetadataClaims) bool {
	return time.Now().Add(term).After(mc.ExpiresAt.Time)
}
