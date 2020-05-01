package models

import (
	"time"
)

type PreserveDetail struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	EventTime  time.Time `json:"event_time" xorm:"not null DATETIME"`
	Uv         int       `json:"uv" xorm:"INT(11)"`
	NewUv      int       `json:"new_uv" xorm:"INT(11)"`
	Day1Active int       `json:"day1_active" xorm:"INT(11)"`
	Day2Active int       `json:"day2_active" xorm:"INT(11)"`
	Day3Active int       `json:"day3_active" xorm:"INT(11)"`
	Day4Active int       `json:"day4_active" xorm:"INT(11)"`
	Day5Active int       `json:"day5_active" xorm:"INT(11)"`
	Day6Active int       `json:"day6_active" xorm:"INT(11)"`
	WeekActive int       `json:"week_active" xorm:"INT(11)"`
	Detail     string    `json:"detail" xorm:"TEXT"`
	LastUpdate time.Time `json:"last_update" xorm:"DATETIME"`
}
type PreserveDetailWeb struct {
	Id         string       `json:"id" `
	EventTime  string `json:"event_time" `
	Uv         string       `json:"uv" `
	NewUv      string       `json:"new_uv" `
	Day1Active string       `json:"day1_active"`
	Day2Active string       `json:"day2_active" `
	Day3Active string       `json:"day3_active" `
	Day4Active string       `json:"day4_active"`
	Day5Active string       `json:"day5_active" `
	Day6Active string       `json:"day6_active"`
	WeekActive string       `json:"week_active" `
	Detail     string    `json:"detail" `
	LastUpdate string `json:"last_update" `
}
func (t PreserveDetail) TableName() string {
	return "preserve_detail"
}

//Cols ...用户web显示使用
func (t PreserveDetail) Cols() []map[string]string {
	var cols []map[string]string
	col_event_day := map[string]string{
		"name": "事件日期",
		"key":  "event_day",
	}
	cols = append(cols, col_event_day)
	col_uv := map[string]string{
		"name": "uv",
		"key":  "uv",
	}
	cols = append(cols, col_uv)

	col_new_uv := map[string]string{
		"name": "新增用户",
		"key":  "new_uv",
	}
	cols = append(cols, col_new_uv)
	//次日留存
	col_day1_active := map[string]string{
		"name": "1日留存",
		"key":  "day1_active",
	}
	cols = append(cols, col_day1_active)

	//二日留存
	col_day2_active := map[string]string{
		"name": "2日留存",
		"key":  "day2_active",
	}
	cols = append(cols, col_day2_active)

	//三日留存
	col_day3_active := map[string]string{
		"name": "3日留存",
		"key":  "day3_active",
	}
	cols = append(cols, col_day3_active)

	//四日留存
	col_day4_active := map[string]string{
		"name": "4日留存",
		"key":  "day4_active",
	}
	cols = append(cols, col_day4_active)

	//五日留存
	col_day5_active := map[string]string{
		"name": "5日留存",
		"key":  "day5_active",
	}
	cols = append(cols, col_day5_active)

	//六日留存
	col_day6_active := map[string]string{
		"name": "6日留存",
		"key":  "day6_active",
	}
	cols = append(cols, col_day6_active)

	//周留存
	col_week_active := map[string]string{
		"name": "周留存",
		"key":  "week_active",
	}
	cols = append(cols, col_week_active)

	col_last_update := map[string]string{
		"name": "更新时间",
		"key":  "last_update",
	}
	cols = append(cols, col_last_update)
	return cols
}
