package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"sync/atomic"
	"time"

	pb "github.com/Lirikl/mafia/pkg/proto/mafia"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MafiaGame struct {
	Roles      map[string]string
	Status     map[string]string //mafia citizen ghost
	Mutex      sync.Mutex
	Day        bool
	MafiaVote  map[string]int
	SherifVote map[string]int
	Names      []string

	MafiaCount  int
	SherifCount int
	AllCount    int
	DayVote     map[string]int
	Input       chan pb.GameCommand
	Output      map[string]chan pb.GameEvent
}

type server struct {
	pb.UnimplementedMafiaServer
	UsersMutex   sync.Mutex
	Names        map[string]bool
	Users        []string
	Sessions     []MafiaGame
	Channel      map[string]chan pb.ConnectionUpdate
	WaitingCount uint32
}

type ConnectMsg struct {
	id int
}

var civ_cnt int = 2
var maf_cnt int = 1
var sher_cnt int = 1
var all_cnt = 4

func get_keys(m *map[string]bool) []string {
	keys := make([]string, 0)
	for k := range *m {
		keys = append(keys, k)
	}
	return keys
}

func max(numbers *map[string]int) string {
	var maxNumber int
	maxKey := make([]string, 0)
	var mK string
	for mK, maxNumber = range *numbers {
		maxKey = append(maxKey, mK)
		break
	}
	for k, v := range *numbers {
		if v > maxNumber {
			maxNumber = v
			maxKey = make([]string, 0)
		}
		if v == maxNumber {
			maxKey = append(maxKey, k)
		}
	}

	return maxKey[rand.Int()%len(maxKey)]
}
func night_round(game *MafiaGame) {
	game.MafiaVote = make(map[string]int)
	game.SherifVote = make(map[string]int)
	cnt := 0
	fmt.Println("start night")
	for cnt != game.MafiaCount+game.SherifCount {
		msg := <-game.Input
		fmt.Println("step")

		fmt.Println(msg)
		if game.Status[msg.Name] == "ghost" {
			continue
		}
		if game.Roles[msg.Name] == "Mafia" {
			game.MafiaVote[msg.Vote] += 1
			cnt += 1
		}
		if game.Roles[msg.Name] == "Sherif" {
			game.SherifVote[msg.Vote] += 1
			cnt += 1
		}
		fmt.Println(cnt)

	}
	if game.SherifCount > 0 {
		name := max(&game.SherifVote)
		role_ans := (game.Roles[name] == "Mafia")
		for player, role := range game.Roles {
			if role == "Sherif" && game.Status[player] == "alive" {
				game.Output[player] <- pb.GameEvent{Suspect: name, CheckResult: role_ans}
			}
		}
	}

	victim := max(&game.MafiaVote)
	game.Status[victim] = "ghost"
	if game.Roles[victim] == "Mafia" {
		game.MafiaCount -= 1
	}
	if game.Roles[victim] == "Sherif" {
		game.SherifCount -= 1
	}
	game.AllCount -= 1

	res := int64(0)
	if MafiaWon(game) {
		res = 1
	}
	if CityWon(game) {
		res = 2
	}
	for player, _ := range game.Roles {
		game.Output[player] <- pb.GameEvent{Victim: victim, Winner: res}
	}

}

func day_round(game *MafiaGame) {

	game.DayVote = make(map[string]int)
	cnt := 0

	fmt.Println("start day")
	for cnt != game.AllCount {
		msg := <-game.Input
		if game.Status[msg.Name] == "ghost" {
			continue
		}
		game.DayVote[msg.Vote] += 1
		cnt += 1
	}
	victim := max(&game.DayVote)
	if game.Roles[victim] == "Mafia" {
		game.MafiaCount -= 1
	}
	if game.Roles[victim] == "Sherif" {
		game.SherifCount -= 1
	}

	res := int64(0)
	if MafiaWon(game) {
		res = 1
	}
	if CityWon(game) {
		res = 2
	}
	for player, _ := range game.Roles {
		game.Output[player] <- pb.GameEvent{Victim: victim, Winner: res}
	}
}

func MafiaWon(game *MafiaGame) bool {
	return game.AllCount <= 2*game.MafiaCount
}

func CityWon(game *MafiaGame) bool {
	return game.MafiaCount == 0
}

func game_runner(game *MafiaGame) {
	for {
		night_round(game)
		day_round(game)
	}
}

func (s *server) GameSession(stream pb.Mafia_GameSessionServer) error {
	con, _ := stream.Recv()
	s.UsersMutex.Lock()
	var game *MafiaGame = &s.Sessions[con.SessionID]
	s.UsersMutex.Unlock()
	var last_check pb.GameEvent
	for {
		//night
		if game.Roles[con.Name] != "Civ" && game.Status[con.Name] == "alive" {
			msg, _ := stream.Recv()
			game.Input <- *msg
			if game.Roles[con.Name] == "Sherif" {
				last_check = <-game.Output[con.Name]
				fmt.Println("send check")
				stream.Send(&last_check)
			}
		}
		fmt.Println("get res", con.Name)
		round_res := <-game.Output[con.Name]

		fmt.Println("send res", con.Name)
		stream.Send(&round_res)
		if round_res.Winner > 0 {
			return nil
		}
		//day
		fmt.Println(game.Status[con.Name], con.Name)
		if game.Status[con.Name] == "alive" {
			for {
				msg, _ := stream.Recv()
				game.Input <- *msg
				if msg.Type == "vote" {
					break
				}
			}
		}
		round_res = <-game.Output[con.Name]
		stream.Send(&round_res)
		if round_res.Winner > 0 {
			return nil
		}
	}
}

func (s *server) init_game() MafiaGame {
	atomic.AddUint32(&s.WaitingCount, -uint32(all_cnt))
	mf := MafiaGame{AllCount: all_cnt, MafiaCount: maf_cnt, SherifCount: sher_cnt}
	mf.SherifVote = make(map[string]int)
	mf.MafiaVote = make(map[string]int)
	mf.DayVote = make(map[string]int)
	mf.Input = make(chan pb.GameCommand, 15)
	mf.Output = make(map[string]chan pb.GameEvent)
	mf.Roles = make(map[string]string)
	mf.Status = make(map[string]string)

	s.UsersMutex.Lock()
	names := s.Users[:all_cnt]
	mf.Names = names
	for _, v := range mf.Names {
		delete(s.Names, v)
	}
	s.Users = s.Users[all_cnt:]
	s.UsersMutex.Unlock()
	roles := make([]string, 0)
	for i := 0; i < maf_cnt; i++ {
		roles = append(roles, "Mafia")
	}
	for i := 0; i < sher_cnt; i++ {
		roles = append(roles, "Sherif")
	}
	for i := 0; i < all_cnt-(sher_cnt+maf_cnt); i++ {
		roles = append(roles, "Civ")
	}
	rand.Shuffle(len(roles), func(i, j int) {
		roles[i], roles[j] = roles[j], roles[i]
	})
	for i, v := range names {
		mf.Output[v] = make(chan pb.GameEvent, 15)
		mf.Roles[v] = roles[i]
		mf.Status[v] = "alive"
	}
	return mf
	//for i := 0; i < all_cnt; i++ {
	//	mf.Output[i] = make(chan pb.GameEvent, 15)
	//}
}

func (s *server) DeleteName(name string) {
	for i, v := range s.Users {
		if v == name {
			s.Users[i], s.Users[len(s.Users)-1] = s.Users[len(s.Users)-1], s.Users[i]
			s.Users = s.Users[:len(s.Users)-1]
			break
		}
	}
}

func (s *server) Connect(req *pb.ConnectionRequest, stream pb.Mafia_ConnectServer) error {
	//req, err := stream.Recv()
	fmt.Println(string(req.Name))
	timer := time.Tick(time.Second)
	s.UsersMutex.Lock()
	if s.Names == nil {
		s.Names = make(map[string]bool)
	}

	if s.Channel == nil {
		s.Channel = make(map[string]chan pb.ConnectionUpdate)
	}
	if s.Names[req.Name] {
		s.UsersMutex.Unlock()
		return status.Error(codes.InvalidArgument, "Name already in use")
	}
	channel := make(chan pb.ConnectionUpdate, 10)

	s.Channel[req.Name] = channel
	s.Names[req.Name] = true
	s.Users = append(s.Users, req.Name)
	s.UsersMutex.Unlock()

	atomic.AddUint32(&s.WaitingCount, 1)
	if atomic.LoadUint32(&s.WaitingCount) == uint32(all_cnt) {
		mf := s.init_game()
		sz := len(s.Sessions)
		s.Sessions = append(s.Sessions, mf)
		for k, v := range mf.Roles {
			s.Channel[k] <- pb.ConnectionUpdate{SessionID: int64(sz), Role: v}
		}
		go game_runner(&mf)
	}
	for {

		select {
		case <-timer:
			//<-timer
			s.UsersMutex.Lock()
			users := get_keys(&s.Names)
			s.UsersMutex.Unlock()
			if err := stream.Send(&pb.ConnectionUpdate{Connect: pb.ConnectionStatus_None, Users: users}); err != nil {
				s.UsersMutex.Lock()
				s.DeleteName(req.Name)
				delete(s.Names, req.Name)
				s.UsersMutex.Unlock()
				return err
			}
		case msg := <-channel:
			fmt.Println("start ses")
			stream.Send(&pb.ConnectionUpdate{Connect: pb.ConnectionStatus_Start, Users: s.Sessions[msg.SessionID].Names, SessionID: msg.SessionID, Role: msg.Role})
			return nil
			//s.GameSession(s)
		}

	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterMafiaServer(srv, &server{})
	log.Fatalln(srv.Serve(lis))
}

func stringReverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
