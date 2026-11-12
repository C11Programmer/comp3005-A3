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
	// connects you to the database
	conn, err := pgx.Connect(ctxType, "postgres://postgres:NewPGpass@127.0.0.1:5432/postgres")

	if err != nil {
		fmt.Printf("There was an error connecting to the db\n")
		os.Exit(1)
	}

	students := getAllStudents(conn, ctxType)
	fmt.Println(students)
	conn.Close(context.Background())
}


// gets every student and prints out a table for each one
func getAllStudents(conn *pgx.Conn, ctx context.Context) []students{
	rows, _ := conn.Query(ctx, "SELECT student_id, first_name, last_name, email, enrollment_date FROM public.students")
	students, err := pgx.CollectRows(rows, pgx.RowToStructByName[students])
	if err != nil {
		fmt.Println(err)
		log.Fatal("collecting the rows failed")
	}

	fmt.Printf("|%3s|%12s|%12s|%30s|%17s|\n", "ID", "First Name", "Last Name", "Email", "Enrollment Date")
	for _, v := range students {
		fmt.Printf("|%3d|%12s|%12s|%30s|%17s|\n", v.StudentID, v.FirstName, v.LastName, v.Email, v.EnrollmentDate.Format("Jan 2, 2006"))
	} 

	return students
}

// adds a new student to the database
func addStudent(conn *pgx.Conn, ctx context.Context, firstName, lastName, email string, enrollmentDate time.Time) bool{
	_, err := conn.Exec(ctx, `insert into students (first_name, last_name, email, enrollment_date) values (@first_name, @last_name, @email, @enrollment_date)`,
		pgx.NamedArgs{
			"first_name":      firstName,
			"last_name":       lastName,
			"email":           email,
			"enrollment_date": enrollmentDate,
		})

	return err == nil
}

func updateStudentEmail(conn *pgx.Conn, ctx context.Context, studentId int, newEmail string) bool{
	_, err := conn.Exec(ctx, "update students set email=$1 where student_id=$2", newEmail, studentId)
	return err == nil
}

func deleteStudent(conn *pgx.Conn, ctx context.Context, studentId int) bool {
	_, err := conn.Exec(ctx, "delete from students where student_id=$1", studentId)
	return err == nil
}
