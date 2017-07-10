package main

    import (
        "database/sql"
        "fmt"
        _"github.com/lib/pq"
    )

    const (
        DB_USER     = "postgres"
        DB_PASSWORD = "123456"
        DB_NAME     = "DataMahasiswa"
    )

    func inputData(){
        var nama string
        var jurusan string

        dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
        db, err := sql.Open("postgres", dbinfo)
        checkErr(err)
        defer db.Close()

        print("Masukkan Nama : ")
        fmt.Scanf("%s", &nama)

        print("Masukkan Jurusan : ")
        fmt.Scanf("%s", &jurusan)

        var lastInsertId int
        err = db.QueryRow("INSERT INTO userinfo(nama,jurusan) VALUES($1,$2) returning uid;", nama, jurusan).Scan(&lastInsertId)
        checkErr(err)
        fmt.Println("Insert Success\n")
    }

    func readData(){
        dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
        db, err := sql.Open("postgres", dbinfo)
        checkErr(err)
        defer db.Close()

        rows, err := db.Query("SELECT * FROM userinfo")
        checkErr(err)

        for rows.Next() {
            var uid int
            var nama string
            var jurusan string
            err = rows.Scan(&uid, &nama, &jurusan)
            checkErr(err)
            fmt.Println("uid | Nama | Jurusan")
            fmt.Printf("%3v | %3v | %3v\n", uid, nama, jurusan)
        }
        fmt.Println("Read Data Succes\n")
    }

    var input string

    func main() {

        for input != "3" {
            print("Program Data Mahasiswa\n")
            print("1. Input Data\n")
            print("2. Read Data\n")
            print("3. End\n")
            print("Masukkan Pilihan : ")
            fmt.Scanf("%s", &input)
            print("\n")

            switch input {
                case "1":
                    inputData()
                    break
                case "2":
                    readData()
                    break
                case "3":
                    break
            }
        }

        /*
        fmt.Println("# Updating")
        stmt, err := db.Prepare("update userinfo set nama=$1 where uid=$2")
        checkErr(err)

        res, err := stmt.Exec("kadekupdate", lastInsertId)
        checkErr(err)

        affect, err := res.RowsAffected()
        checkErr(err)

        fmt.Println(affect, "rows changed")



        fmt.Println("# Deleting")
        stmt, err = db.Prepare("delete from userinfo where uid=$1")
        checkErr(err)

        res, err = stmt.Exec(lastInsertId)
        checkErr(err)

        affect, err = res.RowsAffected()
        checkErr(err)

        fmt.Println(affect, "rows changed")
        */
    }

    func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }
