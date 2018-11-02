package entity

import (
	"fmt"
)

type Meeting struct {
	Title         string
	Sponsor       string
	Participators []string
	SDate, EDate  Date
}

func (m Meeting) getSponsor() string {
	return m.Sponsor
}

func (m Meeting) getParticipators() []string {
	return m.Participators
}

func (m Meeting) getTitle() string {
	return m.Title
}

func (m Meeting) getStart() Date {
	return m.SDate
}

func (m Meeting) getEnd() Date {
	return m.EDate
}

func (m Meeting) printMeeting() {
	fmt.Println("MEETING INFO---Title: " + m.Title + " Sponsor: " + m.Sponsor +
		" Participators: " + getParticipatorsStr(m.Participators) +
		" Start Time: " + date2String(m.SDate) + " End Time: " + date2String(m.EDate))
}

func (m Meeting) isParticipatorOrSponsor(u User) bool {
	name := u.getUsername()
	if m.Sponsor == name {
		return true
	}
	for i := 0; i < len(m.Participators); i++ {
		if m.Participators[i] == name {
			return true
		}
	}
	return false
}

//TAKE PARR ALL AS VALID USERS
func (m Meeting) addParticipators(pArr []string) {
	for _, p := range pArr {
		m.Participators = append(m.Participators, p)
	}
}

//TAKE PARR ALL AS VALID USERS
func (m Meeting) deleteParticipators(pArr []string) {
	for _, p := range pArr {
		for i, q := range m.Participators {
			if q == p {
				m.Participators = append(m.Participators[:i], m.Participators[i+1:]...)
			}
		}
	}

	fmt.Println(m.Participators)
}

func getParticipatorsStr(pArr []string) string {
	pStr := ""
	for _, p := range pArr {
		if pStr == "" {
			pStr = p
		} else {
			pStr = pStr + ", " + p
		}
	}
	return pStr
}
