package Struct

type ItTeacherInfo struct {
	Id          int    `json:"id"`
	Worknum     string `json:"worknum"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Idcard      string `json:"idcard"`
	Workaddress string `json:"workaddress"`
	Entrydate   string `json:"entrydate"`
	Position    string `json:"position"`
	Institution string `json:"institution"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}

//func (m *ItTeacherInfo) TableName() string {
//	return "it_teacher_infos"
//}

func (ItTeacherInfo) TableName() string {
	return "lt_teacher_infos"
}
