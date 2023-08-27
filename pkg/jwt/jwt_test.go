/*
 * @FilePath: /workflow-server/pkg/jwt/jwt_test.go
 * @Author: maggot-code
 * @Date: 2023-08-15 22:17:30
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-16 01:46:57
 * @Description:
 */
package jwt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maggot-code/workflow-server/pkg/jwt"
)

var secretKey = "123456"

type testMetadata struct {
	UID string `json:"uid"`
}

func TestNew(t *testing.T) {
	exp := time.Now().Add(3 * time.Hour)

	jwt.New("127.0.0.1:8080", exp, testMetadata{
		UID: "123456789",
	})
}

func TestIssue(t *testing.T) {
	exp := time.Now().Add(3 * time.Hour)

	claims := jwt.New("127.0.0.1:8080", exp, testMetadata{
		UID: "123456789",
	})

	token, err := jwt.Issue(secretKey, claims)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(token)
}

func TestParse(t *testing.T) {
	exp := time.Now().Add(3 * time.Hour)

	claims := jwt.New("127.0.0.1:8080", exp, testMetadata{
		UID: "123456789",
	})

	token, err := jwt.Issue(secretKey, claims)
	if err != nil {
		t.Error(err)
	}

	mc, err := jwt.Parse(secretKey, token)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(mc)
}

func TestRefresh(t *testing.T) {
	exp := time.Now().Add(3 * time.Hour)

	claims := jwt.New("127.0.0.1:8080", exp, testMetadata{
		UID: "123456789",
	})

	token, err := jwt.Issue(secretKey, claims)
	if err != nil {
		t.Error(err)
	}

	mc, err := jwt.Parse(secretKey, token)
	if err != nil {
		t.Error(err)
	}

	nt, err := jwt.Refresh(secretKey, time.Now().Add(1*time.Hour), mc)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(nt)
}

func TestNeedRefresh(t *testing.T) {
	exp := time.Now().Add(5 * time.Second)

	claims := jwt.New("127.0.0.1:8080", exp, testMetadata{
		UID: "123456789",
	})

	token, err := jwt.Issue(secretKey, claims)
	if err != nil {
		t.Error(err)
	}

	mc, err := jwt.Parse(secretKey, token)
	if err != nil {
		t.Error(err)
	}

	not := jwt.NeedRefresh(time.Second, mc)
	fmt.Println(not)

	time.Sleep(4 * time.Second)

	yes := jwt.NeedRefresh(time.Second, mc)
	fmt.Println(yes)
}
