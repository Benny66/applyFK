package visit

import (
	"applyFK/holiday"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func (v *Apply) ApplyDate(date string, num int) error {
	if num <= 0 {
		return errors.New("num should be greater than 0")
	}
	for i := 0; i < num; i++ {
		v.SetApplyDate(date)
		err := v.Apply()
		if err != nil {
			fmt.Printf("时间%s 第%d次申请失败\n", date, i+1)
			return err
		}
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("时间%s 第%d次申请成功\n", date, i+1)
	}
	return nil
}

func (v *Apply) ApplyDays(applyDayNum int, endDate string) error {
	fmt.Printf("开始申请%d天\n", applyDayNum)
	// 获取已申请日期列表
	hadApplyDays, err := v.GetHadApplyDays(20)
	if err != nil {
		fmt.Printf("获取已申请天数失败%s", err)
		return err
	}
	day := 0 // 第几天
	hadDayNum := 0
	for {
		currentDate := holiday.CurrentDayStr(day)
		if endDate == currentDate {
			fmt.Printf("时间%s 已到达结束日期，退出\n", currentDate)
			break
		}
		day += 1
		applyNum := 3 // 每天申请3次
		if _, ok := hadApplyDays[currentDate]; ok {
			days := hadApplyDays[currentDate]
			if days >= 3 {
				fmt.Printf("时间%s 已申请3天，跳过\n", currentDate)
				continue
			}
			applyNum = 3 - days
			fmt.Printf("时间%s 已申请%d次,还需要%d次\n", currentDate, days, applyNum)
		}
		isHoliday := holiday.IsHoliday(currentDate)
		if isHoliday {
			fmt.Printf("时间%s 跳过当前周期的节假日\n", currentDate)
			continue
		}
		fmt.Printf("时间%s 开始申请\n", currentDate)
		err = v.ApplyDate(currentDate, applyNum)
		if err != nil {
			return err
		}
		hadDayNum += 1
		if hadDayNum >= applyDayNum {
			break
		}
		rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second) // 暂停相应的秒数
	}
	return nil
}
