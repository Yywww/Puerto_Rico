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



func sqlexec(db *sql.DB,query string){
	stmt, _ := db.Prepare(query)
	stmt.Exec()
}


func get_plant_number(x []int)(int,int,int,int,int){
	var corn_plant_number,indigo_plant_number,sugar_plant_number,tobacco_plant_number,coffee_plant_number int
	for _,element := range x {
		if element==6 {
			corn_plant_number+=1
		}
		if element==7 {
			indigo_plant_number+=1
		}
		if element==8 {
			sugar_plant_number+=1
		}
		if element==9 {
			tobacco_plant_number+=1
		}
		if element==10 {
			coffee_plant_number+=1
		}
	}
	return corn_plant_number,indigo_plant_number,sugar_plant_number,tobacco_plant_number,coffee_plant_number
}

func bool2int(bool bool) int {
    if bool {
        return 1
    } else {
        return 0
    }
}

func start(w http.ResponseWriter, r *http.Request){


	// Player Data
	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")
	sqlexec(db,"DELETE from player_data")
	sqlexec(db,"INSERT player_data SET player=1,coin=3,plantation1=2")
	sqlexec(db,"INSERT player_data SET player=2,coin=3,plantation1=2")
	sqlexec(db,"INSERT player_data SET player=3,coin=2,plantation1=1")
	sqlexec(db,"INSERT player_data SET player=4,coin=2,plantation1=1")



	// Boat Data
	sqlexec(db,"DELETE from cargo_boat")
	sqlexec(db,"INSERT cargo_boat SET cargo_boat=1,max_space=5,current_cargo=0")
	sqlexec(db,"INSERT cargo_boat SET cargo_boat=2,max_space=6,current_cargo=0")
	sqlexec(db,"INSERT cargo_boat SET cargo_boat=3,max_space=7,current_cargo=0")
	
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

func craftman(w http.ResponseWriter, r *http.Request){

	url := r.URL
	params := url.Query()
	fmt.Println(params)
	var player_number string = params["player"][0]

	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")


	row := db.QueryRow("select plantation1,plantation2,plantation3,plantation4 from player_data where player=?",player_number)
	
	var (
		plantation1 int
		plantation2 int
		plantation3 int
		plantation4 int
	)

	row.Scan(&plantation1, &plantation2,&plantation3,&plantation4)


	plantation_list :=[]int{plantation1,plantation2,plantation3,plantation4}

	corn_plant_number,indigo_plant_number,sugar_plant_number,tobacco_plant_number,coffee_plant_number:=get_plant_number(plantation_list)



	sqlexec(db,"update player_data SET corn=corn+"+strconv.Itoa(corn_plant_number)+" where player="+player_number)
	sqlexec(db,"update player_data SET indigo=indigo+"+strconv.Itoa(indigo_plant_number)+" where player="+player_number)
	sqlexec(db,"update player_data SET sugar=sugar+"+strconv.Itoa(sugar_plant_number)+" where player="+player_number)
	sqlexec(db,"update player_data SET tobacco=tobacco+"+strconv.Itoa(tobacco_plant_number)+" where player="+player_number)
	sqlexec(db,"update player_data SET coffee=coffee+"+strconv.Itoa(coffee_plant_number)+" where player="+player_number)

	
	db.Close()

}

func settler(w http.ResponseWriter, r *http.Request){
	url := r.URL
	params := url.Query()
	fmt.Println(params)
	var player_number string = params["player"][0]
	var plantation string = params["plantation"][0]
	var plantation_number string = params["plantation_number"][0]
	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")
	stmt, _ := db.Prepare("update player_data set plantation"+plantation_number+"="+plantation+" where player="+player_number)	
	stmt.Exec()
	db.Close()	
}

func mayor(w http.ResponseWriter, r *http.Request){
	url := r.URL
	params := url.Query()
	fmt.Println(params)
	var player_number string = params["player"][0]
	// var plantation_number string = params["plantation_number"][0]
	var building string = params["building"][0]

	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")

	sqlexec(db,"update player_data set "+building+"="+building+"+1 where player="+player_number)

	db.Close()		
}

func trader(w http.ResponseWriter, r *http.Request){

}


func captain(w http.ResponseWriter, r *http.Request){
	
}


func builder(w http.ResponseWriter, r *http.Request){
	buildingcost := map[string]int{"small_sugar_mill": 2}
	// buildingdiscount := map[string]int{"small_sugar_mill": 1}
	
	url := r.URL
	params := url.Query()
	fmt.Println(params)
	var player_number string = params["player"][0]
	var build string = params["build"][0]
	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")
	stmt, _ := db.Prepare("update player_data set "+build+"=1 where player="+player_number)	
	stmt.Exec()
	cost := strconv.Itoa(buildingcost[build])
	stmt, _ = db.Prepare("update player_data set coin=coin-"+cost+" where player="+player_number)	
	stmt.Exec()
	db.Close()
}



func getPlayerInfo(w http.ResponseWriter, r *http.Request){


	url := r.URL
	params := url.Query()
	fmt.Println(params)
	var player_number string = params["player"][0]
	// fmt.Println()
	// w.Write([]byte(`{"Status":"Passed"}`))
	db, _:= sql.Open("mysql","root:eded....@/puerto_rico?charset=utf8")
	row := db.QueryRow("select player,small_indigo_plant,small_sugar_mill,coin,corn from player_data where player=?",player_number)

	var (
		player int
		small_indigo_plant int
		small_sugar_mill int
		coin int
		corn int
	)

	
	row.Scan(&player, &small_indigo_plant,&small_sugar_mill,&coin,&corn)


	w.Write([]byte(	`{"Player":`+strconv.Itoa(player)+
					`,"small_indigo_plant":`+strconv.Itoa(small_indigo_plant)+
					`,"small_sugar_mill":`+strconv.Itoa(small_sugar_mill)+
					`,"coin":`+strconv.Itoa(coin)+
					`,"corn":`+strconv.Itoa(corn)+
					`}`))

	db.Close()

}


func main() {

    r := mux.NewRouter()
    r.HandleFunc("/getPlayerInfo", getPlayerInfo).Methods("GET")
    r.HandleFunc("/start", start).Methods("GET")
    r.HandleFunc("/prospector", prospector).Methods("GET")
    r.HandleFunc("/builder", builder).Methods("GET")
    r.HandleFunc("/settler", settler).Methods("GET")
    r.HandleFunc("/craftman", craftman).Methods("GET")  
    log.Fatal(http.ListenAndServe(":8086",r))
}