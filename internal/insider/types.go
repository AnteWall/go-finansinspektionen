package insider

import "go.uber.org/zap/zapcore"

type Transaction struct {
	PublicationDate                             DateTime `csv:"Publication date"`
	Issuer                                      string   `csv:"Issuer"`
	LEICode                                     string   `csv:"LEI-code"`
	Notifier                                    string   `csv:"Notifier"`
	PersonDischargingManagerialResponsibilities string   `csv:"Person discharging managerial responsibilities"`
	Responsibilities                            string   `csv:"Responsibilities"`
	Position                                    string   `csv:"Position"`
	CloselyAssociated                           string   `csv:"Closely associated"`
	Amendment                                   string   `csv:"Amendment"`
	DetailsOfAmendment                          string   `csv:"Details of amendment"`
	InitialNotification                         string   `csv:"Initial notification"`
	LinkedToShareOptionProgramme                string   `csv:"Linked to share option programme"`
	NatureOfTransaction                         string   `csv:"Nature of transaction"`
	InstrumentType                              string   `csv:"Intrument type"`
	ISIN                                        string   `csv:"ISIN"`
	TransactionDate                             DateTime `csv:"Transaction date"`
	Volume                                      float64  `csv:"Volume"`
	Unit                                        string   `csv:"Unit"`
	Price                                       float64  `csv:"Price"`
	Currency                                    string   `csv:"Currency"`
	TradingVenue                                string   `csv:"Trading venue"`
	Status                                      string   `csv:"Status"`
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
	enc.AddString("Amendment", f.Amendment)
	enc.AddString("DetailsOfAmendment", f.DetailsOfAmendment)
	enc.AddString("InitialNotification", f.InitialNotification)
	enc.AddString("LinkedToShareOptionProgramme", f.LinkedToShareOptionProgramme)
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
