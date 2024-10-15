package message

// 确定一些消息类型
const (
	LoginMsgType            = "LoginMsg"
	LoginResMsgType         = "LoginResMsg"
	RegisterMsgType         = "RegisterMsg"
	RegisterResMsgType      = "RegisterResMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	SmsMsgType              = "SmsMsg"
)

// 定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	// 用户id
	UserId int `json:"userId"`
	// 用户密码
	UserPwd string `json:"userPwd"`
	// 用户名
	UserName string `json:"userName"`
}

type LoginResMsg struct {
	// 返回状态码 500 表示该用户未注册
	Code    int   `json:"code"`
	UserIds []int `json:"userIds"`
	// 返回错误信息
	Error string `json:"error"`
}

type RegisterMsg struct {
	User User `json:"user"`
}

type RegisterResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMsg struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMsg struct {
	Content string `json:"content"`
	User           // 匿名结构体
}
