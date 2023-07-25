package pkg

import (
	"QuickAuth/pkg/utils"
	"fmt"
	"github.com/pkg/errors"
	"strings"
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

func amendPath(p string) string {
	if strings.HasPrefix(p, "/applications/") && len(p) > len("/applications/") {
		for i := len("/applications/"); i < len(p); i++ {
			if p[i] == '/' {
				return "/applications/[appId]" + p[i:]
			}
		}
		return "applications/[appId]/"
	}
	return p
}

func TestTemp(t *testing.T) {
	fmt.Println(amendPath("/applications/123"))
	fmt.Println(amendPath("/applications/123/"))
	fmt.Println(amendPath("/applications/123/messages/"))
	fmt.Println(amendPath("/applications/123/messages"))
	fmt.Println(amendPath("/application/messages/"))
	fmt.Println(amendPath("/applications/1asdfasdfa232rl-asdf/messages/"))
}
