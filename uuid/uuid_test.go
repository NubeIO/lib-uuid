package uuid

import (
	"fmt"
	"testing"
)

func TestSmallUUID(t *testing.T) {
	gotUuid := ShortUUID()
	fmt.Println(gotUuid)
	gotUuid = ShortUUID("hey")
	fmt.Println(gotUuid)
	gotUuid = ShortUUID("hey")
	fmt.Println(gotUuid)
}
