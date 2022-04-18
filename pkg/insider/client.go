package insider

import (
	"compress/gzip"
	"fmt"
	"github.com/AnteWall/go-finansinspektionen/pkg/utils"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"io"
	"time"
)

const DefaultBaseUrl string = "https://marknadssok.fi.se/publiceringsklient"
const LanguageEN string = "en-GB"

func GetDefaultHeaders() map[string]string {
	return map[string]string{
		"Content-Type":    "text/csv.go",
		"Accept-Encoding": "gzip",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:94.0) Gecko/20100101 Firefox/94.0",
	}
}

type insiderClient struct {
	httpClient *resty.Client
	baseUrl    string
	logger     *zap.SugaredLogger
}

func (i *insiderClient) request(date time.Time) ([]*Transaction, error) {
	formattedDate := date.Format("2006-01-02")
	get, err := i.httpClient.R().
		SetHeaders(GetDefaultHeaders()).
		SetQueryParam("button", "export").
		SetQueryParam("SearchFunctionType", "Insyn").
		SetQueryParam("Publiceringsdatum.From", formattedDate).
		SetQueryParam("Publiceringsdatum.To", formattedDate).
		SetDoNotParseResponse(true).
		Get(
			fmt.Sprintf("%s/%s/Search/Search", i.baseUrl, LanguageEN),
		)
	if err != nil {
		return nil, err
	}
	defer get.RawBody().Close()

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch get.Header().Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(get.RawBody())
		defer reader.Close()
	default:
		print(get.Header().Get("Content-Encoding"))
		reader = get.RawBody()
	}
	return i.ReadCSV(i.decodeUTF16(reader))
}

func (i *insiderClient) GetTransactions(day time.Time) ([]*Transaction, error) {
	return i.request(day)
}

func (i *insiderClient) GetTodayTransactions() ([]*Transaction, error) {
	return i.GetTransactions(time.Now())
}

func NewClient(opts ...func(*insiderClient)) *insiderClient {
	httpClient := resty.New()
	logger := utils.NewLogger()
	c := &insiderClient{
		logger:     logger,
		baseUrl:    DefaultBaseUrl,
		httpClient: httpClient,
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

func WithBaseUrl(url string) func(client *insiderClient) {
	return func(client *insiderClient) {
		client.baseUrl = url
	}
}

func WithLogger(logger *zap.SugaredLogger) func(client *insiderClient) {
	return func(client *insiderClient) {
		client.logger = logger
	}
}

func WithDebug(debug bool) func(client *insiderClient) {
	return func(client *insiderClient) {
		client.httpClient.SetDebug(debug)
	}
}
