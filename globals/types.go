package globals

import (
	"time"
)

type YMDFormat time.Time

func (t YMDFormat) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format("2006-01-02")), nil
}

func (t *YMDFormat) UnmarshalJSON(data []byte) error {
	t1, err := time.Parse("2006-01-02", string(data))
	if err != nil {
		return err
	}
	*t = YMDFormat(t1)
	return nil
}

type Seeder interface {
	Seed() string
}
