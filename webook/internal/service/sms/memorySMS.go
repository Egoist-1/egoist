package sms

import (
	"context"
	"fmt"
)

type MemorySMS struct {
}

func (m MemorySMS) Send(ctx context.Context, phone string, biz string, code string) error {
	fmt.Println(phone, biz, code)
	return nil
}

func NewMemorySMS() SMS {
	return &MemorySMS{}
}
