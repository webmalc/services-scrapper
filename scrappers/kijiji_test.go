package scrappers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/services-scrapper/scrappers/mocks"
)

func TestNewKijiji(t *testing.T) {
	log := &mocks.Logger{}
	kijiji := NewKijiji(log)
	assert.Equal(t, kijiji.logger, log)
	assert.Equal(t, kijiji.id, kijijiID)
	assert.Contains(t, kijiji.allowedDomains, "kijiji.ca")
	assert.Contains(t, kijiji.allowedDomains, "www.kijiji.ca")
}
