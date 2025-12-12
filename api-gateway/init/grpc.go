package init

import (
	"class/passenger/api-gateway/global"
	"flag"
	"log"

	carrpc "class/passenger/api-gateway/proto/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitCarRpc() {
	flag.Parse()
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	global.CarRpcClient = carrpc.NewCartClient(conn)
}
