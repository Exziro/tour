package timer

import (
	//"log"
	"time"
)

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai") //此处利用时区设定
	return time.Now().In(location)
}
func GetcalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil

}
