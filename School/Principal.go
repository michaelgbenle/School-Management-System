package School

type Principal struct {
	Sattendance  int
	Tpunctuality bool
}

func (p *Principal) Expel(Sname string) string {
	if p.Sattendance < 30 {
		return Sname + " is expelled"
	}
	return Sname + " does not have any attendance issues"
}

func (p *Principal) Query(tname string) string {
	if p.Tpunctuality == false {
		result := tname + " is hereby queried"
		return result
	}
	return "Kudos " + tname
}
