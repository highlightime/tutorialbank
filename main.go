package main

import (
	"database/sql"
	"firstgo/mysql"
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("index.html")
	//t.Execute(w, "hello world!")

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		id := r.FormValue("id")
		pw := r.FormValue("pw")

		db, err := sql.Open("mysql", "root:970810@/testdb")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = mysql.Open(db)
		if err != nil {
			log.Fatal(err)
		}

		isDup, err := mysql.Find(db, id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(isDup)
		if !isDup {
			err = mysql.Put(db, id, pw)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
			//fmt.Fprintf(w, "ID = %s\n", id)
			//fmt.Fprintf(w, "Password = %s\n", pw)
		} else {
			// login
			ok, err := mysql.Login(db, id, pw)
			if err != nil {
				log.Fatal(err)
			}
			// TODO: Ok 가 아닌 경우에도 로그인되는거로 들어가고있음
			if ok {
				http.ServeFile(w, r, "balance.html")
				bal, err := mysql.GetBal(db, id)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(bal)
				//fmt.Fprintf(w, "ID = %s\n", id)
				//fmt.Fprintf(w, "Password = %s\n", pw)
			} else {
				// reload
				http.ServeFile(w, r, "index.html")
			}
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func balance(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "balance.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		balance := r.FormValue("balance")
		fmt.Fprintf(w, balance)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

var EXTERNAL int = 1

func IsEven(chEven chan bool, input int) (result bool) {
	if input%2 == 0 {
		chEven <- true
		return true
	} else {
		chEven <- false
		return false
	}
}

func main() {
	/*
		go concurrency.IsEven(EXTERNAL)
		EXTERNAL++
		go concurrency.IsEven(EXTERNAL)
		//time.Sleep(time.Second * 1)
	*/

	/*
		chEven := make(chan bool)
		//chEven := make(chan bool, 100)
		defer close(chEven)

		start := time.Now()
		for i := 1; i < 100; i++ {
			go IsEven(chEven, i)
		}
		for i := 1; i < 100; i++ {
			result := <-chEven
			fmt.Println(result)
		}
		elasped := time.Since(start)
		fmt.Println(elasped)
	*/

	http.HandleFunc("/", index)
	http.HandleFunc("/balance", balance)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
