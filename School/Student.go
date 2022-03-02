package School

type Student struct {
	Name   string
	Grade  []string
	Course string
}

func (s *Student) CheckGrades(name string) []string {
	if name == s.Name {
		return s.Grade
	}
	return []string{}
}

func (s *Student) OfferCourse(course string) string {
	if course == s.Course {
		return "You have chosen to offer " + course
	}
	return "Course not available"
}
