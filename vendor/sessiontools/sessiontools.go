package sessiontools

// osluga sessji

import (
	//	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	//"database/sql"

	//"dbtools"

	"mstr"
	"tools"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//DB     *sql.DB
	SpoolU = make(map[string]*mstr.SingleSessionUser)
	//SpoolH = make(map[string]*mystru.SingleSessionUser)
	mutexs = &sync.Mutex{}
)

func Test() {
	//	fmt.Println(dbtools.Sprawdz_haslo("kamil@topsa.com.pl", "a"))

}

func DeleteOlderSession(id int) {
	for k := range SpoolU {
		if SpoolU[k].Id == id {
			delete(SpoolU, k)
		}
	}
	//	fmt.Printf("tes")
}
func NewSessionUser(r *http.Request,
	login, haslo, imie, nazwisko string, id int) (ciacho string) {
	mutexs.Lock()
	for k := range SpoolU {
		if SpoolU[k].Id == id {
			czas := time.Now()
			SpoolU[k].Czas = czas

			ciacho = k
			break
			//fmt.Println("co jest")
		}
	}

	mutexs.Unlock()
	if len(ciacho) > 0 {
		return
	}
	//	fmt.Printf("tes")
	ciacho, _ = tools.GenerateRandomString(12)
	ciacho = strconv.Itoa(id) + ":" + ciacho
	czas := time.Now()
	mutexs.Lock()
	//	fmt.Println("Tworze sesje: ", ciacho)
	SpoolU[ciacho] = mstr.NewSingleSessionUser(id,
		imie,
		nazwisko,
		login,
		haslo,
		strings.Split(r.RemoteAddr, ":")[0],
		czas)
	mutexs.Unlock()
	return
	//cookie := http.Cookie{Name: "sessiontopgoid", Value: ssid}

	//fmt.Println(adressip, email, imie, nazwisko, id)
	//http.SetCookie(w, &cookie)

}

func CheckSessionUser(ciacho string) (value bool) {
	//return true
	//fmt.Println("sprawdzam sesje: ")
	//curl -k -X GET -H 'sessiontopgoid: 1:90iqmk2eAl35L_J7'  https://localhost:8080/version
	//curl -k -X POST -H 'haslo: moje' -d '{"username":"davidwalsh","password":"something"}' https://localhost:8080/logon

	if ciacho != "" {
		czas := time.Now()
		mutexs.Lock()
		if _, ok := SpoolU[ciacho]; ok {
			//do something here
			//fmt.Println("sesje istnieje: ", val.Id, val.Czas.String())
			// aktualizuje czas ostaniej akcji bo moze byc potrzeba usuwania
			// nieaktywnych sesji oraz wylogowywanie po jakims czasie nieaktywnosci
			// na stronie
			SpoolU[ciacho].Czas = czas
			//fmt.Println("sesje istnieje: ", val.Id, val.Czas.String())

			mutexs.Unlock()
			return true

		} else {
			mutexs.Unlock()
			//fmt.Println("ciasto jest ale nie poprawne")

			return false
		}

	} else {
		//fmt.Println("ciastka nie ma")
		return false

	}
}

func GetInfoAboutUsSe(ciacho string) (sesionInfoUser mstr.SingleSessionUser) {

	mutexs.Lock()
	if val, ok := SpoolU[ciacho]; ok {
		sesionInfoUser = val.CopyToNewRecord()

	}
	mutexs.Unlock()
	return

}
