package main

import (
	"log"
	"sms-gateway/internal/infrastructure/queue"
	"sms-gateway/internal/interfaces/smsru"
	"sms-gateway/internal/usecases"
	"time"
)

func main() {
    // Инициализация очереди
    redisQueue := queue.NewRedisQueue("localhost:6379", "", 0)

    // Инициализация клиента sms.ru
    smsClient := smsru.NewSmsRuClient("your-api-id-here")

    // Инициализация usecase
    smsUseCase := usecases.NewSendSMSUseCase(redisQueue, smsClient)

    // Постоянная обработка очереди
    for {
        err := smsUseCase.ProcessSMSQueue()
        if err != nil {
            log.Printf("Error processing SMS queue: %v", err)
        }
        time.Sleep(5 * time.Second) // Задержка между итерациями
    }
}
