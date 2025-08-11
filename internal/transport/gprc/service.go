package transportgrpc

import (
	"fmt"
	"log"
	"net"

	userpb "github.com/AddonVbs/project-protos/proto/user"
	"github.com/AddonVbs/users-service/internal/user"

	"google.golang.org/grpc"
)

func RunGRPC(svc *user.UserService1) error {
	// 1. Открываем порт
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("не удалось открыть порт :50051: %w", err)
	}

	// 2. Создаём gRPC сервер
	grpcSrv := grpc.NewServer()

	// 3. Регистрируем наш сервис
	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(*svc))

	log.Println("gRPC сервер запущен на :50051")

	// 4. Запускаем сервер
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("ошибка запуска gRPC сервера: %w", err)
	}

	return nil
}
