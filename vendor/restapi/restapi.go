package restapi

//restapi ustawienia oraz ładowanie metod

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	//"sessiontools"
	"callhandle"
	"tools"
)

var (
	Ports string             = ":8080"
	REST  *httprouter.Router = httprouter.New()

	//FileaPath string             = "/home/kamil/Pobrane/"
)

func Up() {
	cert, err := tls.LoadX509KeyPair("./server.crt", "./server.key")
	tools.CheckErr(err)

	// Load CA cert
	caCert, err := ioutil.ReadFile("./CA.crt")
	tools.CheckErr(err)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		RootCAs:      caCertPool,
		ClientCAs:    caCertPool,
	}
	config.Rand = rand.Reader
	srv := &http.Server{
		Addr:      Ports,
		TLSConfig: config,
		Handler:   REST,
	}
	srv.SetKeepAlivesEnabled(true)
	fmt.Println("---- SERWER HTTP START LISTENING ----\n")

	err = srv.ListenAndServeTLS("server.crt", "server.key")
	tools.CheckErr(err)
	fmt.Println("---- SERWER HTTP STOP LISTENING ----\n")

}

func CreateFunctionAcces() {
	//logowanie do systemu
	REST.POST("/logon", callhandle.Logon)
	// informacje o access
	REST.GET("/getmyacceslist", callhandle.GetInfoAboutAcces)
	// pobieranie informacji o sesji
	REST.GET("/getmysession", callhandle.GetInfoAboutMeSession)
	REST.POST("/changepasswd", callhandle.ChangePassword)
	REST.POST("/passwd", callhandle.ChangePassword)
}

func CreateFunctionInfo() {
	//pobieranie versi i daty
	REST.GET("/version", callhandle.GetServerVersion)
	REST.GET("/datenow", callhandle.GetServerDate)

}
func CreateFunctionStatic() {
	// Pobieranie statycznych plików
	REST.GET("/static/*file", callhandle.StaticSessionFiles)
}

func CreateFunctionPamietnikHasel() {
	REST.POST("/ph/add", callhandle.AddRecordPassword)
	REST.POST("/ph/update", callhandle.UpdateRecord)
	REST.POST("/ph/remove", callhandle.RemoveRecord)
	REST.GET("/ph/getall", callhandle.GetRecords)
	//Pamiętnik Hasel
}
func CreateFunction() {
	CreateFunctionAcces()
	CreateFunctionInfo()
	//CreateFunctionStatic()

	CreateFunctionPamietnikHasel()
}

//	REST.POST("/passwd", callhandle.ZmianaHasla)
//	REST.POST("/add", callhandle.AddRecord)
//	REST.POST("/update", callhandle.UpdateRecord)
//	REST.POST("/remove", callhandle.RemoveRecord)
//	REST.GET("/getall", callhandle.GetRecords)
//Pobieranie zazwyczaj stałych danych, ale waznych:

func Start() {
	CreateFunction()
	Up()
}
