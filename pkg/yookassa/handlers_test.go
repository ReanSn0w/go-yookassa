package yookassa_test

import (
	"net"
	"testing"

	"github.com/ReanSn0w/go-yookassa/pkg/yookassa"
)

func Test_CheckIPSuccess(t *testing.T) {
	successString := "77.75.153.10"
	ip := net.ParseIP(successString)

	err := yookassa.CheckIP(ip)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func Test_CheckIPFail(t *testing.T) {
	failureString := "77.75.152.10"
	ip := net.ParseIP(failureString)

	err := yookassa.CheckIP(ip)
	if err == nil {
		t.Log("check ip failed")
		t.FailNow()
	}
}

func Test_CheckIPError(t *testing.T) {
	err := yookassa.CheckIP(nil)
	if err == nil {
		t.Log("fail nil is not ip")
		t.Fail()
	}
}
