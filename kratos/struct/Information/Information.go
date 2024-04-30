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
