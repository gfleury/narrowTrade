// +build !integration
// +build unit

package saxo_oauth2

import (
	"testing"

	"gopkg.in/check.v1"
)

type OauthTest struct {
}

func (s *OauthTest) SetUpSuite(c *check.C) {
}

func (s *OauthTest) TearDownSuite(c *check.C) {
}

func (s *OauthTest) SetUpTest(c *check.C) {
}

var _ = check.Suite(&OauthTest{})

func Test(t *testing.T) {
	check.TestingT(t)
}

func (s *OauthTest) TestLoadConfiguration(c *check.C) {
	data := []byte("{}")

	_, err := LoadConfiguration(data)

	c.Assert(err, check.IsNil)
}
