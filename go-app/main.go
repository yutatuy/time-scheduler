package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)



func main() {

	// サンプルAPI
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		data := map[string]string{"message": "Hello, World!"}
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

		// サンプルmysqlへの接続コード
		db, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/time_scheduler")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		rows, err := db.Query("SELECT * FROM sample")
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		fmt.Println("START:START:START")

		for rows.Next() {
			var column1 string
			var column2 string
			var column3 string // 追加した行
			err = rows.Scan(&column1, &column2, &column3)
			if err != nil {
					panic(err.Error())
			}
			fmt.Println(column1, column2, column3)
		}

		err = rows.Err()
		if err != nil {
			panic(err.Error())
		}
	})

	fmt.Println("Starting server on port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

/**
	- Gorm導入
	- ユーザ作成処理
	- ログイン処理
	- Gin導入
	- JWT認証
*/

 /**
    User
 		- name
		- email
		- password
	Schedule
 		- user_id
		- name
	Place
 		- user_id
		- schedule_id
		- name
		- position
		- time
		- before_place_id(default=0)
 */