package visit

import (
	"applyFK/holiday"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Response 定义了整个响应的结构
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

// Data 定义了响应中的数据部分
type Data struct {
	List  []Visit `json:"list"`
	Total int     `json:"total"`
}

// Visit 定义了访问记录的结构
type Visit struct {
	ID             string      `json:"id"`
	City           string      `json:"city"`
	Build          string      `json:"build"`
	VisitAt        string      `json:"visit_at"` // 使用time.Time来处理时间
	VisitNum       int         `json:"visit_num"`
	StatusName     string      `json:"status_name"`
	VisitorUsers   interface{} `json:"visitor_users"` // 使用interface{}因为原始数据为null
	ApplicantName  string      `json:"applicant_name"`
	ApplicantPhone string      `json:"applicant_phone"`
	CanSwipe       string      `json:"can_swipe"`
	VisitReason    string      `json:"visit_reason"`
	VisitCompany   string      `json:"visit_company"`
}

var joinListUrl = "https://gw.xxxxx.com/visitor/apply/joinList"

// GetApplyJoinList 获取申请记录列表
func (v *Apply) GetApplyJoinList(page int) (*Response, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	if page <= 0 {
		page = 1
	}
	url := fmt.Sprintf("%s?page=%d", joinListUrl, page)
	// 发起GET请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E217 MicroMessenger/6.8.0(0x16080000) NetType/WIFI Language/en Branch/Br_trunk MiniProgramEnv/Mac")
	req.Header.Add("access-token", v.Token)
	req.Header.Add("x-mgw-ticket", v.Ticket)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch applied records: %v", resp.Status)
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应数据
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetHadApplyDays 获取已经申请过的天数
func (v *Apply) GetHadApplyDays() (map[string]int, error) {
	var retDays = make(map[string]int, 0)
	now := holiday.CurrentDayTime()
	page := 2
	for i := 1; i <= page; i++ {
		response, err := v.GetApplyJoinList(i)
		if err != nil {
			return nil, err
		}
		if response.Code != 0 {
			return nil, errors.New(response.Msg)
		}
		for _, visit := range response.Data.List {
			visitTime, err := holiday.ParsedTime(visit.VisitAt)
			if err != nil {
				return nil, err
			}
			if visitTime.Before(now) {
				continue
			}
			key := visitTime.Format(holiday.DayLayout)
			if _, ok := retDays[key]; ok {
				retDays[key] = retDays[key] + 1
			} else {
				retDays[key] = 1
			}
		}
	}
	return retDays, nil

}
