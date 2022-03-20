package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	pb "github.com/Lirikl/mafia/pkg/proto/mafia"
	"google.golang.org/grpc"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var chat_chan chan string = make(chan string, 100)
var command_chan chan string = make(chan string, 100)

func reading() {
	for {
		str, _ := reader.ReadString('\n')
		str = str[:len(str)-1]
		fmt.Println("read ", str)
		if str[0] == '/' {
			command_chan <- str
		} else {
			chat_chan <- str
		}
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

func print_status(m *map[string]bool) {
	alive := make([]string, 0)
	dead := make([]string, 0)
	for k, v := range *m {
		if v {
			alive = append(alive, k)
		} else {
			dead = append(dead, k)
		}
	}
	fmt.Println("alive players: ", alive)
	fmt.Println("dead players: ", dead)
}

func RunCiv(game pb.Mafia_GameSessionClient, id int64, name string, alive map[string]bool) {
	for {

		fmt.Println("night falls on the city")
		//print_status(&alive)
		msg, _ := game.Recv()

		fmt.Println("city awakens")
		fmt.Println(msg.Victim, " was killed by mafia ")
		if msg.Victim == name {
			fmt.Println(msg.Victim, "You are dead")
		}
		alive[msg.Victim] = false
		print_status(&alive)
		if CheckEnd(msg) {
			break
		}

		var text []string
		if alive[name] {
			for {
				text = strings.Fields(<-command_chan)
				if len(text) != 2 || text[0] != "/vote" || !alive[text[1]] {
					fmt.Println("Wrong command")
					continue
				}
				fmt.Println("Vote accepted")
				break
			}
			game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id, Name: name})
		}
		msg, _ = game.Recv()
		fmt.Println(msg.Victim, " was killed by vote ")
		if msg.Victim == name {
			fmt.Println(msg.Victim, "You are dead")

		}
		alive[msg.Victim] = false
		if CheckEnd(msg) {
			break
		}
	}
}

func RunMafia(game pb.Mafia_GameSessionClient, id int64, name string, alive map[string]bool) {
	for {
		fmt.Println("night falls on the city")
		print_status(&alive)
		var text []string
		if alive[name] {
			for {
				text = strings.Fields(<-command_chan)
				if len(text) != 2 || text[0] != "/vote" || !alive[text[1]] {
					fmt.Println("Wrong command")
					continue
				}
				fmt.Println("Vote accepted")
				break
			}
			game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id, Name: name})
		}
		msg, _ := game.Recv()

		fmt.Println("city awakens")
		fmt.Println(msg.Victim, " was killed by mafia ")
		if msg.Victim == name {
			fmt.Println(msg.Victim, "You are dead")

		}
		alive[msg.Victim] = false

		print_status(&alive)
		if CheckEnd(msg) {
			break
		}
		if alive[name] {
			for {
				text = strings.Fields(<-command_chan)
				if len(text) != 2 || text[0] != "/vote" || !alive[text[1]] {
					fmt.Println("Wrong command")
					continue
				}
				fmt.Println("Vote accepted")
				break
			}
			game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id, Name: name})
		}
		msg, _ = game.Recv()
		fmt.Println(msg.Victim, " was killed by vote ")
		if msg.Victim == name {
			fmt.Println(msg.Victim, "You are dead")

		}
		alive[msg.Victim] = false
		if CheckEnd(msg) {
			break
		}
	}
}

func RunSherif(game pb.Mafia_GameSessionClient, id int64, name string, alive map[string]bool) {
	//supect := ""
	is_mafia := false

	for {
		fmt.Println("night falls on the city")
		print_status(&alive)
		var text []string
		//fmt.Println("start")
		if alive[name] {
			for {
				text = strings.Fields(<-command_chan)
				//fmt.Println(len(text), text)
				if len(text) != 2 || text[0] != "/vote" || !alive[text[1]] {
					fmt.Println("Wrong command")
					continue
				}
				fmt.Println("Vote accepted")
				break
			}
			game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id, Name: name})

			msg, _ := game.Recv()
			if msg.CheckResult {
				fmt.Println(msg.Suspect, "is mafia")
			} else {
				fmt.Println(msg.Suspect, "is not mafia")
			}
		}
		msg, _ := game.Recv()
		fmt.Println(msg.Victim, " was killed by mafia ")
		if msg.Victim == name {
			fmt.Println(msg.Victim, "You are dead")

		}
		alive[msg.Victim] = false
		print_status(&alive)
		if CheckEnd(msg) {
			break
		}

		fmt.Println("city awakens")
		if alive[name] {
			for {
				text = strings.Fields(<-command_chan)
				if len(text) == 1 && text[0] == "/reveal" && is_mafia {
					game.Send(&pb.GameCommand{Type: "reveal", SessionID: id, Name: name})
					continue
				}
				if len(text) == 2 && text[0] == "/vote" && alive[text[1]] {
					fmt.Println("Vote accepted")
					break
				}
				fmt.Println("Wrong command")
			}
			game.Send(&pb.GameCommand{Type: "vote", Vote: text[1], SessionID: id, Name: name})
		}
		msg, err := game.Recv()
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(msg)
		fmt.Println(msg.Victim, " was killed by vote ")
		if msg.Victim == name {
			fmt.Println(msg.Victim, "You are dead")
		}
		alive[msg.Victim] = false
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
	go reading()
	for {
		fmt.Println("Register name with /name <name>")
		ctx := context.Background()
		name := <-command_chan
		txt := strings.Fields(name)
		if len(txt) != 2 || txt[0] != "/name" {
			fmt.Println("Wrong command")
			continue
		}
		name = txt[1]
		stream, err := client.Connect(ctx, &pb.ConnectionRequest{Name: name, Connect: pb.ConnectionStatus_Connect})
		if err != nil {
			fmt.Println("Name already taken")
			continue
		}
		_, err = stream.Recv()
		if err != nil {
			fmt.Println("Name already taken")
			continue
		}
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {

			for {
				in, _ := stream.Recv()
				if in.Connect == pb.ConnectionStatus_Start {
					fmt.Println("Start session as ", in.Role, "!!!")
					role := in.Role
					game, _ := client.GameSession(ctx)
					game.Send(&pb.GameCommand{SessionID: in.SessionID, Name: name})
					m := make(map[string]bool)
					for _, n := range in.Users {
						m[n] = true
					}
					fmt.Println("players: ", in.Users)
					switch role {
					case "Civ":
						RunCiv(game, in.SessionID, name, m)
					case "Mafia":
						RunMafia(game, in.SessionID, name, m)
					case "Sherif":
						RunSherif(game, in.SessionID, name, m)
					}
					wg.Done()
					return
				}
			}
		}()
		wg.Wait()
	}
}
