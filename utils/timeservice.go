package utils

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

//TimeService interface provides time related functions
type TimeService interface {
	Now() time.Time
	NowTimestamp() timestamp.Timestamp
}

type timeUtil struct {
}

func (tu *timeUtil) Now() time.Time {
	return time.Now()
}

func (tu *timeUtil) NowTimestamp() timestamp.Timestamp {
	t := tu.Now()
	seconds := t.Unix()
	nanos := int32(t.Sub(time.Unix(seconds, 0)))
	ts := timestamp.Timestamp{
		Seconds: seconds,
		Nanos:   nanos,
	}
	return ts
}

func NewTimeService() TimeService {
	return new(timeUtil)
}
