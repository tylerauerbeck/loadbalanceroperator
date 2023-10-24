package srv

import (
	"context"
	"time"

	lbapi "go.infratographer.com/load-balancer-api/pkg/client"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/gidx"
)

const (
	LBPrefix = "loadbal"

	typeLB      = 1
	typeAssocLB = 2
	typeNoLB    = 0
)

type loadBalancer struct {
	loadBalancerID gidx.PrefixedID
	lbData         *lbapi.LoadBalancer
	lbType         int
}

type lbTask struct {
	id    string
	event string
	ts    time.Time
}

type runner struct {
	reader     chan lbTask
	writer     chan lbTask
	quit       chan struct{}
	buffer     []lbTask
	taskRunner func(lbTask)
}

type taskRunner func(lbTask)

type Message interface {
	events.EventMessage | events.ChangeMessage
	GetTraceContext(ctx context.Context) context.Context
	GetSubject() gidx.PrefixedID
	GetAddSubjects() []gidx.PrefixedID
}
