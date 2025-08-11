package main

import (
	"log"

	"github.com/AddonVbs/users-service/internal/database"
	transportgrpc "github.com/AddonVbs/users-service/internal/transport/gprc"
	"github.com/AddonVbs/users-service/internal/user"
)

func main() {
	// 1. Инициализация базы
	database.InitDB()

	// 2. Репозиторий
	repo := user.NewUserRepository(database.DB)

	// 3. Сервис
	svc := user.NewUserService(repo)

	// 4. Запуск gRPC
	if err := transportgrpc.RunGRPC(&svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
