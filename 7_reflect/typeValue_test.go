package __reflect

import (
	"testing"
	"time"
)

func TestGetType(t *testing.T) {
	GetType()
}

func TestGetValue(t *testing.T) {
	GetValue()
}

func TestGetMethod(t *testing.T) {
	GetMethod(time.Hour)
}

func TestSetValue(t *testing.T) {
	SetValue()
}
