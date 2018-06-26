package custom

import "encoding/xml"

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Ser     string   `xml:"xmlns:ser,attr"`
	Body    SOAPBody
}

type SOAPBody struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Data    []byte   `xml:",innerxml"`
}

type DoGetItemsListRequest struct {
	XMLName       xml.Name       `xml:"ser:DoGetItemsListRequest"`
	WebAPIKey     string         `xml:"ser:webapiKey"`
	CountryID     int            `xml:"ser:countryId"`
	FilterOptions *FilterOptions `xml:",omitempty"`
	SortOptions   *SortOptions   `xml:",omitempty"`
}

type FilterOptions struct {
	XMLName xml.Name `xml:"ser:filterOptions"`
	Items   interface{}
}

type Item struct {
	XMLName xml.Name `xml:"ser:item"`
}

type FilterOptionsItem struct {
	Item
	FilterID string `xml:"ser:filterId"`
}

type FilterOptionsValue struct {
	FilterOptionsItem
	Items []interface{} `xml:"ser:filterValueId>ser:item"`
}

type FilterOptionsValueRange struct {
	FilterOptionsItem
	Min int `xml:"ser:filterValueRange>ser:rangeValueMin"`
	Max int `xml:"ser:filterValueRange>ser:rangeValueMax"`
}

type SortOptions struct {
	XMLName   xml.Name `xml:"ser:sortOptions"`
	SortType  string   `xml:"ser:sortType"`
	SortOrder string   `xml:"ser:sortOrder"`
}
