package insider

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Transaction struct {
	PublicationDate                             DateTime  `csv:"Publication date"`
	Issuer                                      string    `csv:"Issuer"`
	LEICode                                     string    `csv:"LEI-code"`
	Notifier                                    string    `csv:"Notifier"`
	PersonDischargingManagerialResponsibilities string    `csv:"Person discharging managerial responsibilities"`
	Responsibilities                            string    `csv:"Responsibilities"`
	Position                                    string    `csv:"Position"`
	CloselyAssociated                           string    `csv:"Closely associated"`
	Amendment                                   YesNoBool `csv:"Amendment"`
	DetailsOfAmendment                          string    `csv:"Details of amendment"`
	InitialNotification                         YesNoBool `csv:"Initial notification"`
	LinkedToShareOptionProgramme                YesNoBool `csv:"Linked to share option programme"`
	NatureOfTransaction                         string    `csv:"Nature of transaction"`
	InstrumentType                              string    `csv:"Intrument type"`
	ISIN                                        string    `csv:"ISIN"`
	TransactionDate                             DateTime  `csv:"Transaction date"`
	Volume                                      float64   `csv:"Volume"`
	Unit                                        string    `csv:"Unit"`
	Price                                       float64   `csv:"Price"`
	Currency                                    string    `csv:"Currency"`
	TradingVenue                                string    `csv:"Trading venue"`
	Status                                      string    `csv:"Status"`
}

func (f *Transaction) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("PublicationDate", f.PublicationDate.GetTime().String())
	enc.AddString("Issuer", f.Issuer)
	enc.AddString("LEICode", f.LEICode)
	enc.AddString("Notifier", f.Notifier)
	enc.AddString("PersonDischargingManagerialResponsibilities", f.PersonDischargingManagerialResponsibilities)
	enc.AddString("Responsibilities", f.Responsibilities)
	enc.AddString("Position", f.Position)
	enc.AddString("CloselyAssociated", f.CloselyAssociated)
	enc.AddBool("Amendment", bool(f.Amendment))
	enc.AddString("DetailsOfAmendment", f.DetailsOfAmendment)
	enc.AddBool("InitialNotification", bool(f.InitialNotification))
	enc.AddBool("LinkedToShareOptionProgramme", bool(f.LinkedToShareOptionProgramme))
	enc.AddString("NatureOfTransaction", f.NatureOfTransaction)
	enc.AddString("InstrumentType", f.InstrumentType)
	enc.AddString("ISIN", f.ISIN)
	enc.AddString("TransactionDate", f.TransactionDate.GetTime().String())
	enc.AddFloat64("Volume", f.Volume)
	enc.AddString("Unit", f.Unit)
	enc.AddFloat64("Price", f.Price)
	enc.AddString("Currency", f.Currency)
	enc.AddString("TradingVenue", f.TradingVenue)
	enc.AddString("Status", f.Status)
	return nil
}

type YesNoBool bool

// UnmarshalCSV is an implementation of the Unmarshaler interface, converts a string record to a native
// value for this type.
func (ynb *YesNoBool) UnmarshalCSV(data []byte) error {
	if ynb == nil {
		return fmt.Errorf("cannot unmarshal into nil pointer")
	}
	switch strings.ToLower(string(data)) {
	case "yes":
		*ynb = YesNoBool(true)
		return nil
	case "no":
		*ynb = YesNoBool(false)
		return nil
	}
	*ynb = YesNoBool(false)
	return nil
}
