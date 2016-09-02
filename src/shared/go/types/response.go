package types 

import (
	"errors"

	"shared/go/log"
)

type Result int32
type ErrorCode byte


const (
	ErrorCodeError ErrorCode = 127
	ErrorCodeNoError ErrorCode = 0
	ResponseTML	TML = 7
)


type CalculationResponse struct {
	TML TML
	RequestID RequestID
	ErrorCode ErrorCode
	Result 	Result 
}

func result_toBytes(n Result) []byte {
	return []byte{byte(n>>24), byte(n>>16), byte(n>>8), byte(n)}
}

func bytes_toResult(n []byte) (Result, error) {
	if len(n) != 4 {
		return 0, errors.New("int32 needs 4 bytes!")
	}
	return Result(n[0]) << 24 + Result(n[1]) << 16 + Result(n[2]) << 8 + Result(n[3]) , nil;
}

func CalcResponseFromBytes(in []byte) (*CalculationResponse, error) {
	if len(in) != 7 {
		return nil, errors.New("Calc response needs 7 bytes!")
	}

	result, err := bytes_toResult(in[3:7])
	if err != nil {
		return nil, err
	}

	return &CalculationResponse{
		TML: TML(in[0]),
		RequestID: RequestID(in[1]),
		ErrorCode: ErrorCode(in[2]),
		Result: result,
	}, nil
}

func (resp *CalculationResponse) ToBytes() []byte {
	out := make([]byte, 0, int(resp.TML))
	out = append(out, byte(resp.TML))
	out = append(out, byte(resp.RequestID))
	out = append(out, byte(resp.ErrorCode))
	out = append(out, result_toBytes(resp.Result)...)
	log.Trace("Request bytes: %+v", out)
	return out
}

func BuildResponse(reqID RequestID, result Result) *CalculationResponse {
	return &CalculationResponse{
		TML: ResponseTML,
		RequestID: reqID,
		ErrorCode: ErrorCodeNoError,
		Result: result,
	}	
}

func ErrResponse(reqID RequestID) *CalculationResponse {
	return &CalculationResponse {
		TML: ResponseTML,
		RequestID: reqID,
		ErrorCode: ErrorCodeError,
	}
}
