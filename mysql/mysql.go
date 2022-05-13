package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Open(db *sql.DB) error {
	fmt.Println("connect success", db)

	// CREATE TABLE
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test(id VARCHAR(20), pw VARCHAR(20),balance INT, primary key(id))")
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("create")
	/*
		// INSERT 실행
		_, err = db.Exec("INSERT IGNORE INTO test(id, pw) VALUES ('myid', 'mypw')")
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("insert")
		// SELECT 쿼리
		rows, err := db.Query("SELECT id, pw FROM test")
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer rows.Close() //반드시 닫는다 (지연하여 닫기)
		fmt.Println("select")
		var id string
		var name string
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
				return err
			}
			fmt.Println(id, name)
		}
	*/
	return nil
}

func Login(db *sql.DB, id, pw string) (bool, error) {

	return false, nil
}

func Put(db *sql.DB, id, pw string) error {
	// INSERT 실행
	money := 10000
	row := fmt.Sprintf(`INSERT INTO test(id, pw, balance) VALUES ('%s', '%s', %d)`, id, pw, money)
	_, err := db.Exec(row)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("insert")
	return nil
}

func Find(db *sql.DB, id string) (bool, error) {
	// SELECT 쿼리
	var s sql.NullString
	row := fmt.Sprintf(`SELECT id, pw FROM test WHERE id='%s'`, id)
	err := db.QueryRow(row).Scan(&s)
	if err != nil {
		return false, err
	}
	fmt.Println("select")
	if s.Valid {
		// use s.String
		fmt.Println("Duplicated")
		return true, nil
	}
	return false, nil
}

func Signup(db *sql.DB, id string, pw string) (bool, error) {
	isDup, _ := Find(db, id)
	if !isDup {
		err := Put(db, id, pw)
		if err != nil {
			log.Fatal(err)
			return false, nil
		}
		return true, nil
	} else {
		return false, nil
	}
}

func GetBal(db *sql.DB, id string) (int, error) {
	row := fmt.Sprintf(`SELECT balance FROM test WHERE id='%s'`, id)
	rows, err := db.Query(row)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
	fmt.Println("select")
	var bal int
	for rows.Next() {
		err := rows.Scan(&bal)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		return bal, nil
	}
	return 0, err
}

func GetAll() error {
	db, err := sql.Open("mysql", "root:970810@/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// SELECT 쿼리
	rows, err := db.Query("SELECT id, pw, balance FROM test")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
	fmt.Println("select")
	var id string
	var name string
	var bal int
	for rows.Next() {
		err := rows.Scan(&id, &name, &bal)
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println(id, name, bal)
	}
	return nil
}
