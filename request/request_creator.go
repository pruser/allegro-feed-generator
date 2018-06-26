package request

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/pruser/allegro-feed-generator/model"

	api "github.com/pruser/allegro-feed-generator/allegro_service/custom"
)

type RequestCreator interface {
	CreateRequest(key model.WebAPIKey, sortSettings model.SortSettings, filterSettings model.FilterSettings) (api.SOAPEnvelope, error)
}

type SoapRequestCreator struct{}

func (*SoapRequestCreator) CreateRequest(key model.WebAPIKey, sortSettings model.SortSettings, filterSettings model.FilterSettings) (api.SOAPEnvelope, error) {
	sortOptions := convertSortSettings(sortSettings)
	filterOptions, _ := convertFilterSettings(filterSettings)

	doGetItemsListRequest := api.DoGetItemsListRequest{WebAPIKey: string(key), CountryID: 1, FilterOptions: &filterOptions, SortOptions: &sortOptions}

	content, err := xml.MarshalIndent(doGetItemsListRequest, " ", " ")
	if err != nil {
		return api.SOAPEnvelope{}, err
	}

	return api.SOAPEnvelope{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		Ser:     "https://webapi.allegro.pl.allegrosandbox.pl/service.php",
		Body: api.SOAPBody{
			Data: content,
		},
	}, nil
}

func convertSortSettings(s model.SortSettings) api.SortOptions {
	return api.SortOptions{SortType: string(s.Type), SortOrder: string(s.Order)}
}

func convertFilterSettings(s model.FilterSettings) (api.FilterOptions, error) {
	items := make([]interface{}, 0)
	for _, f := range s.Filters {
		i, err := processFilter(f)
		if err != nil {
			return api.FilterOptions{}, err
		}
		items = append(items, i)
	}
	return api.FilterOptions{Items: items}, nil
}

func processFilter(filter interface{}) (interface{}, error) {
	if f, ok := filter.(model.ValueFilter); ok {
		vals := make([]interface{}, 0)
		for _, v := range f.Values {
			vals = append(vals, v)
		}
		return api.FilterOptionsValue{FilterOptionsItem: api.FilterOptionsItem{FilterID: f.Filter.Name}, Items: vals}, nil
	}
	if f, ok := filter.(model.RangeFilter); ok {
		min, err := strconv.Atoi(f.Min)
		if err != nil {
			return nil, err
		}
		max, err := strconv.Atoi(f.Max)
		if err != nil {
			return nil, err
		}
		return api.FilterOptionsValueRange{FilterOptionsItem: api.FilterOptionsItem{FilterID: f.Filter.Name}, Min: min, Max: max}, nil
	}
	return nil, fmt.Errorf("No matching filter type")
}
