package main

import (
	"fmt"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Auth/pkg/config"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Auth/pkg/db"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Auth/pkg/pb"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Auth/pkg/services"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Auth/pkg/utils"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(
		c.DB_USERNAME,
		c.DB_PASSWORD,
		c.DB_HOSTNAME,
		c.DB_PORT,
		c.DB_NAME,
	)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
