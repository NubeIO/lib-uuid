package uuid

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	u := UUID()
	fmt.Println(u)
	u = UUID("uid")
	fmt.Println(u)
}

func TestShortUUID(t *testing.T) {
	u := ShortUUID()
	fmt.Println(u)
	u = ShortUUID("uid")
	fmt.Println(u)
}

func TestUniqueNumber(t *testing.T) {
	u := UniqueNumber()
	fmt.Println(u)
	u = UniqueNumber()
	fmt.Println(u)
}
