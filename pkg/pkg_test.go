package pkg

import (
	"QuickAuth/pkg/utils"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestDebug(t *testing.T) {
	fmt.Println(utils.GetNoLineUUID())
}

func TestError(t *testing.T) {
	fmt.Println(errors.Wrap(errors.Wrap(errors.New("first err"), "second err"), "third err"))
}
