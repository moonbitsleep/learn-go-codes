package teacher

// Teacher for saving info
type Teacher struct {
	// 工号
	TeacherID string
	Name      string
	Sex       rune
	Age       int
	// 部门
	Branch string
}

// NewTeacher return a teacher
func NewTeacher(ti string, n string, s rune, a int, b string) *Teacher {
	return &Teacher{
		TeacherID: ti,
		Name:      n,
		Sex:       s,
		Age:       a,
		Branch:    b,
	}
}
