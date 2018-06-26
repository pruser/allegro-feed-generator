// file generated from gowsdl tool with manual adjustments

package generated

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type DoAddPackageInfoToPostBuyFormRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoAddPackageInfoToPostBuyFormRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	TransactionId int64 `xml:"transactionId,omitempty"`

	PackageInfo *ArrayOfPackageinfostruct `xml:"packageInfo,omitempty"`
}

type DoAddPackageInfoToPostBuyFormResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doAddPackageInfoToPostBuyFormResponse"`

	PostBuyFormPackageInfo *PostBuyFormPackageInfoStruct `xml:"postBuyFormPackageInfo,omitempty"`
}

type DoAddToBlackListRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoAddToBlackListRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	UsersBlackListArray *ArrayOfUserblackliststruct `xml:"usersBlackListArray,omitempty"`
}

type DoAddToBlackListResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doAddToBlackListResponse"`

	UserBlackListResultsArr *ArrayOfUserblacklistaddresultstruct `xml:"userBlackListResultsArr,omitempty"`
}

type DoBidItemRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoBidItemRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	BidItId int64 `xml:"bidItId,omitempty"`

	BidUserPrice float32 `xml:"bidUserPrice,omitempty"`

	BidQuantity int64 `xml:"bidQuantity,omitempty"`

	BidBuyNow int64 `xml:"bidBuyNow,omitempty"`

	PharmacyRecipientData *PharmacyRecipientDataStruct `xml:"pharmacyRecipientData,omitempty"`

	VariantId string `xml:"variantId,omitempty"`
}

type DoBidItemResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doBidItemResponse"`

	BidPrice string `xml:"bidPrice,omitempty"`
}

type DoCancelBidItemRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoCancelBidItemRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	CancelItemId int64 `xml:"cancelItemId,omitempty"`

	CancelBidsArray *ArrayOfInt `xml:"cancelBidsArray,omitempty"`

	CancelBidsReason string `xml:"cancelBidsReason,omitempty"`

	CancelAddToBlackList int64 `xml:"cancelAddToBlackList,omitempty"`
}

type DoCancelBidItemResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doCancelBidItemResponse"`

	CancelBidValue int32 `xml:"cancelBidValue,omitempty"`

	CancelledBids *ArrayOfInt `xml:"cancelledBids,omitempty"`

	NotCancelledBids *ArrayOfInt `xml:"notCancelledBids,omitempty"`
}

type DoCancelRefundFormRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoCancelRefundFormRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	RefundId int32 `xml:"refundId,omitempty"`
}

type DoCancelRefundFormResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doCancelRefundFormResponse"`

	CancellationResult bool `xml:"cancellationResult,omitempty"`
}

type DoCancelRefundWarningRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoCancelRefundWarningRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	RefundId int32 `xml:"refundId,omitempty"`
}

type DoCancelRefundWarningResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doCancelRefundWarningResponse"`

	CancellationResult bool `xml:"cancellationResult,omitempty"`
}

type DoCancelTransactionRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoCancelTransactionRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	TransactionId int64 `xml:"transactionId,omitempty"`
}

type DoCancelTransactionResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doCancelTransactionResponse"`

	CancellationResult int32 `xml:"cancellationResult,omitempty"`
}

type DoChangeItemFieldsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoChangeItemFieldsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	FieldsToModify *ArrayOfFieldsvalue `xml:"fieldsToModify,omitempty"`

	FieldsToRemove *ArrayOfInt `xml:"fieldsToRemove,omitempty"`

	PreviewOnly int32 `xml:"previewOnly,omitempty"`

	Variants *ArrayOfVariantstruct `xml:"variants,omitempty"`

	Tags *ArrayOfTagnamestruct `xml:"tags,omitempty"`

	AfterSalesServiceConditions *AfterSalesServiceConditionsStruct `xml:"afterSalesServiceConditions,omitempty"`

	AdditionalServicesGroup string `xml:"additionalServicesGroup,omitempty"`
}

type DoChangeItemFieldsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doChangeItemFieldsResponse"`

	ChangedItem *ChangedItemStruct `xml:"changedItem,omitempty"`
}

type DoChangePriceItemRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoChangePriceItemRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	NewStartingPrice float32 `xml:"newStartingPrice,omitempty"`

	NewReservePrice float32 `xml:"newReservePrice,omitempty"`

	NewBuyNowPrice float32 `xml:"newBuyNowPrice,omitempty"`

	NewAdvertisementPrice float32 `xml:"newAdvertisementPrice,omitempty"`
}

type DoChangePriceItemResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doChangePriceItemResponse"`

	ItemInfo string `xml:"itemInfo,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`
}

type DoChangeQuantityItemRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoChangeQuantityItemRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	NewItemQuantity int32 `xml:"newItemQuantity,omitempty"`
}

type DoChangeQuantityItemResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doChangeQuantityItemResponse"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemInfo string `xml:"itemInfo,omitempty"`

	ItemQuantityLeft int32 `xml:"itemQuantityLeft,omitempty"`

	ItemQuantitySold int32 `xml:"itemQuantitySold,omitempty"`
}

type DoCheckItemDescriptionRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoCheckItemDescriptionRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	DescriptionContent string `xml:"descriptionContent,omitempty"`
}

type DoCheckItemDescriptionResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doCheckItemDescriptionResponse"`

	ItemDescription *ItemDescriptionStruct `xml:"itemDescription,omitempty"`
}

type DoCheckNewAuctionExtRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoCheckNewAuctionExtRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	Fields *ArrayOfFieldsvalue `xml:"fields,omitempty"`

	Variants *ArrayOfVariantstruct `xml:"variants,omitempty"`

	Tags *ArrayOfTagnamestruct `xml:"tags,omitempty"`

	AfterSalesServiceConditions *AfterSalesServiceConditionsStruct `xml:"afterSalesServiceConditions,omitempty"`
}

type DoCheckNewAuctionExtResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doCheckNewAuctionExtResponse"`

	ItemPrice string `xml:"itemPrice,omitempty"`

	ItemPriceDesc string `xml:"itemPriceDesc,omitempty"`

	ItemIsAllegroStandard int32 `xml:"itemIsAllegroStandard,omitempty"`
}

type DoFinishItemRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoFinishItemRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	FinishItemId int64 `xml:"finishItemId,omitempty"`

	FinishCancelAllBids int32 `xml:"finishCancelAllBids,omitempty"`

	FinishCancelReason string `xml:"finishCancelReason,omitempty"`
}

type DoFinishItemResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doFinishItemResponse"`

	FinishValue int32 `xml:"finishValue,omitempty"`
}

type DoFinishItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoFinishItemsRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	FinishItemsList *ArrayOfFinishitemsstruct `xml:"finishItemsList,omitempty"`
}

type DoFinishItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doFinishItemsResponse"`

	FinishItemsSucceed *ArrayOfLong `xml:"finishItemsSucceed,omitempty"`

	FinishItemsFailed *ArrayOfFinishfailurestruct `xml:"finishItemsFailed,omitempty"`
}

type DoGetArchiveRefundsListRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetArchiveRefundsListRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`
}

type DoGetArchiveRefundsListResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetArchiveRefundsListResponse"`

	RefundsCount int32 `xml:"refundsCount,omitempty"`

	RefundsList *ArrayOfArchiverefundslisttypestruct `xml:"refundsList,omitempty"`
}

type DoGetBidItem2Request struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetBidItem2Request"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`
}

type DoGetBidItem2Response struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetBidItem2Response"`

	BiditemList *ArrayOfBidliststruct2 `xml:"biditemList,omitempty"`
}

type DoGetBlackListUsersRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetBlackListUsersRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`
}

type DoGetBlackListUsersResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetBlackListUsersResponse"`

	BlackListUsers *ArrayOfBlackliststruct `xml:"blackListUsers,omitempty"`
}

type DoGetCategoryPathRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetCategoryPathRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`
}

type DoGetCategoryPathResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetCategoryPathResponse"`

	CategoryPath *ArrayOfCategorydata `xml:"categoryPath,omitempty"`
}

type DoGetCatsDataRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetCatsDataRequest"`

	CountryId int32 `xml:"countryId,omitempty"`

	LocalVersion int64 `xml:"localVersion,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`

	OnlyLeaf bool `xml:"onlyLeaf,omitempty"`
}

type DoGetCatsDataResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetCatsDataResponse"`

	CatsList *ArrayOfCatinfotype `xml:"catsList,omitempty"`

	VerKey int64 `xml:"verKey,omitempty"`

	VerStr string `xml:"verStr,omitempty"`
}

type DoGetCatsDataCountRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetCatsDataCountRequest"`

	CountryId int32 `xml:"countryId,omitempty"`

	LocalVersion int64 `xml:"localVersion,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`

	OnlyLeaf bool `xml:"onlyLeaf,omitempty"`
}

type DoGetCatsDataCountResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetCatsDataCountResponse"`

	CatsCount int32 `xml:"catsCount,omitempty"`

	VerKey int64 `xml:"verKey,omitempty"`

	VerStr string `xml:"verStr,omitempty"`
}

type DoGetCatsDataLimitRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetCatsDataLimitRequest"`

	CountryId int32 `xml:"countryId,omitempty"`

	LocalVersion int64 `xml:"localVersion,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`

	Offset int32 `xml:"offset,omitempty"`

	PackageElement int32 `xml:"packageElement,omitempty"`

	OnlyLeaf bool `xml:"onlyLeaf,omitempty"`
}

type DoGetCatsDataLimitResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetCatsDataLimitResponse"`

	CatsList *ArrayOfCatinfotype `xml:"catsList,omitempty"`

	VerKey int64 `xml:"verKey,omitempty"`

	VerStr string `xml:"verStr,omitempty"`
}

type DoGetCountriesRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetCountriesRequest"`

	CountryCode int32 `xml:"countryCode,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoGetCountriesResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetCountriesResponse"`

	CountryArray *ArrayOfCountryinfotype `xml:"countryArray,omitempty"`
}

type DoGetDealsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetDealsRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	BuyerId int32 `xml:"buyerId,omitempty"`
}

type DoGetDealsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetDealsResponse"`

	DealsList *ArrayOfDealsstruct `xml:"dealsList,omitempty"`
}

type DoGetFilledPostBuyFormsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetFilledPostBuyFormsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	PaymentType int32 `xml:"paymentType,omitempty"`

	UserRole int32 `xml:"userRole,omitempty"`

	FillingTimeFrom int64 `xml:"fillingTimeFrom,omitempty"`

	FillingTimeTo int64 `xml:"fillingTimeTo,omitempty"`
}

type DoGetFilledPostBuyFormsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetFilledPostBuyFormsResponse"`

	FilledPostBuyForms *FilledPostBuyFormsStruct `xml:"filledPostBuyForms,omitempty"`
}

type DoGetFreeDeliveryAmountRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetFreeDeliveryAmountRequest"`

	UserId int32 `xml:"userId,omitempty"`

	CountryId int32 `xml:"countryId,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoGetFreeDeliveryAmountResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetFreeDeliveryAmountResponse"`

	FreeDeliveryAmount float32 `xml:"freeDeliveryAmount,omitempty"`

	ActiveFlag int32 `xml:"activeFlag,omitempty"`
}

type DoGetItemFieldsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetItemFieldsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`
}

type DoGetItemFieldsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetItemFieldsResponse"`

	ItemFields *ArrayOfFieldsvalue `xml:"itemFields,omitempty"`

	AfterSalesServiceConditions *AfterSalesServiceConditionsStruct `xml:"afterSalesServiceConditions,omitempty"`

	AdditionalServicesGroup string `xml:"additionalServicesGroup,omitempty"`
}

type DoGetItemsImagesRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetItemsImagesRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemsArray *ArrayOfItemgetimage `xml:"itemsArray,omitempty"`

	ImageType int32 `xml:"imageType,omitempty"`
}

type DoGetItemsImagesResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetItemsImagesResponse"`

	Get_items_images_result *ArrayOfItemimagesstruct `xml:"get_items_images_result,omitempty"`
}

type DoGetItemsInfoRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetItemsInfoRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemsIdArray *ArrayOfLong `xml:"itemsIdArray,omitempty"`

	GetDesc int32 `xml:"getDesc,omitempty"`

	GetImageUrl int32 `xml:"getImageUrl,omitempty"`

	GetAttribs int32 `xml:"getAttribs,omitempty"`

	GetPostageOptions int32 `xml:"getPostageOptions,omitempty"`

	GetCompanyInfo int32 `xml:"getCompanyInfo,omitempty"`

	GetProductInfo int32 `xml:"getProductInfo,omitempty"`

	GetAfterSalesServiceConditions int32 `xml:"getAfterSalesServiceConditions,omitempty"`

	GetEan int32 `xml:"getEan,omitempty"`

	GetAdditionalServicesGroup int32 `xml:"getAdditionalServicesGroup,omitempty"`
}

type DoGetItemsInfoResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetItemsInfoResponse"`

	ArrayItemListInfo *ArrayOfIteminfostruct `xml:"arrayItemListInfo,omitempty"`

	ArrayItemsNotFound *ArrayOfLong `xml:"arrayItemsNotFound,omitempty"`

	ArrayItemsAdminKilled *ArrayOfLong `xml:"arrayItemsAdminKilled,omitempty"`
}

type DoGetItemsListRequest struct {
	XMLName xml.Name `xml:"s:DoGetItemsListRequest"`
	XMLNS   string   `xml:"xmlns:s,attr"`

	WebapiKey string `xml:"webapiKey,omitempty"`

	CountryId int32 `xml:"countryId,omitempty"`

	FilterOptions *ArrayOfFilteroptionstype `xml:"filterOptions,omitempty"`

	SortOptions *SortOptionsType `xml:"sortOptions,omitempty"`

	ResultSize int32 `xml:"resultSize,omitempty"`

	ResultOffset int32 `xml:"resultOffset,omitempty"`

	ResultScope int32 `xml:"resultScope,omitempty"`
}

type DoGetItemsListResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetItemsListResponse"`

	ItemsCount int32 `xml:"itemsCount,omitempty"`

	ItemsFeaturedCount int32 `xml:"itemsFeaturedCount,omitempty"`

	ItemsList *ArrayOfItemslisttype `xml:"itemsList,omitempty"`

	CategoriesList *CategoriesListType `xml:"categoriesList,omitempty"`

	FiltersList *ArrayOfFilterslisttype `xml:"filtersList,omitempty"`

	FiltersRejected *ArrayOfString `xml:"filtersRejected,omitempty"`
}

type DoGetMyAddressesRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyAddressesRequest"`

	SessionId string `xml:"sessionId,omitempty"`
}

type DoGetMyAddressesResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyAddressesResponse"`

	AddressesInfo *ArrayOfAddressinfostruct `xml:"addressesInfo,omitempty"`
}

type DoGetMyBidItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyBidItemsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	SortOptions *SortOptionsStruct `xml:"sortOptions,omitempty"`

	SearchValue string `xml:"searchValue,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`

	PageSize int32 `xml:"pageSize,omitempty"`

	PageNumber int32 `xml:"pageNumber,omitempty"`
}

type DoGetMyBidItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyBidItemsResponse"`

	BidItemsCounter int32 `xml:"bidItemsCounter,omitempty"`

	BidItemsList *ArrayOfBiditemstruct `xml:"bidItemsList,omitempty"`
}

type DoGetMyCurrentShipmentPriceTypeRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyCurrentShipmentPriceTypeRequest"`

	SessionId string `xml:"sessionId,omitempty"`
}

type DoGetMyCurrentShipmentPriceTypeResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyCurrentShipmentPriceTypeResponse"`

	ShipmentPriceTypeId int32 `xml:"shipmentPriceTypeId,omitempty"`
}

type DoGetMyDataRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyDataRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`
}

type DoGetMyDataResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyDataResponse"`

	UserData *UserDataStruct `xml:"userData,omitempty"`

	InvoiceData *InvoiceDataStruct `xml:"invoiceData,omitempty"`

	CompanyExtraData *CompanyExtraDataStruct `xml:"companyExtraData,omitempty"`

	CompanySecondAddress *CompanySecondAddressStruct `xml:"companySecondAddress,omitempty"`

	PharmacyData *PharmacyDataStruct `xml:"pharmacyData,omitempty"`

	AlcoholData *AlcoholDataStruct `xml:"alcoholData,omitempty"`

	RelatedPersons *RelatedPersonsStruct `xml:"relatedPersons,omitempty"`
}

type DoGetMyFutureItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyFutureItemsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	SortOptions *SortOptionsStruct `xml:"sortOptions,omitempty"`

	FilterOptions *FutureFilterOptionsStruct `xml:"filterOptions,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`

	PageSize int32 `xml:"pageSize,omitempty"`

	PageNumber int32 `xml:"pageNumber,omitempty"`
}

type DoGetMyFutureItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyFutureItemsResponse"`

	FutureItemsCounter int32 `xml:"futureItemsCounter,omitempty"`

	FutureItemsList *ArrayOfFutureitemstruct `xml:"futureItemsList,omitempty"`
}

type DoGetMyIncomingPaymentsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyIncomingPaymentsRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	BuyerId int32 `xml:"buyerId,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	TransRecvDateFrom int64 `xml:"transRecvDateFrom,omitempty"`

	TransRecvDateTo int64 `xml:"transRecvDateTo,omitempty"`

	TransPageLimit int32 `xml:"transPageLimit,omitempty"`

	TransOffset int32 `xml:"transOffset,omitempty"`

	StrictedSearch int32 `xml:"strictedSearch,omitempty"`
}

type DoGetMyIncomingPaymentsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyIncomingPaymentsResponse"`

	PayTransIncome *ArrayOfUserincomingpaymentstruct `xml:"payTransIncome,omitempty"`
}

type DoGetMyIncomingPaymentsRefundsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyIncomingPaymentsRefundsRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	BuyerId int32 `xml:"buyerId,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	Limit int32 `xml:"limit,omitempty"`

	Offset int32 `xml:"offset,omitempty"`
}

type DoGetMyIncomingPaymentsRefundsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyIncomingPaymentsRefundsResponse"`

	PayTransIncomeRefunds *ArrayOfUserincomingpaymentrefundsstruct `xml:"payTransIncomeRefunds,omitempty"`
}

type DoGetMyNotSoldItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyNotSoldItemsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	SortOptions *SortOptionsStruct `xml:"sortOptions,omitempty"`

	FilterOptions *NotSoldFilterOptionsStruct `xml:"filterOptions,omitempty"`

	SearchValue string `xml:"searchValue,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`

	PageSize int32 `xml:"pageSize,omitempty"`

	PageNumber int32 `xml:"pageNumber,omitempty"`
}

type DoGetMyNotSoldItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyNotSoldItemsResponse"`

	NotSoldItemsCounter int32 `xml:"notSoldItemsCounter,omitempty"`

	NotSoldItemsList *ArrayOfNotsolditemstruct `xml:"notSoldItemsList,omitempty"`
}

type DoGetMyNotWonItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyNotWonItemsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	SortOptions *SortOptionsStruct `xml:"sortOptions,omitempty"`

	SearchValue string `xml:"searchValue,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`

	PageSize int32 `xml:"pageSize,omitempty"`

	PageNumber int32 `xml:"pageNumber,omitempty"`
}

type DoGetMyNotWonItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyNotWonItemsResponse"`

	NotWonItemsCounter int32 `xml:"notWonItemsCounter,omitempty"`

	NotWonItemsList *ArrayOfNotwonitemstruct `xml:"notWonItemsList,omitempty"`
}

type DoGetMyPaymentsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyPaymentsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	SellerId int32 `xml:"sellerId,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	PaymentTimeFrom int64 `xml:"paymentTimeFrom,omitempty"`

	PaymentTimeTo int64 `xml:"paymentTimeTo,omitempty"`

	PageSize int32 `xml:"pageSize,omitempty"`

	PageNumber int32 `xml:"pageNumber,omitempty"`

	StrictedSearch int32 `xml:"strictedSearch,omitempty"`
}

type DoGetMyPaymentsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyPaymentsResponse"`

	PayTransPayment *ArrayOfUserpaymentstruct `xml:"payTransPayment,omitempty"`
}

type DoGetMyPaymentsInfoRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyPaymentsInfoRequest"`

	SessionId string `xml:"sessionId,omitempty"`
}

type DoGetMyPaymentsInfoResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyPaymentsInfoResponse"`

	PaymentsInfo *PaymentsInfoStruct `xml:"paymentsInfo,omitempty"`
}

type DoGetMyPaymentsRefundsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyPaymentsRefundsRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	SellerId int32 `xml:"sellerId,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	Limit int32 `xml:"limit,omitempty"`

	Offset int32 `xml:"offset,omitempty"`
}

type DoGetMyPaymentsRefundsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyPaymentsRefundsResponse"`

	PayTransPaymentRefunds *ArrayOfUserpaymentrefundsstruct `xml:"payTransPaymentRefunds,omitempty"`
}

type DoGetMyPayoutsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyPayoutsRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	TransCreateDateFrom int64 `xml:"transCreateDateFrom,omitempty"`

	TransCreateDateTo int64 `xml:"transCreateDateTo,omitempty"`

	TransPageLimit int32 `xml:"transPageLimit,omitempty"`

	TransOffset int32 `xml:"transOffset,omitempty"`
}

type DoGetMyPayoutsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyPayoutsResponse"`

	PayTransPayout *ArrayOfUserpayoutstruct `xml:"payTransPayout,omitempty"`
}

type DoGetMyPayoutsDetailsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyPayoutsDetailsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	PayoutId int32 `xml:"payoutId,omitempty"`

	Limit int32 `xml:"limit,omitempty"`

	Offset int32 `xml:"offset,omitempty"`
}

type DoGetMyPayoutsDetailsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyPayoutsDetailsResponse"`

	PaymentsCount int32 `xml:"paymentsCount,omitempty"`

	Payments *ArrayOfPayoutpaymentsstruct `xml:"payments,omitempty"`

	RefundsFromCount int32 `xml:"refundsFromCount,omitempty"`

	RefundFrom *ArrayOfPayoutrefundfromstruct `xml:"refundFrom,omitempty"`

	RefundsToCount int32 `xml:"refundsToCount,omitempty"`

	RefundTo *ArrayOfPayoutrefundtostruct `xml:"refundTo,omitempty"`
}

type DoGetMySellItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMySellItemsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	SortOptions *SortOptionsStruct `xml:"sortOptions,omitempty"`

	FilterOptions *SellFilterOptionsStruct `xml:"filterOptions,omitempty"`

	SearchValue string `xml:"searchValue,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`

	PageSize int32 `xml:"pageSize,omitempty"`

	PageNumber int32 `xml:"pageNumber,omitempty"`
}

type DoGetMySellItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMySellItemsResponse"`

	SellItemsCounter int32 `xml:"sellItemsCounter,omitempty"`

	SellItemsList *ArrayOfSellitemstruct `xml:"sellItemsList,omitempty"`
}

type DoGetMySoldItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMySoldItemsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	SortOptions *SortOptionsStruct `xml:"sortOptions,omitempty"`

	FilterOptions *SoldFilterOptionsStruct `xml:"filterOptions,omitempty"`

	SearchValue string `xml:"searchValue,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`

	PageSize int32 `xml:"pageSize,omitempty"`

	PageNumber int32 `xml:"pageNumber,omitempty"`
}

type DoGetMySoldItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMySoldItemsResponse"`

	SoldItemsCounter int32 `xml:"soldItemsCounter,omitempty"`

	SoldItemsList *ArrayOfSolditemstruct `xml:"soldItemsList,omitempty"`
}

type DoGetMyWonItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetMyWonItemsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	SortOptions *SortOptionsStruct `xml:"sortOptions,omitempty"`

	SearchValue string `xml:"searchValue,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`

	PageSize int32 `xml:"pageSize,omitempty"`

	PageNumber int32 `xml:"pageNumber,omitempty"`
}

type DoGetMyWonItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetMyWonItemsResponse"`

	WonItemsCounter int32 `xml:"wonItemsCounter,omitempty"`

	WonItemsList *ArrayOfWonitemstruct `xml:"wonItemsList,omitempty"`
}

type DoGetPaymentMethodsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetPaymentMethodsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`
}

type DoGetPaymentMethodsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetPaymentMethodsResponse"`

	PaymentMethods *ArrayOfPaymentmethodstruct `xml:"paymentMethods,omitempty"`
}

type DoGetPostBuyDataRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetPostBuyDataRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemsArray *ArrayOfLong `xml:"itemsArray,omitempty"`

	BuyerFilterArray *ArrayOfLong `xml:"buyerFilterArray,omitempty"`
}

type DoGetPostBuyDataResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetPostBuyDataResponse"`

	ItemsPostBuyData *ArrayOfItempostbuydatastruct `xml:"itemsPostBuyData,omitempty"`
}

type DoGetPostBuyFormsDataForBuyersRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetPostBuyFormsDataForBuyersRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	TransactionsIdsArray *ArrayOfLong `xml:"transactionsIdsArray,omitempty"`
}

type DoGetPostBuyFormsDataForBuyersResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetPostBuyFormsDataForBuyersResponse"`

	PostBuyFormDataForBuyers *ArrayOfPostbuyformforbuyersdatastruct `xml:"postBuyFormDataForBuyers,omitempty"`
}

type DoGetPostBuyFormsDataForSellersRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetPostBuyFormsDataForSellersRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	TransactionsIdsArray *ArrayOfLong `xml:"transactionsIdsArray,omitempty"`
}

type DoGetPostBuyFormsDataForSellersResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetPostBuyFormsDataForSellersResponse"`

	PostBuyFormData *ArrayOfPostbuyformdatastruct `xml:"postBuyFormData,omitempty"`
}

type DoGetPostBuyFormsIdsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetPostBuyFormsIdsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	FilterOptions *ArrayOfFilteroptionstype `xml:"filterOptions,omitempty"`

	ResultSize int32 `xml:"resultSize,omitempty"`

	ResultOffset int32 `xml:"resultOffset,omitempty"`
}

type DoGetPostBuyFormsIdsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetPostBuyFormsIdsResponse"`

	FormsCount int32 `xml:"formsCount,omitempty"`

	FormsIds *ArrayOfLong `xml:"formsIds,omitempty"`

	FiltersList *ArrayOfFilterslisttype `xml:"filtersList,omitempty"`
}

type DoGetPostBuyItemInfoRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetPostBuyItemInfoRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`
}

type DoGetPostBuyItemInfoResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetPostBuyItemInfoResponse"`

	ItemPostBuyFormInfo *ArrayOfPostbuyiteminfostruct `xml:"itemPostBuyFormInfo,omitempty"`
}

type DoGetRefundsDealsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetRefundsDealsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	FilterOptions *ArrayOfFilteroptionstype `xml:"filterOptions,omitempty"`

	SortOrder string `xml:"sortOrder,omitempty"`

	ResultSize int32 `xml:"resultSize,omitempty"`

	ResultOffset int32 `xml:"resultOffset,omitempty"`
}

type DoGetRefundsDealsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetRefundsDealsResponse"`

	DealsCount int32 `xml:"dealsCount,omitempty"`

	DealsList *ArrayOfRefundsdealslisttype `xml:"dealsList,omitempty"`

	FiltersList *ArrayOfFilterslisttype `xml:"filtersList,omitempty"`
}

type DoGetRefundsListRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetRefundsListRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	FilterOptions *ArrayOfFilteroptionstype `xml:"filterOptions,omitempty"`

	ResultSize int32 `xml:"resultSize,omitempty"`

	ResultOffset int32 `xml:"resultOffset,omitempty"`
}

type DoGetRefundsListResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetRefundsListResponse"`

	RefundsCount int32 `xml:"refundsCount,omitempty"`

	RefundsList *ArrayOfRefundlisttype `xml:"refundsList,omitempty"`

	FiltersList *ArrayOfFilterslisttype `xml:"filtersList,omitempty"`
}

type DoGetRefundsReasonsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetRefundsReasonsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	DealId int32 `xml:"dealId,omitempty"`
}

type DoGetRefundsReasonsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetRefundsReasonsResponse"`

	ReasonsCount int32 `xml:"reasonsCount,omitempty"`

	ReasonsList *ArrayOfReasoninfotype `xml:"reasonsList,omitempty"`
}

type DoGetRelatedItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetRelatedItemsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`
}

type DoGetRelatedItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetRelatedItemsResponse"`

	RelatedItems *RelatedItemsStruct `xml:"relatedItems,omitempty"`
}

type DoGetSellFormFieldsForCategoryRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetSellFormFieldsForCategoryRequest"`

	WebapiKey string `xml:"webapiKey,omitempty"`

	CountryId int32 `xml:"countryId,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`
}

type DoGetSellFormFieldsForCategoryResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetSellFormFieldsForCategoryResponse"`

	SellFormFieldsForCategory *SellFormFieldsForCategoryStruct `xml:"sellFormFieldsForCategory,omitempty"`
}

type DoGetSessionHandleForWidgetRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetSessionHandleForWidgetRequest"`

	WebapiKey string `xml:"webapiKey,omitempty"`

	CountryCode int32 `xml:"countryCode,omitempty"`
}

type DoGetSessionHandleForWidgetResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetSessionHandleForWidgetResponse"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ServerTime int64 `xml:"serverTime,omitempty"`
}

type DoGetShipmentDataRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetShipmentDataRequest"`

	CountryId int32 `xml:"countryId,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoGetShipmentDataResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetShipmentDataResponse"`

	ShipmentDataList *ArrayOfShipmentdatastruct `xml:"shipmentDataList,omitempty"`

	LocalVersion int64 `xml:"localVersion,omitempty"`
}

type DoGetShipmentDataForRelatedItemsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetShipmentDataForRelatedItemsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	ItemIds *ArrayOfLong `xml:"itemIds,omitempty"`
}

type DoGetShipmentDataForRelatedItemsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetShipmentDataForRelatedItemsResponse"`

	RelatedItemsShipmentData *RelatedItemsShipmentDataStruct `xml:"relatedItemsShipmentData,omitempty"`
}

type DoGetShipmentPriceTypesRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetShipmentPriceTypesRequest"`

	CountryId int32 `xml:"countryId,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoGetShipmentPriceTypesResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetShipmentPriceTypesResponse"`

	ShipmentPriceTypes *ArrayOfShipmentpricetypestruct `xml:"shipmentPriceTypes,omitempty"`
}

type DoGetSiteJournalRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetSiteJournalRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	StartingPoint int64 `xml:"startingPoint,omitempty"`

	InfoType int32 `xml:"infoType,omitempty"`
}

type DoGetSiteJournalResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetSiteJournalResponse"`

	SiteJournalArray *ArrayOfSitejournal `xml:"siteJournalArray,omitempty"`
}

type DoGetSiteJournalDealsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetSiteJournalDealsRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	JournalStart int64 `xml:"journalStart,omitempty"`
}

type DoGetSiteJournalDealsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetSiteJournalDealsResponse"`

	SiteJournalDeals *ArrayOfSitejournaldealsstruct `xml:"siteJournalDeals,omitempty"`
}

type DoGetSiteJournalDealsInfoRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetSiteJournalDealsInfoRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	JournalStart int64 `xml:"journalStart,omitempty"`
}

type DoGetSiteJournalDealsInfoResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetSiteJournalDealsInfoResponse"`

	SiteJournalDealsInfo *SiteJournalDealsInfoStruct `xml:"siteJournalDealsInfo,omitempty"`
}

type DoGetSiteJournalInfoRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetSiteJournalInfoRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	StartingPoint int64 `xml:"startingPoint,omitempty"`

	InfoType int32 `xml:"infoType,omitempty"`
}

type DoGetSiteJournalInfoResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetSiteJournalInfoResponse"`

	SiteJournalInfo *SiteJournalInfo `xml:"siteJournalInfo,omitempty"`
}

type DoGetStatesInfoRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetStatesInfoRequest"`

	CountryCode int32 `xml:"countryCode,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoGetStatesInfoResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetStatesInfoResponse"`

	StatesInfoArray *ArrayOfStateinfostruct `xml:"statesInfoArray,omitempty"`
}

type DoGetSystemTimeRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetSystemTimeRequest"`

	CountryId int32 `xml:"countryId,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoGetSystemTimeResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetSystemTimeResponse"`

	ServerTime int64 `xml:"serverTime,omitempty"`
}

type DoGetTransactionsIDsRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetTransactionsIDsRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemsIdArray *ArrayOfLong `xml:"itemsIdArray,omitempty"`

	UserRole string `xml:"userRole,omitempty"`

	ShipmentIdArray *ArrayOfLong `xml:"shipmentIdArray,omitempty"`
}

type DoGetTransactionsIDsResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetTransactionsIDsResponse"`

	TransactionsIdsArray *ArrayOfLong `xml:"transactionsIdsArray,omitempty"`
}

type DoGetUserIDRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetUserIDRequest"`

	CountryId int32 `xml:"countryId,omitempty"`

	UserLogin string `xml:"userLogin,omitempty"`

	UserEmail string `xml:"userEmail,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoGetUserIDResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetUserIDResponse"`

	UserId int32 `xml:"userId,omitempty"`
}

type DoGetUserLoginRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoGetUserLoginRequest"`

	CountryId int32 `xml:"countryId,omitempty"`

	UserId int32 `xml:"userId,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoGetUserLoginResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doGetUserLoginResponse"`

	UserLogin string `xml:"userLogin,omitempty"`
}

type DoLoginRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoLoginRequest"`

	UserLogin string `xml:"userLogin,omitempty"`

	UserPassword string `xml:"userPassword,omitempty"`

	CountryCode int32 `xml:"countryCode,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`

	LocalVersion int64 `xml:"localVersion,omitempty"`
}

type DoLoginResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doLoginResponse"`

	SessionHandlePart string `xml:"sessionHandlePart,omitempty"`

	UserId int64 `xml:"userId,omitempty"`

	ServerTime int64 `xml:"serverTime,omitempty"`
}

type DoLoginEncRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoLoginEncRequest"`

	UserLogin string `xml:"userLogin,omitempty"`

	UserHashPassword string `xml:"userHashPassword,omitempty"`

	CountryCode int32 `xml:"countryCode,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`

	LocalVersion int64 `xml:"localVersion,omitempty"`
}

type DoLoginEncResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doLoginEncResponse"`

	SessionHandlePart string `xml:"sessionHandlePart,omitempty"`

	UserId int64 `xml:"userId,omitempty"`

	ServerTime int64 `xml:"serverTime,omitempty"`
}

type DoLoginWithAccessTokenRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoLoginWithAccessTokenRequest"`

	AccessToken string `xml:"accessToken,omitempty"`

	CountryCode int32 `xml:"countryCode,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoLoginWithAccessTokenResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doLoginWithAccessTokenResponse"`

	SessionHandlePart string `xml:"sessionHandlePart,omitempty"`

	UserId int64 `xml:"userId,omitempty"`

	ServerTime int64 `xml:"serverTime,omitempty"`
}

type DoMyAccount2Request struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoMyAccount2Request"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	AccountType string `xml:"accountType,omitempty"`

	Offset int32 `xml:"offset,omitempty"`

	ItemsArray *ArrayOfLong `xml:"itemsArray,omitempty"`

	Limit int32 `xml:"limit,omitempty"`
}

type DoMyAccount2Response struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doMyAccount2Response"`

	MyaccountList *ArrayOfMyaccountstruct2 `xml:"myaccountList,omitempty"`
}

type DoMyAccountItemsCountRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoMyAccountItemsCountRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	AccountType string `xml:"accountType,omitempty"`

	ItemsArray *ArrayOfLong `xml:"itemsArray,omitempty"`
}

type DoMyAccountItemsCountResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doMyAccountItemsCountResponse"`

	MyaccountItemsCount int32 `xml:"myaccountItemsCount,omitempty"`
}

type DoMyBillingRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoMyBillingRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`
}

type DoMyBillingResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doMyBillingResponse"`

	MyBilling string `xml:"myBilling,omitempty"`
}

type DoMyBillingItemRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoMyBillingItemRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	Option string `xml:"option,omitempty"`
}

type DoMyBillingItemResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doMyBillingItemResponse"`

	EntryFees *ArrayOfItembilling `xml:"entryFees,omitempty"`

	EndingFees *ArrayOfItembilling `xml:"endingFees,omitempty"`
}

type DoMyContactRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoMyContactRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	AuctionIdList *ArrayOfLong `xml:"auctionIdList,omitempty"`

	Offset int32 `xml:"offset,omitempty"`

	Desc int32 `xml:"desc,omitempty"`
}

type DoMyContactResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doMyContactResponse"`

	MycontactList *ArrayOfMycontactlist `xml:"mycontactList,omitempty"`
}

type DoNewAuctionExtRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoNewAuctionExtRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	Fields *ArrayOfFieldsvalue `xml:"fields,omitempty"`

	ItemTemplateId int32 `xml:"itemTemplateId,omitempty"`

	LocalId int32 `xml:"localId,omitempty"`

	ItemTemplateCreate *ItemTemplateCreateStruct `xml:"itemTemplateCreate,omitempty"`

	Variants *ArrayOfVariantstruct `xml:"variants,omitempty"`

	Tags *ArrayOfTagnamestruct `xml:"tags,omitempty"`

	AfterSalesServiceConditions *AfterSalesServiceConditionsStruct `xml:"afterSalesServiceConditions,omitempty"`

	AdditionalServicesGroup string `xml:"additionalServicesGroup,omitempty"`
}

type DoNewAuctionExtResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doNewAuctionExtResponse"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemInfo string `xml:"itemInfo,omitempty"`

	ItemIsAllegroStandard int32 `xml:"itemIsAllegroStandard,omitempty"`
}

type DoQueryAllSysStatusRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoQueryAllSysStatusRequest"`

	CountryId int32 `xml:"countryId,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoQueryAllSysStatusResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doQueryAllSysStatusResponse"`

	SysCountryStatus *ArrayOfSysstatustype `xml:"sysCountryStatus,omitempty"`
}

type DoQuerySysStatusRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoQuerySysStatusRequest"`

	Sysvar int32 `xml:"sysvar,omitempty"`

	CountryId int32 `xml:"countryId,omitempty"`

	WebapiKey string `xml:"webapiKey,omitempty"`
}

type DoQuerySysStatusResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doQuerySysStatusResponse"`

	Info string `xml:"info,omitempty"`

	VerKey int64 `xml:"verKey,omitempty"`
}

type DoRemoveFromBlackListRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoRemoveFromBlackListRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	UsersIdArray *ArrayOfLong `xml:"usersIdArray,omitempty"`
}

type DoRemoveFromBlackListResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doRemoveFromBlackListResponse"`

	BlackListResult *ArrayOfBlacklistresstruct `xml:"blackListResult,omitempty"`
}

type DoRequestCancelBidRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoRequestCancelBidRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	RequestItemId int64 `xml:"requestItemId,omitempty"`

	RequestCancelReason string `xml:"requestCancelReason,omitempty"`
}

type DoRequestCancelBidResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doRequestCancelBidResponse"`

	RequestValue int32 `xml:"requestValue,omitempty"`
}

type DoRequestPayoutRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoRequestPayoutRequest"`

	SessionId string `xml:"sessionId,omitempty"`
}

type DoRequestPayoutResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doRequestPayoutResponse"`

	RequestPayout *RequestPayoutStruct `xml:"requestPayout,omitempty"`
}

type DoRequestSurchargeRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoRequestSurchargeRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	TransactionId int64 `xml:"transactionId,omitempty"`

	SurchargeValue float32 `xml:"surchargeValue,omitempty"`

	SurchargeMessageToBuyer string `xml:"surchargeMessageToBuyer,omitempty"`
}

type DoRequestSurchargeResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doRequestSurchargeResponse"`

	SurchargeResult int32 `xml:"surchargeResult,omitempty"`
}

type DoSellSomeAgainRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoSellSomeAgainRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	SellItemsArray *ArrayOfLong `xml:"sellItemsArray,omitempty"`

	SellStartingTime int64 `xml:"sellStartingTime,omitempty"`

	SellAuctionDuration int32 `xml:"sellAuctionDuration,omitempty"`

	SellOptions int32 `xml:"sellOptions,omitempty"`

	LocalIds *ArrayOfInt `xml:"localIds,omitempty"`

	SellProlongOptions int32 `xml:"sellProlongOptions,omitempty"`
}

type DoSellSomeAgainResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doSellSomeAgainResponse"`

	ItemsSellAgain *ArrayOfStructsellagain `xml:"itemsSellAgain,omitempty"`

	ItemsSellFailed *ArrayOfStructsellfailed `xml:"itemsSellFailed,omitempty"`

	ItemsSellNotFound *ArrayOfLong `xml:"itemsSellNotFound,omitempty"`
}

type DoSellSomeAgainInShopRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoSellSomeAgainInShopRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	SellItemsArray *ArrayOfLong `xml:"sellItemsArray,omitempty"`

	SellStartingTime int64 `xml:"sellStartingTime,omitempty"`

	SellShopDuration int32 `xml:"sellShopDuration,omitempty"`

	SellShopOptions int32 `xml:"sellShopOptions,omitempty"`

	SellProlongOptions int32 `xml:"sellProlongOptions,omitempty"`

	SellShopCategory int64 `xml:"sellShopCategory,omitempty"`

	LocalIds *ArrayOfInt `xml:"localIds,omitempty"`
}

type DoSellSomeAgainInShopResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doSellSomeAgainInShopResponse"`

	ItemsSellAgain *ArrayOfStructsellagain `xml:"itemsSellAgain,omitempty"`

	ItemsSellFailed *ArrayOfStructsellfailed `xml:"itemsSellFailed,omitempty"`

	ItemsSellNotFound *ArrayOfLong `xml:"itemsSellNotFound,omitempty"`
}

type DoSendEmailToUserRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoSendEmailToUserRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	MailToUserItemId int64 `xml:"mailToUserItemId,omitempty"`

	MailToUserReceiverId int64 `xml:"mailToUserReceiverId,omitempty"`

	MailToUserSubjectId int32 `xml:"mailToUserSubjectId,omitempty"`

	MailToUserOption int32 `xml:"mailToUserOption,omitempty"`

	MailToUserMessage string `xml:"mailToUserMessage,omitempty"`
}

type DoSendEmailToUserResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doSendEmailToUserResponse"`

	MailToUserReceiverId int64 `xml:"mailToUserReceiverId,omitempty"`

	MailToUserResult string `xml:"mailToUserResult,omitempty"`
}

type DoSendPostBuyFormRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoSendPostBuyFormRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	NewPostBuyFormSeller *ArrayOfNewpostbuyformsellerstruct `xml:"newPostBuyFormSeller,omitempty"`

	NewPostBuyFormCommon *NewPostBuyFormCommonStruct `xml:"newPostBuyFormCommon,omitempty"`
}

type DoSendPostBuyFormResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doSendPostBuyFormResponse"`

	PostBuyForm *PostBuyFormStruct `xml:"postBuyForm,omitempty"`
}

type DoSendRefundFormRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoSendRefundFormRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	DealId int32 `xml:"dealId,omitempty"`

	ReasonId int32 `xml:"reasonId,omitempty"`

	RefundQuantity int32 `xml:"refundQuantity,omitempty"`
}

type DoSendRefundFormResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doSendRefundFormResponse"`

	RefundId int32 `xml:"refundId,omitempty"`
}

type DoSetFreeDeliveryAmountRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoSetFreeDeliveryAmountRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	ActiveFlag int32 `xml:"activeFlag,omitempty"`

	FreeDeliveryAmount float32 `xml:"freeDeliveryAmount,omitempty"`
}

type DoSetFreeDeliveryAmountResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doSetFreeDeliveryAmountResponse"`

	ResponseStatus bool `xml:"responseStatus,omitempty"`
}

type DoSetShipmentPriceTypeRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoSetShipmentPriceTypeRequest"`

	SessionId string `xml:"sessionId,omitempty"`

	ShipmentPriceTypeId int32 `xml:"shipmentPriceTypeId,omitempty"`
}

type DoSetShipmentPriceTypeResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doSetShipmentPriceTypeResponse"`

	OperationResult int32 `xml:"operationResult,omitempty"`
}

type DoShowItemInfoExtRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoShowItemInfoExtRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	GetDesc int32 `xml:"getDesc,omitempty"`

	GetImageUrl int32 `xml:"getImageUrl,omitempty"`

	GetAttribs int32 `xml:"getAttribs,omitempty"`

	GetPostageOptions int32 `xml:"getPostageOptions,omitempty"`

	GetCompanyInfo int32 `xml:"getCompanyInfo,omitempty"`

	GetProductInfo int32 `xml:"getProductInfo,omitempty"`

	GetAfterSalesServiceConditions int32 `xml:"getAfterSalesServiceConditions,omitempty"`

	GetEan int32 `xml:"getEan,omitempty"`

	GetAdditionalServicesGroup int32 `xml:"getAdditionalServicesGroup,omitempty"`
}

type DoShowItemInfoExtResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doShowItemInfoExtResponse"`

	ItemListInfoExt *ItemInfoExt `xml:"itemListInfoExt,omitempty"`

	ItemCatPath *ArrayOfItemcatlist `xml:"itemCatPath,omitempty"`

	ItemImgList *ArrayOfItemimagelist `xml:"itemImgList,omitempty"`

	ItemAttribList *ArrayOfAttribstruct `xml:"itemAttribList,omitempty"`

	ItemPostageOptions *ArrayOfPostagestruct `xml:"itemPostageOptions,omitempty"`

	ItemPaymentOptions *ItemPaymentOptions `xml:"itemPaymentOptions,omitempty"`

	ItemCompanyInfo *CompanyInfoStruct `xml:"itemCompanyInfo,omitempty"`

	ItemProductInfo *ProductStruct `xml:"itemProductInfo,omitempty"`

	ItemVariants string `xml:"itemVariants,omitempty"`

	ItemAfterSalesServiceConditions *AfterSalesServiceConditionsStruct `xml:"itemAfterSalesServiceConditions,omitempty"`

	ItemAdditionalServicesGroup string `xml:"itemAdditionalServicesGroup,omitempty"`
}

type DoShowUserRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoShowUserRequest"`

	WebapiKey string `xml:"webapiKey,omitempty"`

	CountryId int32 `xml:"countryId,omitempty"`

	UserId int64 `xml:"userId,omitempty"`

	UserLogin string `xml:"userLogin,omitempty"`
}

type DoShowUserResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doShowUserResponse"`

	UserId int64 `xml:"userId,omitempty"`

	UserLogin string `xml:"userLogin,omitempty"`

	UserCountry int32 `xml:"userCountry,omitempty"`

	UserCreateDate int64 `xml:"userCreateDate,omitempty"`

	UserLoginDate int64 `xml:"userLoginDate,omitempty"`

	UserRating int32 `xml:"userRating,omitempty"`

	UserIsNewUser int32 `xml:"userIsNewUser,omitempty"`

	UserNotActivated int32 `xml:"userNotActivated,omitempty"`

	UserClosed int32 `xml:"userClosed,omitempty"`

	UserBlocked int32 `xml:"userBlocked,omitempty"`

	UserTerminated int32 `xml:"userTerminated,omitempty"`

	UserHasPage int32 `xml:"userHasPage,omitempty"`

	UserIsSseller int32 `xml:"userIsSseller,omitempty"`

	UserIsEco int32 `xml:"userIsEco,omitempty"`

	UserPositiveFeedback *ShowUserFeedbacks `xml:"userPositiveFeedback,omitempty"`

	UserNegativeFeedback *ShowUserFeedbacks `xml:"userNegativeFeedback,omitempty"`

	UserNeutralFeedback *ShowUserFeedbacks `xml:"userNeutralFeedback,omitempty"`

	UserJuniorStatus int32 `xml:"userJuniorStatus,omitempty"`

	UserHasShop int32 `xml:"userHasShop,omitempty"`

	UserCompanyIcon int32 `xml:"userCompanyIcon,omitempty"`

	UserSellRatingCount int32 `xml:"userSellRatingCount,omitempty"`

	UserSellRatingAverage *ArrayOfSellratingaveragestruct `xml:"userSellRatingAverage,omitempty"`

	UserIsAllegroStandard int32 `xml:"userIsAllegroStandard,omitempty"`

	UserIsB2cSeller int32 `xml:"userIsB2cSeller,omitempty"`
}

type DoVerifyItemRequest struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DoVerifyItemRequest"`

	SessionHandle string `xml:"sessionHandle,omitempty"`

	LocalId int32 `xml:"localId,omitempty"`
}

type DoVerifyItemResponse struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php doVerifyItemResponse"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemListed int32 `xml:"itemListed,omitempty"`

	ItemStartingTime int64 `xml:"itemStartingTime,omitempty"`
}

type PackageInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PackageInfoStruct"`

	OperatorId int32 `xml:"operatorId,omitempty"`

	PackageId string `xml:"packageId,omitempty"`

	OperatorName string `xml:"operatorName,omitempty"`
}

type ArrayOfPackageinfostruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPackageinfostruct"`

	Item []*PackageInfoStruct `xml:"item,omitempty"`
}

type ArrayOfString struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfString"`

	Item []string `xml:"item,omitempty"`
}

type PostBuyFormPackageInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormPackageInfoStruct"`

	PackageIdsAdded *ArrayOfString `xml:"packageIdsAdded,omitempty"`

	PackageIdsNotAddedIncorrectOperatorId *ArrayOfString `xml:"packageIdsNotAddedIncorrectOperatorId,omitempty"`

	PackageIdsNotAddedIncorrectPackageId *ArrayOfString `xml:"packageIdsNotAddedIncorrectPackageId,omitempty"`
}

type UserBlackListStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserBlackListStruct"`

	UserId int32 `xml:"userId,omitempty"`

	UserBlackListNote string `xml:"userBlackListNote,omitempty"`
}

type ArrayOfUserblackliststruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfUserblackliststruct"`

	Item []*UserBlackListStruct `xml:"item,omitempty"`
}

type UserBlackListAddResultStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserBlackListAddResultStruct"`

	UserId int32 `xml:"userId,omitempty"`

	AddToBlackListResult int32 `xml:"addToBlackListResult,omitempty"`

	AddToBlackListErrCode string `xml:"addToBlackListErrCode,omitempty"`

	AddToBlackListErrMsg string `xml:"addToBlackListErrMsg,omitempty"`
}

type ArrayOfUserblacklistaddresultstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfUserblacklistaddresultstruct"`

	Item []*UserBlackListAddResultStruct `xml:"item,omitempty"`
}

type PharmacyRecipientDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PharmacyRecipientDataStruct"`

	RecipientFirstName string `xml:"recipientFirstName,omitempty"`

	RecipientLastName string `xml:"recipientLastName,omitempty"`

	RecipientAddress string `xml:"recipientAddress,omitempty"`

	RecipientPostcode string `xml:"recipientPostcode,omitempty"`

	RecipientCity string `xml:"recipientCity,omitempty"`
}

type ArrayOfInt struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfInt"`

	Item []int32 `xml:"item,omitempty"`
}

type RangeIntValueStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RangeIntValueStruct"`

	FvalueRangeIntMin int32 `xml:"fvalueRangeIntMin,omitempty"`

	FvalueRangeIntMax int32 `xml:"fvalueRangeIntMax,omitempty"`
}

type RangeFloatValueStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RangeFloatValueStruct"`

	FvalueRangeFloatMin float32 `xml:"fvalueRangeFloatMin,omitempty"`

	FvalueRangeFloatMax float32 `xml:"fvalueRangeFloatMax,omitempty"`
}

type RangeDateValueStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RangeDateValueStruct"`

	FvalueRangeDateMin string `xml:"fvalueRangeDateMin,omitempty"`

	FvalueRangeDateMax string `xml:"fvalueRangeDateMax,omitempty"`
}

type FieldsValue struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FieldsValue"`

	Fid int32 `xml:"fid,omitempty"`

	FvalueString string `xml:"fvalueString,omitempty"`

	FvalueInt int32 `xml:"fvalueInt,omitempty"`

	FvalueFloat float32 `xml:"fvalueFloat,omitempty"`

	FvalueImage []byte `xml:"fvalueImage,omitempty"`

	FvalueDatetime int64 `xml:"fvalueDatetime,omitempty"`

	FvalueDate string `xml:"fvalueDate,omitempty"`

	FvalueRangeInt *RangeIntValueStruct `xml:"fvalueRangeInt,omitempty"`

	FvalueRangeFloat *RangeFloatValueStruct `xml:"fvalueRangeFloat,omitempty"`

	FvalueRangeDate *RangeDateValueStruct `xml:"fvalueRangeDate,omitempty"`
}

type ArrayOfFieldsvalue struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfFieldsvalue"`

	Item []*FieldsValue `xml:"item,omitempty"`
}

type VariantQuantityStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php VariantQuantityStruct"`

	Mask int32 `xml:"mask,omitempty"`

	Quantity int32 `xml:"quantity,omitempty"`
}

type ArrayOfVariantquantitystruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfVariantquantitystruct"`

	Item []*VariantQuantityStruct `xml:"item,omitempty"`
}

type VariantStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php VariantStruct"`

	Fid int32 `xml:"fid,omitempty"`

	Quantities *ArrayOfVariantquantitystruct `xml:"quantities,omitempty"`
}

type ArrayOfVariantstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfVariantstruct"`

	Item []*VariantStruct `xml:"item,omitempty"`
}

type TagNameStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php TagNameStruct"`

	TagName string `xml:"tagName,omitempty"`
}

type ArrayOfTagnamestruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfTagnamestruct"`

	Item []*TagNameStruct `xml:"item,omitempty"`
}

type AfterSalesServiceConditionsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php AfterSalesServiceConditionsStruct"`

	ImpliedWarranty string `xml:"impliedWarranty,omitempty"`

	ReturnPolicy string `xml:"returnPolicy,omitempty"`

	Warranty string `xml:"warranty,omitempty"`
}

type AmountStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php AmountStruct"`

	AmountValue float32 `xml:"amountValue,omitempty"`

	AmountCurrencySign string `xml:"amountCurrencySign,omitempty"`
}

type ItemSurchargeStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemSurchargeStruct"`

	SurchargeDescription string `xml:"surchargeDescription,omitempty"`

	SurchargeAmount *AmountStruct `xml:"surchargeAmount,omitempty"`
}

type ArrayOfItemsurchargestruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfItemsurchargestruct"`

	Item []*ItemSurchargeStruct `xml:"item,omitempty"`
}

type ChangedItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ChangedItemStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemFields *ArrayOfFieldsvalue `xml:"itemFields,omitempty"`

	ItemSurcharge *ArrayOfItemsurchargestruct `xml:"itemSurcharge,omitempty"`
}

type ItemDescriptionStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemDescriptionStruct"`

	DescriptionResult string `xml:"descriptionResult,omitempty"`
}

type FinishItemsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FinishItemsStruct"`

	FinishItemId int64 `xml:"finishItemId,omitempty"`

	FinishCancelAllBids int32 `xml:"finishCancelAllBids,omitempty"`

	FinishCancelReason string `xml:"finishCancelReason,omitempty"`
}

type ArrayOfFinishitemsstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfFinishitemsstruct"`

	Item []*FinishItemsStruct `xml:"item,omitempty"`
}

type ArrayOfLong struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfLong"`

	Item []int64 `xml:"item,omitempty"`
}

type FinishFailureStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FinishFailureStruct"`

	FinishItemId int64 `xml:"finishItemId,omitempty"`

	FinishErrorCode string `xml:"finishErrorCode,omitempty"`

	FinishErrorMessage string `xml:"finishErrorMessage,omitempty"`
}

type ArrayOfFinishfailurestruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfFinishfailurestruct"`

	Item []*FinishFailureStruct `xml:"item,omitempty"`
}

type ArchiveRefundsListTypeStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArchiveRefundsListTypeStruct"`

	RefundId int32 `xml:"refundId,omitempty"`

	BuyerId int32 `xml:"buyerId,omitempty"`

	BuyerLogin string `xml:"buyerLogin,omitempty"`
}

type ArrayOfArchiverefundslisttypestruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfArchiverefundslisttypestruct"`

	Item []*ArchiveRefundsListTypeStruct `xml:"item,omitempty"`
}

type BidListStruct2 struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php BidListStruct2"`

	BidsArray *ArrayOfString `xml:"bidsArray,omitempty"`
}

type ArrayOfBidliststruct2 struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfBidliststruct2"`

	Item []*BidListStruct2 `xml:"item,omitempty"`
}

type BlackListStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php BlackListStruct"`

	UserId int64 `xml:"userId,omitempty"`

	UserLogin string `xml:"userLogin,omitempty"`

	UserRating int32 `xml:"userRating,omitempty"`

	UserCountry int32 `xml:"userCountry,omitempty"`
}

type ArrayOfBlackliststruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfBlackliststruct"`

	Item []*BlackListStruct `xml:"item,omitempty"`
}

type CategoryData struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php CategoryData"`

	CatId int32 `xml:"catId,omitempty"`

	CatParent int32 `xml:"catParent,omitempty"`

	CatCountry int32 `xml:"catCountry,omitempty"`

	CatLevel int32 `xml:"catLevel,omitempty"`

	CatIsLeaf int32 `xml:"catIsLeaf,omitempty"`

	CatName string `xml:"catName,omitempty"`

	CatOptions int32 `xml:"catOptions,omitempty"`
}

type ArrayOfCategorydata struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfCategorydata"`

	Item []*CategoryData `xml:"item,omitempty"`
}

type CatInfoType struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php CatInfoType"`

	CatId int32 `xml:"catId,omitempty"`

	CatName string `xml:"catName,omitempty"`

	CatParent int32 `xml:"catParent,omitempty"`

	CatPosition int32 `xml:"catPosition,omitempty"`

	CatIsProductCatalogueEnabled int32 `xml:"catIsProductCatalogueEnabled,omitempty"`

	CatIsLeaf bool `xml:"catIsLeaf,omitempty"`
}

type ArrayOfCatinfotype struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfCatinfotype"`

	Item []*CatInfoType `xml:"item,omitempty"`
}

type CountryInfoType struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php CountryInfoType"`

	CountryId int32 `xml:"countryId,omitempty"`

	CountryName string `xml:"countryName,omitempty"`
}

type ArrayOfCountryinfotype struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfCountryinfotype"`

	Item []*CountryInfoType `xml:"item,omitempty"`
}

type DealsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DealsStruct"`

	DealId int64 `xml:"dealId,omitempty"`

	DealDate int64 `xml:"dealDate,omitempty"`

	DealQuantity int32 `xml:"dealQuantity,omitempty"`

	DealAmountOriginal float32 `xml:"dealAmountOriginal,omitempty"`

	DealAmountDiscounted float32 `xml:"dealAmountDiscounted,omitempty"`
}

type ArrayOfDealsstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfDealsstruct"`

	Item []*DealsStruct `xml:"item,omitempty"`
}

type FilledPostBuyFormsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FilledPostBuyFormsStruct"`

	TransactionIds *ArrayOfLong `xml:"transactionIds,omitempty"`
}

type ItemGetImage struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemGetImage"`

	ItId int64 `xml:"itId,omitempty"`

	ItSellerId int64 `xml:"itSellerId,omitempty"`

	ItFotoCount int32 `xml:"itFotoCount,omitempty"`
}

type ArrayOfItemgetimage struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfItemgetimage"`

	Item []*ItemGetImage `xml:"item,omitempty"`
}

type ItemImageList struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemImageList"`

	ImageType int32 `xml:"imageType,omitempty"`

	ImageUrl string `xml:"imageUrl,omitempty"`
}

type ArrayOfItemimagelist struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfItemimagelist"`

	Item []*ItemImageList `xml:"item,omitempty"`
}

type ItemImagesStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemImagesStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemImages *ArrayOfItemimagelist `xml:"itemImages,omitempty"`
}

type ArrayOfItemimagesstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfItemimagesstruct"`

	Item []*ItemImagesStruct `xml:"item,omitempty"`
}

type DurationInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php DurationInfoStruct"`

	DurationType int32 `xml:"durationType,omitempty"`
}

type ItemInfo struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemInfo"`

	ItId int64 `xml:"itId,omitempty"`

	ItCountry int32 `xml:"itCountry,omitempty"`

	ItName string `xml:"itName,omitempty"`

	ItPrice float32 `xml:"itPrice,omitempty"`

	ItBidCount int32 `xml:"itBidCount,omitempty"`

	ItEndingTime int64 `xml:"itEndingTime,omitempty"`

	ItSellerId int64 `xml:"itSellerId,omitempty"`

	ItSellerLogin string `xml:"itSellerLogin,omitempty"`

	ItSellerRating int32 `xml:"itSellerRating,omitempty"`

	ItStartingTime int64 `xml:"itStartingTime,omitempty"`

	ItStartingPrice float32 `xml:"itStartingPrice,omitempty"`

	ItQuantity int32 `xml:"itQuantity,omitempty"`

	ItFotoCount int32 `xml:"itFotoCount,omitempty"`

	ItReservePrice float32 `xml:"itReservePrice,omitempty"`

	ItLocation string `xml:"itLocation,omitempty"`

	ItBuyNowPrice float32 `xml:"itBuyNowPrice,omitempty"`

	ItBuyNowActive int32 `xml:"itBuyNowActive,omitempty"`

	ItAdvertisementPrice float32 `xml:"itAdvertisementPrice,omitempty"`

	ItAdvertisementActive int32 `xml:"itAdvertisementActive,omitempty"`

	ItHighBidder int32 `xml:"itHighBidder,omitempty"`

	ItHighBidderLogin string `xml:"itHighBidderLogin,omitempty"`

	ItDescription string `xml:"itDescription,omitempty"`

	ItStandardizedDescription string `xml:"itStandardizedDescription,omitempty"`

	ItOptions int32 `xml:"itOptions,omitempty"`

	ItState int32 `xml:"itState,omitempty"`

	ItIsEco int32 `xml:"itIsEco,omitempty"`

	ItHitCount int64 `xml:"itHitCount,omitempty"`

	ItPostcode string `xml:"itPostcode,omitempty"`

	ItVatInvoice int32 `xml:"itVatInvoice,omitempty"`

	ItVatMarginInvoice int32 `xml:"itVatMarginInvoice,omitempty"`

	ItWithoutVatInvoice int32 `xml:"itWithoutVatInvoice,omitempty"`

	ItBankAccount1 string `xml:"itBankAccount1,omitempty"`

	ItBankAccount2 string `xml:"itBankAccount2,omitempty"`

	ItStartingQuantity int32 `xml:"itStartingQuantity,omitempty"`

	ItIsForGuests int32 `xml:"itIsForGuests,omitempty"`

	ItHasProduct int32 `xml:"itHasProduct,omitempty"`

	ItOrderFulfillmentTime int32 `xml:"itOrderFulfillmentTime,omitempty"`

	ItEndingInfo int32 `xml:"itEndingInfo,omitempty"`

	ItIsAllegroStandard int32 `xml:"itIsAllegroStandard,omitempty"`

	ItIsNewUsed int32 `xml:"itIsNewUsed,omitempty"`

	ItIsBrandZone int32 `xml:"itIsBrandZone,omitempty"`

	ItDurationInfo *DurationInfoStruct `xml:"itDurationInfo,omitempty"`

	ItIsFulfillmentTimeActive int32 `xml:"itIsFulfillmentTimeActive,omitempty"`

	ItEan string `xml:"itEan,omitempty"`

	ItContact string `xml:"itContact,omitempty"`
}

type ItemCatList struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemCatList"`

	CatLevel int32 `xml:"catLevel,omitempty"`

	CatId int64 `xml:"catId,omitempty"`

	CatName string `xml:"catName,omitempty"`
}

type ArrayOfItemcatlist struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfItemcatlist"`

	Item []*ItemCatList `xml:"item,omitempty"`
}

type AttribStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php AttribStruct"`

	AttribName string `xml:"attribName,omitempty"`

	AttribValues *ArrayOfString `xml:"attribValues,omitempty"`
}

type ArrayOfAttribstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfAttribstruct"`

	Item []*AttribStruct `xml:"item,omitempty"`
}

type FulfillmentTimeStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FulfillmentTimeStruct"`

	FulfillmentTimeFrom int32 `xml:"fulfillmentTimeFrom,omitempty"`

	FulfillmentTimeTo int32 `xml:"fulfillmentTimeTo,omitempty"`
}

type PostageStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostageStruct"`

	PostageAmount float32 `xml:"postageAmount,omitempty"`

	PostageAmountAdd float32 `xml:"postageAmountAdd,omitempty"`

	PostagePackSize int32 `xml:"postagePackSize,omitempty"`

	PostageId int32 `xml:"postageId,omitempty"`

	PostageFreeShipping int32 `xml:"postageFreeShipping,omitempty"`

	PostageFreeReturn int32 `xml:"postageFreeReturn,omitempty"`

	PostageFulfillmentTime *FulfillmentTimeStruct `xml:"postageFulfillmentTime,omitempty"`
}

type ArrayOfPostagestruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPostagestruct"`

	Item []*PostageStruct `xml:"item,omitempty"`
}

type ItemPaymentOptions struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemPaymentOptions"`

	PayOptionTransfer int32 `xml:"payOptionTransfer,omitempty"`

	PayOptionOnDelivery int32 `xml:"payOptionOnDelivery,omitempty"`

	PayOptionAllegroPay int32 `xml:"payOptionAllegroPay,omitempty"`

	PayOptionSeeDesc int32 `xml:"payOptionSeeDesc,omitempty"`

	PayOptionPayu int32 `xml:"payOptionPayu,omitempty"`

	PayOptionEscrow int32 `xml:"payOptionEscrow,omitempty"`

	PayOptionQiwi int32 `xml:"payOptionQiwi,omitempty"`
}

type CompanyInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php CompanyInfoStruct"`

	CompanyFirstName string `xml:"companyFirstName,omitempty"`

	CompanyLastName string `xml:"companyLastName,omitempty"`

	CompanyName string `xml:"companyName,omitempty"`

	CompanyAddress string `xml:"companyAddress,omitempty"`

	CompanyPostcode string `xml:"companyPostcode,omitempty"`

	CompanyCity string `xml:"companyCity,omitempty"`

	CompanyNip string `xml:"companyNip,omitempty"`
}

type ProductParametersStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ProductParametersStruct"`

	ProductParameterName string `xml:"productParameterName,omitempty"`

	ProductParameterValues *ArrayOfString `xml:"productParameterValues,omitempty"`

	ProductParameterDescription string `xml:"productParameterDescription,omitempty"`
}

type ArrayOfProductparametersstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfProductparametersstruct"`

	Item []*ProductParametersStruct `xml:"item,omitempty"`
}

type ProductParametersGroupStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ProductParametersGroupStruct"`

	ProductParametersGroupName string `xml:"productParametersGroupName,omitempty"`

	ProductParametersList *ArrayOfProductparametersstruct `xml:"productParametersList,omitempty"`
}

type ArrayOfProductparametersgroupstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfProductparametersgroupstruct"`

	Item []*ProductParametersGroupStruct `xml:"item,omitempty"`
}

type ProductStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ProductStruct"`

	ProductId int64 `xml:"productId,omitempty"`

	ProductName string `xml:"productName,omitempty"`

	ProductDescription string `xml:"productDescription,omitempty"`

	ProductImagesList *ArrayOfString `xml:"productImagesList,omitempty"`

	ProductParametersGroupList *ArrayOfProductparametersgroupstruct `xml:"productParametersGroupList,omitempty"`
}

type ItemInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemInfoStruct"`

	ItemInfo *ItemInfo `xml:"itemInfo,omitempty"`

	ItemCats *ArrayOfItemcatlist `xml:"itemCats,omitempty"`

	ItemImages *ArrayOfItemimagelist `xml:"itemImages,omitempty"`

	ItemAttribs *ArrayOfAttribstruct `xml:"itemAttribs,omitempty"`

	ItemPostageOptions *ArrayOfPostagestruct `xml:"itemPostageOptions,omitempty"`

	ItemPaymentOptions *ItemPaymentOptions `xml:"itemPaymentOptions,omitempty"`

	ItemCompanyInfo *CompanyInfoStruct `xml:"itemCompanyInfo,omitempty"`

	ItemProductInfo *ProductStruct `xml:"itemProductInfo,omitempty"`

	ItemAfterSalesServiceConditions *AfterSalesServiceConditionsStruct `xml:"itemAfterSalesServiceConditions,omitempty"`

	ItemAdditionalServicesGroup string `xml:"itemAdditionalServicesGroup,omitempty"`
}

type ArrayOfIteminfostruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfIteminfostruct"`

	Item []*ItemInfoStruct `xml:"item,omitempty"`
}

type RangeValueType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RangeValueType"`

	RangeValueMin string `xml:"rangeValueMin,omitempty"`

	RangeValueMax string `xml:"rangeValueMax,omitempty"`
}

type FilterOptionsType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FilterOptionsType"`

	FilterId string `xml:"filterId,omitempty"`

	FilterValueId *ArrayOfString `xml:"filterValueId,omitempty"`

	FilterValueRange *RangeValueType `xml:"filterValueRange,omitempty"`
}

type ArrayOfFilteroptionstype struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfFilteroptionstype"`

	Item []*FilterOptionsType `xml:"item,omitempty"`
}

type SortOptionsType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SortOptionsType"`

	SortType string `xml:"sortType,omitempty"`

	SortOrder string `xml:"sortOrder,omitempty"`
}

type UserInfoType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserInfoType"`

	UserId int32 `xml:"userId,omitempty"`

	UserLogin string `xml:"userLogin,omitempty"`

	UserRating int32 `xml:"userRating,omitempty"`

	UserIcons int32 `xml:"userIcons,omitempty"`

	CountryId int32 `xml:"countryId,omitempty"`
}

type PriceInfoType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PriceInfoType"`

	PriceType string `xml:"priceType,omitempty"`

	PriceValue float32 `xml:"priceValue,omitempty"`
}

type ArrayOfPriceinfotype struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPriceinfotype"`

	Item []*PriceInfoType `xml:"item,omitempty"`
}

type PhotoInfoType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PhotoInfoType"`

	PhotoSize string `xml:"photoSize,omitempty"`

	PhotoUrl string `xml:"photoUrl,omitempty"`

	PhotoIsMain bool `xml:"photoIsMain,omitempty"`
}

type ArrayOfPhotoinfotype struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPhotoinfotype"`

	Item []*PhotoInfoType `xml:"item,omitempty"`
}

type ParameterInfoType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ParameterInfoType"`

	ParameterName string `xml:"parameterName,omitempty"`

	ParameterValue *ArrayOfString `xml:"parameterValue,omitempty"`

	ParameterUnit string `xml:"parameterUnit,omitempty"`

	ParameterIsRange bool `xml:"parameterIsRange,omitempty"`
}

type ArrayOfParameterinfotype struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfParameterinfotype"`

	Item []*ParameterInfoType `xml:"item,omitempty"`
}

type AdvertInfoType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php AdvertInfoType"`

	ServiceId string `xml:"serviceId,omitempty"`

	AdvertId string `xml:"advertId,omitempty"`
}

type ItemsListType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemsListType"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	LeftCount int32 `xml:"leftCount,omitempty"`

	BidsCount int32 `xml:"bidsCount,omitempty"`

	BiddersCount int32 `xml:"biddersCount,omitempty"`

	QuantityType string `xml:"quantityType,omitempty"`

	EndingTime time.Time `xml:"endingTime,omitempty"`

	TimeToEnd string `xml:"timeToEnd,omitempty"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	ConditionInfo string `xml:"conditionInfo,omitempty"`

	PromotionInfo int32 `xml:"promotionInfo,omitempty"`

	AdditionalInfo int32 `xml:"additionalInfo,omitempty"`

	SellerInfo *UserInfoType `xml:"sellerInfo,omitempty"`

	PriceInfo *ArrayOfPriceinfotype `xml:"priceInfo,omitempty"`

	PhotosInfo *ArrayOfPhotoinfotype `xml:"photosInfo,omitempty"`

	ParametersInfo *ArrayOfParameterinfotype `xml:"parametersInfo,omitempty"`

	AdvertInfo *AdvertInfoType `xml:"advertInfo,omitempty"`
}

type ArrayOfItemslisttype struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfItemslisttype"`

	Item []*ItemsListType `xml:"item,omitempty"`
}

type CategoryTreeType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php CategoryTreeType"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	CategoryName string `xml:"categoryName,omitempty"`

	CategoryParentId int32 `xml:"categoryParentId,omitempty"`

	CategoryItemsCount int32 `xml:"categoryItemsCount,omitempty"`
}

type ArrayOfCategorytreetype struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfCategorytreetype"`

	Item []*CategoryTreeType `xml:"item,omitempty"`
}

type CategoryPathType struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php CategoryPathType"`

	CategoryId int32 `xml:"categoryId,omitempty"`

	CategoryName string `xml:"categoryName,omitempty"`

	CategoryParentId int32 `xml:"categoryParentId,omitempty"`
}

type ArrayOfCategorypathtype struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfCategorypathtype"`

	Item []*CategoryPathType `xml:"item,omitempty"`
}

type CategoriesListType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php CategoriesListType"`

	CategoriesTree *ArrayOfCategorytreetype `xml:"categoriesTree,omitempty"`

	CategoriesPath *ArrayOfCategorypathtype `xml:"categoriesPath,omitempty"`
}

type FilterValueType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FilterValueType"`

	FilterValueId string `xml:"filterValueId,omitempty"`

	FilterValueName string `xml:"filterValueName,omitempty"`

	FilterValueCount int32 `xml:"filterValueCount,omitempty"`
}

type ArrayOfFiltervaluetype struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfFiltervaluetype"`

	Item []*FilterValueType `xml:"item,omitempty"`
}

type FilterRelationType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FilterRelationType"`

	RelationAnd *ArrayOfString `xml:"relationAnd,omitempty"`

	RelationOr *ArrayOfString `xml:"relationOr,omitempty"`

	RelationExclude *ArrayOfString `xml:"relationExclude,omitempty"`
}

type FiltersListType struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FiltersListType"`

	FilterId string `xml:"filterId,omitempty"`

	FilterName string `xml:"filterName,omitempty"`

	FilterType string `xml:"filterType,omitempty"`

	FilterControlType string `xml:"filterControlType,omitempty"`

	FilterDataType string `xml:"filterDataType,omitempty"`

	FilterIsRange bool `xml:"filterIsRange,omitempty"`

	FilterArraySize int32 `xml:"filterArraySize,omitempty"`

	FilterValues *ArrayOfFiltervaluetype `xml:"filterValues,omitempty"`

	FilterRelations *FilterRelationType `xml:"filterRelations,omitempty"`
}

type ArrayOfFilterslisttype struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfFilterslisttype"`

	Item []*FiltersListType `xml:"item,omitempty"`
}

type AddressUserDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php AddressUserDataStruct"`

	UserCompany string `xml:"userCompany,omitempty"`

	UserFullName string `xml:"userFullName,omitempty"`

	UserAddress string `xml:"userAddress,omitempty"`

	UserPostcode string `xml:"userPostcode,omitempty"`

	UserCity string `xml:"userCity,omitempty"`
}

type AddressInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php AddressInfoStruct"`

	AddressType int32 `xml:"addressType,omitempty"`

	AddressUserData *AddressUserDataStruct `xml:"addressUserData,omitempty"`
}

type ArrayOfAddressinfostruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfAddressinfostruct"`

	Item []*AddressInfoStruct `xml:"item,omitempty"`
}

type SortOptionsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SortOptionsStruct"`

	SortType int32 `xml:"sortType,omitempty"`

	SortOrder int32 `xml:"sortOrder,omitempty"`
}

type ItemPriceStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemPriceStruct"`

	PriceType int32 `xml:"priceType,omitempty"`

	PriceValue float32 `xml:"priceValue,omitempty"`
}

type ArrayOfItempricestruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfItempricestruct"`

	Item []*ItemPriceStruct `xml:"item,omitempty"`
}

type UserInfoStruct struct {
	// XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserInfoStruct"`

	UserId int32 `xml:"userId,omitempty"`

	UserLogin string `xml:"userLogin,omitempty"`

	UserRating int32 `xml:"userRating,omitempty"`

	UserIcons int32 `xml:"userIcons,omitempty"`

	UserCountry int32 `xml:"userCountry,omitempty"`
}

type BidItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php BidItemStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	ItemThumbnailUrl string `xml:"itemThumbnailUrl,omitempty"`

	ItemPrice *ArrayOfItempricestruct `xml:"itemPrice,omitempty"`

	ItemBidQuantity int32 `xml:"itemBidQuantity,omitempty"`

	ItemLeftQuantity int32 `xml:"itemLeftQuantity,omitempty"`

	ItemQuantityType int32 `xml:"itemQuantityType,omitempty"`

	ItemEndTime int64 `xml:"itemEndTime,omitempty"`

	ItemEndTimeLeft string `xml:"itemEndTimeLeft,omitempty"`

	ItemSeller *UserInfoStruct `xml:"itemSeller,omitempty"`

	ItemBiddersCounter int32 `xml:"itemBiddersCounter,omitempty"`

	ItemHighestBidder *UserInfoStruct `xml:"itemHighestBidder,omitempty"`

	ItemCategoryId int32 `xml:"itemCategoryId,omitempty"`

	ItemViewsCounter int32 `xml:"itemViewsCounter,omitempty"`

	ItemNote string `xml:"itemNote,omitempty"`

	ItemSpecialInfo int32 `xml:"itemSpecialInfo,omitempty"`

	ItemShopInfo int32 `xml:"itemShopInfo,omitempty"`

	ItemProductInfo int64 `xml:"itemProductInfo,omitempty"`

	ItemPayuInfo int32 `xml:"itemPayuInfo,omitempty"`
}

type ArrayOfBiditemstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfBiditemstruct"`

	Item []*BidItemStruct `xml:"item,omitempty"`
}

type UserDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserDataStruct"`

	UserId int64 `xml:"userId,omitempty"`

	UserLogin string `xml:"userLogin,omitempty"`

	UserRating int32 `xml:"userRating,omitempty"`

	UserFirstName string `xml:"userFirstName,omitempty"`

	UserLastName string `xml:"userLastName,omitempty"`

	UserMaidenName string `xml:"userMaidenName,omitempty"`

	UserCompany string `xml:"userCompany,omitempty"`

	UserCountryId int32 `xml:"userCountryId,omitempty"`

	UserStateId int32 `xml:"userStateId,omitempty"`

	UserPostcode string `xml:"userPostcode,omitempty"`

	UserCity string `xml:"userCity,omitempty"`

	UserAddress string `xml:"userAddress,omitempty"`

	UserEmail string `xml:"userEmail,omitempty"`

	UserPhone string `xml:"userPhone,omitempty"`

	UserPhone2 string `xml:"userPhone2,omitempty"`

	UserSsStatus int32 `xml:"userSsStatus,omitempty"`

	SiteCountryId int32 `xml:"siteCountryId,omitempty"`

	UserJuniorStatus int32 `xml:"userJuniorStatus,omitempty"`

	UserBirthDate int64 `xml:"userBirthDate,omitempty"`

	UserHasShop int32 `xml:"userHasShop,omitempty"`

	UserCompanyIcon int32 `xml:"userCompanyIcon,omitempty"`

	UserIsAllegroStandard int32 `xml:"userIsAllegroStandard,omitempty"`
}

type InvoiceDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php InvoiceDataStruct"`

	CompanyName string `xml:"companyName,omitempty"`

	CompanyNip string `xml:"companyNip,omitempty"`
}

type CompanyExtraDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php CompanyExtraDataStruct"`

	CompanyType string `xml:"companyType,omitempty"`

	CompanyNip string `xml:"companyNip,omitempty"`

	CompanyRegon string `xml:"companyRegon,omitempty"`

	CompanyKrs string `xml:"companyKrs,omitempty"`

	CompanyActivitySort string `xml:"companyActivitySort,omitempty"`
}

type CompanySecondAddressStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php CompanySecondAddressStruct"`

	CompanyWorkerFirstName string `xml:"companyWorkerFirstName,omitempty"`

	CompanyWorkerLastName string `xml:"companyWorkerLastName,omitempty"`

	CompanyAddress string `xml:"companyAddress,omitempty"`

	CompanyPostcode string `xml:"companyPostcode,omitempty"`

	CompanyCity string `xml:"companyCity,omitempty"`

	CompanyCountryId int32 `xml:"companyCountryId,omitempty"`

	CompanyStateId int32 `xml:"companyStateId,omitempty"`
}

type PharmacyDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PharmacyDataStruct"`

	PharmacyOpeningDate string `xml:"pharmacyOpeningDate,omitempty"`

	PharmacyExpirationDate string `xml:"pharmacyExpirationDate,omitempty"`

	PharmacyName string `xml:"pharmacyName,omitempty"`

	PharmacyPharmacistFullName string `xml:"pharmacyPharmacistFullName,omitempty"`

	PharmacyAddress string `xml:"pharmacyAddress,omitempty"`

	PharmacyPostcode string `xml:"pharmacyPostcode,omitempty"`

	PharmacyCity string `xml:"pharmacyCity,omitempty"`

	PharmacyCommune string `xml:"pharmacyCommune,omitempty"`

	PharmacyCountryId int32 `xml:"pharmacyCountryId,omitempty"`

	PharmacyStateId int32 `xml:"pharmacyStateId,omitempty"`

	PharmacyPermitNumber string `xml:"pharmacyPermitNumber,omitempty"`

	PharmacyPermitInfo string `xml:"pharmacyPermitInfo,omitempty"`
}

type AlcoholDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php AlcoholDataStruct"`

	AlcoholOpeningDate string `xml:"alcoholOpeningDate,omitempty"`

	AlcoholExpirationDate string `xml:"alcoholExpirationDate,omitempty"`

	AlcoholShopName string `xml:"alcoholShopName,omitempty"`

	AlcoholShopAddress string `xml:"alcoholShopAddress,omitempty"`

	AlcoholShopPostcode string `xml:"alcoholShopPostcode,omitempty"`

	AlcoholShopCity string `xml:"alcoholShopCity,omitempty"`

	AlcoholShopCommune string `xml:"alcoholShopCommune,omitempty"`

	AlcoholShopCountryId int32 `xml:"alcoholShopCountryId,omitempty"`

	AlcoholShopStateId int32 `xml:"alcoholShopStateId,omitempty"`

	AlcoholPermitNumber string `xml:"alcoholPermitNumber,omitempty"`

	AlcoholPermitAuthority string `xml:"alcoholPermitAuthority,omitempty"`

	AlcoholPermitInfo string `xml:"alcoholPermitInfo,omitempty"`
}

type RelatedPersonsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RelatedPersonsStruct"`

	SpouseFirstName string `xml:"spouseFirstName,omitempty"`

	SpouseLastName string `xml:"spouseLastName,omitempty"`
}

type FutureFilterOptionsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FutureFilterOptionsStruct"`

	FilterFormat int32 `xml:"filterFormat,omitempty"`
}

type FutureItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FutureItemStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	ItemThumbnailUrl string `xml:"itemThumbnailUrl,omitempty"`

	ItemPrice *ArrayOfItempricestruct `xml:"itemPrice,omitempty"`

	ItemStartQuantity int32 `xml:"itemStartQuantity,omitempty"`

	ItemQuantityType int32 `xml:"itemQuantityType,omitempty"`

	ItemStartTime int64 `xml:"itemStartTime,omitempty"`

	ItemCategoryId int32 `xml:"itemCategoryId,omitempty"`

	ItemNote string `xml:"itemNote,omitempty"`

	ItemSpecialInfo int32 `xml:"itemSpecialInfo,omitempty"`

	ItemShopInfo int32 `xml:"itemShopInfo,omitempty"`

	ItemProductInfo int64 `xml:"itemProductInfo,omitempty"`

	ItemPayuInfo int32 `xml:"itemPayuInfo,omitempty"`

	ItemStatus int32 `xml:"itemStatus,omitempty"`
}

type ArrayOfFutureitemstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfFutureitemstruct"`

	Item []*FutureItemStruct `xml:"item,omitempty"`
}

type PaymentDetailsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PaymentDetailsStruct"`

	PayTransDetailsItId int64 `xml:"payTransDetailsItId,omitempty"`

	PayTransDetailsPrice float32 `xml:"payTransDetailsPrice,omitempty"`

	PayTransDetailsCount int32 `xml:"payTransDetailsCount,omitempty"`
}

type ArrayOfPaymentdetailsstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPaymentdetailsstruct"`

	Item []*PaymentDetailsStruct `xml:"item,omitempty"`
}

type UserIncomingPaymentStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserIncomingPaymentStruct"`

	PayTransId int64 `xml:"payTransId,omitempty"`

	PayTransItId int64 `xml:"payTransItId,omitempty"`

	PayTransBuyerId int32 `xml:"payTransBuyerId,omitempty"`

	PayTransType string `xml:"payTransType,omitempty"`

	PayTransStatus string `xml:"payTransStatus,omitempty"`

	PayTransAmount float32 `xml:"payTransAmount,omitempty"`

	PayTransCreateDate int64 `xml:"payTransCreateDate,omitempty"`

	PayTransRecvDate int64 `xml:"payTransRecvDate,omitempty"`

	PayTransPrice float32 `xml:"payTransPrice,omitempty"`

	PayTransCount int32 `xml:"payTransCount,omitempty"`

	PayTransPostageAmount float32 `xml:"payTransPostageAmount,omitempty"`

	PayTransDetails *ArrayOfPaymentdetailsstruct `xml:"payTransDetails,omitempty"`

	PayTransIncomplete int32 `xml:"payTransIncomplete,omitempty"`

	PayTransMainId int64 `xml:"payTransMainId,omitempty"`
}

type ArrayOfUserincomingpaymentstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfUserincomingpaymentstruct"`

	Item []*UserIncomingPaymentStruct `xml:"item,omitempty"`
}

type UserIncomingPaymentRefundsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserIncomingPaymentRefundsStruct"`

	PayRefundTransId int64 `xml:"payRefundTransId,omitempty"`

	PayRefundItId int64 `xml:"payRefundItId,omitempty"`

	PayRefundBuyerId int32 `xml:"payRefundBuyerId,omitempty"`

	PayRefundValue float32 `xml:"payRefundValue,omitempty"`

	PayRefundReason string `xml:"payRefundReason,omitempty"`

	PayRefundDate int64 `xml:"payRefundDate,omitempty"`
}

type ArrayOfUserincomingpaymentrefundsstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfUserincomingpaymentrefundsstruct"`

	Item []*UserIncomingPaymentRefundsStruct `xml:"item,omitempty"`
}

type FilterPriceStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php FilterPriceStruct"`

	FilterPriceFrom float32 `xml:"filterPriceFrom,omitempty"`

	FilterPriceTo float32 `xml:"filterPriceTo,omitempty"`
}

type NotSoldFilterOptionsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php NotSoldFilterOptionsStruct"`

	FilterFormat int32 `xml:"filterFormat,omitempty"`

	FilterFromEnd int32 `xml:"filterFromEnd,omitempty"`

	FilterAutoListing int32 `xml:"filterAutoListing,omitempty"`

	FilterPrice *FilterPriceStruct `xml:"filterPrice,omitempty"`

	FilterDurationType int32 `xml:"filterDurationType,omitempty"`
}

type NotSoldItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php NotSoldItemStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	ItemThumbnailUrl string `xml:"itemThumbnailUrl,omitempty"`

	ItemPrice *ArrayOfItempricestruct `xml:"itemPrice,omitempty"`

	ItemStartQuantity int32 `xml:"itemStartQuantity,omitempty"`

	ItemQuantityType int32 `xml:"itemQuantityType,omitempty"`

	ItemStartTime int64 `xml:"itemStartTime,omitempty"`

	ItemEndTime int64 `xml:"itemEndTime,omitempty"`

	ItemBiddersCounter int32 `xml:"itemBiddersCounter,omitempty"`

	ItemHighestBidder *UserInfoStruct `xml:"itemHighestBidder,omitempty"`

	ItemCategoryId int32 `xml:"itemCategoryId,omitempty"`

	ItemWatchersCounter int32 `xml:"itemWatchersCounter,omitempty"`

	ItemViewsCounter int32 `xml:"itemViewsCounter,omitempty"`

	ItemNote string `xml:"itemNote,omitempty"`

	ItemSpecialInfo int32 `xml:"itemSpecialInfo,omitempty"`

	ItemShopInfo int32 `xml:"itemShopInfo,omitempty"`

	ItemProductInfo int64 `xml:"itemProductInfo,omitempty"`

	ItemPayuInfo int32 `xml:"itemPayuInfo,omitempty"`
}

type ArrayOfNotsolditemstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfNotsolditemstruct"`

	Item []*NotSoldItemStruct `xml:"item,omitempty"`
}

type NotWonItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php NotWonItemStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	ItemThumbnailUrl string `xml:"itemThumbnailUrl,omitempty"`

	ItemPrice *ArrayOfItempricestruct `xml:"itemPrice,omitempty"`

	ItemLeftQuantity int32 `xml:"itemLeftQuantity,omitempty"`

	ItemQuantityType int32 `xml:"itemQuantityType,omitempty"`

	ItemEndTime int64 `xml:"itemEndTime,omitempty"`

	ItemEndTimeLeft string `xml:"itemEndTimeLeft,omitempty"`

	ItemSeller *UserInfoStruct `xml:"itemSeller,omitempty"`

	ItemBiddersCounter int32 `xml:"itemBiddersCounter,omitempty"`

	ItemHighestBidder *UserInfoStruct `xml:"itemHighestBidder,omitempty"`

	ItemCategoryId int32 `xml:"itemCategoryId,omitempty"`

	ItemViewsCounter int32 `xml:"itemViewsCounter,omitempty"`

	ItemNote string `xml:"itemNote,omitempty"`

	ItemSpecialInfo int32 `xml:"itemSpecialInfo,omitempty"`

	ItemShopInfo int32 `xml:"itemShopInfo,omitempty"`

	ItemProductInfo int64 `xml:"itemProductInfo,omitempty"`

	ItemPayuInfo int32 `xml:"itemPayuInfo,omitempty"`
}

type ArrayOfNotwonitemstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfNotwonitemstruct"`

	Item []*NotWonItemStruct `xml:"item,omitempty"`
}

type PaymentItemsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PaymentItemsStruct"`

	PayTransItId int64 `xml:"payTransItId,omitempty"`

	PayTransItName string `xml:"payTransItName,omitempty"`

	PayTransItCount int32 `xml:"payTransItCount,omitempty"`

	PayTransItPrice float32 `xml:"payTransItPrice,omitempty"`
}

type ArrayOfPaymentitemsstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPaymentitemsstruct"`

	Item []*PaymentItemsStruct `xml:"item,omitempty"`
}

type PaymentSellersStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PaymentSellersStruct"`

	PayTransSellerId int32 `xml:"payTransSellerId,omitempty"`

	PayTransSellerName string `xml:"payTransSellerName,omitempty"`

	PayTransItems *ArrayOfPaymentitemsstruct `xml:"payTransItems,omitempty"`

	PayTransSellerPostageAmount float32 `xml:"payTransSellerPostageAmount,omitempty"`
}

type ArrayOfPaymentsellersstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPaymentsellersstruct"`

	Item []*PaymentSellersStruct `xml:"item,omitempty"`
}

type UserPaymentStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserPaymentStruct"`

	PayTransId int64 `xml:"payTransId,omitempty"`

	PayTransSellers *ArrayOfPaymentsellersstruct `xml:"payTransSellers,omitempty"`

	PayTransType string `xml:"payTransType,omitempty"`

	PayTransStatus string `xml:"payTransStatus,omitempty"`

	PayTransAmount float32 `xml:"payTransAmount,omitempty"`

	PayTransCreateDate int64 `xml:"payTransCreateDate,omitempty"`

	PayTransPrice float32 `xml:"payTransPrice,omitempty"`

	PayTransPostageAmount float32 `xml:"payTransPostageAmount,omitempty"`

	PayTransIncomplete int32 `xml:"payTransIncomplete,omitempty"`
}

type ArrayOfUserpaymentstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfUserpaymentstruct"`

	Item []*UserPaymentStruct `xml:"item,omitempty"`
}

type PaymentsUserDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PaymentsUserDataStruct"`

	UserFullName string `xml:"userFullName,omitempty"`

	UserAddress string `xml:"userAddress,omitempty"`

	UserPostcode string `xml:"userPostcode,omitempty"`

	UserCity string `xml:"userCity,omitempty"`

	UserCountry int32 `xml:"userCountry,omitempty"`

	UserPhone string `xml:"userPhone,omitempty"`
}

type PayoutAutoFrequencyStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PayoutAutoFrequencyStruct"`

	FrequencyType int16 `xml:"frequencyType,omitempty"`

	FrequencyWeekValue int16 `xml:"frequencyWeekValue,omitempty"`

	FrequencyMonthValue int16 `xml:"frequencyMonthValue,omitempty"`
}

type PayoutAutoSettingsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PayoutAutoSettingsStruct"`

	PayoutAutoAmount float32 `xml:"payoutAutoAmount,omitempty"`

	PayoutAutoFrequency *PayoutAutoFrequencyStruct `xml:"payoutAutoFrequency,omitempty"`
}

type PaymentsPayoutStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PaymentsPayoutStruct"`

	PayoutType int32 `xml:"payoutType,omitempty"`

	PayoutAutoSettings *PayoutAutoSettingsStruct `xml:"payoutAutoSettings,omitempty"`
}

type PaymentsInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PaymentsInfoStruct"`

	PaymentsBalance float32 `xml:"paymentsBalance,omitempty"`

	PaymentsBankAccount string `xml:"paymentsBankAccount,omitempty"`

	PaymentsUserData *PaymentsUserDataStruct `xml:"paymentsUserData,omitempty"`

	PaymentsPayout *PaymentsPayoutStruct `xml:"paymentsPayout,omitempty"`

	PaymentsNotifications int32 `xml:"paymentsNotifications,omitempty"`
}

type UserPaymentRefundsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserPaymentRefundsStruct"`

	PayRefundTransId int64 `xml:"payRefundTransId,omitempty"`

	PayRefundItId int64 `xml:"payRefundItId,omitempty"`

	PayRefundSellerId int32 `xml:"payRefundSellerId,omitempty"`

	PayRefundValue float32 `xml:"payRefundValue,omitempty"`

	PayRefundReason string `xml:"payRefundReason,omitempty"`

	PayRefundDate int64 `xml:"payRefundDate,omitempty"`
}

type ArrayOfUserpaymentrefundsstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfUserpaymentrefundsstruct"`

	Item []*UserPaymentRefundsStruct `xml:"item,omitempty"`
}

type UserPayoutStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserPayoutStruct"`

	PayTransId int64 `xml:"payTransId,omitempty"`

	PayTransStatus string `xml:"payTransStatus,omitempty"`

	PayTransAmount float32 `xml:"payTransAmount,omitempty"`

	PayTransCreateDate int64 `xml:"payTransCreateDate,omitempty"`

	PayTransRecvDate int64 `xml:"payTransRecvDate,omitempty"`

	PayTransCancelDate int64 `xml:"payTransCancelDate,omitempty"`

	PayTransReport string `xml:"payTransReport,omitempty"`
}

type ArrayOfUserpayoutstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfUserpayoutstruct"`

	Item []*UserPayoutStruct `xml:"item,omitempty"`
}

type PayoutPaymentsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PayoutPaymentsStruct"`

	TranasctionId int64 `xml:"tranasctionId,omitempty"`

	UserName string `xml:"userName,omitempty"`

	UserId int64 `xml:"userId,omitempty"`

	Amount float32 `xml:"amount,omitempty"`

	TransportAmount float32 `xml:"transportAmount,omitempty"`

	TotalAmount float32 `xml:"totalAmount,omitempty"`

	PaidDate string `xml:"paidDate,omitempty"`
}

type ArrayOfPayoutpaymentsstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPayoutpaymentsstruct"`

	Item []*PayoutPaymentsStruct `xml:"item,omitempty"`
}

type PayoutRefundFromStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PayoutRefundFromStruct"`

	TranasctionId int64 `xml:"tranasctionId,omitempty"`

	RefundId int64 `xml:"refundId,omitempty"`

	RefundReason string `xml:"refundReason,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	FromUserId int64 `xml:"fromUserId,omitempty"`

	Amount float32 `xml:"amount,omitempty"`

	PaidDate string `xml:"paidDate,omitempty"`
}

type ArrayOfPayoutrefundfromstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPayoutrefundfromstruct"`

	Item []*PayoutRefundFromStruct `xml:"item,omitempty"`
}

type PayoutRefundToStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PayoutRefundToStruct"`

	TranasctionId int64 `xml:"tranasctionId,omitempty"`

	RefundId int64 `xml:"refundId,omitempty"`

	RefundReason string `xml:"refundReason,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	ToUserId int64 `xml:"toUserId,omitempty"`

	Amount float32 `xml:"amount,omitempty"`

	PaidDate string `xml:"paidDate,omitempty"`
}

type ArrayOfPayoutrefundtostruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPayoutrefundtostruct"`

	Item []*PayoutRefundToStruct `xml:"item,omitempty"`
}

type SellFilterOptionsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SellFilterOptionsStruct"`

	FilterFormat int32 `xml:"filterFormat,omitempty"`

	FilterBids int32 `xml:"filterBids,omitempty"`

	FilterToEnd int32 `xml:"filterToEnd,omitempty"`

	FilterFromStart int32 `xml:"filterFromStart,omitempty"`

	FilterAutoListing int32 `xml:"filterAutoListing,omitempty"`

	FilterPrice *FilterPriceStruct `xml:"filterPrice,omitempty"`

	FilterDurationType int32 `xml:"filterDurationType,omitempty"`
}

type SellItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SellItemStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	ItemThumbnailUrl string `xml:"itemThumbnailUrl,omitempty"`

	ItemPrice *ArrayOfItempricestruct `xml:"itemPrice,omitempty"`

	ItemStartQuantity int32 `xml:"itemStartQuantity,omitempty"`

	ItemSoldQuantity int32 `xml:"itemSoldQuantity,omitempty"`

	ItemQuantityType int32 `xml:"itemQuantityType,omitempty"`

	ItemStartTime int64 `xml:"itemStartTime,omitempty"`

	ItemEndTime int64 `xml:"itemEndTime,omitempty"`

	ItemEndTimeLeft string `xml:"itemEndTimeLeft,omitempty"`

	ItemBiddersCounter int32 `xml:"itemBiddersCounter,omitempty"`

	ItemHighestBidder *UserInfoStruct `xml:"itemHighestBidder,omitempty"`

	ItemCategoryId int32 `xml:"itemCategoryId,omitempty"`

	ItemWatchersCounter int32 `xml:"itemWatchersCounter,omitempty"`

	ItemViewsCounter int32 `xml:"itemViewsCounter,omitempty"`

	ItemNote string `xml:"itemNote,omitempty"`

	ItemSpecialInfo int32 `xml:"itemSpecialInfo,omitempty"`

	ItemShopInfo int32 `xml:"itemShopInfo,omitempty"`

	ItemProductInfo int64 `xml:"itemProductInfo,omitempty"`

	ItemPayuInfo int32 `xml:"itemPayuInfo,omitempty"`

	ItemDurationInfo *DurationInfoStruct `xml:"itemDurationInfo,omitempty"`
}

type ArrayOfSellitemstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfSellitemstruct"`

	Item []*SellItemStruct `xml:"item,omitempty"`
}

type SoldFilterOptionsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SoldFilterOptionsStruct"`

	FilterFormat int32 `xml:"filterFormat,omitempty"`

	FilterFromEnd int32 `xml:"filterFromEnd,omitempty"`

	FilterAutoListing int32 `xml:"filterAutoListing,omitempty"`

	FilterPrice *FilterPriceStruct `xml:"filterPrice,omitempty"`

	FilterDurationType int32 `xml:"filterDurationType,omitempty"`
}

type SoldItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SoldItemStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	ItemThumbnailUrl string `xml:"itemThumbnailUrl,omitempty"`

	ItemPrice *ArrayOfItempricestruct `xml:"itemPrice,omitempty"`

	ItemStartQuantity int32 `xml:"itemStartQuantity,omitempty"`

	ItemSoldQuantity int32 `xml:"itemSoldQuantity,omitempty"`

	ItemQuantityType int32 `xml:"itemQuantityType,omitempty"`

	ItemStartTime int64 `xml:"itemStartTime,omitempty"`

	ItemEndTime int64 `xml:"itemEndTime,omitempty"`

	ItemEndTimeLeft string `xml:"itemEndTimeLeft,omitempty"`

	ItemBiddersCounter int32 `xml:"itemBiddersCounter,omitempty"`

	ItemHighestBidder *UserInfoStruct `xml:"itemHighestBidder,omitempty"`

	ItemCategoryId int32 `xml:"itemCategoryId,omitempty"`

	ItemWatchersCounter int32 `xml:"itemWatchersCounter,omitempty"`

	ItemViewsCounter int32 `xml:"itemViewsCounter,omitempty"`

	ItemNote string `xml:"itemNote,omitempty"`

	ItemSpecialInfo int32 `xml:"itemSpecialInfo,omitempty"`

	ItemShopInfo int32 `xml:"itemShopInfo,omitempty"`

	ItemProductInfo int64 `xml:"itemProductInfo,omitempty"`

	ItemPayuInfo int32 `xml:"itemPayuInfo,omitempty"`

	ItemDurationInfo *DurationInfoStruct `xml:"itemDurationInfo,omitempty"`
}

type ArrayOfSolditemstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfSolditemstruct"`

	Item []*SoldItemStruct `xml:"item,omitempty"`
}

type WonItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php WonItemStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	ItemThumbnailUrl string `xml:"itemThumbnailUrl,omitempty"`

	ItemPrice *ArrayOfItempricestruct `xml:"itemPrice,omitempty"`

	ItemBoughtQuantity int32 `xml:"itemBoughtQuantity,omitempty"`

	ItemLeftQuantity int32 `xml:"itemLeftQuantity,omitempty"`

	ItemQuantityType int32 `xml:"itemQuantityType,omitempty"`

	ItemEndTime int64 `xml:"itemEndTime,omitempty"`

	ItemEndTimeLeft string `xml:"itemEndTimeLeft,omitempty"`

	ItemSeller *UserInfoStruct `xml:"itemSeller,omitempty"`

	ItemBiddersCounter int32 `xml:"itemBiddersCounter,omitempty"`

	ItemCategoryId int32 `xml:"itemCategoryId,omitempty"`

	ItemViewsCounter int32 `xml:"itemViewsCounter,omitempty"`

	ItemNote string `xml:"itemNote,omitempty"`

	ItemSpecialInfo int32 `xml:"itemSpecialInfo,omitempty"`

	ItemShopInfo int32 `xml:"itemShopInfo,omitempty"`

	ItemProductInfo int64 `xml:"itemProductInfo,omitempty"`

	ItemPayuInfo int32 `xml:"itemPayuInfo,omitempty"`
}

type ArrayOfWonitemstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfWonitemstruct"`

	Item []*WonItemStruct `xml:"item,omitempty"`
}

type PaymentMethodStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PaymentMethodStruct"`

	PaymentMethodId string `xml:"paymentMethodId,omitempty"`

	PaymentMethodName string `xml:"paymentMethodName,omitempty"`

	PaymentMethodImage string `xml:"paymentMethodImage,omitempty"`

	PaymentMethodUsage int32 `xml:"paymentMethodUsage,omitempty"`
}

type ArrayOfPaymentmethodstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPaymentmethodstruct"`

	Item []*PaymentMethodStruct `xml:"item,omitempty"`
}

type UserSentToDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserSentToDataStruct"`

	UserId int64 `xml:"userId,omitempty"`

	UserFirstName string `xml:"userFirstName,omitempty"`

	UserLastName string `xml:"userLastName,omitempty"`

	UserCompany string `xml:"userCompany,omitempty"`

	UserCountryId int32 `xml:"userCountryId,omitempty"`

	UserPostcode string `xml:"userPostcode,omitempty"`

	UserCity string `xml:"userCity,omitempty"`

	UserAddress string `xml:"userAddress,omitempty"`
}

type UserPostBuyDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php UserPostBuyDataStruct"`

	UserData *UserDataStruct `xml:"userData,omitempty"`

	UserSentToData *UserSentToDataStruct `xml:"userSentToData,omitempty"`

	UserBankAccounts *ArrayOfString `xml:"userBankAccounts,omitempty"`

	CompanySecondAddress *CompanySecondAddressStruct `xml:"companySecondAddress,omitempty"`
}

type ArrayOfUserpostbuydatastruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfUserpostbuydatastruct"`

	Item []*UserPostBuyDataStruct `xml:"item,omitempty"`
}

type ItemPostBuyDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemPostBuyDataStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	UsersPostBuyData *ArrayOfUserpostbuydatastruct `xml:"usersPostBuyData,omitempty"`
}

type ArrayOfItempostbuydatastruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfItempostbuydatastruct"`

	Item []*ItemPostBuyDataStruct `xml:"item,omitempty"`
}

type PostBuyFormItemDealsVariantStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormItemDealsVariantStruct"`

	VariantName string `xml:"variantName,omitempty"`

	VariantValue string `xml:"variantValue,omitempty"`
}

type PostBuyFormItemDealsAdditionalServiceStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormItemDealsAdditionalServiceStruct"`

	AdditionalServiceDefinitionId string `xml:"additionalServiceDefinitionId,omitempty"`

	AdditionalServiceName string `xml:"additionalServiceName,omitempty"`

	AdditionalServiceQuantity int32 `xml:"additionalServiceQuantity,omitempty"`

	AdditionalServicePrice float32 `xml:"additionalServicePrice,omitempty"`
}

type ArrayOfPostbuyformitemdealsadditionalservicestruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPostbuyformitemdealsadditionalservicestruct"`

	Item []*PostBuyFormItemDealsAdditionalServiceStruct `xml:"item,omitempty"`
}

type PostBuyFormItemDealsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormItemDealsStruct"`

	DealId int64 `xml:"dealId,omitempty"`

	DealFinalPrice float32 `xml:"dealFinalPrice,omitempty"`

	DealQuantity int32 `xml:"dealQuantity,omitempty"`

	DealDate time.Time `xml:"dealDate,omitempty"`

	DealWasDiscounted bool `xml:"dealWasDiscounted,omitempty"`

	DealVariant *PostBuyFormItemDealsVariantStruct `xml:"dealVariant,omitempty"`

	DealAdditionalServices *ArrayOfPostbuyformitemdealsadditionalservicestruct `xml:"dealAdditionalServices,omitempty"`
}

type ArrayOfPostbuyformitemdealsstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPostbuyformitemdealsstruct"`

	Item []*PostBuyFormItemDealsStruct `xml:"item,omitempty"`
}

type PostBuyFormItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormItemStruct"`

	PostBuyFormItQuantity int32 `xml:"postBuyFormItQuantity,omitempty"`

	PostBuyFormItAmount float32 `xml:"postBuyFormItAmount,omitempty"`

	PostBuyFormItId int64 `xml:"postBuyFormItId,omitempty"`

	PostBuyFormItTitle string `xml:"postBuyFormItTitle,omitempty"`

	PostBuyFormItCountry int32 `xml:"postBuyFormItCountry,omitempty"`

	PostBuyFormItPrice float32 `xml:"postBuyFormItPrice,omitempty"`

	PostBuyFormItDeals *ArrayOfPostbuyformitemdealsstruct `xml:"postBuyFormItDeals,omitempty"`
}

type ArrayOfPostbuyformitemstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPostbuyformitemstruct"`

	Item []*PostBuyFormItemStruct `xml:"item,omitempty"`
}

type PostBuyFormShipmentTrackingStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormShipmentTrackingStruct"`

	PostBuyFormOperatorId int32 `xml:"postBuyFormOperatorId,omitempty"`

	PostBuyFormOperatorName string `xml:"postBuyFormOperatorName,omitempty"`

	PostBuyFormPackageId string `xml:"postBuyFormPackageId,omitempty"`

	PostBuyFormPackageStatus string `xml:"postBuyFormPackageStatus,omitempty"`
}

type ArrayOfPostbuyformshipmenttrackingstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPostbuyformshipmenttrackingstruct"`

	Item []*PostBuyFormShipmentTrackingStruct `xml:"item,omitempty"`
}

type PostBuyFormAddressStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormAddressStruct"`

	PostBuyFormAdrCountry int32 `xml:"postBuyFormAdrCountry,omitempty"`

	PostBuyFormAdrStreet string `xml:"postBuyFormAdrStreet,omitempty"`

	PostBuyFormAdrPostcode string `xml:"postBuyFormAdrPostcode,omitempty"`

	PostBuyFormAdrCity string `xml:"postBuyFormAdrCity,omitempty"`

	PostBuyFormAdrFullName string `xml:"postBuyFormAdrFullName,omitempty"`

	PostBuyFormAdrCompany string `xml:"postBuyFormAdrCompany,omitempty"`

	PostBuyFormAdrPhone string `xml:"postBuyFormAdrPhone,omitempty"`

	PostBuyFormAdrNip string `xml:"postBuyFormAdrNip,omitempty"`

	PostBuyFormCreatedDate string `xml:"postBuyFormCreatedDate,omitempty"`

	PostBuyFormAdrType int32 `xml:"postBuyFormAdrType,omitempty"`

	PostBuyFormAdrId string `xml:"postBuyFormAdrId,omitempty"`
}

type PostBuyFormSellersStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormSellersStruct"`

	PostBuyFormSellerId int32 `xml:"postBuyFormSellerId,omitempty"`

	PostBuyFormSellerName string `xml:"postBuyFormSellerName,omitempty"`

	PostBuyFormItems *ArrayOfPostbuyformitemstruct `xml:"postBuyFormItems,omitempty"`

	PostBuyFormShipmentId int32 `xml:"postBuyFormShipmentId,omitempty"`

	PostBuyFormPostageAmount float32 `xml:"postBuyFormPostageAmount,omitempty"`

	PostBuyFormMsgToSeller string `xml:"postBuyFormMsgToSeller,omitempty"`

	PostBuyFormAmount float32 `xml:"postBuyFormAmount,omitempty"`

	PostBuyFormSurchargesList *ArrayOfLong `xml:"postBuyFormSurchargesList,omitempty"`

	PostBuyFormShipmentTracking *ArrayOfPostbuyformshipmenttrackingstruct `xml:"postBuyFormShipmentTracking,omitempty"`

	PostBuyFormGdAddress *PostBuyFormAddressStruct `xml:"postBuyFormGdAddress,omitempty"`

	PostBuyFormGdAdditionalInfo string `xml:"postBuyFormGdAdditionalInfo,omitempty"`

	PostBuyFormSentBySeller int32 `xml:"postBuyFormSentBySeller,omitempty"`
}

type ArrayOfPostbuyformsellersstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPostbuyformsellersstruct"`

	Item []*PostBuyFormSellersStruct `xml:"item,omitempty"`
}

type PostBuyFormForBuyersDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormForBuyersDataStruct"`

	PostBuyFormId int64 `xml:"postBuyFormId,omitempty"`

	PostBuyFormBuyerId int32 `xml:"postBuyFormBuyerId,omitempty"`

	PostBuyFormSellers *ArrayOfPostbuyformsellersstruct `xml:"postBuyFormSellers,omitempty"`

	PostBuyFormTotalAmount float32 `xml:"postBuyFormTotalAmount,omitempty"`

	PostBuyFormInvoiceOption int32 `xml:"postBuyFormInvoiceOption,omitempty"`

	PostBuyFormInvoiceData *PostBuyFormAddressStruct `xml:"postBuyFormInvoiceData,omitempty"`

	PostBuyFormShipmentAddress *PostBuyFormAddressStruct `xml:"postBuyFormShipmentAddress,omitempty"`

	PostBuyFormPayType string `xml:"postBuyFormPayType,omitempty"`

	PostBuyFormPayId int64 `xml:"postBuyFormPayId,omitempty"`

	PostBuyFormPayStatus string `xml:"postBuyFormPayStatus,omitempty"`

	PostBuyFormDateInit string `xml:"postBuyFormDateInit,omitempty"`

	PostBuyFormDateRecv string `xml:"postBuyFormDateRecv,omitempty"`

	PostBuyFormDateCancel string `xml:"postBuyFormDateCancel,omitempty"`

	PostBuyFormPaymentAmount float32 `xml:"postBuyFormPaymentAmount,omitempty"`
}

type ArrayOfPostbuyformforbuyersdatastruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPostbuyformforbuyersdatastruct"`

	Item []*PostBuyFormForBuyersDataStruct `xml:"item,omitempty"`
}

type PostBuyFormDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormDataStruct"`

	PostBuyFormId int64 `xml:"postBuyFormId,omitempty"`

	PostBuyFormItems *ArrayOfPostbuyformitemstruct `xml:"postBuyFormItems,omitempty"`

	PostBuyFormBuyerId int64 `xml:"postBuyFormBuyerId,omitempty"`

	PostBuyFormAmount float32 `xml:"postBuyFormAmount,omitempty"`

	PostBuyFormPostageAmount float32 `xml:"postBuyFormPostageAmount,omitempty"`

	PostBuyFormInvoiceOption int32 `xml:"postBuyFormInvoiceOption,omitempty"`

	PostBuyFormMsgToSeller string `xml:"postBuyFormMsgToSeller,omitempty"`

	PostBuyFormInvoiceData *PostBuyFormAddressStruct `xml:"postBuyFormInvoiceData,omitempty"`

	PostBuyFormShipmentAddress *PostBuyFormAddressStruct `xml:"postBuyFormShipmentAddress,omitempty"`

	PostBuyFormPayType string `xml:"postBuyFormPayType,omitempty"`

	PostBuyFormPayId int64 `xml:"postBuyFormPayId,omitempty"`

	PostBuyFormPayStatus string `xml:"postBuyFormPayStatus,omitempty"`

	PostBuyFormDateInit string `xml:"postBuyFormDateInit,omitempty"`

	PostBuyFormDateRecv string `xml:"postBuyFormDateRecv,omitempty"`

	PostBuyFormDateCancel string `xml:"postBuyFormDateCancel,omitempty"`

	PostBuyFormShipmentId int32 `xml:"postBuyFormShipmentId,omitempty"`

	PostBuyFormGdAddress *PostBuyFormAddressStruct `xml:"postBuyFormGdAddress,omitempty"`

	PostBuyFormShipmentTracking *ArrayOfPostbuyformshipmenttrackingstruct `xml:"postBuyFormShipmentTracking,omitempty"`

	PostBuyFormSurchargesList *ArrayOfLong `xml:"postBuyFormSurchargesList,omitempty"`

	PostBuyFormGdAdditionalInfo string `xml:"postBuyFormGdAdditionalInfo,omitempty"`

	PostBuyFormPaymentAmount float32 `xml:"postBuyFormPaymentAmount,omitempty"`

	PostBuyFormSentBySeller int32 `xml:"postBuyFormSentBySeller,omitempty"`

	PostBuyFormBuyerLogin string `xml:"postBuyFormBuyerLogin,omitempty"`

	PostBuyFormBuyerEmail string `xml:"postBuyFormBuyerEmail,omitempty"`

	PostBuyFormAdditionalServicesAmount float32 `xml:"postBuyFormAdditionalServicesAmount,omitempty"`
}

type ArrayOfPostbuyformdatastruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPostbuyformdatastruct"`

	Item []*PostBuyFormDataStruct `xml:"item,omitempty"`
}

type PostBuyItemInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyItemInfoStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemPostBuyFormInfo int32 `xml:"itemPostBuyFormInfo,omitempty"`
}

type ArrayOfPostbuyiteminfostruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfPostbuyiteminfostruct"`

	Item []*PostBuyItemInfoStruct `xml:"item,omitempty"`
}

type RefundsDealsListType struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RefundsDealsListType"`

	DealId int64 `xml:"dealId,omitempty"`

	DealDate time.Time `xml:"dealDate,omitempty"`

	TimeLeft int32 `xml:"timeLeft,omitempty"`

	BuyerId int32 `xml:"buyerId,omitempty"`

	BuyerLogin string `xml:"buyerLogin,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	BidsCount int32 `xml:"bidsCount,omitempty"`

	QuantityType string `xml:"quantityType,omitempty"`

	PaymentStatus string `xml:"paymentStatus,omitempty"`
}

type ArrayOfRefundsdealslisttype struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfRefundsdealslisttype"`

	Item []*RefundsDealsListType `xml:"item,omitempty"`
}

type RefundDetailsType struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RefundDetailsType"`

	RefundId int32 `xml:"refundId,omitempty"`

	RefundStatus string `xml:"refundStatus,omitempty"`

	RefundQuantity int32 `xml:"refundQuantity,omitempty"`

	CreatedDate time.Time `xml:"createdDate,omitempty"`

	ConsiderDate time.Time `xml:"considerDate,omitempty"`
}

type RefundListType struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RefundListType"`

	DealId int64 `xml:"dealId,omitempty"`

	DealDate time.Time `xml:"dealDate,omitempty"`

	BuyerId int32 `xml:"buyerId,omitempty"`

	BuyerLogin string `xml:"buyerLogin,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	BidsCount int32 `xml:"bidsCount,omitempty"`

	QuantityType string `xml:"quantityType,omitempty"`

	RefundDetails *RefundDetailsType `xml:"refundDetails,omitempty"`
}

type ArrayOfRefundlisttype struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfRefundlisttype"`

	Item []*RefundListType `xml:"item,omitempty"`
}

type ReasonInfoType struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ReasonInfoType"`

	ReasonId int32 `xml:"reasonId,omitempty"`

	ReasonName string `xml:"reasonName,omitempty"`

	MaxQuantity int32 `xml:"maxQuantity,omitempty"`
}

type ArrayOfReasoninfotype struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfReasoninfotype"`

	Item []*ReasonInfoType `xml:"item,omitempty"`
}

type RelatedItemStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RelatedItemStruct"`

	ItemId int64 `xml:"itemId,omitempty"`

	ItemTitle string `xml:"itemTitle,omitempty"`

	ItemThumbnail string `xml:"itemThumbnail,omitempty"`

	ItemPrice float32 `xml:"itemPrice,omitempty"`

	ItemBoughtCount int32 `xml:"itemBoughtCount,omitempty"`

	ItemAmount float32 `xml:"itemAmount,omitempty"`

	ItemPaymentType int32 `xml:"itemPaymentType,omitempty"`

	ItemSellerId int64 `xml:"itemSellerId,omitempty"`

	ItemSellerName string `xml:"itemSellerName,omitempty"`

	ItemInvoiceInfo int32 `xml:"itemInvoiceInfo,omitempty"`

	ItemCategoryId int32 `xml:"itemCategoryId,omitempty"`
}

type ArrayOfRelateditemstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfRelateditemstruct"`

	Item []*RelatedItemStruct `xml:"item,omitempty"`
}

type RelatedItemsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RelatedItemsStruct"`

	RelatedItemsInfo *ArrayOfRelateditemstruct `xml:"relatedItemsInfo,omitempty"`

	RelatedItemsAmount float32 `xml:"relatedItemsAmount,omitempty"`
}

type SellFormType struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SellFormType"`

	SellFormId int32 `xml:"sellFormId,omitempty"`

	SellFormTitle string `xml:"sellFormTitle,omitempty"`

	SellFormCat int32 `xml:"sellFormCat,omitempty"`

	SellFormType int32 `xml:"sellFormType,omitempty"`

	SellFormResType int32 `xml:"sellFormResType,omitempty"`

	SellFormDefValue int32 `xml:"sellFormDefValue,omitempty"`

	SellFormOpt int32 `xml:"sellFormOpt,omitempty"`

	SellFormPos int32 `xml:"sellFormPos,omitempty"`

	SellFormLength int32 `xml:"sellFormLength,omitempty"`

	SellMinValue string `xml:"sellMinValue,omitempty"`

	SellMaxValue string `xml:"sellMaxValue,omitempty"`

	SellFormDesc string `xml:"sellFormDesc,omitempty"`

	SellFormOptsValues string `xml:"sellFormOptsValues,omitempty"`

	SellFormFieldDesc string `xml:"sellFormFieldDesc,omitempty"`

	SellFormParamId int32 `xml:"sellFormParamId,omitempty"`

	SellFormParamValues string `xml:"sellFormParamValues,omitempty"`

	SellFormParentId int32 `xml:"sellFormParentId,omitempty"`

	SellFormParentValue string `xml:"sellFormParentValue,omitempty"`

	SellFormUnit string `xml:"sellFormUnit,omitempty"`

	SellFormOptions int32 `xml:"sellFormOptions,omitempty"`
}

type ArrayOfSellformtype struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfSellformtype"`

	Item []*SellFormType `xml:"item,omitempty"`
}

type SellFormFieldsForCategoryStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SellFormFieldsForCategoryStruct"`

	SellFormFieldsList *ArrayOfSellformtype `xml:"sellFormFieldsList,omitempty"`

	SellFormFieldsVersionKey int64 `xml:"sellFormFieldsVersionKey,omitempty"`

	SellFormFieldsComponentValue string `xml:"sellFormFieldsComponentValue,omitempty"`
}

type ShipmentTimeStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ShipmentTimeStruct"`

	ShipmentTimeFrom int32 `xml:"shipmentTimeFrom,omitempty"`

	ShipmentTimeTo int32 `xml:"shipmentTimeTo,omitempty"`
}

type ShipmentDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ShipmentDataStruct"`

	ShipmentId int32 `xml:"shipmentId,omitempty"`

	ShipmentName string `xml:"shipmentName,omitempty"`

	ShipmentType int32 `xml:"shipmentType,omitempty"`

	ShipmentTime *ShipmentTimeStruct `xml:"shipmentTime,omitempty"`
}

type ArrayOfShipmentdatastruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfShipmentdatastruct"`

	Item []*ShipmentDataStruct `xml:"item,omitempty"`
}

type ShipmentPaymentInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ShipmentPaymentInfoStruct"`

	ShipmentId int32 `xml:"shipmentId,omitempty"`

	ShipmentName string `xml:"shipmentName,omitempty"`

	ShipmentAmount float32 `xml:"shipmentAmount,omitempty"`

	ShipmentPaymentType int32 `xml:"shipmentPaymentType,omitempty"`

	ShipmentItemIds *ArrayOfLong `xml:"shipmentItemIds,omitempty"`
}

type ArrayOfShipmentpaymentinfostruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfShipmentpaymentinfostruct"`

	Item []*ShipmentPaymentInfoStruct `xml:"item,omitempty"`
}

type SellerPaymentInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SellerPaymentInfoStruct"`

	ShipmentPaymentInfoPza *ArrayOfShipmentpaymentinfostruct `xml:"shipmentPaymentInfoPza,omitempty"`

	ShipmentPaymentInfoNonPza *ArrayOfShipmentpaymentinfostruct `xml:"shipmentPaymentInfoNonPza,omitempty"`
}

type SellerShipmentDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SellerShipmentDataStruct"`

	SellerId int32 `xml:"sellerId,omitempty"`

	SellerPaymentInfo *SellerPaymentInfoStruct `xml:"sellerPaymentInfo,omitempty"`

	SellerOtherShipmentIsActive int32 `xml:"sellerOtherShipmentIsActive,omitempty"`

	GeneralDeliveryPaymentInfo *ArrayOfShipmentpaymentinfostruct `xml:"generalDeliveryPaymentInfo,omitempty"`
}

type ArrayOfSellershipmentdatastruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfSellershipmentdatastruct"`

	Item []*SellerShipmentDataStruct `xml:"item,omitempty"`
}

type RelatedItemsShipmentDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RelatedItemsShipmentDataStruct"`

	SellerShipmentData *ArrayOfSellershipmentdatastruct `xml:"sellerShipmentData,omitempty"`
}

type ShipmentPriceTypeStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ShipmentPriceTypeStruct"`

	ShipmentPriceTypeId int32 `xml:"shipmentPriceTypeId,omitempty"`

	ShipmentPriceTypeName string `xml:"shipmentPriceTypeName,omitempty"`
}

type ArrayOfShipmentpricetypestruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfShipmentpricetypestruct"`

	Item []*ShipmentPriceTypeStruct `xml:"item,omitempty"`
}

type SiteJournal struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SiteJournal"`

	RowId int64 `xml:"rowId,omitempty"`

	ItemId int64 `xml:"itemId,omitempty"`

	ChangeType string `xml:"changeType,omitempty"`

	ChangeDate int64 `xml:"changeDate,omitempty"`

	CurrentPrice float32 `xml:"currentPrice,omitempty"`

	ItemSellerId int64 `xml:"itemSellerId,omitempty"`
}

type ArrayOfSitejournal struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfSitejournal"`

	Item []*SiteJournal `xml:"item,omitempty"`
}

type SiteJournalDealsStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SiteJournalDealsStruct"`

	DealEventId int64 `xml:"dealEventId,omitempty"`

	DealEventType int32 `xml:"dealEventType,omitempty"`

	DealEventTime int64 `xml:"dealEventTime,omitempty"`

	DealId int64 `xml:"dealId,omitempty"`

	DealTransactionId int64 `xml:"dealTransactionId,omitempty"`

	DealSellerId int32 `xml:"dealSellerId,omitempty"`

	DealItemId int64 `xml:"dealItemId,omitempty"`

	DealBuyerId int32 `xml:"dealBuyerId,omitempty"`

	DealQuantity int32 `xml:"dealQuantity,omitempty"`
}

type ArrayOfSitejournaldealsstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfSitejournaldealsstruct"`

	Item []*SiteJournalDealsStruct `xml:"item,omitempty"`
}

type SiteJournalDealsInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SiteJournalDealsInfoStruct"`

	DealEventsCount int32 `xml:"dealEventsCount,omitempty"`

	DealLastEventTime int64 `xml:"dealLastEventTime,omitempty"`
}

type SiteJournalInfo struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SiteJournalInfo"`

	ItemsNumber int32 `xml:"itemsNumber,omitempty"`

	LastItemDate int64 `xml:"lastItemDate,omitempty"`
}

type StateInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php StateInfoStruct"`

	StateId int32 `xml:"stateId,omitempty"`

	StateName string `xml:"stateName,omitempty"`
}

type ArrayOfStateinfostruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfStateinfostruct"`

	Item []*StateInfoStruct `xml:"item,omitempty"`
}

type MyAccountStruct2 struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php MyAccountStruct2"`

	MyAccountArray *ArrayOfString `xml:"myAccountArray,omitempty"`
}

type ArrayOfMyaccountstruct2 struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfMyaccountstruct2"`

	Item []*MyAccountStruct2 `xml:"item,omitempty"`
}

type ItemBilling struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemBilling"`

	BiName string `xml:"biName,omitempty"`

	BiValue string `xml:"biValue,omitempty"`
}

type ArrayOfItembilling struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfItembilling"`

	Item []*ItemBilling `xml:"item,omitempty"`
}

type MyContactList struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php MyContactList"`

	ContactUserId int32 `xml:"contactUserId,omitempty"`

	ContactNick string `xml:"contactNick,omitempty"`

	ContactFirstName string `xml:"contactFirstName,omitempty"`

	ContactLastName string `xml:"contactLastName,omitempty"`

	ContactCompany string `xml:"contactCompany,omitempty"`

	ContactEmail string `xml:"contactEmail,omitempty"`

	ContactStreet string `xml:"contactStreet,omitempty"`

	ContactPostcode string `xml:"contactPostcode,omitempty"`

	ContactCity string `xml:"contactCity,omitempty"`

	ContactCountry string `xml:"contactCountry,omitempty"`

	ContactPhone string `xml:"contactPhone,omitempty"`

	ContactPhone2 string `xml:"contactPhone2,omitempty"`

	ContactRating string `xml:"contactRating,omitempty"`

	ContactBlocked string `xml:"contactBlocked,omitempty"`
}

type ArrayOfMycontactlist struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfMycontactlist"`

	Item []*MyContactList `xml:"item,omitempty"`
}

type ItemTemplateCreateStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemTemplateCreateStruct"`

	ItemTemplateOption int32 `xml:"itemTemplateOption,omitempty"`

	ItemTemplateName string `xml:"itemTemplateName,omitempty"`
}

type SysStatusType struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SysStatusType"`

	CountryId int32 `xml:"countryId,omitempty"`

	ProgramVersion string `xml:"programVersion,omitempty"`

	CatsVersion string `xml:"catsVersion,omitempty"`

	ApiVersion string `xml:"apiVersion,omitempty"`

	AttribVersion string `xml:"attribVersion,omitempty"`

	FormSellVersion string `xml:"formSellVersion,omitempty"`

	SiteVersion string `xml:"siteVersion,omitempty"`

	VerKey int64 `xml:"verKey,omitempty"`
}

type ArrayOfSysstatustype struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfSysstatustype"`

	Item []*SysStatusType `xml:"item,omitempty"`
}

type BlackListResStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php BlackListResStruct"`

	UserId int64 `xml:"userId,omitempty"`

	Result int32 `xml:"result,omitempty"`
}

type ArrayOfBlacklistresstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfBlacklistresstruct"`

	Item []*BlackListResStruct `xml:"item,omitempty"`
}

type RequestPayoutStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php RequestPayoutStruct"`

	PayoutId int64 `xml:"payoutId,omitempty"`

	PayoutAmount float32 `xml:"payoutAmount,omitempty"`

	PayoutTime int64 `xml:"payoutTime,omitempty"`
}

type StructSellAgain struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php StructSellAgain"`

	SellItemId int64 `xml:"sellItemId,omitempty"`

	SellItemInfo string `xml:"sellItemInfo,omitempty"`

	SellItemLocalId int32 `xml:"sellItemLocalId,omitempty"`
}

type ArrayOfStructsellagain struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfStructsellagain"`

	Item []*StructSellAgain `xml:"item,omitempty"`
}

type StructSellFailed struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php StructSellFailed"`

	SellItemId int64 `xml:"sellItemId,omitempty"`

	SellFaultCode string `xml:"sellFaultCode,omitempty"`

	SellFaultString string `xml:"sellFaultString,omitempty"`
}

type ArrayOfStructsellfailed struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfStructsellfailed"`

	Item []*StructSellFailed `xml:"item,omitempty"`
}

type NewPostBuyFormSellerStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php NewPostBuyFormSellerStruct"`

	SellerId int32 `xml:"sellerId,omitempty"`

	SellerItemIds *ArrayOfLong `xml:"sellerItemIds,omitempty"`

	SellerShipmentId int32 `xml:"sellerShipmentId,omitempty"`

	SellerGdId int32 `xml:"sellerGdId,omitempty"`

	SellerShipmentAmount float32 `xml:"sellerShipmentAmount,omitempty"`

	SellerMessageTo string `xml:"sellerMessageTo,omitempty"`
}

type ArrayOfNewpostbuyformsellerstruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfNewpostbuyformsellerstruct"`

	Item []*NewPostBuyFormSellerStruct `xml:"item,omitempty"`
}

type InvoiceInfoStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php InvoiceInfoStruct"`

	InvoiceAddressType int32 `xml:"invoiceAddressType,omitempty"`

	InvoiceAddressData *AddressUserDataStruct `xml:"invoiceAddressData,omitempty"`

	InvoiceNip string `xml:"invoiceNip,omitempty"`
}

type NewPostBuyFormCommonStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php NewPostBuyFormCommonStruct"`

	PaymentMethodId string `xml:"paymentMethodId,omitempty"`

	ShipmentAddressType int32 `xml:"shipmentAddressType,omitempty"`

	ShipmentAddressData *AddressUserDataStruct `xml:"shipmentAddressData,omitempty"`

	ContactPhone string `xml:"contactPhone,omitempty"`

	InvoiceOption int32 `xml:"invoiceOption,omitempty"`

	InvoiceInfo *InvoiceInfoStruct `xml:"invoiceInfo,omitempty"`
}

type ActionDataStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ActionDataStruct"`

	ActionKey string `xml:"actionKey,omitempty"`

	ActionValue string `xml:"actionValue,omitempty"`
}

type ArrayOfActiondatastruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfActiondatastruct"`

	Item []*ActionDataStruct `xml:"item,omitempty"`
}

type TransactionPayByLinkStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php TransactionPayByLinkStruct"`

	ActionHttpMethod string `xml:"actionHttpMethod,omitempty"`

	ActionUrl string `xml:"actionUrl,omitempty"`

	ActionData *ArrayOfActiondatastruct `xml:"actionData,omitempty"`
}

type PostBuyFormStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php PostBuyFormStruct"`

	TransactionId int64 `xml:"transactionId,omitempty"`

	TransactionPackageIds *ArrayOfLong `xml:"transactionPackageIds,omitempty"`

	TransactionPayByLink *TransactionPayByLinkStruct `xml:"transactionPayByLink,omitempty"`
}

type ItemInfoExt struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ItemInfoExt"`

	ItId int64 `xml:"itId,omitempty"`

	ItCountry int32 `xml:"itCountry,omitempty"`

	ItName string `xml:"itName,omitempty"`

	ItPrice float32 `xml:"itPrice,omitempty"`

	ItBidCount int32 `xml:"itBidCount,omitempty"`

	ItEndingTime int64 `xml:"itEndingTime,omitempty"`

	ItSellerId int64 `xml:"itSellerId,omitempty"`

	ItSellerLogin string `xml:"itSellerLogin,omitempty"`

	ItSellerRating int32 `xml:"itSellerRating,omitempty"`

	ItStartingTime int64 `xml:"itStartingTime,omitempty"`

	ItStartingPrice float32 `xml:"itStartingPrice,omitempty"`

	ItQuantity int32 `xml:"itQuantity,omitempty"`

	ItFotoCount int32 `xml:"itFotoCount,omitempty"`

	ItReservePrice float32 `xml:"itReservePrice,omitempty"`

	ItLocation string `xml:"itLocation,omitempty"`

	ItBuyNowPrice float32 `xml:"itBuyNowPrice,omitempty"`

	ItBuyNowActive int32 `xml:"itBuyNowActive,omitempty"`

	ItAdvertisementPrice float32 `xml:"itAdvertisementPrice,omitempty"`

	ItAdvertisementActive int32 `xml:"itAdvertisementActive,omitempty"`

	ItHighBidder int32 `xml:"itHighBidder,omitempty"`

	ItHighBidderLogin string `xml:"itHighBidderLogin,omitempty"`

	ItDescription string `xml:"itDescription,omitempty"`

	ItStandardizedDescription string `xml:"itStandardizedDescription,omitempty"`

	ItOptions int32 `xml:"itOptions,omitempty"`

	ItState int32 `xml:"itState,omitempty"`

	ItWireTransfer float32 `xml:"itWireTransfer,omitempty"`

	ItPostDelivery float32 `xml:"itPostDelivery,omitempty"`

	ItPostInfo string `xml:"itPostInfo,omitempty"`

	ItQuantityType int32 `xml:"itQuantityType,omitempty"`

	ItIsEco int32 `xml:"itIsEco,omitempty"`

	ItHitCount int64 `xml:"itHitCount,omitempty"`

	ItPostcode string `xml:"itPostcode,omitempty"`

	ItVatInvoice int32 `xml:"itVatInvoice,omitempty"`

	ItVatMarginInvoice int32 `xml:"itVatMarginInvoice,omitempty"`

	ItWithoutVatInvoice int32 `xml:"itWithoutVatInvoice,omitempty"`

	ItBankAccount1 string `xml:"itBankAccount1,omitempty"`

	ItBankAccount2 string `xml:"itBankAccount2,omitempty"`

	ItStartingQuantity int32 `xml:"itStartingQuantity,omitempty"`

	ItIsForGuests int32 `xml:"itIsForGuests,omitempty"`

	ItHasProduct int32 `xml:"itHasProduct,omitempty"`

	ItOrderFulfillmentTime int32 `xml:"itOrderFulfillmentTime,omitempty"`

	ItEndingInfo int32 `xml:"itEndingInfo,omitempty"`

	ItIsAllegroStandard int32 `xml:"itIsAllegroStandard,omitempty"`

	ItIsNewUsed int32 `xml:"itIsNewUsed,omitempty"`

	ItIsBrandZone int32 `xml:"itIsBrandZone,omitempty"`

	ItIsFulfillmentTimeActive int32 `xml:"itIsFulfillmentTimeActive,omitempty"`

	ItEan string `xml:"itEan,omitempty"`

	ItContact string `xml:"itContact,omitempty"`
}

type ShowUserFeedbacks struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ShowUserFeedbacks"`

	UserFLastWeek int32 `xml:"userFLastWeek,omitempty"`

	UserFLastMonth int32 `xml:"userFLastMonth,omitempty"`

	UserFAll int32 `xml:"userFAll,omitempty"`

	UserFSoldItems int32 `xml:"userFSoldItems,omitempty"`

	UserFBuyItems int32 `xml:"userFBuyItems,omitempty"`
}

type SellRatingAverageStruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php SellRatingAverageStruct"`

	SellRatingGroupTitle string `xml:"sellRatingGroupTitle,omitempty"`

	SellRatingGroupAverage float32 `xml:"sellRatingGroupAverage,omitempty"`
}

type ArrayOfSellratingaveragestruct struct {
	XMLName xml.Name `xml:"https://webapi.allegro.pl.allegrosandbox.pl/service.php ArrayOfSellratingaveragestruct"`

	Item []*SellRatingAverageStruct `xml:"item,omitempty"`
}

type ServicePort struct {
	client *SOAPClient
}

func NewServicePort(url string, tls bool, auth *BasicAuth) *ServicePort {
	if url == "" {
		url = "https://webapi.allegro.pl.allegrosandbox.pl/service.php"
	}
	client := NewSOAPClient(url, tls, auth)

	return &ServicePort{
		client: client,
	}
}

func NewServicePortWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *ServicePort {
	if url == "" {
		url = "https://webapi.allegro.pl.allegrosandbox.pl/service.php"
	}
	client := NewSOAPClientWithTLSConfig(url, tlsCfg, auth)

	return &ServicePort{
		client: client,
	}
}

func (service *ServicePort) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *ServicePort) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func (service *ServicePort) DoAddPackageInfoToPostBuyForm(request *DoAddPackageInfoToPostBuyFormRequest) (*DoAddPackageInfoToPostBuyFormResponse, error) {
	response := new(DoAddPackageInfoToPostBuyFormResponse)
	err := service.client.Call("#doAddPackageInfoToPostBuyForm", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoAddToBlackList(request *DoAddToBlackListRequest) (*DoAddToBlackListResponse, error) {
	response := new(DoAddToBlackListResponse)
	err := service.client.Call("#doAddToBlackList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoBidItem(request *DoBidItemRequest) (*DoBidItemResponse, error) {
	response := new(DoBidItemResponse)
	err := service.client.Call("#doBidItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoCancelBidItem(request *DoCancelBidItemRequest) (*DoCancelBidItemResponse, error) {
	response := new(DoCancelBidItemResponse)
	err := service.client.Call("#doCancelBidItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoCancelRefundForm(request *DoCancelRefundFormRequest) (*DoCancelRefundFormResponse, error) {
	response := new(DoCancelRefundFormResponse)
	err := service.client.Call("#doCancelRefundForm", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoCancelRefundWarning(request *DoCancelRefundWarningRequest) (*DoCancelRefundWarningResponse, error) {
	response := new(DoCancelRefundWarningResponse)
	err := service.client.Call("#doCancelRefundWarning", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoCancelTransaction(request *DoCancelTransactionRequest) (*DoCancelTransactionResponse, error) {
	response := new(DoCancelTransactionResponse)
	err := service.client.Call("#doCancelTransaction", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoChangeItemFields(request *DoChangeItemFieldsRequest) (*DoChangeItemFieldsResponse, error) {
	response := new(DoChangeItemFieldsResponse)
	err := service.client.Call("#doChangeItemFields", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoChangePriceItem(request *DoChangePriceItemRequest) (*DoChangePriceItemResponse, error) {
	response := new(DoChangePriceItemResponse)
	err := service.client.Call("#doChangePriceItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoChangeQuantityItem(request *DoChangeQuantityItemRequest) (*DoChangeQuantityItemResponse, error) {
	response := new(DoChangeQuantityItemResponse)
	err := service.client.Call("#doChangeQuantityItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoCheckItemDescription(request *DoCheckItemDescriptionRequest) (*DoCheckItemDescriptionResponse, error) {
	response := new(DoCheckItemDescriptionResponse)
	err := service.client.Call("#doCheckItemDescription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoCheckNewAuctionExt(request *DoCheckNewAuctionExtRequest) (*DoCheckNewAuctionExtResponse, error) {
	response := new(DoCheckNewAuctionExtResponse)
	err := service.client.Call("#doCheckNewAuctionExt", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoFinishItem(request *DoFinishItemRequest) (*DoFinishItemResponse, error) {
	response := new(DoFinishItemResponse)
	err := service.client.Call("#doFinishItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoFinishItems(request *DoFinishItemsRequest) (*DoFinishItemsResponse, error) {
	response := new(DoFinishItemsResponse)
	err := service.client.Call("#doFinishItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetArchiveRefundsList(request *DoGetArchiveRefundsListRequest) (*DoGetArchiveRefundsListResponse, error) {
	response := new(DoGetArchiveRefundsListResponse)
	err := service.client.Call("#doGetArchiveRefundsList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetBidItem2(request *DoGetBidItem2Request) (*DoGetBidItem2Response, error) {
	response := new(DoGetBidItem2Response)
	err := service.client.Call("#doGetBidItem2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetBlackListUsers(request *DoGetBlackListUsersRequest) (*DoGetBlackListUsersResponse, error) {
	response := new(DoGetBlackListUsersResponse)
	err := service.client.Call("#doGetBlackListUsers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetCategoryPath(request *DoGetCategoryPathRequest) (*DoGetCategoryPathResponse, error) {
	response := new(DoGetCategoryPathResponse)
	err := service.client.Call("#doGetCategoryPath", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetCatsData(request *DoGetCatsDataRequest) (*DoGetCatsDataResponse, error) {
	response := new(DoGetCatsDataResponse)
	err := service.client.Call("#doGetCatsData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetCatsDataCount(request *DoGetCatsDataCountRequest) (*DoGetCatsDataCountResponse, error) {
	response := new(DoGetCatsDataCountResponse)
	err := service.client.Call("#doGetCatsDataCount", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetCatsDataLimit(request *DoGetCatsDataLimitRequest) (*DoGetCatsDataLimitResponse, error) {
	response := new(DoGetCatsDataLimitResponse)
	err := service.client.Call("#doGetCatsDataLimit", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetCountries(request *DoGetCountriesRequest) (*DoGetCountriesResponse, error) {
	response := new(DoGetCountriesResponse)
	err := service.client.Call("#doGetCountries", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetDeals(request *DoGetDealsRequest) (*DoGetDealsResponse, error) {
	response := new(DoGetDealsResponse)
	err := service.client.Call("#doGetDeals", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetFilledPostBuyForms(request *DoGetFilledPostBuyFormsRequest) (*DoGetFilledPostBuyFormsResponse, error) {
	response := new(DoGetFilledPostBuyFormsResponse)
	err := service.client.Call("#doGetFilledPostBuyForms", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetFreeDeliveryAmount(request *DoGetFreeDeliveryAmountRequest) (*DoGetFreeDeliveryAmountResponse, error) {
	response := new(DoGetFreeDeliveryAmountResponse)
	err := service.client.Call("#doGetFreeDeliveryAmount", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetItemFields(request *DoGetItemFieldsRequest) (*DoGetItemFieldsResponse, error) {
	response := new(DoGetItemFieldsResponse)
	err := service.client.Call("#doGetItemFields", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetItemsImages(request *DoGetItemsImagesRequest) (*DoGetItemsImagesResponse, error) {
	response := new(DoGetItemsImagesResponse)
	err := service.client.Call("#doGetItemsImages", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetItemsInfo(request *DoGetItemsInfoRequest) (*DoGetItemsInfoResponse, error) {
	response := new(DoGetItemsInfoResponse)
	err := service.client.Call("#doGetItemsInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetItemsList(request *DoGetItemsListRequest) (*DoGetItemsListResponse, error) {
	response := new(DoGetItemsListResponse)
	err := service.client.Call("#doGetItemsList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyAddresses(request *DoGetMyAddressesRequest) (*DoGetMyAddressesResponse, error) {
	response := new(DoGetMyAddressesResponse)
	err := service.client.Call("#doGetMyAddresses", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyBidItems(request *DoGetMyBidItemsRequest) (*DoGetMyBidItemsResponse, error) {
	response := new(DoGetMyBidItemsResponse)
	err := service.client.Call("#doGetMyBidItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyCurrentShipmentPriceType(request *DoGetMyCurrentShipmentPriceTypeRequest) (*DoGetMyCurrentShipmentPriceTypeResponse, error) {
	response := new(DoGetMyCurrentShipmentPriceTypeResponse)
	err := service.client.Call("#doGetMyCurrentShipmentPriceType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyData(request *DoGetMyDataRequest) (*DoGetMyDataResponse, error) {
	response := new(DoGetMyDataResponse)
	err := service.client.Call("#doGetMyData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyFutureItems(request *DoGetMyFutureItemsRequest) (*DoGetMyFutureItemsResponse, error) {
	response := new(DoGetMyFutureItemsResponse)
	err := service.client.Call("#doGetMyFutureItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyIncomingPayments(request *DoGetMyIncomingPaymentsRequest) (*DoGetMyIncomingPaymentsResponse, error) {
	response := new(DoGetMyIncomingPaymentsResponse)
	err := service.client.Call("#doGetMyIncomingPayments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyIncomingPaymentsRefunds(request *DoGetMyIncomingPaymentsRefundsRequest) (*DoGetMyIncomingPaymentsRefundsResponse, error) {
	response := new(DoGetMyIncomingPaymentsRefundsResponse)
	err := service.client.Call("#doGetMyIncomingPaymentsRefunds", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyNotSoldItems(request *DoGetMyNotSoldItemsRequest) (*DoGetMyNotSoldItemsResponse, error) {
	response := new(DoGetMyNotSoldItemsResponse)
	err := service.client.Call("#doGetMyNotSoldItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyNotWonItems(request *DoGetMyNotWonItemsRequest) (*DoGetMyNotWonItemsResponse, error) {
	response := new(DoGetMyNotWonItemsResponse)
	err := service.client.Call("#doGetMyNotWonItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyPayments(request *DoGetMyPaymentsRequest) (*DoGetMyPaymentsResponse, error) {
	response := new(DoGetMyPaymentsResponse)
	err := service.client.Call("#doGetMyPayments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyPaymentsInfo(request *DoGetMyPaymentsInfoRequest) (*DoGetMyPaymentsInfoResponse, error) {
	response := new(DoGetMyPaymentsInfoResponse)
	err := service.client.Call("#doGetMyPaymentsInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyPaymentsRefunds(request *DoGetMyPaymentsRefundsRequest) (*DoGetMyPaymentsRefundsResponse, error) {
	response := new(DoGetMyPaymentsRefundsResponse)
	err := service.client.Call("#doGetMyPaymentsRefunds", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyPayouts(request *DoGetMyPayoutsRequest) (*DoGetMyPayoutsResponse, error) {
	response := new(DoGetMyPayoutsResponse)
	err := service.client.Call("#doGetMyPayouts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyPayoutsDetails(request *DoGetMyPayoutsDetailsRequest) (*DoGetMyPayoutsDetailsResponse, error) {
	response := new(DoGetMyPayoutsDetailsResponse)
	err := service.client.Call("#doGetMyPayoutsDetails", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMySellItems(request *DoGetMySellItemsRequest) (*DoGetMySellItemsResponse, error) {
	response := new(DoGetMySellItemsResponse)
	err := service.client.Call("#doGetMySellItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMySoldItems(request *DoGetMySoldItemsRequest) (*DoGetMySoldItemsResponse, error) {
	response := new(DoGetMySoldItemsResponse)
	err := service.client.Call("#doGetMySoldItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetMyWonItems(request *DoGetMyWonItemsRequest) (*DoGetMyWonItemsResponse, error) {
	response := new(DoGetMyWonItemsResponse)
	err := service.client.Call("#doGetMyWonItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetPaymentMethods(request *DoGetPaymentMethodsRequest) (*DoGetPaymentMethodsResponse, error) {
	response := new(DoGetPaymentMethodsResponse)
	err := service.client.Call("#doGetPaymentMethods", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetPostBuyData(request *DoGetPostBuyDataRequest) (*DoGetPostBuyDataResponse, error) {
	response := new(DoGetPostBuyDataResponse)
	err := service.client.Call("#doGetPostBuyData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetPostBuyFormsDataForBuyers(request *DoGetPostBuyFormsDataForBuyersRequest) (*DoGetPostBuyFormsDataForBuyersResponse, error) {
	response := new(DoGetPostBuyFormsDataForBuyersResponse)
	err := service.client.Call("#doGetPostBuyFormsDataForBuyers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetPostBuyFormsDataForSellers(request *DoGetPostBuyFormsDataForSellersRequest) (*DoGetPostBuyFormsDataForSellersResponse, error) {
	response := new(DoGetPostBuyFormsDataForSellersResponse)
	err := service.client.Call("#doGetPostBuyFormsDataForSellers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetPostBuyFormsIds(request *DoGetPostBuyFormsIdsRequest) (*DoGetPostBuyFormsIdsResponse, error) {
	response := new(DoGetPostBuyFormsIdsResponse)
	err := service.client.Call("#doGetPostBuyFormsIds", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetPostBuyItemInfo(request *DoGetPostBuyItemInfoRequest) (*DoGetPostBuyItemInfoResponse, error) {
	response := new(DoGetPostBuyItemInfoResponse)
	err := service.client.Call("#doGetPostBuyItemInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetRefundsDeals(request *DoGetRefundsDealsRequest) (*DoGetRefundsDealsResponse, error) {
	response := new(DoGetRefundsDealsResponse)
	err := service.client.Call("#doGetRefundsDeals", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetRefundsList(request *DoGetRefundsListRequest) (*DoGetRefundsListResponse, error) {
	response := new(DoGetRefundsListResponse)
	err := service.client.Call("#doGetRefundsList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetRefundsReasons(request *DoGetRefundsReasonsRequest) (*DoGetRefundsReasonsResponse, error) {
	response := new(DoGetRefundsReasonsResponse)
	err := service.client.Call("#doGetRefundsReasons", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetRelatedItems(request *DoGetRelatedItemsRequest) (*DoGetRelatedItemsResponse, error) {
	response := new(DoGetRelatedItemsResponse)
	err := service.client.Call("#doGetRelatedItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetSellFormFieldsForCategory(request *DoGetSellFormFieldsForCategoryRequest) (*DoGetSellFormFieldsForCategoryResponse, error) {
	response := new(DoGetSellFormFieldsForCategoryResponse)
	err := service.client.Call("#doGetSellFormFieldsForCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetSessionHandleForWidget(request *DoGetSessionHandleForWidgetRequest) (*DoGetSessionHandleForWidgetResponse, error) {
	response := new(DoGetSessionHandleForWidgetResponse)
	err := service.client.Call("#doGetSessionHandleForWidget", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetShipmentData(request *DoGetShipmentDataRequest) (*DoGetShipmentDataResponse, error) {
	response := new(DoGetShipmentDataResponse)
	err := service.client.Call("#doGetShipmentData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetShipmentDataForRelatedItems(request *DoGetShipmentDataForRelatedItemsRequest) (*DoGetShipmentDataForRelatedItemsResponse, error) {
	response := new(DoGetShipmentDataForRelatedItemsResponse)
	err := service.client.Call("#doGetShipmentDataForRelatedItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetShipmentPriceTypes(request *DoGetShipmentPriceTypesRequest) (*DoGetShipmentPriceTypesResponse, error) {
	response := new(DoGetShipmentPriceTypesResponse)
	err := service.client.Call("#doGetShipmentPriceTypes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetSiteJournal(request *DoGetSiteJournalRequest) (*DoGetSiteJournalResponse, error) {
	response := new(DoGetSiteJournalResponse)
	err := service.client.Call("#doGetSiteJournal", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetSiteJournalDeals(request *DoGetSiteJournalDealsRequest) (*DoGetSiteJournalDealsResponse, error) {
	response := new(DoGetSiteJournalDealsResponse)
	err := service.client.Call("#doGetSiteJournalDeals", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetSiteJournalDealsInfo(request *DoGetSiteJournalDealsInfoRequest) (*DoGetSiteJournalDealsInfoResponse, error) {
	response := new(DoGetSiteJournalDealsInfoResponse)
	err := service.client.Call("#doGetSiteJournalDealsInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetSiteJournalInfo(request *DoGetSiteJournalInfoRequest) (*DoGetSiteJournalInfoResponse, error) {
	response := new(DoGetSiteJournalInfoResponse)
	err := service.client.Call("#doGetSiteJournalInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetStatesInfo(request *DoGetStatesInfoRequest) (*DoGetStatesInfoResponse, error) {
	response := new(DoGetStatesInfoResponse)
	err := service.client.Call("#doGetStatesInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetSystemTime(request *DoGetSystemTimeRequest) (*DoGetSystemTimeResponse, error) {
	response := new(DoGetSystemTimeResponse)
	err := service.client.Call("#doGetSystemTime", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetTransactionsIDs(request *DoGetTransactionsIDsRequest) (*DoGetTransactionsIDsResponse, error) {
	response := new(DoGetTransactionsIDsResponse)
	err := service.client.Call("#doGetTransactionsIDs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetUserID(request *DoGetUserIDRequest) (*DoGetUserIDResponse, error) {
	response := new(DoGetUserIDResponse)
	err := service.client.Call("#doGetUserID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoGetUserLogin(request *DoGetUserLoginRequest) (*DoGetUserLoginResponse, error) {
	response := new(DoGetUserLoginResponse)
	err := service.client.Call("#doGetUserLogin", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoLogin(request *DoLoginRequest) (*DoLoginResponse, error) {
	response := new(DoLoginResponse)
	err := service.client.Call("#doLogin", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoLoginEnc(request *DoLoginEncRequest) (*DoLoginEncResponse, error) {
	response := new(DoLoginEncResponse)
	err := service.client.Call("#doLoginEnc", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoLoginWithAccessToken(request *DoLoginWithAccessTokenRequest) (*DoLoginWithAccessTokenResponse, error) {
	response := new(DoLoginWithAccessTokenResponse)
	err := service.client.Call("#doLoginWithAccessToken", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoMyAccount2(request *DoMyAccount2Request) (*DoMyAccount2Response, error) {
	response := new(DoMyAccount2Response)
	err := service.client.Call("#doMyAccount2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoMyAccountItemsCount(request *DoMyAccountItemsCountRequest) (*DoMyAccountItemsCountResponse, error) {
	response := new(DoMyAccountItemsCountResponse)
	err := service.client.Call("#doMyAccountItemsCount", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoMyBilling(request *DoMyBillingRequest) (*DoMyBillingResponse, error) {
	response := new(DoMyBillingResponse)
	err := service.client.Call("#doMyBilling", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoMyBillingItem(request *DoMyBillingItemRequest) (*DoMyBillingItemResponse, error) {
	response := new(DoMyBillingItemResponse)
	err := service.client.Call("#doMyBillingItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoMyContact(request *DoMyContactRequest) (*DoMyContactResponse, error) {
	response := new(DoMyContactResponse)
	err := service.client.Call("#doMyContact", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoNewAuctionExt(request *DoNewAuctionExtRequest) (*DoNewAuctionExtResponse, error) {
	response := new(DoNewAuctionExtResponse)
	err := service.client.Call("#doNewAuctionExt", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoQueryAllSysStatus(request *DoQueryAllSysStatusRequest) (*DoQueryAllSysStatusResponse, error) {
	response := new(DoQueryAllSysStatusResponse)
	err := service.client.Call("#doQueryAllSysStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoQuerySysStatus(request *DoQuerySysStatusRequest) (*DoQuerySysStatusResponse, error) {
	response := new(DoQuerySysStatusResponse)
	err := service.client.Call("#doQuerySysStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoRemoveFromBlackList(request *DoRemoveFromBlackListRequest) (*DoRemoveFromBlackListResponse, error) {
	response := new(DoRemoveFromBlackListResponse)
	err := service.client.Call("#doRemoveFromBlackList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoRequestCancelBid(request *DoRequestCancelBidRequest) (*DoRequestCancelBidResponse, error) {
	response := new(DoRequestCancelBidResponse)
	err := service.client.Call("#doRequestCancelBid", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoRequestPayout(request *DoRequestPayoutRequest) (*DoRequestPayoutResponse, error) {
	response := new(DoRequestPayoutResponse)
	err := service.client.Call("#doRequestPayout", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoRequestSurcharge(request *DoRequestSurchargeRequest) (*DoRequestSurchargeResponse, error) {
	response := new(DoRequestSurchargeResponse)
	err := service.client.Call("#doRequestSurcharge", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoSellSomeAgain(request *DoSellSomeAgainRequest) (*DoSellSomeAgainResponse, error) {
	response := new(DoSellSomeAgainResponse)
	err := service.client.Call("#doSellSomeAgain", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoSellSomeAgainInShop(request *DoSellSomeAgainInShopRequest) (*DoSellSomeAgainInShopResponse, error) {
	response := new(DoSellSomeAgainInShopResponse)
	err := service.client.Call("#doSellSomeAgainInShop", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoSendEmailToUser(request *DoSendEmailToUserRequest) (*DoSendEmailToUserResponse, error) {
	response := new(DoSendEmailToUserResponse)
	err := service.client.Call("#doSendEmailToUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoSendPostBuyForm(request *DoSendPostBuyFormRequest) (*DoSendPostBuyFormResponse, error) {
	response := new(DoSendPostBuyFormResponse)
	err := service.client.Call("#doSendPostBuyForm", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoSendRefundForm(request *DoSendRefundFormRequest) (*DoSendRefundFormResponse, error) {
	response := new(DoSendRefundFormResponse)
	err := service.client.Call("#doSendRefundForm", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoSetFreeDeliveryAmount(request *DoSetFreeDeliveryAmountRequest) (*DoSetFreeDeliveryAmountResponse, error) {
	response := new(DoSetFreeDeliveryAmountResponse)
	err := service.client.Call("#doSetFreeDeliveryAmount", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoSetShipmentPriceType(request *DoSetShipmentPriceTypeRequest) (*DoSetShipmentPriceTypeResponse, error) {
	response := new(DoSetShipmentPriceTypeResponse)
	err := service.client.Call("#doSetShipmentPriceType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoShowItemInfoExt(request *DoShowItemInfoExtRequest) (*DoShowItemInfoExtResponse, error) {
	response := new(DoShowItemInfoExtResponse)
	err := service.client.Call("#doShowItemInfoExt", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoShowUser(request *DoShowUserRequest) (*DoShowUserResponse, error) {
	response := new(DoShowUserResponse)
	err := service.client.Call("#doShowUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePort) DoVerifyItem(request *DoVerifyItemRequest) (*DoVerifyItemResponse, error) {
	response := new(DoVerifyItemResponse)
	err := service.client.Call("#doVerifyItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Items []interface{} `xml:",omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU  string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tlsCfg  *tls.Config
	auth    *BasicAuth
	headers []interface{}
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: "UsernameToken-" + randStringBytesMaskImprSrc(9)}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, insecureSkipVerify bool, auth *BasicAuth) *SOAPClient {
	tlsCfg := &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	return NewSOAPClientWithTLSConfig(url, tlsCfg, auth)
}

func NewSOAPClientWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:    url,
		tlsCfg: tlsCfg,
		auth:   auth,
	}
}

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: s.tlsCfg,
		Dial:            dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
