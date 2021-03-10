package stuinfo

import "fmt"

// Student for saving info
type Student struct {
	StudentID string
	Name      string
	Sex       rune
	Age       int
	ClassID   int
	Major     string
}

// NewStudent return an pointer to a student
func NewStudent(si string, n string, s rune, a int, c int, m string) *Student {
	return &Student{
		StudentID: si,
		Name:      n,
		Sex:       s,
		Age:       a,
		ClassID:   c,
		Major:     m,
	}
}

// print Student Info
func (s *Student) String() string {
	return fmt.Sprintf("%-20s%-20s%-10q%-10v%-10d%-20s",
		s.StudentID,
		s.Name,
		s.Sex,
		s.Age,
		s.ClassID,
		s.Major,
	)
}
