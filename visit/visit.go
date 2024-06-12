package visit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Apply struct {
	Token  string      `json:"token"`
	Ticket string      `json:"ticket"`
	Body   ApplyRecord `json:"body"`
}
type ApplyRecord struct {
	AppliedUserID string `json:"visited_user_id"` // 被访问者的 ID
	AppliedName   string `json:"visited_name"`    // 被访问者的姓名
	City          string `json:"city"`            // 访问者所在城市
	CityID        int    `json:"city_id"`         // 访问者所在城市 ID
	Build         string `json:"build"`           // 访问者所在楼栋
	BuildID       int    `json:"build_id"`        // 访问者所在楼栋 ID
	ApplyDate     string `json:"visit_date"`      // 这里假设日期格式为 YYYY/MM/DD
	ApplyTime     string `json:"visit_time"`      // 这里假设时间格式为 HH:MM
	ApplyNum      int    `json:"visit_num"`       // 访问次数
	VerifiedNum   string `json:"verified_num"`    //	被验证次数
	AppliedPhone  string `json:"visited_phone"`   // 被访问者的手机号
	ApplyReason   string `json:"visit_reason"`    // 访问原因
	ApplyCompany  string `json:"visit_company"`   // 访问公司
}

func NewApply() *Apply {
	return &Apply{
		Body: ApplyRecord{
			City:         "xxx",
			CityID:       28,
			Build:        "xxx",
			BuildID:      53,
			ApplyTime:    "09:00",
			ApplyNum:     1,
			ApplyReason:  "xxx",
			ApplyCompany: "xxx",
			AppliedPhone: "xxx", // 假设被访问者的手机号为 xxx xxx
		},
	}
}

func (v *Apply) Apply() error {
	fmt.Printf("开始申请%s访客码\n", v.Body.ApplyDate)
	jsonValue, _ := json.Marshal(v.Body)
	payload := bytes.NewBuffer(jsonValue)

	// 替换以下URL为您要发送POST请求的服务器地址
	url := "https://gw.xxxxx.com/visitor/apply/save"

	// 创建请求
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 设置请求头
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E217 MicroMessenger/6.8.0(0x16080000) NetType/WIFI Language/en Branch/Br_trunk MiniProgramEnv/Mac")
	req.Header.Add("access-token", v.Token)
	req.Header.Add("x-mgw-ticket", v.Ticket)
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	// 打印响应状态码
	fmt.Printf("申请成功\n")
	// 打印响应体
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
	return nil
}

func (v *Apply) SetToken(token string) {
	v.Token = token
}
func (v *Apply) SetTicket(ticket string) {
	v.Ticket = ticket
}
func (v *Apply) SetAppliedUserID(userID string) {
	v.Body.AppliedUserID = userID
}

func (v *Apply) SetAppliedName(name string) {
	v.Body.AppliedName = name
}

func (v *Apply) SetCity(city string) {
	v.Body.City = city
}

func (v *Apply) SetCityID(cityID int) {
	v.Body.CityID = cityID
}

func (v *Apply) SetBuild(build string) {
	v.Body.Build = build
}

func (v *Apply) SetBuildID(buildID int) {
	v.Body.BuildID = buildID
}

func (v *Apply) SetApplyDate(date string) {
	v.Body.ApplyDate = date
}

func (v *Apply) SetApplyTime(time string) {
	v.Body.ApplyTime = time
}

func (v *Apply) SetApplyNum(num int) {
	v.Body.ApplyNum = num
}

func (v *Apply) SetVerifiedNum(num string) {
	v.Body.VerifiedNum = num
}

func (v *Apply) SetAppliedPhone(phone string) {
	v.Body.AppliedPhone = phone
}

func (v *Apply) SetApplyReason(reason string) {
	v.Body.ApplyReason = reason
}

func (v *Apply) SetApplyCompany(company string) {
	v.Body.ApplyCompany = company
}
