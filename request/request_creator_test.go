package request

import (
	"allegro/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SortSettingsConverter(t *testing.T) {
	settings := model.SortSettings{Type: model.Price, Order: model.Asc}
	sortOptions := convertSortSettings(settings)
	assert.Equal(t, "price", sortOptions.SortType, "")
	assert.Equal(t, "asc", sortOptions.SortOrder, "")
}
