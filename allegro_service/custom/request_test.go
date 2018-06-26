package custom

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EnvelopeWithBody(t *testing.T) {
	testSOAP := SOAPEnvelope{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		Ser:     "https://webapi.allegro.pl.allegrosandbox.pl/service.php",
		Body:    SOAPBody{},
	}
	res, err := xml.Marshal(testSOAP)

	assert.Nil(t, err, "")
	test := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="https://webapi.allegro.pl.allegrosandbox.pl/service.php"><soapenv:Body></soapenv:Body></soapenv:Envelope>`
	assert.Equal(t, test, string(res), "")
}

func Test_EnvelopeWithBody_DoGetItemsListRequest(t *testing.T) {
	doGetItemRequest := DoGetItemsListRequest{WebAPIKey: "TEST", CountryID: 1}
	content, _ := xml.Marshal(doGetItemRequest)
	testSOAP := SOAPEnvelope{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		Ser:     "https://webapi.allegro.pl.allegrosandbox.pl/service.php",
		Body:    SOAPBody{Data: content},
	}
	res, err := xml.Marshal(testSOAP)

	assert.Nil(t, err, "")
	test := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="https://webapi.allegro.pl.allegrosandbox.pl/service.php"><soapenv:Body><ser:DoGetItemsListRequest><ser:webapiKey>TEST</ser:webapiKey><ser:countryId>1</ser:countryId></ser:DoGetItemsListRequest></soapenv:Body></soapenv:Envelope>`
	assert.Equal(t, test, string(res), "")
}

func Test_FilterOptions_filterValueId(t *testing.T) {
	filterOptionsValID := FilterOptionsValue{Items: []interface{}{1, 2, 3, 4}}
	filterOptions := FilterOptions{Items: []interface{}{filterOptionsValID}}
	res, err := xml.MarshalIndent(filterOptions, "", " ")

	assert.Nil(t, err, "")
	test := `<ser:filterOptions>
 <ser:item>
  <ser:filterId></ser:filterId>
  <ser:filterValueId>
   <ser:item>1</ser:item>
   <ser:item>2</ser:item>
   <ser:item>3</ser:item>
   <ser:item>4</ser:item>
  </ser:filterValueId>
 </ser:item>
</ser:filterOptions>`

	assert.Equal(t, test, string(res), "")
}

func Test_FilterOptions_filterValueRange(t *testing.T) {
	filterOptionsValID := FilterOptionsValueRange{Min: 1, Max: 2}
	filterOptions := FilterOptions{Items: []interface{}{filterOptionsValID}}
	res, err := xml.MarshalIndent(filterOptions, "", " ")

	assert.Nil(t, err, "")
	test := `<ser:filterOptions>
 <ser:item>
  <ser:filterId></ser:filterId>
  <ser:filterValueRange>
   <ser:rangeValueMin>1</ser:rangeValueMin>
   <ser:rangeValueMax>2</ser:rangeValueMax>
  </ser:filterValueRange>
 </ser:item>
</ser:filterOptions>`

	assert.Equal(t, test, string(res), "")
}

func Test_FilterOptions_sortOptions(t *testing.T) {
	sortOptions := SortOptions{SortOrder: "?", SortType: "??"}
	res, err := xml.MarshalIndent(sortOptions, "", " ")

	assert.Nil(t, err, "")
	test := `<ser:sortOptions>
 <ser:sortType>??</ser:sortType>
 <ser:sortOrder>?</ser:sortOrder>
</ser:sortOptions>`

	assert.Equal(t, test, string(res), "")
}
