package entity

type Meeting struct {
	sponsor      string
	participator []string
	title        string
	start, end   Date
}

func (m Meeting) getSponsor() string {
	return m.sponsor
}

func (m Meeting) getParticipator() []string {
	return m.participator
}

func (m Meeting) getTitle() string {
	return m.title
}

func (m Meeting) getStart() Date {
	return m.start
}

func (m Meeting) getEnd() Date {
	return m.end
}

func (m Meeting) IsParticipator(u User) bool {
	name := u.getUsername()
	if m.sponsor == name {
		return true
	}
	for i := 0; i < len(m.participator); i++ {
		if m.participator[i] == name {
			return true
		}
	}
	return false
}
