package test

import (
	"market/app/service/ali_sms"
	_ "market/bootstrap"
	"testing"
)

func TestSms(t *testing.T) {
	err := ali_sms.SendSms("", 9999)
	if err != nil {
		t.Fatal(err)
	}
}
