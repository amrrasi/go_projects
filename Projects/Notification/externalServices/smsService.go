package externalServices

import (
	"fmt"
)

type SmsService struct{}

func (s *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("This %s\n SMS has been Sent to! %s\n", message, receiver)
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
