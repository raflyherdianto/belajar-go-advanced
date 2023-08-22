package main

import (
	"database/sql"
	_ "database/sql"
	"net/http"
	"reflect"

	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type newStudent struct {
	Student_id      int    `json:"student_id" binding : "required"`
	Student_name    string `json:"student_name" binding : "required"`
	Student_age     int    `json:"student_age" binding : "required"`
	Student_address string `json:"student_address" binding : "required"`
	Student_phone   string `json:"student_phone" binding : "required"`
}

func rowToStruct(rows *sql.Rows, dest interface{}) error {
	destv := reflect.ValueOf(dest).Elem()

	args := make([]interface{}, destv.Type().Elem().NumField())

	for rows.Next() {
		rowp := reflect.New(destv.Type().Elem())
		rowv := rowp.Elem()

		for i := 0; i < rowv.NumField(); i++ {
			args[i] = rowv.Field(i).Addr().Interface()
		}

		if err := rows.Scan(args...); err != nil {
			return err
		}

		destv.Set(reflect.Append(destv, rowv))
	}

	return nil
}

func getAllStudent(c *gin.Context, db *sql.DB) {
	var students []newStudent

	row, err := db.Query("SELECT * FROM students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	rowToStruct(row, &students)

	if students == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": students,
	})
}

func getStudentById(c *gin.Context, db *sql.DB) {
	var student []newStudent

	studentId := c.Param("student_id")

	row, err := db.Query("SELECT * FROM students WHERE student_id = $1", studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	rowToStruct(row, &student)
	if student == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func putStudent(c *gin.Context, db *sql.DB) {
	var student newStudent

	studentId := c.Param("student_id")

	if c.Bind(&student) == nil {
		_, err := db.Exec("UPDATE students SET student_name = $1, student_age = $2, student_address = $3, student_phone = $4 WHERE student_id = $5", student.Student_name, student.Student_age, student.Student_address, student.Student_phone, studentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Student updated successfully!",
		})
	}

}

func postStudent(c *gin.Context, db *sql.DB) {
	var student newStudent

	if c.Bind(&student) == nil {
		_, err := db.Exec("INSERT INTO students (student_id, student_name, student_age, student_address, student_phone) VALUES ($1, $2, $3, $4, $5)", student.Student_id, student.Student_name, student.Student_age, student.Student_address, student.Student_phone)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Student added successfully!",
		})
	}
}

func deleteStudent(c *gin.Context, db *sql.DB) {
	studentId := c.Param("student_id")

	_, err := db.Exec("DELETE FROM students WHERE student_id = $1", studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Student deleted successfully!",
	})
}

func setupRouter() *gin.Engine {
	conn := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.POST("/student", func(ctx *gin.Context) {
		postStudent(ctx, db)
	})

	r.GET("/student", func(ctx *gin.Context) {
		getAllStudent(ctx, db)
	})

	r.GET("/student/:student_id", func(ctx *gin.Context) {
		getStudentById(ctx, db)
	})

	r.PUT("/student/:student_id", func(ctx *gin.Context) {
		putStudent(ctx, db)
	})

	r.DELETE("/student/:student_id", func(ctx *gin.Context) {
		deleteStudent(ctx, db)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
