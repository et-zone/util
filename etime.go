package util

import (
	"errors"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
	TIME = tme(0)
)
type tme int

func (tm tme)NewNowTimeFormat(fmt string) string {
	return time.Now().Format(fmt)
}

func  (tm tme)TimeStringToTime(timeString string, fmt string) (time.Time, error) {
	return time.Parse(fmt, timeString)
}

func  (tm tme)TimeToString(t *time.Time, format string) string {
	if t == nil {
		return ""
	}

	return t.Format(format)
}

func  (tm tme)TimeToStringFormat(t *time.Time, format string) string {
	if t == nil {
		return ""
	}
	return t.Format(format)
}

func  (tm tme)TimeStringToUnix(s string, fmt string) (int64, error) {
	t, err := time.Parse(fmt, s)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

//return like  2006-01-02
func  (tm tme)DateTimeToDataString(s string) (string, error) {
	if s == "" {
		return "", errors.New("s not datetime")
	}
	return s[:10], nil
}

//return like  15:04:05
func  (tm tme)DateTimeToTimeString(s string) (string, error) {
	if s == "" {
		return "", errors.New("s not datetime")
	}
	return s[11:], nil
}

/*
	Sunday Weekday = 0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
*/
func  (tm tme)DateTimeToWeek(s string, fmt string) (time.Weekday, error) {
	t, err := time.Parse(fmt, s)
	if err != nil {
		return 0, err
	}
	return t.Weekday(), nil
}
