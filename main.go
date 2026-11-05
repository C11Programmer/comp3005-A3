package main

import (
	"fmt"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)


func main(){
	conn, err := pgx.Connect(context.Background(), "url")

	if err != nil {
		fmt.Printf("There was an error")
		os.Exit(1)
	}

	conn.Close(context.Background())
}
