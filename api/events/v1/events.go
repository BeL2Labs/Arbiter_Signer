// Copyright (c) 2025 The bel2 developers
package v1

import (
	"github.com/BeL2Labs/Arbiter_Signer/utility/events"

	"github.com/gogf/gf/v2/frame/g"
)

type AllEventsReq struct {
	g.Meta `path:"/events" tags:"All Events" method:"get" summary:"Get all events you have processed."`
}

type AllEventsRes struct {
	Events []events.EventInfo `json:"events"`
}
