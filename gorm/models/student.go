package models

type Student struct {
	Student_id      int    `json:"student_id" binding : "required"`
	Student_name    string `json:"student_name" binding : "required"`
	Student_age     int    `json:"student_age" binding : "required"`
	Student_address string `json:"student_address" binding : "required"`
	Student_phone   string `json:"student_phone" binding : "required"`
}
