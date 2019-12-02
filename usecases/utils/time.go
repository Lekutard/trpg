package utils

import (
	"time"
)

//DateFormat Date型の文字列フォーマット
const DateFormat = "2006-01-02"

//DatetimeFormat DateTime型の文字列フォーマット
const DatetimeFormat = "2006-01-02 15:04"

//TimeFormat Time型の文字列フォーマット
const TimeFormat = "15:04"

//DateFormatNoHyphen Date型の文字列フォーマット(yyyyMMddHH)
const DateFormatNoHyphen = "20060102"

//DatetimeSecondFormat DateTimeSeconde型の文字列フォーマット(yyyyMMddHHmmss)
const DatetimeSecondFormat = "20060102150405"

//ParseTime 日付文字列からTimeを返す
func ParseTime(date string, format string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation(format, date, loc)
}

//FormatTime Timeを日付文字列で返す
func FormatTime(t time.Time, format string) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(format)
}

//LoadLocation ロケーションを取得する
func LoadLocation() *time.Location {
	loc, err := time.LoadLocation("Asia/Tokyo") //日本
	if err != nil {
		panic(err)
	}
	return loc
}

//NowTime 現在時刻を返す
func NowTime() time.Time {
	return time.Now().In(LoadLocation())
}

//GetAgeCategoryByBirthDay Get AgeCategory By BirthDay
func GetAgeCategoryByBirthDay(birthday time.Time) uint {
	return (uint(NowTime().Sub(birthday).Hours()/24/365/10) * 10)
}

// AddMonth 指定した分だけの月数を加算する。時間は00:00。
func AddMonth(t time.Time, period int) time.Time {
	year := t.Year()
	day := t.Day()

	newMonth := int(t.Month()) + period
	newLastDay := getLastDay(year, newMonth)

	var newDay int
	if day > newLastDay {
		newDay = newLastDay
	} else {
		newDay = day
	}

	return time.Date(year, time.Month(newMonth), newDay, 0, 0, 0, 0, LoadLocation())
}

// 月の最終日を計算する
func getLastDay(year, month int) int {
	return time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, LoadLocation()).
		AddDate(0, 0, -1).Day()
}
