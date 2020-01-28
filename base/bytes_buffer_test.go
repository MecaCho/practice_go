package base

import "testing"

func TestCreateBuffer(t *testing.T) {
	CreateBuffer()

	BufferWrite()
	BufferWriteToFile()
}
