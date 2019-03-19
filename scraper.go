/*
Programmet saknar errorhandling, om du vill kan du implementera det.
en transaktion (tx) måste antingen slutas med en Commit eller en Rollback

*/
package main

import (
	"fmt"
	"time"
	"net/http"
	_ "github.com/gwenn/gosqlite"
	"database/sql"
	"io/ioutil"
)

func main(){
	fmt.Println("program started")
	//Öppna en databas som heter db.sqlite
	db,_ := sql.Open("sqlite3","./db.sqlite")
	for {
		fmt.Println("Fetching data")

		//Gör ett GET-anrop till https://api.chucknorris.io/jokes/random
		response,_ := http.DefaultClient.Get("https://api.chucknorris.io/jokes/random")
		body,_ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		tx,_ := db.Begin()
		fmt.Println("Inserting data into database")

		//Spara datan till en tabell som heter scrapes
		txStmt,_ := tx.Prepare("INSERT INTO scrapes (http_response_body,timestamp) VALUES(?,?);")
		txStmt.Exec(body,time.Now().Unix())
		tx.Commit()
		//Vänta 2 sekunder
		time.Sleep(2*time.Second)
	}
}
