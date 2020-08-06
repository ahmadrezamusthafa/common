package time

import (
	"encoding/json"
	"strings"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	parsedTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}
