package information

func init() {

}

func Hello(c string) error {
	return nil
}

type PersonInfo struct {
	User_id  int
	Username string `json:"usename"` //输出会小写
	Age      int
	Address  string
	Sex      string
	Work     string
	Email    string
}

type GetBindStruct struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type GetXml struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}
