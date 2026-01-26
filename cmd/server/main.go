package main

import (
	"log"

	"github.com/IAMBURNINGMAN/user-service/internal/database"
	transportgrpc "github.com/IAMBURNINGMAN/user-service/internal/transport/grpc"
	"github.com/IAMBURNINGMAN/user-service/internal/user"
)

func main() {
	// 1. Инициализация БД
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// 2. Автомиграция
	if err := db.AutoMigrate(&user.UserStruct{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	// 3. Создание слоёв
	repo := user.NewUserRepository(db)
	svc := user.NewUserService(repo)

	// 4. Запуск gRPC сервера
	log.Println("Запуск gRPC сервера на :50051...")
	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
