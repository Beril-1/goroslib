//nolint:golint
package actionlib

import (
	"github.com/aler9/goroslib/pkg/msg"
	"time"
)

const (
	TestRequestActionGoal_TERMINATE_SUCCESS   int32 = 0
	TestRequestActionGoal_TERMINATE_ABORTED   int32 = 1
	TestRequestActionGoal_TERMINATE_REJECTED  int32 = 2
	TestRequestActionGoal_TERMINATE_LOSE      int32 = 3
	TestRequestActionGoal_TERMINATE_DROP      int32 = 4
	TestRequestActionGoal_TERMINATE_EXCEPTION int32 = 5
)

type TestRequestActionGoal struct {
	msg.Definitions `ros:"int32 TERMINATE_SUCCESS=0,int32 TERMINATE_ABORTED=1,int32 TERMINATE_REJECTED=2,int32 TERMINATE_LOSE=3,int32 TERMINATE_DROP=4,int32 TERMINATE_EXCEPTION=5"`
	TerminateStatus int32
	IgnoreCancel    bool
	ResultText      string
	TheResult       int32
	IsSimpleClient  bool
	DelayAccept     time.Duration
	DelayTerminate  time.Duration
	PauseStatus     time.Duration
}

type TestRequestActionResult struct {
	TheResult      int32
	IsSimpleServer bool
}

type TestRequestActionFeedback struct {
}

type TestRequestAction struct {
	msg.Package `ros:"actionlib"`
	TestRequestActionGoal
	TestRequestActionResult
	TestRequestActionFeedback
}
