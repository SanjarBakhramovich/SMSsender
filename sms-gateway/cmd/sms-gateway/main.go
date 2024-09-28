package main

import (
	"log"
	"sms-gateway/internal/infrastructure/api"
	"sms-gateway/internal/infrastructure/db"
	"sms-gateway/internal/infrastructure/queue"
	"sms-gateway/internal/interfaces/web"
)

func main() {
    // Инициализация базы данных
    dbConn, err := db.ConnectPostgres()
    if err != nil {
        log.Fatalf("Ошибка подключения к базе данных: %v", err)
    }

    // Инициализация очереди
    queueConn, err := queue.ConnectRedis()
    if err != nil {
        log.Fatalf("Ошибка подключения к Redis: %v", err)
    }

    // Запуск API-сервера
    err = api.StartServer(dbConn, queueConn)
    if err != nil {
        log.Fatalf("Ошибка запуска API-сервера: %v", err)
    }

    // Запуск веб-интерфейса
    webServer := web.NewServer()
    if err := webServer.Run(); err != nil {
        log.Fatalf("Ошибка запуска веб-сервера: %v", err)
    }
}
