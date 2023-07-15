package pkg

import (
	"QuickAuth/pkg/tools/utils"
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"testing"
)

func TestDebug(t *testing.T) {
	fmt.Println(utils.GetNoLineUUID())
}

func TestError(t *testing.T) {
	err := errors.Wrap(errors.Wrap(errors.New("first err"), "second err"), "third err")
	fmt.Println(err)
	fmt.Println(errors.Unwrap(err))
}

func TestTemp(t *testing.T) {
	fmt.Println(url.QueryUnescape("https%3A%2F%2Fdouyin.com"))
}
