package pkg

import (
	"QuickAuth/pkg/utils"
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	fmt.Println(utils.GetNoLineUUID())
}
