package insider

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-finansinspektionen/test"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type InsiderCSVTestSuite struct {
	suite.Suite
}

func (suite *InsiderCSVTestSuite) TestDecodeUTF16() {
	client := NewClient()
	open, err := os.Open(test.GetFile("test/mocks/Insyn2022-04-06.csv"))
	assert.Nil(suite.T(), err)
	decoded := client.decodeUTF16(open)
	data, err := ioutil.ReadAll(decoded)
	assert.Nil(suite.T(), err)
	assert.Contains(suite.T(), string(data), "då")
}

func (suite *InsiderCSVTestSuite) TestReadCSV() {
	client := NewClient()
	open, err := os.Open(test.GetFile("test/mocks/Insyn2022-04-06.csv"))
	assert.Nil(suite.T(), err)
	decoded := client.decodeUTF16(open)
	data, err := client.ReadCSV(decoded)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 6, len(data))
}

func (suite *InsiderCSVTestSuite) TestReadCSV_ParsesFieldsCorrectly() {
	client := NewClient()
	open, err := os.Open(test.GetFile("test/mocks/Insyn2022-04-06.csv"))
	assert.Nil(suite.T(), err)
	decoded := client.decodeUTF16(open)
	data, err := client.ReadCSV(decoded)
	assert.Nil(suite.T(), err)
	tzSto, err := time.LoadLocation("Europe/Stockholm")
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), "Biofrigas Sweden AB (publ) ", data[0].Issuer)
	assert.Equal(suite.T(), "549300433GW2P6KWVJ09", data[0].LEICode)
	assert.Equal(suite.T(), "Marcus Benzon", data[0].PersonDischargingManagerialResponsibilities)
	assert.Equal(suite.T(), "Marcus Benzon", data[0].Notifier)
	assert.Equal(suite.T(), "", data[0].Responsibilities)
	assert.Equal(suite.T(), "Other\u00a0senior\u00a0executive", data[0].Position)
	assert.Equal(suite.T(), "", data[0].CloselyAssociated)
	assert.Equal(suite.T(), false, bool(data[0].Amendment))
	assert.Equal(suite.T(), "", data[0].DetailsOfAmendment)
	assert.Equal(suite.T(), true, bool(data[0].InitialNotification))
	assert.Equal(suite.T(), true, bool(data[0].LinkedToShareOptionProgramme))
	assert.Equal(suite.T(), "Acquisition", data[0].NatureOfTransaction)
	assert.Equal(suite.T(), "Share", data[0].InstrumentType)
	assert.Equal(suite.T(), 1250.0, data[0].Volume)
	assert.Equal(suite.T(), "", data[0].Responsibilities)
	assert.Equal(suite.T(), time.Date(2022, 4, 6, 0, 0, 0, 0, tzSto), data[0].TransactionDate.GetTime())
	assert.Equal(suite.T(), 5.67, data[0].Price)
	assert.Equal(suite.T(), "Quantity", data[0].Unit)
	assert.Equal(suite.T(), "SEK", data[0].Currency)
	assert.Equal(suite.T(), "Outside a trading venue", data[0].TradingVenue)
	assert.Equal(suite.T(), "Current", data[0].Status)
	assert.Equal(suite.T(), time.Date(2022, 4, 6, 11, 55, 5, 0, tzSto), data[0].PublicationDate.GetTime())
	assert.Equal(suite.T(), true, bool(data[0].LinkedToShareOptionProgramme))
	assert.Equal(suite.T(), true, bool(data[3].Amendment))
	assert.Equal(suite.T(), "Teckningsoptionerna har dragits in då programmet avbrutits och optionerna ska makuleras ", data[3].DetailsOfAmendment)
}

func TestInsiderCSVTestSuite(t *testing.T) {
	suite.Run(t, new(InsiderCSVTestSuite))
}
