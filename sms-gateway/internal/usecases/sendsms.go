package usecases

import (
	"log"
	"sms-gateway/internal/infrastructure/queue"
	"sms-gateway/internal/interfaces/smsru"
)

type SendSMSUseCase struct {
    queue   *queue.RedisQueue
    client  *smsru.SmsRuClient
}

func NewSendSMSUseCase(queue *queue.RedisQueue, client *smsru.SmsRuClient) *SendSMSUseCase {
    return &SendSMSUseCase{queue: queue, client: client}
}

func (s *SendSMSUseCase) ProcessSMSQueue() error {
    phoneNumber, message, err := s.queue.DequeueSMS()
    if err != nil {
        return err
    }

    err = s.client.SendSMS(phoneNumber, message)
    if err != nil {
        log.Printf("Error sending SMS: %v", err)
        // Возможно, вернуть сообщение в очередь
        s.queue.EnqueueSMS(phoneNumber, message)
        return err
    }

    log.Printf("SMS sent successfully to %s", phoneNumber)
    return nil
}
