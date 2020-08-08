# ELearning

E learning
user adalah murid dan guru

-tidak ada login (ceritanya program gratis)
-program untuk peserta sekolah (dalam hal ini, mari batasi untuk e-learning bahasa inggris dan bahasa mandarin kelas 1-3sd)
-user bisa input nama, pilih kelas dan pilih pelajaran, kemudian soal muncul berupa pg
-soalnya sama setiap kali dipilih (bikin mock 5 nomor aja gitu), tapi abcd pgnya beda-beda
-setiap selesai mengerjakan, hasilnya keluar, generate certificate (.txt aja)

Apa yang bikin API kita berbeda dengan e-learning yg ada di luar sana:
- mari buat fitur user guru bisa bikin soal sendiri, user guru input soal dan abcd di client, simpan di server, lalu soal ditampilkan dan bisa diwrite to file
- soal yang dibuat guru bisa diread oleh client, lalu diolah menjadi soal pg, lalu dikerjakan oleh murid


compile proto dengan ini
protoc   --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:.   --go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:.   --go_opt=paths=source_relative   --go-grpc_opt=paths=source_relative   siswaservice.proto
