package callhandle

// uchwyty
import (
	"encoding/json"
	"fmt"
	"mstr"
	//"os/user"
	"io/ioutil"
	"net/http"
	"sessiontools"
	"strconv"
	//"log"
	"dbtools"
	"tools"

	"github.com/julienschmidt/httprouter"

	"time"
)

func ChangePassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {
		//Zmien_haslo(id int, stare_haslo string, nowe_haslo string) (status bool)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		nowe_haslo := mstr.ZmianaHasla{}
		json.Unmarshal(bodyBytes, &nowe_haslo)
		return_update := mstr.ReturnUpdate{}
		return_update.Affected = dbtools.DB.Model(&mstr.User{}).Where("ID_UZYTKOWNIKA= ?", nowe_haslo.Id).Where("HASLO= ?", nowe_haslo.Stare).Update("HASLO", nowe_haslo.Nowe).RowsAffected
		//Zmien_haslo(nowe_haslo.Id, nowe_haslo.Stare, nowe_haslo.Nowe)
		return_update_json, err := json.Marshal(return_update)
		if err == nil {
			fmt.Fprintf(w, string(return_update_json))
			//fmt.Println(string(return_update_json))
			return
		}
		tools.CheckErr(err)
	}
	http.Redirect(w, r, "/", 301)

}

func GetInfoAboutAcces(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {

		info_session, err := json.Marshal(dbtools.GetAccesDataUser(strconv.Itoa(sessiontools.GetInfoAboutUsSe(r.Header.Get("sessiontopgoid")).Id)))
		if err == nil {
			fmt.Fprintf(w, string(info_session))
			return
		}
		tools.CheckErr(err)
	}
	http.Redirect(w, r, "/", 301)
}

func GetInfoAboutMeSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {
		info_session, err := json.Marshal(sessiontools.GetInfoAboutUsSe(r.Header.Get("sessiontopgoid")))
		if err == nil {
			fmt.Fprintf(w, string(info_session))
			return
		}
		tools.CheckErr(err)
	}
	http.Redirect(w, r, "/", 301)
}

func Logon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	user := mstr.SingleSessionUser{}
	json.Unmarshal(bodyBytes, &user)
	//fmt.Println(user)
	usin := mstr.User{}
	//fmt.Println(user.Login)
	status := dbtools.DB.Where("EMAIL = ?", user.Login).Where("HASLO = ?", user.Haslo).First(&usin).RowsAffected
	//status := dbtools.DB.Where(mstr.User{EMAIL: user.Imie}).Where(mstr.User{HASLO: user.Haslo}).First(&usin).RowsAffected
	//fmt.Println(usin)
	if status == 1 {
		ssid, err := json.Marshal(map[string]string{"Sessiontopgoid": sessiontools.NewSessionUser(r,
			user.Login,
			user.Haslo,
			usin.IMIE,
			usin.NAZWISKO,
			usin.ID_UZYTKOWNIKA)})
		if err == nil {
			fmt.Fprintf(w, string(ssid))
			return
		}
		tools.CheckErr(err)
	}
	http.Redirect(w, r, "/", 301)

}
func GetServerDate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {
		v, err := json.Marshal(map[string]time.Time{"datetime": time.Now()})
		if err == nil {
			fmt.Fprintf(w, string(v))
			return
		}
		//tools.CheckErr(err)
	}
	http.Redirect(w, r, "/", 301)
}

func GetAccesList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {
		v, err := json.Marshal(map[string]time.Time{"datetime": time.Now()})
		if err == nil {
			fmt.Fprintf(w, string(v))
			return
		}
		//tools.CheckErr(err)
	}
	http.Redirect(w, r, "/", 301)
}

func GetServerVersion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {
		v, err := json.Marshal(mstr.Version{Numer: 0.1, Data: time.Date(
			2017, 12, 10, 23, 32, 00, 00, time.UTC)})
		if err == nil {
			fmt.Fprintf(w, string(v))
			return
		}
		//tools.CheckErr(err)
	}
	http.Redirect(w, r, "/", 301)
}
