package scrappers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/services-scrapper/common/db"
	"github.com/webmalc/services-scrapper/models"
	"github.com/webmalc/services-scrapper/scrappers/mocks"
)

// newTestServer return a new test server
func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
<title>Kijiji</title>
</head>
<body>
		<div class="pagination"><a href="/2" title="Next">Next</a></div>
</body>
</html>
		`))
	})

	mux.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
<title>Kijiji</title>
</head>
<body>
		<a href="/service" class="title"></a>
</body>
</html>
		`))
	})

	mux.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
<title>Kijiji</title>
</head>
<body>
	<div id="ViewItemPage">
		<h1 itemProp="name">Kijiji service</h1>
		<ul class="crumbList-100"><li>Ontario</li><li>Invalid</li></ul>
		<span itemProp="address">Kijiji address</span>
		<div itemProp="description">Kijiji description</div>
	</div>
</body>
</html>
		`))
	})

	return httptest.NewServer(mux)
}

// Should parse a kijiji website.
func Test_processKijijiURL(t *testing.T) {
	ts := newTestServer()
	conn := db.NewConnection()
	defer ts.Close()
	defer conn.Close()
	models.Migrate(conn)
	repo := models.NewServiceRepository(conn)
	log := &mocks.Logger{}
	kijiji := NewKijiji(log, repo)
	kijiji.allowedDomains = []string{}
	colly := kijiji.getCollector()

	log.On("Infof", mock.Anything, mock.Anything).Return(nil).Times(3)
	kijiji.processURL(ts.URL, colly, repo)
	log.AssertExpectations(t)
	var count int
	var service models.Service
	conn.Find(&models.Service{}).Count(&count)
	assert.Equal(t, 1, count)
	conn.Find(&service)
	assert.Equal(t, "Kijiji service", service.Title)
	assert.Equal(t, "Kijiji description", service.Description)
	assert.Equal(t, "Canada", service.Country)
	assert.Equal(t, "Ontario", service.City)
	assert.Equal(t, "Kijiji address", service.Address)
}

// Should create a new kijiji scrapper instance
func TestNewKijiji(t *testing.T) {
	log := &mocks.Logger{}
	repo := &mocks.ServiceRepository{}
	kijiji := NewKijiji(log, repo)
	assert.Equal(t, kijiji.logger, log)
	assert.Equal(t, kijiji.id, kijijiID)
	assert.Contains(t, kijiji.allowedDomains, "kijiji.ca")
	assert.Contains(t, kijiji.allowedDomains, "www.kijiji.ca")
}
