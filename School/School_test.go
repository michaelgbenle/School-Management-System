package School

import (
	"reflect"
	"testing"
)

func TestPrincipal_Expel(t *testing.T) {
	testExpel := []struct {
		input0 Principal
		input1 string

		output string
	}{
		{Principal{29, true}, "Vladimir Putin", "Vladimir Putin is expelled"},
		{Principal{65, true}, "Joe Biden", "Joe Biden does not have any attendance issues"},
		{Principal{70, true}, "Donald Trump", "Donald Trump does not have any attendance issues"},
	}
	for _, value := range testExpel {
		result := value.input0.Expel(value.input1)
		if result != value.output {
			t.Errorf("expected %v, got %v", result, value.output)
		}
	}

}

func TestPrincipal_Query(t *testing.T) {
	testQuery := []struct {
		input0 Principal
		input1 string

		output string
	}{
		{Principal{29, false}, "Vladimir Putin", "Vladimir Putin is hereby queried"},
		{Principal{65, true}, "Goodluck Jonathan", "Kudos Goodluck Jonathan"},
		{Principal{70, true}, "Joseph Stalin", "Kudos Joseph Stalin"},
	}
	for _, value := range testQuery {
		result := value.input0.Query(value.input1)
		if result != value.output {
			t.Errorf("expected %v, got %v", result, value.output)
		}
	}
}

func TestApplicant_ApplyToSchool(t *testing.T) {
	testApplyToSchool := []struct {
		input0 Applicant
		input1 string

		output string
	}{
		{Applicant{65, 25}, "Adolf Hitler", "Adolf Hitler has been offered admission"},
		{Applicant{49, 19}, "Pablo Rudolf", "Pablo Rudolf has been offered Supplementary admission by the Principal"},
		{Applicant{44, 23}, "Leo Tolstoy", "Leo Tolstoy did not meet the cut-off point"},
	}
	for _, value := range testApplyToSchool {
		result := value.input0.ApplyToSchool(value.input1)
		if result != value.output {
			t.Errorf("expected %v, got %v", result, value.output)
		}
	}
}

func TestTeacher_Gradestudent(t *testing.T) {
	testGradestudent := []struct {
		input0 Teacher
		input1 string
		input2 string

		output string
	}{
		{Teacher{70, 55}, "Miles Teller", "Statistics", "Miles Teller passed Statistics"},
		{Teacher{39, 30}, "Paris Hilton", "Mathematics", "Paris Hilton did not pass Mathematics"},
		{Teacher{59, 23}, "Steve Harvey", "English", "Steve Harvey passed English"},
	}
	for _, value := range testGradestudent {
		result := value.input0.Gradestudent(value.input1, value.input2)
		if result != value.output {
			t.Errorf("expected %v, got %v", result, value.output)
		}
	}
}

func TestTeacher_Sattendance(t *testing.T) {
	testSattendance := []struct {
		input0 Teacher
		input1 string
		output string
	}{
		{Teacher{70, 80}, "Miles Teller", "Miles Teller has complete attendance"},
		{Teacher{39, 30}, "Paris Hilton", "Paris Hilton does not have complete attendance"},
		{Teacher{59, 72}, "Steve Harvey", "Steve Harvey has complete attendance"},
	}
	for _, value := range testSattendance {
		result := value.input0.Sattendance(value.input1)
		if result != value.output {
			t.Errorf("expected %v, got %v", result, value.output)
		}
	}
}

func TestStudent_CheckGrades(t *testing.T) {
	testCheckGrades := []struct {
		input0 Student
		input1 string
		output []string
	}{
		{Student{"Joe Goldberg", []string{"Mathematics: 79", "English: 88", "French: 91"}, "English"}, "Joe Goldberg", []string{"Mathematics: 79", "English: 88", "French: 91"}},
		{Student{"Quinn Pecker", []string{"Mathematics: 79", "English: 88", "French: 91"}, "English"}, "Joe Goldberg", []string{}},
		{Student{"Jordan Peterson", []string{"Mathematics: 79", "English: 88", "French: 91"}, "English"}, "Jordan Peterson", []string{"Mathematics: 79", "English: 88", "French: 91"}},
	}
	for _, value := range testCheckGrades {
		result := value.input0.CheckGrades(value.input1)
		if reflect.DeepEqual(result, value.output) != true {
			t.Errorf("expected %v, got %v", result, value.output)
		}
	}
}

func TestStudent_OfferCourse(t *testing.T) {
	testOfferCourse := []struct {
		input0 Student
		input1 string
		output string
	}{
		{Student{"Joe Goldberg", []string{"Mathematics: 79", "English: 88", "French: 91"}, "English"}, "English", "You have chosen to offer English"},
		{Student{"Quinn Pecker", []string{"Mathematics: 79", "English: 88", "French: 91"}, "English"}, "Biology", "Course not available"},
		{Student{"Jordan Peterson", []string{"Mathematics: 79", "English: 88", "French: 91"}, "Psychology"}, "Psychology", "You have chosen to offer Psychology"},
	}
	for _, value := range testOfferCourse {
		result := value.input0.OfferCourse(value.input1)
		if result != value.output {
			t.Errorf("expected %v, got %v", result, value.output)
		}
	}
}
