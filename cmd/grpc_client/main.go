package main

import (
	"context"
	"log"
	"time"

	"github.com/dmtrybogdanov/chat-server/pkg/chat_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50052"

var usernames = []string{"mollit, sint, incididunt, nulla"}


func main()  {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed connect to server: %v", err)
	}

	defer conn.Close()

	c := chat_v1.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Create(ctx, &chat_v1.CreateRequest{Usernames: usernames})
	if err != nil {
		log.Fatalf("Failed create usernames: %v", err)
	}

	log.Printf(color.RedString("chat info: \n"), color.GreenString("%+v", r.Id))
}
