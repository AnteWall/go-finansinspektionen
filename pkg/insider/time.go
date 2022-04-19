package insider

import (
	"time"
)

type DateTime time.Time

func (f DateTime) MarshalCSV() ([]byte, error) {
	return []byte(f.GetTime().Format(time.RFC3339)), nil
}

func (f *DateTime) MarshalJSON() ([]byte, error) {
	return f.GetTime().MarshalJSON()
}

func (f *DateTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(`"2006-01-02T15:04:05.999999999Z07:00"`, string(data))
	*f = DateTime(t)
	return err
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
