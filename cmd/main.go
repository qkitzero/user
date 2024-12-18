package main

import (
	"fmt"
	"log"
	"net"

	application_user "user/internal/application/user"
	"user/internal/infrastructure/db"
	infrastructure_user "user/internal/infrastructure/persistence/user"
	interface_user "user/internal/interface/grpc/user"
	"user/pb"

	"google.golang.org/grpc"
)

func main() {
	db, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	port := 50051
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	userRepository := infrastructure_user.NewUserRepository(db)

	userService := application_user.NewUserService(userRepository)

	userHandler := interface_user.NewUserHandler(*userService)

	pb.RegisterUserServiceServer(server, userHandler)

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
