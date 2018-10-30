package entity

import (
	"fmt"
)

type Meeting struct {
	title         string
	sponsor       string
	participators []string
	sDate, eDate  Date
}

func (m Meeting) getSponsor() string {
	return m.sponsor
}

func (m Meeting) getParticipators() []string {
	return m.participators
}

func (m Meeting) getTitle() string {
	return m.title
}

func (m Meeting) getStart() Date {
	return m.sDate
}

func (m Meeting) getEnd() Date {
	return m.eDate
}

func (m Meeting) printMeeting() {
	fmt.Println("MEETING INFO---Title: " + m.title + " Sponsor: " + m.sponsor +
		" Participators: " + getParticipatorsStr(m.participators) +
		" Start Time: " + date2String(m.sDate) + " End Time: " + date2String(m.eDate))
}

func (m Meeting) isParticipatorOrSponsor(u User) bool {
	name := u.getUsername()
	if m.sponsor == name {
		return true
	}
	for i := 0; i < len(m.participators); i++ {
		if m.participators[i] == name {
			return true
		}
	}
	return false
}

//TAKE PARR ALL AS VALID USERS
func (m Meeting) addParticipators(pArr []string) {
	for _, p := range pArr {
		m.participators = append(m.participators, p)
	}
}

//TAKE PARR ALL AS VALID USERS
func (m Meeting) deleteParticipators(pArr []string) {
	for _, p := range pArr {
		for i, q := range m.participators {
			if q == p {
				m.participators = append(m.participators[:i], m.participators[i+1:]...)
			}
		}
	}

	fmt.Println(m.participators)
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
