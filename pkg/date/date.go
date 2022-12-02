package date

import (
	"time"
)

const DateLayout = "2006/1/2"
const JPDateLayout = "2006年1月"
const TimeLayout = "2006/1/2 15:04:05"

var (
	JST *time.Location
)

func init() {
	JST, _ = time.LoadLocation("Asia/Tokyo")
}

// /////////////////////////////////////////////////////////
// タイムゾーンを変換する関数///////////////////////////////////
// /////////////////////////////////////////////////////////
// TZをJSTに変換
func ToJST(t time.Time) time.Time {
	return t.In(JST)
}

// TZをUTCに変換
func ToUTC(t time.Time) time.Time {
	return t.In(time.UTC)
}

///////////////////////////////////////////////////////////
//日時を操作する関数//////////////////////////////////////////
///////////////////////////////////////////////////////////

// 引数の24時間前を返す
func Yesterday(t time.Time) time.Time {
	return t.AddDate(0, 0, -1)
}

// 引数の24時間後を返す
func Tomorrow(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

// 引数の時刻を00:00:00に変換する
func StartOfTheDay(datetime time.Time) time.Time {
	return time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, datetime.Location())
}

// 引数の時刻を23:59:59に変換する
func EndOfTheDay(datetime time.Time) time.Time {
	return StartOfTheDay(datetime).Add(24*time.Hour - 1*time.Nanosecond)
}

// 引数の日時の前日23:59:59を返す
func EndOfYesterday(datetime time.Time) time.Time {
	return EndOfTheDay(Yesterday(datetime))
}

// 引数の月の一日00:00:00を返す。e.g. 2022/06/11 12:34:50 => 2022/06/01 00:00:00
func StartOfTheMonth(datetime time.Time) time.Time {
	return time.Date(datetime.Year(), datetime.Month(), 1, 0, 0, 0, 0, datetime.Location())
}

// 引数の月の最終日23:59:59を返す。e.g. 2022/06/11 12:34:50 => 2022/06/30 23:59:59
func EndOfTheMonth(datetime time.Time) time.Time {
	return StartOfTheMonth(datetime).AddDate(0, 1, 0).Add(-1 * time.Nanosecond)
}

// 引数の月の翌月の一日00:00:00を返す。e.g. 2022/06/11 12:34:50 => 2022/07/01 00:00:00
func StartOfNextMonth(datetime time.Time) time.Time {
	return time.Date(datetime.Year(), datetime.Month()+1, 1, 0, 0, 0, 0, datetime.Location())
}

///////////////////////////////////////////////////////////
//その他便利な関数（関数が増えてきたら分別する）///////////////////
///////////////////////////////////////////////////////////

// 引数の日付が今日と同じかどうかをチェックする
func IsToday(datetime time.Time) bool {
	tz := datetime.Location()
	now := ToUTC(time.Now()).In(tz)
	return datetime.Year() == now.Year() && datetime.Month() == now.Month() && datetime.Day() == now.Day()
}

// 引数の日時が未来かどうかをチェックする
func IsFuture(datetime time.Time) bool {
	return ToUTC(datetime).After(ToUTC(time.Now()))
}
