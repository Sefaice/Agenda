package entity

import (
	"strconv"
	"strings"
)

type Date struct {
	Year, Month, Day, Hour, Minute int
}

func isVaildDate(d Date) bool {
	if d.Year < 2000 || d.Year > 9999 || d.Month < 1 ||
		d.Month > 12 || d.Day < 1 || d.Hour > 23 ||
		d.Hour < 0 || d.Minute < 0 || d.Minute > 59 {
		return false
	}
	if d.Month == 1 || d.Month == 3 || d.Month == 5 ||
		d.Month == 7 || d.Month == 8 || d.Month == 10 ||
		d.Month == 12 {
		if d.Day > 31 {
			return false
		}
	} else if d.Month != 2 {
		if d.Day > 30 {
			return false
		}
	} else {
		if (d.Year%4 == 0 && d.Year%100 != 0) || d.Year%400 == 0 {
			if d.Day > 29 {
				return false
			}
		} else if d.Day > 28 {
			return false
		}
	}
	return true
}

//1949-10-1-12-0, TAKE INPUT AS VALID
func date2String(d Date) string {
	str := strconv.Itoa(d.Year) + "-" + strconv.Itoa(d.Month) + "-" + strconv.Itoa(d.Day) +
		"-" + strconv.Itoa(d.Hour) + "-" + strconv.Itoa(d.Minute)
	return str
}

func string2ValidDate(s string) (bool, Date) {
	//format check with four'-'
	s1 := strings.Split(s, "-")
	if len(s1) != 5 {
		return false, Date{}
	}
	//every str should be int
	Year, yErr := strconv.Atoi(s1[0])
	Month, mErr := strconv.Atoi(s1[1])
	Day, dErr := strconv.Atoi(s1[2])
	Hour, hErr := strconv.Atoi(s1[3])
	Minute, fErr := strconv.Atoi(s1[4])
	if yErr != nil || mErr != nil || dErr != nil || hErr != nil || fErr != nil {
		return false, Date{}
	}
	d := Date{Year, Month, Day, Hour, Minute}
	//check valid date
	if !isVaildDate(d) {
		return false, Date{}
	}
	return true, d
}

//d1 < d2 true, d1 >= d2 false, TAKE INPUT AS VALID
func compareDate(d1 Date, d2 Date) bool {
	if d1.Year > d2.Year {
		return false
	} else if d1.Year == d2.Year {
		if d1.Month > d2.Month {
			return false
		} else if d1.Month == d2.Month {
			if d1.Day > d2.Day {
				return false
			} else if d1.Day == d2.Day {
				if d1.Hour > d2.Hour {
					return false
				} else if d1.Hour == d2.Hour {
					if d1.Minute >= d2.Minute {
						return false
					}
				}
			}
		}
	}
	return true
}

func equalDate(d1 Date, d2 Date) bool {
	return d1.Year == d2.Year && d1.Month == d2.Month && d1.Day == d2.Day &&
		d1.Hour == d2.Hour && d1.Minute == d2.Minute
}
