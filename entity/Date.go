package entity

import (
	"strconv"
	"strings"
)

type Date struct {
	year, month, day, hour, minute int
}

func isVaildDate(d Date) bool {
	if d.year < 2000 || d.year > 9999 || d.month < 1 ||
		d.month > 12 || d.day < 1 || d.hour > 23 ||
		d.hour < 0 || d.minute < 0 || d.minute > 59 {
		return false
	}
	if d.month == 1 || d.month == 3 || d.month == 5 ||
		d.month == 7 || d.month == 8 || d.month == 10 ||
		d.month == 12 {
		if d.day > 31 {
			return false
		}
	} else if d.month != 2 {
		if d.day > 30 {
			return false
		}
	} else {
		if (d.year%4 == 0 && d.year%100 != 0) || d.year%400 == 0 {
			if d.day > 29 {
				return false
			}
		} else if d.day > 28 {
			return false
		}
	}
	return true
}

//1949-10-1-12-0, TAKE INPUT AS VALID
func date2String(d Date) string {
	str := strconv.Itoa(d.year) + "-" + strconv.Itoa(d.month) + "-" + strconv.Itoa(d.day) +
		"-" + strconv.Itoa(d.hour) + "-" + strconv.Itoa(d.minute)
	return str
}

func string2ValidDate(s string) (bool, Date) {
	//format check with four'-'
	s1 := strings.Split(s, "-")
	if len(s1) != 5 {
		return false, Date{}
	}
	//every str should be int
	year, yErr := strconv.Atoi(s1[0])
	month, mErr := strconv.Atoi(s1[1])
	day, dErr := strconv.Atoi(s1[2])
	hour, hErr := strconv.Atoi(s1[3])
	minute, fErr := strconv.Atoi(s1[4])
	if yErr != nil || mErr != nil || dErr != nil || hErr != nil || fErr != nil {
		return false, Date{}
	}
	d := Date{year, month, day, hour, minute}
	//check valid date
	if !isVaildDate(d) {
		return false, Date{}
	}
	return true, d
}

//d1 < d2 true, d1 >= d2 false, TAKE INPUT AS VALID
func compareDate(d1 Date, d2 Date) bool {
	if d1.year > d2.year {
		return false
	} else if d1.year == d2.year {
		if d1.month > d2.month {
			return false
		} else if d1.month == d2.month {
			if d1.day > d2.day {
				return false
			} else if d1.day == d2.day {
				if d1.hour > d2.hour {
					return false
				} else if d1.hour == d2.hour {
					if d1.minute >= d2.minute {
						return false
					}
				}
			}
		}
	}
	return true
}

func equalDate(d1 Date, d2 Date) bool {
	return d1.year == d2.year && d1.month == d2.month && d1.day == d2.day &&
		d1.hour == d2.hour && d1.minute == d2.minute
}
