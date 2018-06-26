package model

type FilterSettings struct {
	Filters []interface{}
}

type Filter struct {
	Name string
}

type ValueFilter struct {
	Filter
	Values []string
}

func NewValueFilter(name string, vals ...string) ValueFilter {
	return ValueFilter{Filter: Filter{Name: name}, Values: vals}
}

type RangeFilter struct {
	Filter
	Min string
	Max string
}

func NewRangeFilter(name string, min string, max string) RangeFilter {
	return RangeFilter{Filter: Filter{Name: name}, Min: min, Max: max}
}
