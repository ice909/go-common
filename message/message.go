package message

// 确定一些消息类型
const (
	LoginMsgType    = "LoginMsg"
	LoginResMsgType = "LoginResMsg"
	RegisterMsgType = "RegisterMsg"
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
	Code int `json:"code"`
	// 返回错误信息
	Error string `json:"error"`
}

type RegisterMsg struct{}
