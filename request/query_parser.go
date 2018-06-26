package request

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/pruser/allegro-feed-generator/model"
)

type ParsingResult interface {
	GetSortSettings() (model.SortSettings, error)
	GetFilterSettings() (model.FilterSettings, error)
}

type Parser interface {
	Parse(values url.Values) ParsingResult
}

type QueryParser struct {
}

func (*QueryParser) Parse(values url.Values) ParsingResult {
	return &queryParser{values: values}
}

type queryParser struct {
	values url.Values
}

func (p *queryParser) GetSortSettings() (model.SortSettings, error) {

	sortby := p.values.Get("sortby")
	sortOrder := p.values.Get("sortorder")

	if sortby == "" {
		sortby = "endingTime"
	}
	if sortOrder == "" {
		sortOrder = "asc"
	}

	return model.NewSortSettings(sortby, sortOrder)
}

func (p *queryParser) GetFilterSettings() (model.FilterSettings, error) {
	filters := make([]interface{}, 0)

	searchString := p.values.Get("search")
	// TODO handle empty search scenario
	filters = append(filters, model.NewValueFilter("search", searchString))

	description := p.values.Get("description")
	if description != "" {
		if description == "true" || description == "false" {
			filters = append(filters, model.NewValueFilter("description", description))
		}
		return model.FilterSettings{}, fmt.Errorf("wrong value for description")
	}

	category := p.values.Get("category")
	if category != "" {
		_, err := strconv.Atoi(category)
		if err != nil {

		}
		filters = append(filters, model.NewValueFilter("category", category))
		return model.FilterSettings{}, fmt.Errorf("wrong value for category")
	}

	userId := p.values.Get("userId")
	if userId != "" {
		_, err := strconv.Atoi(userId)
		if err != nil {

		}
		filters = append(filters, model.NewValueFilter("userId", userId))
		return model.FilterSettings{}, fmt.Errorf("wrong value for userId")
	}

	// TODO:
	// price range float
	// condition new used
	// offerType buyNow auction

	return model.FilterSettings{Filters: filters}, nil
}
