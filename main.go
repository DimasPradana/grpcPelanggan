package main

// {{{ import
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/DimasPradana/kantor/grpcPelanggan/config"
	pb "github.com/DimasPradana/kantor/grpcPelanggan/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

// }}}

// {{{ others
var unit, alamat, namapelang, wkb_geometry, no_langgan, no_sambung string

// server is used to implement grpcPelanggan.PelangganServer
type server struct {
	pb.UnimplementedGetPelangganServer
}

//func NewServer() *server {
//	return &server{}
//}
// }}}

// {{{ init
func init() {
	config.Getenvfile()
}

// }}}

// {{{ CheckError
func CheckError(err error, pesan string, baris string) {
	if err != nil {
		log.Fatalf("errornya: %v, %s, %s", err, pesan, baris)
		// panic(err)
	}
}

// }}}

// {{{ GetPelangganApi
func (s *server) GetPelangganApi(ctx context.Context, in *pb.PelangganRequest) (*pb.PelangganResponse, error) {
	// connection string
	//  TODO snub on Tue 28 Dec 2021 15:57:12 : pake env belum bisa => done
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("dbport"), os.Getenv("username"),
		os.Getenv("password"), os.Getenv("dbname"))

	// check db
	db, err := sql.Open("postgres", psqlConn)
	CheckError(err, "gagal koneksi ke database", "di baris 64")

	// ping db
	err = db.Ping()
	CheckError(err, "gagal ping ke database", "di baris 68")

	defer db.Close()

	// lo.Printf("Connected to database")
	log.Printf("Pesan dari Client: %v", in.GetNolanggan())

	qry := fmt.Sprintf(`SELECT "unit", "alamat", "namapelang", st_astext("wkb_geometry", 4326), "no_langgan", 
"no_sambung" from "pelanggan" WHERE no_langgan = '%s'`, in.GetNolanggan())

	// get data from tables pelanggan
	rows, err := db.Query(qry)
	CheckError(err, "gagal query", "di baris 80")

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&unit, &alamat, &namapelang, &wkb_geometry, &no_langgan, &no_sambung)
		CheckError(err, "gagal masukkan data ke variable", "di baris 86")

	}

	hasil := &pb.PelangganResponse{
		Unit:       unit,
		NoSambung:  no_sambung,
		NoLanggan:  no_langgan,
		Namapelang: namapelang,
		Alamat:     alamat,
		Geometry:   wkb_geometry,
	}
	unit, alamat, namapelang, wkb_geometry, no_langgan, no_sambung = "", "", "", "", "", ""

	return hasil, nil
}

// }}}

// {{{ main
func main() {
/**
 TODO snub on Tue 28 Dec 2021 23:12:46
  ᚛ ambil semua record dari tabel pelanggan
  ᚛ test untuk branch di git 
  ᚛ 
*/

	// {{{ GRPC
	// listen, err := net.Listen("tcp", port)
	listen, err := net.Listen("tcp", os.Getenv("port"))
	CheckError(err, "failed to listen", "di baris 116")

	ser := grpc.NewServer()
	pb.RegisterGetPelangganServer(ser, &server{})

	log.Printf("serving grpc on port: %v", listen.Addr())

	go func() {
		log.Fatalln(ser.Serve(listen))
	}()
	// }}}

	// {{{ GRPC gateway
	conn, err := grpc.DialContext(
		context.Background(),
		//"localhost"+port,
		"localhost"+os.Getenv("port"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	CheckError(err, "failed to dial server from gateway", "di baris 136")

	gwmux := runtime.NewServeMux()

	err = pb.RegisterGetPelangganHandler(context.Background(), gwmux, conn)
	CheckError(err, "failed to register gateway", "di baris 141")

	gwServer := &http.Server{
		// Addr:    portgw,
		Addr:    os.Getenv("portgw"),
		Handler: gwmux,
	}

	log.Printf("serving grpc-gateway on port: %v", gwServer.Addr)
	log.Fatalln(gwServer.ListenAndServe())
	// }}}
}

// }}}

// vim:fileencoding=utf-8:ft=go:foldmethod=marker
