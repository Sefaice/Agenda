package entity

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
		d.month == 7 || d.month == 9 || d.month == 10 ||
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

/*
func string2Date(s string) Date {

}
*/
