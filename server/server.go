package main

import (
	pb "E-Learning/proto"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var kumpulanSoal = map[string]map[string]map[string][]Soal{
	"Kelas 1": map[string]map[string][]Soal{},
	"Kelas 2": map[string]map[string][]Soal{},
	"Kelas 3": map[string]map[string][]Soal{},
	"Kelas 4": map[string]map[string][]Soal{},
	"Kelas 5": map[string]map[string][]Soal{},
	"Kelas 6": map[string]map[string][]Soal{},
}

//Soal dengan struct
type Soal struct {
	Pertanyaan string
	Jawaban    string
	Pilihan    []string
}

type server struct {
	pb.UnimplementedConnectServer
	pb.UnimplementedGuruServer
}

// type serverGuru struct {
// 	pb.UnimplementedGuruServer
// }

func certified() {
	f, err := os.Create("Certified.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"CERTIFIED",
		"TINGKAT, KELAS, MAPEL, TEMA",
		"NILAI",
		"NAMA"}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

func (s *server) Soal(ctx context.Context, ps *pb.PilihanSoal) (*pb.SoalList, error) {

	var result []Soal

	if v, found := kumpulanSoal[ps.GetKelas()][ps.GetMapel()][ps.GetTemaSoal()]; found {
		result = v
	}

	//fmt.Println(result)

	var result2 []*pb.Soal
	for i := 0; i < len(result); i++ {
		tempsoal := pb.Soal{Pertanyaan: result[i].Pertanyaan, Jawaban: result[i].Jawaban, Pilihan: result[i].Pilihan}
		result2 = append(result2, &tempsoal)
	}
	//fmt.Println(result2)

	result3 := pb.SoalList{SoalList: result2}

	return &result3, nil
}

func (s *server) Hasil(ctx context.Context, kj *pb.KumpulanJawaban) (*pb.Result, error) {
	var result []Soal
	var nilai int = 0
	var correct int = 0
	var notice string

	if v, found := kumpulanSoal[kj.GetKelas()][kj.GetMapel()][kj.GetTemaSoal()]; found {
		result = v
	}

	//fmt.Println(result)

	for i := 0; i < len(kj.GetJawabanList()); i++ {
		fmt.Println(result[i].Jawaban)
		fmt.Println(kj.GetJawabanList()[i])
		if kj.GetJawabanList()[i] == result[i].Jawaban {
			correct++
		}
	}

	nilai = (100 / len(kj.GetJawabanList())) * correct

	if nilai >= 70 {
		notice = "Nilai " + string(nilai) + "\nSELAMAT\nKamu berhasil lulus tema soal ini.\nSilahkan cek sertifikatmu ya."
		certified()
	} else {
		notice = "Nilai " + string(nilai) + "\nJANGAN PATAH SEMANGAT\nTetaplah berusaha."
	}
	fmt.Println(notice)

	notice2 := pb.Result{Result: notice}

	return &notice2, nil
}

func (s *server) KirimSoalKeServer(ctx context.Context, req *pb.PaketSoal) (*pb.ServerGuruReply, error) {
	// fmt.Println(req.GetTingkat())
	// fmt.Println(req.GetKelas())
	// fmt.Println(req.GetMatpel())
	// fmt.Println(req.GetTema())
	// fmt.Println(req.GetSoal()[0])

	var soal []Soal
	for i := 0; i < len(req.GetSoal()); i++ {
		tempsoal := Soal{Pertanyaan: req.GetSoal()[i].Pertanyaan, Jawaban: req.GetSoal()[i].Jawaban, Pilihan: req.GetSoal()[i].Pilihan}
		soal = append(soal, tempsoal)
	}

	//check if soal exists in tingkat-kelas-matpel-tema
	if _, found := kumpulanSoal[req.GetKelas()]; found {
		if _, found := kumpulanSoal[req.GetKelas()][req.GetMatpel()]; found {
			if _, found := kumpulanSoal[req.GetKelas()][req.GetMatpel()]; found {
				if _, found := kumpulanSoal[req.GetKelas()][req.GetMatpel()][req.GetTema()]; found {
					//beri nama lain
					newname := req.GetTema() + " (new)"
					//save
					saveSoal(req.GetTingkat(), req.GetKelas(), req.GetMatpel(), newname, soal)
					return &pb.ServerGuruReply{Message: ("Paket Soal dengan Tema ini sudah tersedia\nSoal disimpan dengan judul " + newname)}, nil
				}
			}
		}
	}

	saveSoal(req.GetTingkat(), req.GetKelas(), req.GetMatpel(), req.GetTema(), soal)
	return &pb.ServerGuruReply{Message: ("Paket Soal dengan Tema ini disimpan dengan judul " + req.GetTema())}, nil

}
func main() {
	loadSoal()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterConnectServer(s, &server{})
	pb.RegisterGuruServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// pb.RegisterGuruServer(s, &server{})
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}

func saveSoal(ti string, k string, m string, te string, s []Soal) {
	fmt.Println(ti, k, m, te, s)

	// kumpulanSoal[k][m] = map[string][]Soal{}
	kumpulanSoal[k][m][te] = s

	fmt.Println(kumpulanSoal)

}

func loadSoal() {

	//Kelas 1 - Bahasa Inggris - Bab 1
	tempsoal1 := Soal{"Bahasa Inggris \"Selamat pagi\" adalah ...", "Good morning", []string{"Good morning", "Good afternoon", "Good night", "Good day"}}
	tempsoal2 := Soal{"How are you?", "I am fine, thank you.", []string{"I am fine, thank you.", "Hello!", "Goodbye!", "How about you?"}}
	tempsoal3 := Soal{"What is your name?", "My name is Budi.", []string{"My name is Budi.", "I am fine, thank you.", "I am 6 years old.", "Goodbye!"}}
	tempsoal4 := Soal{"Bahasa Inggris \"Selamat malam\" adalah ...", "Good night", []string{"Good morning", "Good afternoon", "Good night", "Good day"}}
	tempsoal5 := Soal{"How old are you?", "I am 6 years old.", []string{"My name is Budi.", "I am fine, thank you.", "I am 6 years old.", "Goodbye!"}}

	arraySoal := make([]Soal, 0)
	arraySoal = append(arraySoal, tempsoal1)
	arraySoal = append(arraySoal, tempsoal2)
	arraySoal = append(arraySoal, tempsoal3)
	arraySoal = append(arraySoal, tempsoal4)
	arraySoal = append(arraySoal, tempsoal5)

	kumpulanSoal["Kelas 1"] = map[string]map[string][]Soal{}
	kumpulanSoal["Kelas 1"]["Bahasa Inggris"] = map[string][]Soal{}
	kumpulanSoal["Kelas 1"]["Bahasa Inggris"]["Bab 1"] = arraySoal

	//Kelas 1 - Bahasa Inggris - Bab 2
	// tempsoal1 = Soal{"What are you doing?", "I am studying English.", []string{"I am studying English.", "I eat bread for breakfast.", "I like my bicycle.", "I am a student."}}
	// tempsoal2 = Soal{"Mom ... cooking dinner right now.", "is", []string{"is", "am", "are", "was"}}
	// tempsoal3 = Soal{"My friends ... staying at home.", "are", []string{"are", "were", "is", "am"}}
	// tempsoal4 = Soal{"Dad is (write)... a letter for his friend.", "writing", []string{"writing", "writeing", "writying", "writiing"}}
	// tempsoal5 = Soal{"Budi and Bambang ... best friends", "are", []string{"am", "is", "was", "are"}}

	// arraySoal = make([]Soal, 0)
	// arraySoal = append(arraySoal, tempsoal1)
	// arraySoal = append(arraySoal, tempsoal2)
	// arraySoal = append(arraySoal, tempsoal3)
	// arraySoal = append(arraySoal, tempsoal4)
	// arraySoal = append(arraySoal, tempsoal5)

	// kumpulanSoal["Kelas 1"]["Bahasa Inggris"]["Bab 2"] = arraySoal

	// //Kelas 2 - Bahasa Inggris - Bab 1
	// tempsoal1 = Soal{"I like airplanes, I want to be a ... ", "Pilot", []string{"Pilot", "Nurse", "Teacher", "Youtuber"}}
	// tempsoal2 = Soal{"Mary likes animals, she wants to work in a ... ", "Zoo", []string{"Zoo", "School", "Hospital", "Office"}}
	// tempsoal3 = Soal{"Bob likes to draw, he wants to be a ... ", "Painter", []string{"Programmer", "Engineer", "Painter", "Musician"}}
	// tempsoal4 = Soal{"Gordon likes to cook, he wants to be a ...", "Chef", []string{"Chef", "Racer", "Athlete", "Programmer"}}
	// tempsoal5 = Soal{"Patrick likes to act, he wants to be an ...", "Actor", []string{"Actor", "Actress", "Athlete", "Attorney"}}

	// arraySoal = make([]Soal, 0)
	// arraySoal = append(arraySoal, tempsoal1)
	// arraySoal = append(arraySoal, tempsoal2)
	// arraySoal = append(arraySoal, tempsoal3)
	// arraySoal = append(arraySoal, tempsoal4)
	// arraySoal = append(arraySoal, tempsoal5)

	// kumpulanSoal["Kelas 2"] = map[string]map[string][]Soal{}
	// kumpulanSoal["Kelas 2"]["Bahasa Inggris"] = map[string][]Soal{}
	// kumpulanSoal["Kelas 2"]["Bahasa Inggris"]["Bab 1"] = arraySoal

	// //Kelas 2 - Bahasa Inggris - Bab 2
	// tempsoal1 = Soal{"Sam yawns all day, he is ...", "Sleepy", []string{"Sleepy", "Hungry", "Angry", "Sad"}}
	// tempsoal2 = Soal{"I skipped the breakfast this morning, now I am very ...", "Hungry", []string{"Hungry", "Sleepy", "Bored", "Tired"}}
	// tempsoal3 = Soal{"Mom is ... at me because I lost her tupperware", "Angry", []string{"Angry", "Happy", "Sad", "Bored"}}
	// tempsoal4 = Soal{"Alice forget her water bottle, now she is ...", "Thirsty", []string{"Thirsty", "Hungry", "Sleepy", "Bored"}}
	// tempsoal5 = Soal{"Psyduck uses Confuse, now I am ...", "Confused", []string{"Confused", "Sleepy", "Hungry", "Angry"}}

	// arraySoal = make([]Soal, 0)
	// arraySoal = append(arraySoal, tempsoal1)
	// arraySoal = append(arraySoal, tempsoal2)
	// arraySoal = append(arraySoal, tempsoal3)
	// arraySoal = append(arraySoal, tempsoal4)
	// arraySoal = append(arraySoal, tempsoal5)

	// kumpulanSoal["Kelas 2"]["Bahasa Inggris"]["Bab 2"] = arraySoal

	fmt.Println(kumpulanSoal)
}
