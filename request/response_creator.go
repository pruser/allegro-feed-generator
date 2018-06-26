package request

import (
	"fmt"
	"time"

	api_gen "github.com/pruser/allegro-feed-generator/allegro_service/generated"

	"github.com/gorilla/feeds"
)

type ResponseCreator interface {
	CreateResponse(title string, resp *api_gen.DoGetItemsListResponse) ([]byte, error)
}

type TimeProvider interface {
	Now() time.Time
}

type RealTimeProvider struct {
}

func (*RealTimeProvider) Now() time.Time {
	return time.Now()
}

type DefaultResponseCreator struct {
	timeProvider TimeProvider
	urlBase      string
}

func (rf *DefaultResponseCreator) CreateResponse(title string, res *api_gen.DoGetItemsListResponse) ([]byte, error) {
	items := res.ItemsList.Item
	now := rf.timeProvider.Now()
	feed := &feeds.Feed{
		Title:   "AllegroRSS - " + title,
		Link:    &feeds.Link{},
		Author:  &feeds.Author{Name: "Allegro RSS Generator"},
		Updated: now,
	}

	for _, i := range items {
		itm := convertItemListToFeedItem(rf.urlBase, i)
		feed.Items = append(feed.Items, itm)
	}

	atom, err := feed.ToAtom()
	if err != nil {
		return nil, err
	}
	return []byte(atom), nil
}

func convertItemListToFeedItem(urlbase string, i *api_gen.ItemsListType) *feeds.Item {
	itm := &feeds.Item{
		Id:    fmt.Sprintf("%d", i.ItemId),
		Title: i.ItemTitle,
		Link:  &feeds.Link{Href: fmt.Sprintf("%s/i%d.html", urlbase, i.ItemId)},
	}
	if i.PhotosInfo != nil {
		for _, img := range i.PhotosInfo.Item {
			if img.PhotoIsMain {
				itm.Enclosure = &feeds.Enclosure{Url: img.PhotoUrl, Length: img.PhotoSize, Type: "image/jpeg"}
			}
		}
	}
	desc := ""
	if i.PriceInfo != nil {
		desc = desc + "Price: "
		for _, v := range i.PriceInfo.Item {
			desc = desc + fmt.Sprintf("%s:%.2fzl ", v.PriceType, v.PriceValue)
		}
	}
	desc = desc + fmt.Sprintf("EndTime: %s", i.EndingTime)
	itm.Description = desc

	return itm
}
