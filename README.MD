Requirements:
    go 1.26.5 or above

setup and run instructions:

    step 0: connect to a vm with the correct code on it. Here is an example to connect to my database. `ssh -N -L 5432:localhost:5432 student@134.117.132.132`
    step 1: change directory to where main.go is
    step 2: run command `go run .`



db startup script:
CREATE TABLE students (
	student_id      SERIAL PRIMARY KEY,
    first_name      TEXT NOT NULL,
    last_name       TEXT NOT NULL,
    email           TEXT NOT NULL UNIQUE,
    enrollment_date DATE NOT NULL
);

