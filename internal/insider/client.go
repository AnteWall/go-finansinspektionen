package insider

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jszwec/csvutil"
	"go-finansinspektionen/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/text/encoding/unicode"
	"io"
	"log"
	"time"
)

const DefaultBaseUrl string = "https://marknadssok.fi.se/publiceringsklient"
const LanguageEN string = "en-GB"

func GetDefaultHeaders() map[string]string {
	return map[string]string{
		"Content-Type":    "text/csv",
		"Accept-Encoding": "gzip",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:94.0) Gecko/20100101 Firefox/94.0",
	}
}

type insiderClient struct {
	httpClient *resty.Client
	baseUrl    string
	logger     *zap.SugaredLogger
}

func (i *insiderClient) request() error {
	formattedDate := time.Now().String()
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
		return err
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
	utf16Decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()

	csvReader := csv.NewReader(utf16Decoder.Reader(reader))
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}
	var transactions []Transaction
	for {
		var t Transaction
		if err := dec.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		i.logger.Infow("hello", zap.Object("transaction", &t))
		transactions = append(transactions, t)
	}

	return nil
}

func (i *insiderClient) GetLatestTrades() error {
	err := i.request()
	if err != nil {
		return err
	}
	return nil
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
