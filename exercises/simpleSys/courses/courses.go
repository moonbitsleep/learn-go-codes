package courses

import "learning-go/exercises/simpleSys/teacher"

// Course for saving info
type Course struct {
	Name     string
	CourseID int
	Tea      teacher.Teacher
}
