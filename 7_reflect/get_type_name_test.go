package __reflect

import "testing"

type testStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestGetTypeName(t *testing.T) {
	value := testStruct{}
	value1 := new(testStruct)

	value2 := int64(10)

	// for k, value := range values{
	//
	// }
	GetTypeName(value)
	GetTypeName(value1)
	GetTypeName(value2)
}
