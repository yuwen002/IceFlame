package system_master

type CreateEmployeeInput struct {
	Password     string
	Name         string
	Tel          string
	WxUnionId    string
	WxOpenId     string
	InviteCode   string
	InviteQrcode string
	RealNameType uint8
	Gender       uint8
	Avatar       string
	Nickname     string
}
