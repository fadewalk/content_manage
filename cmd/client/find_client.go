package main

import (
	"content_manage/api/operate"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
)

func main() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := operate.NewAppClient(conn)
	reply, err := client.FindContent(context.Background(), &operate.FindContentReq{
		Page:     2,
		PageSize: 2,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] find %+v\n", reply)
}
