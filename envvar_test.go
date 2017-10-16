package bashir

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func TestEnvVarMaker(t *testing.T) { TestingT(t) }

type EnvVarMakerSuite struct{}

var _ = Suite(&EnvVarMakerSuite{})

func (s *EnvVarMakerSuite) TestGetEnvVarArgument(c *C) {
	maker := NewEnvVarMaker()
	err := os.Setenv("FAKE_KEY", "FAKE_VALUE")
	value, err := maker.GetEnvVarValue("FAKE_KEY")

	c.Assert(err, IsNil)
	c.Assert(value, Equals, "FAKE_KEY=FAKE_VALUE")
}

func (s *EnvVarMakerSuite) TestGetEnvVarArgumentStatic(c *C) {
	maker := NewEnvVarMaker()
	value, err := maker.GetEnvVarValue("FAKE_KKEY=FAKE_VVALUE")

	c.Assert(err, IsNil)
	c.Assert(value, Equals, "FAKE_KKEY=FAKE_VVALUE")
}
