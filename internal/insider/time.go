package insider

import (
	"github.com/jszwec/csvutil"
	"time"
)

type DateTime time.Time


func (f DateTime) MarshalCSV() ([]byte, error) {
	return csvutil.Marshal(f.GetTime())
}

func (f *DateTime) UnmarshalCSV(data []byte) error {
	// 31/03/2022 00:00:00
	stoTz, err := time.LoadLocation("Europe/Stockholm")
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("02/01/2006 15:04:05", string(data), stoTz)
	if err != nil {
		return err
	}
	*f = DateTime(t)
	return nil
}

func (f DateTime) GetTime() time.Time {
	return time.Time(f)
}