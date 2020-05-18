package __complexDataType

import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	res := fmt.Sprintf("print: %v", c.Daemon)
	// 	Sprintf format %v with arg c causes recursive String method call
	// res := fmt.Sprintf("print: %v", c)
	fmt.Println(res)
	return res
}

func GetString() {

	c := &ConfigOne{"test daemon"}
	c.String()

}
