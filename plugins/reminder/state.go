package main

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
)

const (
	taskAddStateSelectType = iota
	taskAddStateSelectYear
	taskAddStateSelectMonth
	taskAddStateSelectWeekDay
	taskAddStateSelectDay
	taskAddStateSelectHour
	taskAddStateSelectMinute
	taskAddStateSelectSecond
	taskAddStateSelectName
	taskAddStateSelectFinish
)

const (
	taskDelStateSelectTask = iota
	taskDelStateFinish
)

const (
	taskTypeOnce    = "Once"
	taskTypeYearly  = "Yearly"
	taskTypeMonthly = "Monthly"
	taskTypeWeekly  = "Weekly"
	taskTypeDaily   = "Dialy"
	taskTypeHourly  = "Hourly"
)

const (
	weekDayMonday    = "Monday"
	weekDayTuesday   = "Tuesday"
	weekDayWednesday = "Wednesday"
	weekDayThursday  = "Thursday"
	weekDayFriday    = "Friday"
	weekDaySaturday  = "Saturday"
	weekDaySunday    = "Sunday"
)

const (
	monthJan = "Januray"
	monthFeb = "February"
	monthMar = "March"
	monthApr = "April"
	monthMay = "May"
	monthJun = "June"
	monthJul = "July"
	monthAug = "August"
	monthSep = "September"
	monthOct = "October"
	monthNov = "November"
	monthDec = "December"
)

const (
	globalCancel = "Cancel"
)

var (
	taskTypes            = []string{taskTypeOnce, taskTypeYearly, taskTypeMonthly, taskTypeWeekly, taskTypeDaily, taskTypeHourly}
	weekDays             = []string{weekDayMonday, weekDayTuesday, weekDayWednesday, weekDayThursday, weekDayFriday, weekDaySaturday, weekDaySunday}
	months               = []string{monthJan, monthFeb, monthMar, monthApr, monthMay, monthJun, monthJul, monthAug, monthSep, monthOct, monthNov, monthDec}
	errUserStateNotFound = errors.New("user state not found in database")
)

// ReminderUserState struct for store user command queue state
type ReminderUserState struct {
	gorm.Model
	ChatID    int64
	Command   string
	Type      string
	Year      int
	Month     int
	Day       int
	Hour      int
	Minute    int
	Second    int
	State     int
	WeekDay   int
	Name      string
	Text      string
	Timestamp time.Time
}

// initReminderUserState function initialize user state object
func initReminderUserState(chatID int64, cmd string) (rus *ReminderUserState) {
	rus = &ReminderUserState{
		ChatID:  chatID,
		Command: cmd,
	}
	return
}

// Save function store user state to database
func (rus *ReminderUserState) Save() (err error) {
	db := ctx.GetDB()
	temp := &ReminderUserState{
		ChatID:  rus.ChatID,
		Command: rus.Command,
	}
	rus.Timestamp = time.Now()
	if err = db.First(temp).Error; err == nil && err == gorm.ErrRecordNotFound {
		return ctx.GetDB().Create(&rus).Error
	}
	return ctx.GetDB().Save(&rus).Error
}

// getReminderUserState function return user state from database
func getReminderUserState(chatID int64, cmd string) (rus *ReminderUserState, err error) {
	rus = initReminderUserState(chatID, cmd)
	if err = ctx.GetDB().Find(rus).Error; err != nil {
		return nil, nil
	} else if err != nil {
		return
	}
	if time.Since(rus.Timestamp) > defaultUserStateTimeout {
		rus = nil
	}

	return
}

func getReminderUserStates(chatID int64) (rusA, rusD *ReminderUserState, err error) {
	if rusA, err = getReminderUserState(chatID, taskAddCommand); err != nil {
		return
	}
	if rusD, err = getReminderUserState(chatID, taskDelCommand); err != nil {
		return
	}
	return
}

// Delete function remove user state from database after timeout or complete command
func (rus *ReminderUserState) Delete() (err error) {
	err = ctx.GetDB().Delete(rus).Error
	return
}

// ToUserTask convert user state to user task object
func (rus *ReminderUserState) ToUserTask() (ut *UserTask) {
	ut = &UserTask{
		ChatID:  rus.ChatID,
		Name:    rus.Name,
		Type:    rus.Type,
		Year:    rus.Year,
		Month:   rus.Month,
		Day:     rus.Day,
		Hour:    rus.Hour,
		Minute:  rus.Minute,
		Second:  rus.Second,
		WeekDay: rus.WeekDay,
		Text:    rus.Text,
	}
	return
}

// GetLastDay function return last day of a month in year
func (rus *ReminderUserState) GetLastDay() int {
	return getLastDay(rus.Year, rus.Month)
}

func deleteUserStates(ctx *context.MultiBotContext, chatID int64) (err error) {
	if rusA, rusD, err := getReminderUserStates(chatID); err != nil {
		ctx.Log().Errorf("Unable to get user states: %s", err)
	} else {
		if rusA != nil {
			if err = ctx.GetDB().Delete(&rusA).Error; err != nil {
				return err
			}
		}
		if rusD != nil {
			if err = ctx.GetDB().Delete(&rusD).Error; err != nil {
				return err
			}
		}
	}
	return
}
