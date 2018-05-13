package dbtools

// podstawowa  obsluga ORM bazy danych
// selecty z bazy
import (
	"log"
	//"time"

	"mstr"
)

func GetAccesDataUser(Id string) (results []mstr.AccesInfo) {
	rows, err := DB.Table("UPRAWNIENIA_DOMYSLNE").Select(`UPRAWNIENIA_DOMYSLNE.ID_UPRAWNIENIA AS ID_UPRAWNIENIA,
		 UPRAWNIENIA_DOMYSLNE.ID_RODZICA AS ID_RODZICA,
		 UPRAWNIENIA_DOMYSLNE.FUNKCJA AS FUNKCJA,
	coalesce(UPRAWNIENIA_NADANE.WARTOSC,UPRAWNIENIA_DOMYSLNE.WARTOSC_DOMYSLNA) AS WARTOSC,
	UPRAWNIENIA_DOMYSLNE.OPIS AS OPIS,
	UPRAWNIENIA_DOMYSLNE.WARTOSC_DOMYSLNA AS WARTOSC_DOMYSLNA ,
	coalesce(UPRAWNIENIA_NADANE.WARTOSC,-1) AS STATUS`).Joins(`left join UPRAWNIENIA_NADANE on UPRAWNIENIA_DOMYSLNE.ID_UPRAWNIENIA=UPRAWNIENIA_NADANE.ID_UPRAWNIENIA_FK and UPRAWNIENIA_NADANE.ID_UZYTKOWNIKA_FK=1`).Rows()
	if err != nil {
		log.Println("Połączenie zerwane / lub nie wiersza do zwrocenia")
		log.Println(err.Error())
	}
	acces := mstr.AccesInfo{}
	for rows.Next() {
		rows.Scan(&acces.Id, &acces.IdR, &acces.Funkcja, &acces.Wartosc, &acces.Opis, &acces.WartoscDomyslna, &acces.Status)
		results = append(results, acces)
		//log.Println("user pobierz dane:", acces)

	}
	return

}
