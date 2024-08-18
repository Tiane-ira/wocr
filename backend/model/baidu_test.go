package model

import (
	"fmt"
	"testing"
)

func TestClearDate(t *testing.T) {
	got := clearDate("202403-1118：48：30")
	fmt.Println(got)
}
