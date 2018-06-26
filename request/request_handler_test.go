package request

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var generatedFeed = `<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
  <title>AllegroRSS - netflix</title>
  <id></id>
  <updated>2018-06-24T14:32:28+02:00</updated>
  <link href=""></link>
  <author>
    <name>Allegro RSS Generator</name>
  </author>
  <entry>
    <title>Netflix 30 dni test / rrr</title>
    <updated></updated>
    <id>6205370452</id>
    <link href="https://allegro.pl.allegrosandbox.pl/i6205370452.html" rel="alternate"></link>
    <link href="https://4.allegroimg.allegrosandbox.pl/s400/03c737/fef4750449d8acd79dcfef48a194" rel="enclosure" type="image/jpeg" length="large"></link>
    <summary type="html">Price: buyNow:5.00zl withDelivery:5.00zl EndTime: 2018-06-28 09:34:59 +0200 CEST</summary>
  </entry>
  <entry>
    <title>Netflix  kosz</title>
    <updated></updated>
    <id>6205370352</id>
    <link href="https://allegro.pl.allegrosandbox.pl/i6205370352.html" rel="alternate"></link>
    <link href="https://0.allegroimg.allegrosandbox.pl/s400/0310a7/e5e189014dda87860ce513904a60" rel="enclosure" type="image/jpeg" length="large"></link>
    <summary type="html">Price: buyNow:8.00zl withDelivery:8.00zl EndTime: 2018-06-28 09:34:59 +0200 CEST</summary>
  </entry>
</feed>`

var apiResponse = `
<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope 
    xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" 
    xmlns:ns1="https://webapi.allegro.pl.allegrosandbox.pl/service.php">
    <SOAP-ENV:Body>
        <ns1:doGetItemsListResponse>
            <ns1:itemsCount>2</ns1:itemsCount>
            <ns1:itemsFeaturedCount>1</ns1:itemsFeaturedCount>
            <ns1:itemsList>
                <ns1:item>
                    <ns1:itemId>6205370452</ns1:itemId>
                    <ns1:itemTitle>Netflix 30 dni test / rrr</ns1:itemTitle>
                    <ns1:leftCount>99995</ns1:leftCount>
                    <ns1:bidsCount>5</ns1:bidsCount>
                    <ns1:biddersCount>3</ns1:biddersCount>
                    <ns1:quantityType>pieces</ns1:quantityType>
                    <ns1:endingTime>2018-06-28T09:34:59+02:00</ns1:endingTime>
                    <ns1:timeToEnd>3 dni</ns1:timeToEnd>
                    <ns1:categoryId>77925</ns1:categoryId>
                    <ns1:conditionInfo>new</ns1:conditionInfo>
                    <ns1:promotionInfo>1</ns1:promotionInfo>
                    <ns1:additionalInfo>0</ns1:additionalInfo>
                    <ns1:sellerInfo>
                        <ns1:userId>43966094</ns1:userId>
                        <ns1:userLogin>ottako</ns1:userLogin>
                        <ns1:userRating>0</ns1:userRating>
                        <ns1:userIcons>0</ns1:userIcons>
                        <ns1:countryId>1</ns1:countryId>
                    </ns1:sellerInfo>
                    <ns1:priceInfo>
                        <ns1:item>
                            <ns1:priceType>buyNow</ns1:priceType>
                            <ns1:priceValue>5</ns1:priceValue>
                        </ns1:item>
                        <ns1:item>
                            <ns1:priceType>withDelivery</ns1:priceType>
                            <ns1:priceValue>5</ns1:priceValue>
                        </ns1:item>
                    </ns1:priceInfo>
                    <ns1:photosInfo>
                        <ns1:item>
                            <ns1:photoSize>small</ns1:photoSize>
                            <ns1:photoUrl>https://4.allegroimg.allegrosandbox.pl/s64/03c737/fef4750449d8acd79dcfef48a194</ns1:photoUrl>
                            <ns1:photoIsMain>true</ns1:photoIsMain>
                        </ns1:item>
                        <ns1:item>
                            <ns1:photoSize>small</ns1:photoSize>
                            <ns1:photoUrl>https://4.allegroimg.allegrosandbox.pl/s64/03c737/fef4750449d8acd79dcfef48a194_1</ns1:photoUrl>
                            <ns1:photoIsMain>false</ns1:photoIsMain>
                        </ns1:item>
                        <ns1:item>
                            <ns1:photoSize>medium</ns1:photoSize>
                            <ns1:photoUrl>https://4.allegroimg.allegrosandbox.pl/s128/03c737/fef4750449d8acd79dcfef48a194</ns1:photoUrl>
                            <ns1:photoIsMain>true</ns1:photoIsMain>
                        </ns1:item>
                        <ns1:item>
                            <ns1:photoSize>medium</ns1:photoSize>
                            <ns1:photoUrl>https://4.allegroimg.allegrosandbox.pl/s128/03c737/fef4750449d8acd79dcfef48a194_1</ns1:photoUrl>
                            <ns1:photoIsMain>false</ns1:photoIsMain>
                        </ns1:item>
                        <ns1:item>
                            <ns1:photoSize>large</ns1:photoSize>
                            <ns1:photoUrl>https://4.allegroimg.allegrosandbox.pl/s400/03c737/fef4750449d8acd79dcfef48a194</ns1:photoUrl>
                            <ns1:photoIsMain>true</ns1:photoIsMain>
                        </ns1:item>
                        <ns1:item>
                            <ns1:photoSize>large</ns1:photoSize>
                            <ns1:photoUrl>https://4.allegroimg.allegrosandbox.pl/s400/03c737/fef4750449d8acd79dcfef48a194_1</ns1:photoUrl>
                            <ns1:photoIsMain>false</ns1:photoIsMain>
                        </ns1:item>
                    </ns1:photosInfo>
                    <ns1:parametersInfo>
                        <ns1:item>
                            <ns1:parameterName>Stan</ns1:parameterName>
                            <ns1:parameterValue>
                                <ns1:item>Nowy</ns1:item>
                            </ns1:parameterValue>
                            <ns1:parameterUnit></ns1:parameterUnit>
                            <ns1:parameterIsRange>false</ns1:parameterIsRange>
                        </ns1:item>
                    </ns1:parametersInfo>
                </ns1:item>
                <ns1:item>
                    <ns1:itemId>6205370352</ns1:itemId>
                    <ns1:itemTitle>Netflix  kosz</ns1:itemTitle>
                    <ns1:leftCount>1000</ns1:leftCount>
                    <ns1:bidsCount>0</ns1:bidsCount>
                    <ns1:biddersCount>0</ns1:biddersCount>
                    <ns1:quantityType>pieces</ns1:quantityType>
                    <ns1:endingTime>2018-06-28T09:34:59+02:00</ns1:endingTime>
                    <ns1:timeToEnd>3 dni</ns1:timeToEnd>
                    <ns1:categoryId>15822</ns1:categoryId>
                    <ns1:conditionInfo>undefined</ns1:conditionInfo>
                    <ns1:promotionInfo>0</ns1:promotionInfo>
                    <ns1:additionalInfo>0</ns1:additionalInfo>
                    <ns1:sellerInfo>
                        <ns1:userId>43966094</ns1:userId>
                        <ns1:userLogin>ottako</ns1:userLogin>
                        <ns1:userRating>0</ns1:userRating>
                        <ns1:userIcons>0</ns1:userIcons>
                        <ns1:countryId>1</ns1:countryId>
                    </ns1:sellerInfo>
                    <ns1:priceInfo>
                        <ns1:item>
                            <ns1:priceType>buyNow</ns1:priceType>
                            <ns1:priceValue>8</ns1:priceValue>
                        </ns1:item>
                        <ns1:item>
                            <ns1:priceType>withDelivery</ns1:priceType>
                            <ns1:priceValue>8</ns1:priceValue>
                        </ns1:item>
                    </ns1:priceInfo>
                    <ns1:photosInfo>
                        <ns1:item>
                            <ns1:photoSize>small</ns1:photoSize>
                            <ns1:photoUrl>https://0.allegroimg.allegrosandbox.pl/s64/0310a7/e5e189014dda87860ce513904a60</ns1:photoUrl>
                            <ns1:photoIsMain>true</ns1:photoIsMain>
                        </ns1:item>
                        <ns1:item>
                            <ns1:photoSize>medium</ns1:photoSize>
                            <ns1:photoUrl>https://0.allegroimg.allegrosandbox.pl/s128/0310a7/e5e189014dda87860ce513904a60</ns1:photoUrl>
                            <ns1:photoIsMain>true</ns1:photoIsMain>
                        </ns1:item>
                        <ns1:item>
                            <ns1:photoSize>large</ns1:photoSize>
                            <ns1:photoUrl>https://0.allegroimg.allegrosandbox.pl/s400/0310a7/e5e189014dda87860ce513904a60</ns1:photoUrl>
                            <ns1:photoIsMain>true</ns1:photoIsMain>
                        </ns1:item>
                    </ns1:photosInfo>
                </ns1:item>
            </ns1:itemsList>
            <ns1:categoriesList>
                <ns1:categoriesTree>
                    <ns1:item>
                        <ns1:categoryId>2</ns1:categoryId>
                        <ns1:categoryName>Komputery</ns1:categoryName>
                        <ns1:categoryParentId>0</ns1:categoryParentId>
                        <ns1:categoryItemsCount>2</ns1:categoryItemsCount>
                    </ns1:item>
                    <ns1:item>
                        <ns1:categoryId>77753</ns1:categoryId>
                        <ns1:categoryName>Pozostałe</ns1:categoryName>
                        <ns1:categoryParentId>2</ns1:categoryParentId>
                        <ns1:categoryItemsCount>1</ns1:categoryItemsCount>
                    </ns1:item>
                    <ns1:item>
                        <ns1:categoryId>15821</ns1:categoryId>
                        <ns1:categoryName>Internet</ns1:categoryName>
                        <ns1:categoryParentId>2</ns1:categoryParentId>
                        <ns1:categoryItemsCount>1</ns1:categoryItemsCount>
                    </ns1:item>
                </ns1:categoriesTree>
            </ns1:categoriesList>
            <ns1:filtersList>
                <ns1:item>
                    <ns1:filterId>price</ns1:filterId>
                    <ns1:filterName>Cena</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>textbox</ns1:filterControlType>
                    <ns1:filterDataType>float</ns1:filterDataType>
                    <ns1:filterIsRange>true</ns1:filterIsRange>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>state</ns1:filterId>
                    <ns1:filterName>Województwo</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>combobox</ns1:filterControlType>
                    <ns1:filterDataType>int</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>1</ns1:filterValueId>
                            <ns1:filterValueName>dolnośląskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>2</ns1:filterValueId>
                            <ns1:filterValueName>kujawsko-pomorskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>3</ns1:filterValueId>
                            <ns1:filterValueName>lubelskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>4</ns1:filterValueId>
                            <ns1:filterValueName>lubuskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>5</ns1:filterValueId>
                            <ns1:filterValueName>łódzkie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>6</ns1:filterValueId>
                            <ns1:filterValueName>małopolskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>7</ns1:filterValueId>
                            <ns1:filterValueName>mazowieckie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>8</ns1:filterValueId>
                            <ns1:filterValueName>opolskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>9</ns1:filterValueId>
                            <ns1:filterValueName>podkarpackie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>10</ns1:filterValueId>
                            <ns1:filterValueName>podlaskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>11</ns1:filterValueId>
                            <ns1:filterValueName>pomorskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>12</ns1:filterValueId>
                            <ns1:filterValueName>śląskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>13</ns1:filterValueId>
                            <ns1:filterValueName>świętokrzyskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>14</ns1:filterValueId>
                            <ns1:filterValueName>warmińsko-mazurskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>15</ns1:filterValueId>
                            <ns1:filterValueName>wielkopolskie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>16</ns1:filterValueId>
                            <ns1:filterValueName>zachodniopomorskie</ns1:filterValueName>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>city</ns1:item>
                            <ns1:item>distance</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                            <ns1:item>postCode</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>userId</ns1:filterId>
                    <ns1:filterName>Identyfikator użytkownika</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>textbox</ns1:filterControlType>
                    <ns1:filterDataType>int</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>category</ns1:filterId>
                    <ns1:filterName>Kategoria</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>textbox</ns1:filterControlType>
                    <ns1:filterDataType>int</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterRelations>
                        <ns1:relationExclude>
                            <ns1:item>mainPage</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>search</ns1:filterId>
                    <ns1:filterName>Szukaj w tytule</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>textbox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>description</ns1:filterId>
                    <ns1:filterName>Szukaj też w opisach i parametrach</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>checkbox</ns1:filterControlType>
                    <ns1:filterDataType>boolean</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterRelations>
                        <ns1:relationAnd>
                            <ns1:item>search</ns1:item>
                        </ns1:relationAnd>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                            <ns1:item>similar</ns1:item>
                            <ns1:item>special</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>closed</ns1:filterId>
                    <ns1:filterName>Szukaj w zakończonych</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>checkbox</ns1:filterControlType>
                    <ns1:filterDataType>boolean</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterRelations>
                        <ns1:relationAnd>
                            <ns1:item>search</ns1:item>
                        </ns1:relationAnd>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                            <ns1:item>similar</ns1:item>
                            <ns1:item>special</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>similar</ns1:filterId>
                    <ns1:filterName>Znajdź podobne oferty</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>checkbox</ns1:filterControlType>
                    <ns1:filterDataType>boolean</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterRelations>
                        <ns1:relationAnd>
                            <ns1:item>search</ns1:item>
                        </ns1:relationAnd>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>closed</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>description</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                            <ns1:item>special</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>department</ns1:filterId>
                    <ns1:filterName>Dział</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>combobox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>electronics</ns1:filterValueId>
                            <ns1:filterValueName>Elektronika</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>fashionBeauty</ns1:filterValueId>
                            <ns1:filterValueName>Moda i uroda</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>householdHealth</ns1:filterValueId>
                            <ns1:filterValueName>Dom i zdrowie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>baby</ns1:filterValueId>
                            <ns1:filterValueName>Dziecko</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>cultureEntertainment</ns1:filterValueId>
                            <ns1:filterValueName>Kultura i rozrywka</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>sportsLeisure</ns1:filterValueId>
                            <ns1:filterValueName>Sport i wypoczynek</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>automotive</ns1:filterValueId>
                            <ns1:filterValueName>Motoryzacja</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>collectionsArt</ns1:filterValueId>
                            <ns1:filterValueName>Kolekcje i sztuka</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>companyServices</ns1:filterValueId>
                            <ns1:filterValueName>Firma</ns1:filterValueName>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                            <ns1:item>special</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>condition</ns1:filterId>
                    <ns1:filterName>Stan</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>checkbox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>new</ns1:filterValueId>
                            <ns1:filterValueName>Nowe</ns1:filterValueName>
                            <ns1:filterValueCount>1</ns1:filterValueCount>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>used</ns1:filterValueId>
                            <ns1:filterValueName>Używane</ns1:filterValueName>
                            <ns1:filterValueCount>0</ns1:filterValueCount>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>offerType</ns1:filterId>
                    <ns1:filterName>Rodzaj oferty</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>checkbox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>buyNow</ns1:filterValueId>
                            <ns1:filterValueName>Kup Teraz</ns1:filterValueName>
                            <ns1:filterValueCount>2</ns1:filterValueCount>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>auction</ns1:filterValueId>
                            <ns1:filterValueName>Licytacje</ns1:filterValueName>
                            <ns1:filterValueCount>0</ns1:filterValueCount>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>distance</ns1:filterId>
                    <ns1:filterName>Dystans</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>combobox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>10km</ns1:filterValueId>
                            <ns1:filterValueName>10 km</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>25km</ns1:filterValueId>
                            <ns1:filterValueName>25 km</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>50km</ns1:filterValueId>
                            <ns1:filterValueName>50 km</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>75km</ns1:filterValueId>
                            <ns1:filterValueName>75 km</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>100km</ns1:filterValueId>
                            <ns1:filterValueName>100 km</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>150km</ns1:filterValueId>
                            <ns1:filterValueName>150 km</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>200km</ns1:filterValueId>
                            <ns1:filterValueName>200 km</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>350km</ns1:filterValueId>
                            <ns1:filterValueName>350 km</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>500km</ns1:filterValueId>
                            <ns1:filterValueName>500 km</ns1:filterValueName>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationAnd>
                            <ns1:item>postCode</ns1:item>
                        </ns1:relationAnd>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>city</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                            <ns1:item>state</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>shippingTime</ns1:filterId>
                    <ns1:filterName>Czas realizacji</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>combobox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>24h</ns1:filterValueId>
                            <ns1:filterValueName>24 godziny</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>2d</ns1:filterValueId>
                            <ns1:filterValueName>do 2 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>3d</ns1:filterValueId>
                            <ns1:filterValueName>do 3 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>4d</ns1:filterValueId>
                            <ns1:filterValueName>do 4 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>5d</ns1:filterValueId>
                            <ns1:filterValueName>do 5 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>7d</ns1:filterValueId>
                            <ns1:filterValueName>do 7 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>10d</ns1:filterValueId>
                            <ns1:filterValueName>do 10 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>14d</ns1:filterValueId>
                            <ns1:filterValueName>do 14 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>21d</ns1:filterValueId>
                            <ns1:filterValueName>do 21 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>more</ns1:filterValueId>
                            <ns1:filterValueName>dłużej</ns1:filterValueName>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>postCode</ns1:filterId>
                    <ns1:filterName>Kod pocztowy</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>textbox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterRelations>
                        <ns1:relationAnd>
                            <ns1:item>distance</ns1:item>
                        </ns1:relationAnd>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>city</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                            <ns1:item>state</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>city</ns1:filterId>
                    <ns1:filterName>Miejscowość</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>textbox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>distance</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                            <ns1:item>postCode</ns1:item>
                            <ns1:item>state</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>offerOptions</ns1:filterId>
                    <ns1:filterName>Opcje oferty</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>checkbox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>freeShipping</ns1:filterValueId>
                            <ns1:filterValueName>Wysyłka gratis</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>freeReturn</ns1:filterValueId>
                            <ns1:filterValueName>Allegro InPost</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>generalDelivery</ns1:filterValueId>
                            <ns1:filterValueName>odbiór w Paczkomacie</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>installmentAvailable</ns1:filterValueId>
                            <ns1:filterValueName>Raty PayU</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>vatInvoice</ns1:filterValueId>
                            <ns1:filterValueName>Faktura VAT</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>personalReceipt</ns1:filterValueId>
                            <ns1:filterValueName>Odbiór osobisty</ns1:filterValueName>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>endingTime</ns1:filterId>
                    <ns1:filterName>Kończące się w ciągu</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>combobox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>1h</ns1:filterValueId>
                            <ns1:filterValueName>1 godziny</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>2h</ns1:filterValueId>
                            <ns1:filterValueName>2 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>3h</ns1:filterValueId>
                            <ns1:filterValueName>3 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>4h</ns1:filterValueId>
                            <ns1:filterValueName>4 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>5h</ns1:filterValueId>
                            <ns1:filterValueName>5 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>12h</ns1:filterValueId>
                            <ns1:filterValueName>12 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>24h</ns1:filterValueId>
                            <ns1:filterValueName>24 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>2d</ns1:filterValueId>
                            <ns1:filterValueName>2 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>3d</ns1:filterValueId>
                            <ns1:filterValueName>3 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>4d</ns1:filterValueId>
                            <ns1:filterValueName>4 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>5d</ns1:filterValueId>
                            <ns1:filterValueName>5 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>6d</ns1:filterValueId>
                            <ns1:filterValueName>6 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>7d</ns1:filterValueId>
                            <ns1:filterValueName>7 dni</ns1:filterValueName>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>startingTime</ns1:filterId>
                    <ns1:filterName>Wystawione w ciągu</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>combobox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>1h</ns1:filterValueId>
                            <ns1:filterValueName>1 godziny</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>2h</ns1:filterValueId>
                            <ns1:filterValueName>2 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>3h</ns1:filterValueId>
                            <ns1:filterValueName>3 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>4h</ns1:filterValueId>
                            <ns1:filterValueName>4 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>5h</ns1:filterValueId>
                            <ns1:filterValueName>5 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>12h</ns1:filterValueId>
                            <ns1:filterValueName>12 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>24h</ns1:filterValueId>
                            <ns1:filterValueName>24 godzin</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>2d</ns1:filterValueId>
                            <ns1:filterValueName>2 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>3d</ns1:filterValueId>
                            <ns1:filterValueName>3 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>4d</ns1:filterValueId>
                            <ns1:filterValueName>4 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>5d</ns1:filterValueId>
                            <ns1:filterValueName>5 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>6d</ns1:filterValueId>
                            <ns1:filterValueName>6 dni</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>7d</ns1:filterValueId>
                            <ns1:filterValueName>7 dni</ns1:filterValueName>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>search</ns1:item>
                            <ns1:item>special</ns1:item>
                            <ns1:item>userId</ns1:item>
                        </ns1:relationOr>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>special</ns1:filterId>
                    <ns1:filterName>Typ listingu</ns1:filterName>
                    <ns1:filterType>country</ns1:filterType>
                    <ns1:filterControlType>combobox</ns1:filterControlType>
                    <ns1:filterDataType>string</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>new</ns1:filterValueId>
                            <ns1:filterValueName>Najnowsze</ns1:filterValueName>
                        </ns1:item>
                        <ns1:item>
                            <ns1:filterValueId>ending</ns1:filterValueId>
                            <ns1:filterValueName>Kończące się</ns1:filterValueName>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationExclude>
                            <ns1:item>categoryPage</ns1:item>
                            <ns1:item>closed</ns1:item>
                            <ns1:item>department</ns1:item>
                            <ns1:item>departmentPage</ns1:item>
                            <ns1:item>description</ns1:item>
                            <ns1:item>mainPage</ns1:item>
                            <ns1:item>similar</ns1:item>
                        </ns1:relationExclude>
                    </ns1:filterRelations>
                </ns1:item>
                <ns1:item>
                    <ns1:filterId>11323</ns1:filterId>
                    <ns1:filterName>Stan</ns1:filterName>
                    <ns1:filterType>category</ns1:filterType>
                    <ns1:filterControlType>checkbox</ns1:filterControlType>
                    <ns1:filterDataType>int</ns1:filterDataType>
                    <ns1:filterIsRange>false</ns1:filterIsRange>
                    <ns1:filterValues>
                        <ns1:item>
                            <ns1:filterValueId>1</ns1:filterValueId>
                            <ns1:filterValueName>Nowy</ns1:filterValueName>
                            <ns1:filterValueCount>1</ns1:filterValueCount>
                        </ns1:item>
                    </ns1:filterValues>
                    <ns1:filterRelations>
                        <ns1:relationOr>
                            <ns1:item>category</ns1:item>
                            <ns1:item>search</ns1:item>
                        </ns1:relationOr>
                    </ns1:filterRelations>
                </ns1:item>
            </ns1:filtersList>
        </ns1:doGetItemsListResponse>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>
`

type FakeWriter struct {
	*bytes.Buffer
}

func (*FakeWriter) Header() http.Header {
	return http.Header{}
}

func (*FakeWriter) WriteHeader(statusCode int) {
}

type FakeClient struct {
	resp *http.Response
}

func (fc *FakeClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return fc.resp, nil
}

type BufferCloser struct {
	*bytes.Buffer
}

func (*BufferCloser) Close() error {
	return nil
}

type FakeTimeProvider struct {
	time.Time
}

func (tp *FakeTimeProvider) Now() time.Time {
	return tp.Time
}

func Test_RequestHandler_BasicRequest(t *testing.T) {

	bufferCloser := BufferCloser{bytes.NewBuffer([]byte(apiResponse))}
	resp := http.Response{Body: &bufferCloser}

	fakeTime, err := time.Parse(time.RFC3339, "2018-06-24T14:32:28+02:00")
	assert.Nil(t, err, "")

	handler := &RequestHandler{
		client:         &FakeClient{&resp},
		parser:         &QueryParser{},
		requestCreator: &SoapRequestCreator{},
		responseCreator: &DefaultResponseCreator{
			urlBase:      "https://allegro.pl.allegrosandbox.pl",
			timeProvider: &FakeTimeProvider{Time: fakeTime},
		},
		webAPIKey: "",
	}

	request, err := http.NewRequest("GET", "http://localhost:8080/feed?search=netflix", nil)
	assert.Nil(t, err)

	response := bytes.Buffer{}
	handler.CreateFeed(&FakeWriter{Buffer: &response}, request)

	body, err := ioutil.ReadAll(&response)
	assert.Equal(t, generatedFeed, string(body), "")
}
