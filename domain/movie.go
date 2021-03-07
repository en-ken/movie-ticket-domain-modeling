package domain

import (
	"time"

	"github.com/yut-kt/goholiday"
)

type TimeCategory int

func init() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	time.Local = loc
}

const (
	TimeCategoryWeekDayGeneral TimeCategory = iota + 1
	TimeCategoryWeekDayLate
	TimeCategoryWeekendGeneral
	TimeCategoryWeekendLate
	TimeCategoryMovieDay
)

type Movie struct {
	StartAt      time.Time
	TimeCategory TimeCategory
}

// 平日か否かの判定
func isWeekDay(t time.Time) bool {
	//土日
	if wd := t.Weekday(); wd == time.Sunday || wd == time.Saturday {
		return false
	}

	//祝日
	if goholiday.IsNationalHoliday(t) {
		return false
	}

	return true
}

func isMovieDay(t time.Time) bool {
	return t.Day() == 1 //映画の日か
}

func timeCategory(t time.Time) TimeCategory {
	if isMovieDay(t) {
		return TimeCategoryMovieDay
	}

	//朝4時まではレイトショーとする
	//金曜や日曜の日付超えたレイトショーは？
	if isWeekDay(t) {
		if t.Hour() < 4 {
			return TimeCategoryWeekDayLate
		} else if t.Hour() < 20 {
			return TimeCategoryWeekDayGeneral
		} else {
			return TimeCategoryWeekDayLate
		}
	}

	if t.Hour() < 4 {
		return TimeCategoryWeekendLate
	} else if t.Hour() < 20 {
		return TimeCategoryWeekendGeneral
	} else {
		return TimeCategoryWeekendLate
	}
}

func NewMovie(startAt time.Time) *Movie {
	startAt = startAt.In(time.Local)
	return &Movie{
		StartAt:      startAt,
		TimeCategory: timeCategory(startAt),
	}
}
