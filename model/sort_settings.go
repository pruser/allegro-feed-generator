package model

import (
	"fmt"
)

type sortType string
type sortOrder string

const (
	EndingTime    = sortType("endingTime")
	StartingTime  = sortType("startingTime")
	Price         = sortType("price")
	PriceDelivery = sortType("priceDelivery")
	Popularity    = sortType("popularity")
	Name          = sortType("name")
	Relevance     = sortType("relevance")
)

const (
	Asc  = sortOrder("asc")
	Desc = sortOrder("desc")
)

type SortSettings struct {
	Type  sortType
	Order sortOrder
}

func NewSortSettings(t, o string) (SortSettings, error) {
	if !checkSortType(t) || !checkSortOrder(o) {
		// TODO split
		return SortSettings{}, fmt.Errorf("wrong sort parameters")
	}
	return SortSettings{Type: sortType(t), Order: sortOrder(o)}, nil
}

func checkSortType(s string) bool {
	switch sortType(s) {
	case EndingTime:
		fallthrough
	case StartingTime:
		fallthrough
	case Price:
		fallthrough
	case PriceDelivery:
		fallthrough
	case Popularity:
		fallthrough
	case Name:
		fallthrough
	case Relevance:
		return true
	default:
		return false
	}
}

func checkSortOrder(s string) bool {
	switch sortOrder(s) {
	case Asc:
		fallthrough
	case Desc:
		return true
	default:
		return false
	}
}
