package pwengine

import (
	"context"
	"time"

	"pathwar.land/go/pkg/pwversion"
)

func (e *engine) GetInfo(context.Context, *Void) (*Info, error) {
	return &Info{
		Version: pwversion.Version,
		Commit:  pwversion.Commit,
		BuiltAt: pwversion.Date,
		BuiltBy: pwversion.BuiltBy,
		Uptime:  int32(time.Now().Sub(e.startedAt).Seconds()),
	}, nil
}
