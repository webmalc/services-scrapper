package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/services-scrapper/cmd/mocks"
	"github.com/webmalc/services-scrapper/common/test"
)

// Should run the root command and log an error.
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.ErrorLogger{}
	cr := NewCommandRouter(m)
	os.Args = []string{"invalid", "invalid"}
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	m := &mocks.ErrorLogger{}
	cr := NewCommandRouter(m)
	assert.Equal(t, m, cr.logger)
	assert.NotNil(t, cr.rootCmd)
}

func TestCommandRouter_server(t *testing.T) {
	cr := NewCommandRouter(&mocks.ErrorLogger{})
	cr.scrap(&cobra.Command{}, make([]string, 0))
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
