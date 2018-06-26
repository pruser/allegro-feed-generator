package request

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"k8s.io/kubernetes/pkg/kubelet/kubeletconfig/util/log"

	api_gen "github.com/pruser/allegro-feed-generator/allegro_service/generated"
	"github.com/pruser/allegro-feed-generator/model"
)

type PostHttpClient interface {
	Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
}

type RequestHandler struct {
	client          PostHttpClient
	parser          Parser
	requestCreator  RequestCreator
	responseCreator ResponseCreator
	webAPIKey       model.WebAPIKey
	webApiUrl       string
	urlbase         string
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	log.Errorf("error occurred: %+v", err)
	w.WriteHeader(statusCode)
}

func (rh *RequestHandler) CreateFeed(w http.ResponseWriter, r *http.Request) {
	searchString := r.URL.Query().Get("search")

	parsingResult := rh.parser.Parse(r.URL.Query())
	sortSettings, err := parsingResult.GetSortSettings()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("can't process sort query params, err %v", err))
		return
	}
	filterSettings, err := parsingResult.GetFilterSettings()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("can't process filter query params, err %v", err))
		return
	}

	// request creation
	soap, err := rh.requestCreator.CreateRequest(sortSettings, filterSettings)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("can't create request message, err %v", err))
		return
	}

	res, err := xml.MarshalIndent(soap, " ", " ")
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("can't marshall request message, err %v", err))
		return
	}

	// sending request
	resp, err := rh.client.Post(rh.webApiUrl, "application/xml", bytes.NewBuffer(res))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("can't send request message, err %v", err))
		return
	}

	// output processing
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("can't read response, err %v", err))
		return
	}

	result := api_gen.SOAPEnvelope{Body: api_gen.SOAPBody{Content: &api_gen.DoGetItemsListResponse{ItemsList: &api_gen.ArrayOfItemslisttype{}}}}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal response, err %v", err))
		return
	}

	getItemsListResponse, ok := result.Body.Content.(*api_gen.DoGetItemsListResponse)
	if !ok {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("wrong response format"))
		return
	}

	response, err := rh.responseCreator.CreateResponse(searchString, getItemsListResponse)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("Can't generate feed data, err %v", err))
		return
	}
	w.Write(response)
}

func (rh *RequestHandler) Health(w http.ResponseWriter, r *http.Request) {}

func NewRequestHandler(apikey model.WebAPIKey, webApiUrl, urlbase string) *RequestHandler {
	return &RequestHandler{
		client:          http.DefaultClient,
		parser:          &QueryParser{},
		requestCreator:  &SoapRequestCreator{webApiKey: apikey, webApiUrl: webApiUrl},
		responseCreator: &DefaultResponseCreator{urlBase: urlbase, timeProvider: &RealTimeProvider{}},
		webAPIKey:       apikey,
		webApiUrl:       webApiUrl,
	}
}
