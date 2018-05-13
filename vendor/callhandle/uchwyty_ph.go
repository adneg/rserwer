package callhandle

import (
	"encoding/json"
	"fmt"
	"mstr"
	//"os/user"
	"io/ioutil"
	"net/http"
	"sessiontools"
	//"strconv"
	//"log"
	"dbtools"
	"tools"

	"github.com/julienschmidt/httprouter"
	//"time"
)

func AddRecordPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {
		//fmt.Println("dupas")
		//Zmien_haslo(id int, stare_haslo string, nowe_haslo string) (status bool)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		nowe_haslo := mstr.PasswordRecord{}
		json.Unmarshal(bodyBytes, &nowe_haslo)
		return_update := mstr.Returner{}
		////fmt.Println(nowe_haslo)
		return_update.Affected = dbtools.DB.Save(&nowe_haslo).RowsAffected
		return_update.Id = nowe_haslo.Id
		//fmt.Println(nowe_haslo.Id)
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

func UpdateRecord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {
		//Zmien_haslo(id int, stare_haslo string, nowe_haslo string) (status bool)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		nowe_haslo := mstr.PasswordRecord{}
		json.Unmarshal(bodyBytes, &nowe_haslo)
		return_update := mstr.Returner{}
		//fmt.Println(nowe_haslo)
		return_update.Affected = dbtools.DB.Model(&mstr.PasswordRecord{}).Where("ID= ?", nowe_haslo.Id).Updates(nowe_haslo).RowsAffected
		//("HASLO", nowe_haslo.Nowe).RowsAffected
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

func RemoveRecord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		nowe_haslo := mstr.PasswordRecord{}
		json.Unmarshal(bodyBytes, &nowe_haslo)
		return_update := mstr.Returner{}
		//fmt.Println(nowe_haslo)
		return_update.Affected = dbtools.DB.Delete(&nowe_haslo).RowsAffected
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

func GetRecords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if sessiontools.CheckSessionUser(r.Header.Get("sessiontopgoid")) {
		passwords := []mstr.PasswordRecord{}
		dbtools.DB.Find(&passwords)
		return_update_json, err := json.Marshal(&passwords)
		if err == nil {
			fmt.Fprintf(w, string(return_update_json))
			//fmt.Println(string(return_update_json))
			return
		}
		tools.CheckErr(err)
	}
	http.Redirect(w, r, "/", 301)

}
