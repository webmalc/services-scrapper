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

// newYellowpagesServer return a new test server
func newYellowpagesServer() *httptest.Server {
	return newTestServer(
		`<!DOCTYPE html>
<html>
<head>
<title>Yellowpages</title>
</head>
<body>
		<span class="pageCount"></span>
		<a href="/next" class="pageButton">Next</a>
</body>
</html>
		`,
		`<!DOCTYPE html>
<html>
<head>
<title>Yellowpages</title>
</head>
<body>
		<a class="listing__link" href="/service"></a>
</body>
</html>
		`,
		`<!DOCTYPE html>
<html>
<head>
<title>Yellowpages</title>
</head>
<body>
	<div class="merchant__header--root">
		<h1 class="merchant__title">
			<span itemprop="name">Yellowpages service</span>
		</h1>
		<div id="businessSection">Yellowpages description</div>
		<div class="merchant__details__section--address">
			<div class="merchant__address">
				Yellowpages <span itemprop="addressLocality">address</span>
			</div>
		</div>
		<a class="merchant-logo-link">
			<img src="https://example.com/image.png">
		</a>
		<div class="middle_section">
			<ul>
				<li class="mlr__item--phone">
					<span class="mlr__sub-text">1111</span>
					<span class="mlr__sub-text">2222</span>
				</li>
			</ul>
			<ul>
				<li class="mlr__item--website">
					<span class="mlr__sub-text">example.com</span>
					<span class="mlr__sub-text">https://github.com</span>
				</li>
			</ul>
		</div>
		// images
	</div>
<script type="application/json" class="jsImagesConfigHandlebar">
{
"medias" :
	[
	{
	"isPhoto" : true,
	"url" : "https://example.com/image.jpg"
	},
	{
	"isPhoto" : false,
	"url" : "https://example.com/invalid"
	}
	]
}
</script>

</body>
</html>
		`,
	)
}

// Should parse a yellowpages website.
func Test_processYellowpagesURL(t *testing.T) {
	ts := newYellowpagesServer()
	conn := db.NewConnection()
	defer ts.Close()
	defer conn.Close()
	models.Migrate(conn)
	repo := models.NewServiceRepository(conn)
	log := &mocks.Logger{}
	yellowpages := NewYellowpages(log, repo)
	yellowpages.allowedDomains = []string{}
	colly := yellowpages.getCollector()

	log.On("Infof", mock.Anything, mock.Anything).Return(nil).Times(3)
	yellowpages.processURL(ts.URL, colly, repo)
	log.AssertExpectations(t)
	var count int
	var service models.Service
	conn.Find(&models.Service{}).Count(&count)
	assert.Equal(t, 1, count)
	conn.Preload("Phones").Preload("Links").Preload("Images").Find(&service)
	assert.Equal(t, "Yellowpages service", service.Title)
	assert.Equal(t, "Yellowpages description", service.Description)
	assert.Equal(t, "Canada", service.Country)
	assert.Equal(t, "address", service.City)
	assert.Equal(t, "Yellowpages address", service.Address)
	assert.Len(t, service.Phones, 2)
	assert.Len(t, service.Images, 1)
	assert.Len(t, service.Links, 2)
	assert.Equal(t, service.Links[0].URL, "http://example.com")
	assert.Equal(t, service.Phones[0].Phone, "1111")
	assert.Equal(t, service.Images[0].Src, "https://example.com/image.jpg")
}

// Should create a new yellowpages scrapper instance
func TestNewYellowpages(t *testing.T) {
	log := &mocks.Logger{}
	repo := &mocks.ServiceRepository{}
	kijiji := NewYellowpages(log, repo)
	assert.Equal(t, kijiji.logger, log)
	assert.Equal(t, kijiji.id, yellowpagesID)
	assert.Contains(t, kijiji.allowedDomains, "yellowpages.ca")
	assert.Contains(t, kijiji.allowedDomains, "www.yellowpages.ca")
}
