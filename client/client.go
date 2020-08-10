package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "E-Learning/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

//Soal dengan struct
type Soal struct {
	Pertanyaan string
	Jawaban    string
	Pilihan    []string
}

//PaketSoal dengan struct
type PaketSoal struct {
	Tingkat string
	Kelas   string
	Matpel  string
	Tema    string
	Soal    []Soal
}

var kumpulanJawaban []string

func main() {
	menu()
}

//=================================METHOD UNTUK CONNECT GRPC==================

//ConnectServer dial ke service siswa
func ConnectServer() pb.ConnectClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pb.NewConnectClient(conn)
}

//ConnectServerGuru dial ke service guru
func ConnectServerGuru() pb.GuruClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	return pb.NewGuruClient(conn)
}

//============================================================================

func menu() {
	var menu int
	fmt.Println("WELCOME TO RUANG GULU")
	fmt.Println("Pilih service apa yang anda ingin gunakan: ")
	fmt.Println("1. Siswa")
	fmt.Println("2. Guru")
	fmt.Scan(&menu)
	switch menu {
	case 1:
		menuSiswa()
	case 2:
		menuGuru()
	default:
		fmt.Println("Mohon pilih menu yang benar")
	}
}

//=================================METHOD UNTUK SISWA==================

//Hasil untuk mendapatkan nilai
func Hasil(tingkat string, kelas string, mapel string, temaSoal string, kumpulanJawaban []string) *pb.Result {
	conn := ConnectServer()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	hasilJawaban := pb.KumpulanJawaban{Tingkat: tingkat, Kelas: kelas, Mapel: mapel,
		TemaSoal: temaSoal, JawabanList: kumpulanJawaban}

	hasil, err := conn.Hasil(ctx, &hasilJawaban)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return hasil
}

func getSoal(tingkat string, kelas string, mapel string, temaSoal string) *pb.SoalList {
	conn := ConnectServer()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pilihanSoal := pb.PilihanSoal{Tingkat: tingkat, Kelas: kelas, Mapel: mapel, TemaSoal: temaSoal}

	soalList, err := conn.Soal(ctx, &pilihanSoal)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return soalList
}

func menuSiswa() {
	var menuTingkat int = 0

	for menuTingkat != 5 {
		var menuKelasSD int = 0
		var menuKelasSMP int = 0
		var menuKelasSMA int = 0
		var menuKelasUmum int = 0
		var pilihJawaban string

		fmt.Println("\n====================")
		fmt.Println("\tE-Learning")
		fmt.Println("1. SD")
		fmt.Println("2. SMP")
		fmt.Println("3. SMA")
		fmt.Println("4. Umum")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih tingkat :\t")
		fmt.Scanf("%d", &menuTingkat)
		fmt.Println("\nTingkat yang dipilih", menuTingkat)
		fmt.Println()

		switch menuTingkat {
		case 1:
			for menuKelasSD != 7 {
				var menuMapelSD int = 0
				tingkat := "SD"

				fmt.Println("\n\tSD")
				fmt.Println("1. Kelas I")
				fmt.Println("2. Kelas II")
				fmt.Println("3. Kelas III")
				fmt.Println("4. Kelas IV")
				fmt.Println("5. Kelas V")
				fmt.Println("6. Kelas VI")
				fmt.Println("7. Kembali")
				fmt.Print("Pilih kelas :\t")
				fmt.Scanf("%d", &menuKelasSD)
				fmt.Println("\nKelas yang dipilih", menuKelasSD)
				fmt.Println()

				switch menuKelasSD {
				case 1:
					for menuMapelSD != 3 {
						var menuTemaSoalSD int = 0
						kelas := "Kelas 1"
						fmt.Println("\n\tMata Pelajaran")
						fmt.Println("1. Bahasa Inggris")
						fmt.Println("2. Bahasa Indonesia")
						fmt.Println("3. Kembali")
						fmt.Print("Pilih mapel :\t")
						fmt.Scanf("%d", &menuMapelSD)
						fmt.Println("\nMapel yang dipilih", menuMapelSD)
						fmt.Println()

						switch menuMapelSD {
						case 1:
							for menuTemaSoalSD != 3 {
								mapel := "Bahasa Inggris"
								fmt.Println("\n\tTema Soal")
								fmt.Println("1. BAB I")
								fmt.Println("2. BAB II")
								fmt.Println("3. Kembali")
								fmt.Print("Pilih mapel :\t")
								fmt.Scanf("%d", &menuTemaSoalSD)
								fmt.Println("\nMapel yang dipilih", menuTemaSoalSD)
								fmt.Println()

								switch menuTemaSoalSD {
								case 1:
									temaSoal := "Bab 1"

									soal := getSoal(tingkat, kelas, mapel, temaSoal)

									//fmt.Println(soal)
									//fmt.Println()

									for i := 0; i < len(soal.SoalList); i++ {
										fmt.Println("Pertanyaan no", i+1)
										fmt.Println(soal.SoalList[i].Pertanyaan)

										for j := 0; j < len(soal.SoalList[i].Pilihan); j++ {
											if j == 0 {
												fmt.Println("a.", soal.SoalList[i].Pilihan[j])
											} else if j == 1 {
												fmt.Println("b.", soal.SoalList[i].Pilihan[j])
											} else if j == 2 {
												fmt.Println("c.", soal.SoalList[i].Pilihan[j])
											} else if j == 3 {
												fmt.Println("d.", soal.SoalList[i].Pilihan[j])
											}
										}
										fmt.Print("Pilih jawaban :\t")
										fmt.Scanf("%s", &pilihJawaban)
										if pilihJawaban == "a" {
											jawaban := soal.SoalList[i].Pilihan[0]
											kumpulanJawaban = append(kumpulanJawaban, jawaban)
										} else if pilihJawaban == "b" {
											jawaban := soal.SoalList[i].Pilihan[1]
											kumpulanJawaban = append(kumpulanJawaban, jawaban)
										} else if pilihJawaban == "c" {
											jawaban := soal.SoalList[i].Pilihan[2]
											kumpulanJawaban = append(kumpulanJawaban, jawaban)
										} else if pilihJawaban == "d" {
											jawaban := soal.SoalList[i].Pilihan[3]
											kumpulanJawaban = append(kumpulanJawaban, jawaban)
										}
									}

									hasil := Hasil(tingkat, kelas, mapel, temaSoal, kumpulanJawaban)

									fmt.Println(hasil.Result)

									kumpulanJawaban = []string{}
									main()
								}
							}
						}
					}
				}
			}
		case 2:
			for menuKelasSMP != 4 {
				fmt.Println("\n\tSMP")
				fmt.Println("1. Kelas I")
				fmt.Println("2. Kelas II")
				fmt.Println("3. Kelas III")
				fmt.Println("4. Kembali")
				fmt.Print("Pilih kelas :\t")
				fmt.Scanf("%d", &menuKelasSMP)
				fmt.Println("\nKelas yang dipilih", menuKelasSMP)
				fmt.Println()
			}
		case 3:
			for menuKelasSMA != 4 {
				fmt.Println("\n\tSMA")
				fmt.Println("1. Kelas I")
				fmt.Println("2. Kelas II")
				fmt.Println("3. Kelas III")
				fmt.Println("4. Kembali")
				fmt.Print("Pilih kelas :\t")
				fmt.Scanf("%d", &menuKelasSMA)
				fmt.Println("\nKelas yang dipilih", menuKelasSMA)
				fmt.Println()
			}
		case 4:
			for menuKelasUmum != 2 {
				fmt.Println("\n\tSMA")
				fmt.Println("1. Umum")
				fmt.Println("2. Kembali")
				fmt.Print("Pilih kelas :\t")
				fmt.Scanf("%d", &menuKelasUmum)
				fmt.Println("\nKelas yang dipilih", menuKelasUmum)
				fmt.Println()
			}
		case 5:
			fmt.Println("Terimakasih")
		default:
			fmt.Print("Maaf, no pilihan yang anda masukkan salah.")
		}
	}
}

//============================================================================

//=================================MEMBUAT SOAL============================

func menuGuru() {
	var input int
	fmt.Println("1. Buat Soal")
	fmt.Scan(&input)

	switch input {
	case 1:
		buatSoal()
	default:
		fmt.Println("Mohon input menu guru yang sesuai")
	}
}

//KirimSoalBuatan kirim Paket Soal yang dibuat untuk disimpan di server
func KirimSoalBuatan(paketSoal PaketSoal) string {
	guru := ConnectServerGuru()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	convertedPS := &pb.PaketSoal{
		Tingkat: "",
		Kelas:   "",
		Matpel:  "",
		Tema:    "",
		Soal:    []*pb.Soal{},
	}
	convertedPS.Tingkat = paketSoal.Tingkat
	convertedPS.Kelas = paketSoal.Kelas
	convertedPS.Matpel = paketSoal.Matpel
	convertedPS.Tema = paketSoal.Tema

	var result2 []*pb.Soal
	for i := 0; i < len(paketSoal.Soal); i++ {
		tempsoal := pb.Soal{Pertanyaan: paketSoal.Soal[i].Pertanyaan, Jawaban: paketSoal.Soal[i].Jawaban, Pilihan: paketSoal.Soal[i].Pilihan}
		result2 = append(result2, &tempsoal)
	}
	convertedPS.Soal = result2

	msg, err := guru.KirimSoalKeServer(ctx, convertedPS)
	if err != nil {
		log.Fatalf("Tidak bisa terhubung: %v", err)
	}
	fmt.Println("Laporan", msg.GetMessage())
	return msg.GetMessage()
}

func buatArraySoal() []Soal {
	reader := bufio.NewReader(os.Stdin)
	//loop input soal. stop ketika guru input "selesai"
	var kumpulansoal []Soal
	var soal, jawaban, pilihan string
	var arraypilihan []string
	var jumlahsoal int

	for jumlahsoal >= 0 {
		fmt.Println("Jumlah soal yang sudah anda buat :", jumlahsoal, "Ketik \"selesai\" jika sudah")
		fmt.Println("Input soal")
		soal, _ = reader.ReadString('\n')
		soal = strings.Replace(soal, "\n", "", -1)
		if soal == "selesai" {
			fmt.Println("Input selesai")
			return kumpulansoal
		}
		fmt.Println("Input jawaban")
		jawaban, _ = reader.ReadString('\n')
		jawaban = strings.Replace(jawaban, "\n", "", -1)
		arraypilihan = append(arraypilihan, jawaban)

		for i := 0; i < 3; i++ {

			fmt.Println("Input pilihan ", i+1)
			pilihan, _ = reader.ReadString('\n')
			pilihan = strings.Replace(pilihan, "\n", "", -1)

			//handle kalau pilihan yang diinput sama dengan yang sebelum"nya
		CHECK:
			for _, a := range arraypilihan {
				if a == pilihan {
					fmt.Println("Input pilihan ada yang sama. Mohon ganti pilihan anda.")
					pilihan, _ = reader.ReadString('\n')
					pilihan = strings.Replace(pilihan, "\n", "", -1)
					goto CHECK
				}
			}
			arraypilihan = append(arraypilihan, pilihan)
		}
		tempsoal := Soal{soal, jawaban, arraypilihan}
		arraypilihan = []string{}
		kumpulansoal = append(kumpulansoal, tempsoal)
		jumlahsoal++
	}
	return kumpulansoal
}

func buatSoal() {
	reader := bufio.NewReader(os.Stdin)
	var tingkat, kelas, mapel, tema string
	var menuTingkat, menuKelas, menuMapel int
	fmt.Println("Pilih tingkat")
	fmt.Println("1. SD")
	fmt.Println("2. SMP")
	fmt.Println("3. SMA")
	fmt.Println("4. Umum")
	fmt.Println("4. Umum")
	fmt.Scan(&menuTingkat)

	switch menuTingkat {
	case 1:
		tingkat = "SD"
		fmt.Println("Pilih Kelas")
		fmt.Println("1. Kelas I")
		fmt.Println("2. Kelas II")
		fmt.Println("3. Kelas III")
		fmt.Println("4. Kelas IV")
		fmt.Println("5. Kelas V")
		fmt.Println("6. Kelas VI")
		fmt.Scan(&menuKelas)

		switch menuKelas {
		case 1:
			kelas = "Kelas 1"
			fmt.Println("Pilih Matpel")
			fmt.Println("1. Bahasa Inggris")
			fmt.Println("2. Bahasa Mandarin")
			fmt.Scan(&menuMapel)

			switch menuMapel {
			case 1:
				mapel = "Bahasa Inggris"

				fmt.Println("Input tema:")
				tema, _ = reader.ReadString('\n')
				tema = strings.Replace(tema, "\n", "", -1)
			}
		}
	}

	fmt.Println(tingkat, kelas, mapel, tema)
	kumpulansoal := buatArraySoal()
	fmt.Println(kumpulansoal)
	paketSoal := PaketSoal{tingkat, kelas, mapel, tema, kumpulansoal}
	KirimSoalBuatan(paketSoal)

}

//============================================================================
