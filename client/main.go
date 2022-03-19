package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/Lirikl/mafia/pkg/proto/mafia"
	"google.golang.org/grpc"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var chat_chan chan string = make(chan string, 100)
var command_chan chan string = make(chan string, 100)

func reading() {
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	if str[0] == '/' {
		command_chan <- str
	} else {
		chat_chan <- str
	}
}

func CheckEnd(msg *pb.GameEvent) bool {
	if msg.Winner == 1 {
		fmt.Println("Mafia Won")
	}
	if msg.Winner == 2 {
		fmt.Println("Civilians Won")
	}
	return msg.Winner > 0
}

func RunCiv(game pb.Mafia_GameSessionClient, id int64) {
	for {
		msg, _ := game.Recv()
		if CheckEnd(msg) {
			break
		}
		var text []string
		for {
			text = strings.Fields(<-command_chan)
			if len(text) != 2 || text[0] != "/vote" {
				continue
			}
			break
		}
		game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id})
		msg, _ = game.Recv()
		if CheckEnd(msg) {
			break
		}
	}
}

func RunMafia(game pb.Mafia_GameSessionClient, id int64) {
	for {
		var text []string
		for {
			text = strings.Fields(<-command_chan)
			if len(text) != 2 || text[0] != "/vote" {
				continue
			}
			break
		}
		game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id})
		msg, _ := game.Recv()
		if CheckEnd(msg) {
			break
		}

		for {
			text = strings.Fields(<-command_chan)
			if len(text) != 2 || text[0] != "/vote" {
				continue
			}
			break
		}
		game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id})
		msg, _ = game.Recv()
		if CheckEnd(msg) {
			break
		}
	}
}

func RunSherif(game pb.Mafia_GameSessionClient, id int64) {
	//supect := ""
	is_mafia := false
	for {
		var text []string
		for {
			text = strings.Fields(<-command_chan)
			if len(text) != 2 || text[0] != "/vote" {
				continue
			}
			break
		}
		game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id})
		msg, _ := game.Recv()
		if CheckEnd(msg) {
			break
		}

		for {
			text = strings.Fields(<-command_chan)
			if len(text) == 1 && text[0] == "/reveal" && is_mafia {
				game.Send(&pb.GameCommand{Type: "reveal", SessionID: id})
				continue
			}
			if len(text) == 2 && text[0] == "/vote" {
				break
			}
		}
		game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id})
		msg, _ = game.Recv()
		if CheckEnd(msg) {
			break
		}
	}
}

func main() {
	log.Println("Client running ...")
	conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := pb.NewMafiaClient(conn)
	request := &pb.Request{Message: "This is a test"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Do(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}
	go reading()
	ctx = context.Background()
	name := <-command_chan
	name = strings.Fields(name)[1]
	stream, err := client.Connect(ctx, &pb.ConnectionRequest{Name: name, Connect: pb.ConnectionStatus_Connect})
	go func() {
		for {
			in, _ := stream.Recv()
			if in.Connect == pb.ConnectionStatus_Start {
				role := in.Role
				game, _ := client.GameSession(ctx)
				game.Send(&pb.GameCommand{SessionID: in.SessionID})

				switch role {
				case "Civ":
					RunCiv(game, in.SessionID)
				case "Mafia":
					RunMafia(game, in.SessionID)
				case "Sherif":
					RunMafia(game, in.SessionID)
				}
			}
		}
	}()
	log.Println("Response:", response.GetMessage())
}
