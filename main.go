package main

import (
	"fmt"
	"context"
	"os"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type students struct {
	StudentID int32 `db:"student_id"`
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
	Email string `db:"email"`
	EnrollmentDate time.Time `db:"enrollment_date"`
}

func main(){
	ctxType := context.Background()
	conn, err := pgx.Connect(ctxType, "postgres://postgres:NewPGpass@127.0.0.1:5432/postgres")

	if err != nil {
		fmt.Printf("There was an error\n")
		os.Exit(1)
	}

	students := getAllStudents(conn, ctxType)
	fmt.Println(students)
	conn.Close(context.Background())
}



func getAllStudents(conn *pgx.Conn, ctx context.Context) []students{
	rows, _ := conn.Query(ctx, "SELECT student_id, first_name, last_name, email, enrollment_date FROM public.students")
	students, err := pgx.CollectRows(rows, pgx.RowToStructByName[students])
	if err != nil {
		fmt.Println(err)
		log.Fatal("collecting the rows failed")
	}

	fmt.Println()
	for _, v := range students {
		fmt.Printf("|%3d|%10s|%10s|%30s|%12s|\n", v.StudentID, v.FirstName, v.LastName, v.Email, v.EnrollmentDate.Format("Jan 2, 2006"))
	} 

	return students
}

func addStudent(firstName, lastName, email, enrollmentDate string) bool{


	return false
}

func updateStudentEmail(studentId int, newEmail string) bool{

	return false
}

func deleteStudent(studentId int) bool {
	return false 
}
