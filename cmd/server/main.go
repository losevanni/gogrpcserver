package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	// _ "github.com/lib/pq"
	db "goserver/internal/dbpool"
	user "goserver/internal/user/service"
	pbuser "goserver/protos/pbuser"
)

const portNumber = "50051"


func main() {
	// DB
	connStr:="postgresql://gsdbuser:kms1030@localhost:5432/goserverdb"
	// db, err := sql.Open("postgres", "user=userid password=pass dbname=mydb sslmode=disable")
	dbpool,err := db.ConnectDB(connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbpool.Close()

	// 서버
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//grpc서버에 등록하기.
	grpcServer := grpc.NewServer()
	// bookpd.RegisterBookServer(grpcServer, &bookServer{})
	pbuser.RegisterServiceJoinServer(grpcServer,user.NewUserJoin(dbpool))
	pbuser.RegisterServiceLoginServer(grpcServer,user.NewUserLogin(dbpool))
	pbuser.RegisterServiceInfoServer(grpcServer,user.NewUserInfo(dbpool))
	// 동적으로 서버 정보 얻을수있음 
	// reflection.Register(grpcServer)
	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
// grpcurl localhost:50051 list
