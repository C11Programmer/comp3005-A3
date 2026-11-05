package main

import (
	"fmt"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)


func main(){
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:NewPGpass@127.0.0.1:5432/postgres")

	if err != nil {
		fmt.Printf("There was an error\n")
		os.Exit(1)
	}

	fmt.Printf("it's all good\n")
	conn.Close(context.Background())
}



func getAllStudents(){

}

func addStudent(firstName, lastName, email, enrollmentDate string){

}
