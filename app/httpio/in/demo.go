package in

type ExamPostRegister struct {
	Username string `json:"username" valid:"length(5|20)~姓名不能超过5-20字符串"`
	Password string `json:"password" valid:"matches([a-zA-Z0-9~\!\+\.]{6,20})~密码格式不正确"` // a-z0-9A-Z!~+.
}
