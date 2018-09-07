package main

import (
	"github.com/gorilla/mux"
	"log"
      _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
	"net/http"
	"strconv"
)

func start(w http.ResponseWriter, r *http.Request){
	// fmt.Println()
	// w.Write([]byte(`{"Status":"Passed"}`))
	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")
	stmt, _ := db.Prepare("INSERT player_data SET player=1,coin=3,plantation1=2")
	stmt.Exec()
	stmt, _ = db.Prepare("INSERT player_data SET player=2,coin=3,plantation1=2")
	stmt.Exec()
	stmt, _ = db.Prepare("INSERT player_data SET player=3,coin=2,plantation1=1")
	stmt.Exec()
	stmt, _ = db.Prepare("INSERT player_data SET player=4,coin=2,plantation1=1")
	stmt.Exec()
	
	db.Close()

}

func prospector(w http.ResponseWriter, r *http.Request){

	url := r.URL
	params := url.Query()
	fmt.Println(params)
	var player_number string = params["player"][0]

	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")
	stmt, _ := db.Prepare("update player_data SET coin=coin+1 where player="+player_number)
	stmt.Exec()
	
	db.Close()

}

// func craftman(w http.ResponseWriter, r *http.Request){

// 	// url := r.URL
// 	// params := url.Query()
// 	// fmt.Println(params)
// 	// var player_number string = params["player"][0]

// 	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")
// 	stmt, _ := db.Prepare("update player_data SET corn=corn+ where player="+player_number)
// 	stmt.Exec()
	
// 	db.Close()

// }



func getPlayerInfo(w http.ResponseWriter, r *http.Request){


	url := r.URL
	params := url.Query()
	fmt.Println(params)
	var player_number string = params["player"][0]
	// fmt.Println()
	// w.Write([]byte(`{"Status":"Passed"}`))
	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")
	rows, err := db.Query("select player,small_indigo_plant,coin from player_data where player=?",player_number)
	var (
		id int
		small_indigo int
		coin int
	)
	for rows.Next() {
		err = rows.Scan(&id, &small_indigo,&coin)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
    	log.Println(id,small_indigo)
	}

	w.Write([]byte(`{"Player":`+strconv.Itoa(id)+`,"small_indigo_plant":`+strconv.Itoa(small_indigo)+`,"coin":`+strconv.Itoa(coin)+`}`))

	db.Close()

}


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/getPlayerInfo", getPlayerInfo).Methods("GET")
    r.HandleFunc("/start", start).Methods("GET")
    r.HandleFunc("/prospector", prospector).Methods("GET")
    // r.HandleFunc("/craftman", craftman).Methods("GET")    
    log.Fatal(http.ListenAndServe(":8086",r))
}