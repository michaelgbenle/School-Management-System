package main

import (
	"fmt"
	"main.go/School"
)

func main() {
	//Principal
	//g := School.Principal{50, false}
	//fmt.Println(g.Query("Donald Trump"))
	//p := School.Principal{29, true}
	//fmt.Println(p.Expel("Vladimir Putin"))

	//Applicant
	//a := School.Applicant{65, 25}
	//fmt.Println(a.ApplyToSchool("Adolf Hitler"))
	//b := School.Applicant{49, 19}
	//fmt.Println(b.ApplyToSchool("Pablo Rudolf"))
	//c := School.Applicant{44, 23}
	//fmt.Println(c.ApplyToSchool("Leo Tolstoy"))

	//Teacher
	//x := School.Teacher{70, 55}
	//fmt.Println(x.Gradestudent("Miles Teller", "Statistics"))
	//y := School.Teacher{70, 55}
	//fmt.Println(y.Sattendance("Miles Teller"))

	//Student
	i := School.Student{"Joe Goldberg", []string{"Mathematics: 79", "English: 88", "French: 91"}, "Joe Goldberg"}
	fmt.Println(i.CheckGrades("Joe Goldberg"))

	j := School.Student{"Joe Goldberg", []string{"Mathematics: 79", "English: 88", "French: 91"}, "English"}
	fmt.Println(j.OfferCourse("Maths"))

}
