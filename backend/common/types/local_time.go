package types

import "time"

const localDateTimeFormat string = "2006-01-02 15:04:05"
const localDateFormat string = "2006-01-02"

type LocalTime time.Time

func (l LocalTime) MarshalJSON() ([]byte, error) {
	if time.Time(l).IsZero() {
		return []byte(`""`), nil
	}
	b := make([]byte, 0, len(localDateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(l).AppendFormat(b, localDateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (l *LocalTime) UnmarshalJSON(b []byte) error {
	if string(b) == `""` { // `""` 并不是完全的空串
		return nil
	}
	format := localDateTimeFormat
	if len(b) == 12 { // 加上引号的长度
		format = localDateFormat
	}
	now, err := time.ParseInLocation(`"`+format+`"`, string(b), time.Local)
	*l = LocalTime(now)
	return err
}

func (l *LocalTime) Time() time.Time {
	return time.Time(*l)
}

func (l *LocalTime) TimePtr() *time.Time {
	t := time.Time(*l)
	return &t
}
