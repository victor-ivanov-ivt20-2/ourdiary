package scheduler

// 1
import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

var s *gocron.Scheduler

func GetNextWeekDay(t time.Time) string {

	weekDay := make(map[time.Weekday]string)

	weekDay[time.Monday] = "Понедельник"
	weekDay[time.Tuesday] = "Вторник"
	weekDay[time.Wednesday] = "Среда"
	weekDay[time.Thursday] = "Четверг"
	weekDay[time.Friday] = "Пятница"
	weekDay[time.Saturday] = "Суббота"
	weekDay[time.Sunday] = "Воскресенье"

	return weekDay[t.Weekday()]
}

func GetEvenOddWeek(t time.Time) bool {
	_, startWeek := time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC).ISOWeek()
	_, nowWeek := t.ISOWeek()
	return (nowWeek-startWeek)%2 != 0
}

func CreateSchedule() *gocron.Scheduler {
	loc, err := time.LoadLocation("Asia/Yakutsk")
	if err != nil {
		fmt.Println(err)
		s = gocron.NewScheduler(time.UTC)
	} else {
		time.Local = loc
		s = gocron.NewScheduler(loc)
	}
	return s
}

func GetScheduler() *gocron.Scheduler {
	return s
}
