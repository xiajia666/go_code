package Struct

type ItTeacherInfo struct {
	Id          int
	Worknum     string
	Name        string
	Gender      string
	Age         int
	Idcard      string
	Workaddress string
	Entrydate   string
	Position    string
	Institution string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

//func (m *ItTeacherInfo) TableName() string {
//	return "it_teacher_infos"
//}

func (ItTeacherInfo) TableName() string {
	return "lt_teacher_infos"
}
