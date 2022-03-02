package School

type Applicant struct {
	Grade int
	Age   int
}

func (a *Applicant) ApplyToSchool(name string) string {

	if a.Grade > 50 {
		return name + " has been offered admission"
	} else if a.Age < 20 && a.Grade > 40 && a.Grade < 50 {
		return name + " has been offered Supplementary admission by the Principal"
	}
	return name + " did not meet the cut-off point"
}
