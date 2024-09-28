// internal/infrastructure/queue/redis.go
package queue

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisQueue struct {
    client *redis.Client
}

func NewRedisQueue(client *redis.Client) *RedisQueue {
    return &RedisQueue{client: client}
}

func (r *RedisQueue) StartWorker(ctx context.Context, workerCount int, processFunc func(string) error) {
    for i := 0; i < workerCount; i++ {
        go func(workerID int) {
            for {
                select {
                case <-ctx.Done():
                    return
                default:
                    // Используем BLPOP для ожидания сообщений
                    msg, err := r.client.BLPop(ctx, 5*time.Second, "sms_queue").Result()
                    if err == redis.Nil {
                        log.Println("Очередь пуста, ждем новых сообщений...")
                    } else if err != nil {
                        log.Printf("Ошибка при получении сообщения: %v\n", err)
                    } else {
                        // Обработка сообщения
                        if err := processFunc(msg[1]); err != nil {
                            log.Printf("Ошибка обработки сообщения: %v\n", err)
                            // Логика повторной обработки может быть добавлена здесь
                        }
                    }
                }
            }
        }(i)
    }
}
