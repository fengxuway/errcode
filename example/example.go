package example

//go:generate errcode -type=Err,Err2 --linecomment -unknowncode=-1 -codefunc=Code

// errcode install command: go install github.com/fengxuway/errcode/cmd/errcode

// Err my custorm error enum
type Err int

// define error enum consts
const (
	ErrPlanIsEmpty        Err = 10 + iota // plan is empty
	ErrManualCreateFailed                 // manual create operation failed
	ErrRPCFailed                          // rpc request failed
	ErrStageNotMatched                    // stage not matched
	ErrNoAvailable                        // not available
	ErrUnmarshalFailed                    // json unmarshal failed
	_
	ErrGetFailed      // get data failed
	ErrPlanNotSupport // plan not support
	ErrDBFailed       // db operation failed
)

type Err2 int32

const (
	Err2Name Err2 = 100001 // invalid name
	Err2Addr Err2 = 200001 // invalid addr
)
