package scrappers

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/services-scrapper/common/db"
	"github.com/webmalc/services-scrapper/models"
	"github.com/webmalc/services-scrapper/scrappers/mocks"
)

// newUslugioServer return a new test server
func newUslugioServer() *httptest.Server {
	return newTestServer(

		`<!DOCTYPE html>
<html>
<head>
<title>Uslugio</title>
</head>
<body>
		<div class="title showone" data-id="-1041031">1</div>
		<span class="pageCount"></span>
		<a href="/next">Next</a>
</body>
</html>
		`,
		`<!DOCTYPE html>
<html>
<head>
<title>Yellowpages</title>
</head>
<body>
		<a href="/service"></a>
</body>
</html>
		`, "",
	)
}

func Test_processUslugioURL(t *testing.T) {
	ts := newUslugioServer()
	conn := db.NewConnection()
	defer ts.Close()
	defer conn.Close()
	models.Migrate(conn)
	repo := models.NewServiceRepository(conn)
	log := &mocks.Logger{}
	uslugio := NewUslugio(log, repo)
	uslugio.allowedDomains = []string{}
	colly := uslugio.getCollector()

	log.On("Infof", mock.Anything, mock.Anything).Return(nil).Times(3)
	uslugio.processURL(ts.URL, colly, repo)
	log.AssertExpectations(t)
}

func TestNewUslugio(t *testing.T) {
	log := &mocks.Logger{}
	repo := &mocks.ServiceRepository{}
	uslugio := NewUslugio(log, repo)
	assert.Equal(t, uslugio.logger, log)
	assert.Equal(t, uslugio.id, uslugioID)
	assert.Contains(t, uslugio.allowedDomains, "uslugio.com")
	assert.Contains(t, uslugio.allowedDomains, "www.uslugio.com")
}
