package queue

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type StompTestSuite struct {
	baseSuite
}

func (suite *StompTestSuite) SetupTest() {
	host, _ := os.LookupEnv("TESTS_HOST")
	b, err := NewStompBackend(host+":61613", stomp.ConnOpt.Host("/"))

	if err != nil {
		fmt.Printf("stomp err: %v\n", err)
	}

	Use(b)
}

func TestStompTestSuite(t *testing.T) {
	suite.Run(t, new(StompTestSuite))
}

func TestInvalidStompUrl(t *testing.T) {
	b, err := NewStompBackend("https://google.com")
	assert.NotNil(t, err)
	assert.Nil(t, b)
}
