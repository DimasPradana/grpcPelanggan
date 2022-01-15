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
// server is used to implement grpcPelanggan.PelangganServer
type server struct {
	pb.UnimplementedGetPelangganServer
} // }}}

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

	pel := pb.PelangganResponse{}

	// connection string
	//  TODO: (snub on Tue 28 Dec 2021 15:57:12) : pake env belum bisa => done
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("dbport"), os.Getenv("username"),
		os.Getenv("password"), os.Getenv("dbname"))

	// check db
	db, err := sql.Open("postgres", psqlConn)
	CheckError(err, "gagal koneksi ke database", "di baris 60")

	// ping db
	err = db.Ping()
	CheckError(err, "gagal ping ke database", "di baris 64")

	defer db.Close()

	// lo.Printf("Connected to database")
	log.Printf("Pesan dari Client: %v", in.GetNolanggan())

	qry := fmt.Sprintf(`SELECT "unit", "alamat", "namapelang", st_astext("wkb_geometry", 4326), "no_langgan", 
"no_sambung" from "pelanggan" WHERE "no_langgan" = '%s'`, in.GetNolanggan())

	// get data from tables pelanggan
	rows, err := db.Query(qry)
	CheckError(err, "gagal query", "di baris 76")

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&pel.Unit, &pel.Alamat, &pel.Namapelang, &pel.Geometry, &pel.NoLanggan, &pel.NoSambung)
		CheckError(err, "gagal masukkan data ke variable", "di baris 82")
	}

	hasil := &pb.PelangganResponse{
		Unit:       pel.Unit,
		NoSambung:  pel.NoSambung,
		NoLanggan:  pel.NoLanggan,
		Namapelang: pel.Namapelang,
		Alamat:     pel.Alamat,
		Geometry:   pel.Geometry,
	}

	pel.Unit = ""
	pel.Alamat = ""
	pel.Namapelang = ""
	pel.Geometry = ""
	pel.NoLanggan = ""
	pel.NoSambung = ""

	return hasil, nil
}

// }}}

// {{{ GetAllPelangganApi
func (s *server) GetAllPelangganApi(ctx context.Context, in *pb.PelangganRequest) (*pb.AllPelangganResponse, error) {
	a := pb.PelangganResponse{}
	// XXX: could no omit type *pb.AllPelangganResponse from declaration by snub on Fri 14 Jan 2022 22:04:09 not sure why
	//var pelangganList *pb.AllPelangganResponse = &pb.AllPelangganResponse{}
	var pelangganList = &pb.AllPelangganResponse{}

	// connection string
	// TODO:(pake env belum bisa): [done] by snub on Tue 28 Dec 2021 15:57:12
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("dbport"), os.Getenv("username"),
		os.Getenv("password"), os.Getenv("dbname"))

	// check db
	db, err := sql.Open("postgres", psqlConn)
	CheckError(err, "gagal koneksi ke database", "di baris 121")

	// ping db
	err = db.Ping()
	CheckError(err, "gagal ping ke database", "di baris 125")

	defer db.Close()

	// lo.Printf("Connected to database")
	log.Printf("Pesan dari Client: allpelanggan\n")

	// qry := fmt.Sprintf(`SELECT "unit", "alamat", coalesce("namapelang",''), st_astext("wkb_geometry", 4326), "no_langgan", "no_sambung" from "pelanggan" LIMIT 5`)
	qry := fmt.Sprintln(`SELECT "unit", "alamat", coalesce("namapelang",''), st_astext("wkb_geometry", 4326), "no_langgan", "no_sambung" from "pelanggan"`)

	// get data from tables pelanggan
	rows, err := db.Query(qry)
	CheckError(err, "gagal query", "di baris 137")

	defer rows.Close()

	// TODO: return as array of pelanggans by snub on Fri 14 Jan 2022 21:50:39 done
	for rows.Next() {
		err = rows.Scan(&a.Unit, &a.Alamat, &a.Namapelang, &a.Geometry, &a.NoLanggan, &a.NoSambung)
		CheckError(err, "gagal masukkan data ke variable", "di baris 144")

		pelangganList.PelangganResponses = append(pelangganList.PelangganResponses, &a)
	}

	return pelangganList, nil

}

// }}}

// {{{ main
func main() {
	/**
	TODO: snub on Tue 28 Dec 2021 23:12:46
		   ᚛ ambil semua record dari tabel pelanggan -> done on Fri 14 Jan 2022 21:50:39
		   ᚛ test untuk branch di git -> done, sudah di merge dengan main on Sat 15 Jan 2022 17:55:00
		   ᚛ make a test file
		   ᚛ split file by its function
	*/

	// {{{ GRPC
	// listen, err := net.Listen("tcp", port)
	listen, err := net.Listen("tcp", os.Getenv("port"))
	CheckError(err, "failed to listen", "di baris 168")

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
	CheckError(err, "failed to dial server from gateway", "di baris 188")

	gwmux := runtime.NewServeMux()

	err = pb.RegisterGetPelangganHandler(context.Background(), gwmux, conn)
	CheckError(err, "failed to register gateway", "di baris 193")

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
