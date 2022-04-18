package insider

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"golang.org/x/text/encoding/unicode"
	"io"
)

func (i *Client) decodeUTF16(reader io.ReadCloser) io.Reader {
	utf16Decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	return utf16Decoder.Reader(reader)
}

func (i *Client) ReadCSV(reader io.Reader) ([]*Transaction, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true
	// csvReader.FieldsPerRecord = -1
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, err
	}
	var transactions []*Transaction
	for {
		var t *Transaction
		if err := dec.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
