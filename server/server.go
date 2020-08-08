package main

import (
	pb "E-Learning/proto"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var kumpulansoal = map[string]map[string]map[string][]Soal{
	"Kelas 1": map[string]map[string][]Soal{},
	"Kelas 2": map[string]map[string][]Soal{},
}

//Soal dengan struct
type Soal struct {
	pertanyaan string
	jawaban    string
	pilihan    [4]string
}

type server struct {
	pb.UnimplementedConnectServer
}

func (s *server) ConnectToServer(ctx context.Context, void *empty.Empty) (*pb.ServerReply, error) {
	return &pb.ServerReply{Message: "Client is connected to Server"}, nil
}
func main() {
	loadSoal()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterConnectServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loadSoal() {

	//Kelas 1 - Bahasa Inggris - Bab 1
	tempsoal1 := Soal{"Bahasa Inggris \"Selamat pagi\" adalah ...", "Good morning", [4]string{"Good morning", "Good afternoon", "Good night", "Good day"}}
	tempsoal2 := Soal{"How are you?", "I am fine, thank you.", [4]string{"I am fine, thank you.", "Hello!", "Goodbye!", "How about you?"}}
	tempsoal3 := Soal{"What is your name?", "My name is Budi.", [4]string{"My name is Budi.", "I am fine, thank you.", "I am 6 years old.", "Goodbye!"}}
	tempsoal4 := Soal{"Bahasa Inggris \"Selamat malam\" adalah ...", "Good night", [4]string{"Good morning", "Good afternoon", "Good night", "Good day"}}
	tempsoal5 := Soal{"How old are you?", "I am 6 years old.", [4]string{"My name is Budi.", "I am fine, thank you.", "I am 6 years old.", "Goodbye!"}}

	arraySoal := make([]Soal, 0)
	arraySoal = append(arraySoal, tempsoal1)
	arraySoal = append(arraySoal, tempsoal2)
	arraySoal = append(arraySoal, tempsoal3)
	arraySoal = append(arraySoal, tempsoal4)
	arraySoal = append(arraySoal, tempsoal5)

	kumpulansoal["Kelas 1"] = map[string]map[string][]Soal{}
	kumpulansoal["Kelas 1"]["Bahasa Inggris"] = map[string][]Soal{}
	kumpulansoal["Kelas 1"]["Bahasa Inggris"]["Bab 1"] = arraySoal

	//Kelas 1 - Bahasa Inggris - Bab 2
	tempsoal1 = Soal{"What are you doing?", "I am studying English.", [4]string{"I am studying English.", "I eat bread for breakfast.", "I like my bicycle.", "I am a student."}}
	tempsoal2 = Soal{"Mom ... cooking dinner right now.", "is", [4]string{"is", "am", "are", "was"}}
	tempsoal3 = Soal{"My friends ... staying at home.", "are", [4]string{"are", "were", "is", "am"}}
	tempsoal4 = Soal{"Dad is (write)... a letter for his friend.", "writing", [4]string{"writing", "writeing", "writying", "writiing"}}
	tempsoal5 = Soal{"Budi and Bambang ... best friends", "are", [4]string{"am", "is", "was", "are"}}

	arraySoal = make([]Soal, 0)
	arraySoal = append(arraySoal, tempsoal1)
	arraySoal = append(arraySoal, tempsoal2)
	arraySoal = append(arraySoal, tempsoal3)
	arraySoal = append(arraySoal, tempsoal4)
	arraySoal = append(arraySoal, tempsoal5)

	kumpulansoal["Kelas 1"]["Bahasa Inggris"]["Bab 2"] = arraySoal

	//Kelas 2 - Bahasa Inggris - Bab 1
	tempsoal1 = Soal{"I like airplanes, I want to be a ... ", "Pilot", [4]string{"Pilot", "Nurse", "Teacher", "Youtuber"}}
	tempsoal2 = Soal{"Mary likes animals, she wants to work in a ... ", "Zoo", [4]string{"Zoo", "School", "Hospital", "Office"}}
	tempsoal3 = Soal{"Bob likes to draw, he wants to be a ... ", "Painter", [4]string{"Programmer", "Engineer", "Painter", "Musician"}}
	tempsoal4 = Soal{"Gordon likes to cook, he wants to be a ...", "Chef", [4]string{"Chef", "Racer", "Athlete", "Programmer"}}
	tempsoal5 = Soal{"Patrick likes to act, he wants to be an ...", "Actor", [4]string{"Actor", "Actress", "Athlete", "Attorney"}}

	arraySoal = make([]Soal, 0)
	arraySoal = append(arraySoal, tempsoal1)
	arraySoal = append(arraySoal, tempsoal2)
	arraySoal = append(arraySoal, tempsoal3)
	arraySoal = append(arraySoal, tempsoal4)
	arraySoal = append(arraySoal, tempsoal5)

	kumpulansoal["Kelas 2"] = map[string]map[string][]Soal{}
	kumpulansoal["Kelas 2"]["Bahasa Inggris"] = map[string][]Soal{}
	kumpulansoal["Kelas 2"]["Bahasa Inggris"]["Bab 1"] = arraySoal

	//Kelas 2 - Bahasa Inggris - Bab 2
	tempsoal1 = Soal{"Sam yawns all day, he is ...", "Sleepy", [4]string{"Sleepy", "Hungry", "Angry", "Sad"}}
	tempsoal2 = Soal{"I skipped the breakfast this morning, now I am very ...", "Hungry", [4]string{"Hungry", "Sleepy", "Bored", "Tired"}}
	tempsoal3 = Soal{"Mom is ... at me because I lost her tupperware", "Angry", [4]string{"Angry", "Happy", "Sad", "Bored"}}
	tempsoal4 = Soal{"Alice forget her water bottle, now she is ...", "Thirsty", [4]string{"Thirsty", "Hungry", "Sleepy", "Bored"}}
	tempsoal5 = Soal{"Psyduck uses Confuse, now I am ...", "Confused", [4]string{"Confused", "Sleepy", "Hungry", "Angry"}}

	arraySoal = make([]Soal, 0)
	arraySoal = append(arraySoal, tempsoal1)
	arraySoal = append(arraySoal, tempsoal2)
	arraySoal = append(arraySoal, tempsoal3)
	arraySoal = append(arraySoal, tempsoal4)
	arraySoal = append(arraySoal, tempsoal5)

	kumpulansoal["Kelas 2"]["Bahasa Inggris"]["Bab 2"] = arraySoal

	fmt.Println(kumpulansoal)
}
