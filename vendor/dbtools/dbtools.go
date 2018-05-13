package dbtools

// podstawowa  obsluga ORM bazy danych
//TWORZNEI MODELU i podstawowe inserty
import (
	"log"
	//"time"

	"mstr"

	//	"mstr"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	//DB, err = gorm.Open("mysql", "root:superhaslo@tcp(127.0.0.1:3306)/dbname?charset=utf8")
	DB, err = gorm.Open("sqlite3", "/Dane/paspas.db")
	DB.Exec("PRAGMA foreign_keys = ON")
	//DB.LogMode(true)

	if err != nil {
		log.Fatalln(err.Error())
	}
	err = DB.DB().Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func DropTable() {
	DB.DropTableIfExists(&mstr.Uprawnienia_Nadane{})
	DB.DropTableIfExists(&mstr.User{}, &mstr.Uprawnienia_Domyslne{}, &mstr.PasswordRecord{})
}
func CreateTable() {

	DB.AutoMigrate(&mstr.User{}, &mstr.Uprawnienia_Domyslne{}, &mstr.Uprawnienia_Nadane{}, &mstr.PasswordRecord{})

}

//func CreateForeignKey() {
//	//DB.Model(&mstr.Uprawnienia_Domyslne{}).AddForeignKey("ID_RODZICA", "UPRAWNIENIA_DOMYSLNE(ID_UPRAWNIENIA)", "CASCADE", "CASCADE")
//	//DB.Model(&mstr.Uprawnienia_Nadane{}).AddForeignKey("ID_UPRAWNIENIA_FK", "UPRAWNIENIA_DOMYSLNE(ID_UPRAWNIENIA)", "RESTRICT", "RESTRICT")
//	//DB.Model(&mstr.Uprawnienia_Nadane{}).AddForeignKey("ID_UZYTKOWNIKA_FK", "UZYTKOWNICY(ID_UZYTKOWNIKA)", "CASCADE", "CASCADE")
//}

func CreateUsersDefault() {
	DB.Save(&mstr.User{EMAIL: "admin", IMIE: "iadmin", NAZWISKO: "nadmin", HASLO: "ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs="})

}

func CreateAccesDefault() {

	DB.Save(&mstr.Uprawnienia_Domyslne{FUNKCJA: "LOGOWANIE", ID_RODZICA: 1, WARTOSC_DOMYSLNA: 1, OPIS: "Pozwala na logowanie do systemu"})
	DB.Save(&mstr.Uprawnienia_Domyslne{FUNKCJA: "POBIERANIE", ID_RODZICA: 1, WARTOSC_DOMYSLNA: 1, OPIS: "Pozwala na POBIERANIE"})
	DB.Save(&mstr.Uprawnienia_Domyslne{FUNKCJA: "WYSYLANIE", ID_RODZICA: 1, WARTOSC_DOMYSLNA: 0, OPIS: "Pozwala na WYSYLANIE"})

	DB.Table("UPRAWNIENIA_NADANE").Save(&mstr.Uprawnienia_Nadane{ID_UPRAWNIENIA_FK: 1, ID_UZYTKOWNIKA_FK: 1, WARTOSC: 1})
	DB.Table("UPRAWNIENIA_NADANE").Save(&mstr.Uprawnienia_Nadane{ID_UPRAWNIENIA_FK: 2, ID_UZYTKOWNIKA_FK: 1, WARTOSC: 0})
}

func Start() {
	if DB.First(&mstr.User{}).RowsAffected == 0 {
		CreateTable()
		CreateUsersDefault()
		CreateAccesDefault()
	}
	//tak jakby init do testow
	//DropTable()
	//	CreateTable()
	//CreateForeignKey()
	//	CreateUsersDefault()
	//	CreateAccesDefault()
}
