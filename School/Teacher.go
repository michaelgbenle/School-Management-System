package School

type Teacher struct {
	Grade      int
	Attendance int
}

func (t *Teacher) Gradestudent(name string, course string) string {

	if t.Grade >= 40 {
		return name + " passed " + course
	}
	return name + " did not pass " + course
}

func (t *Teacher) Sattendance(name string) string {

	if t.Attendance > 70 {
		return name + " has complete attendance"
	}
	return name + " does not have complete attendance"
}
