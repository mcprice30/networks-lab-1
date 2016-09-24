package types 

import (
	"errors"

	"shared/go/log"
)

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


func CalcResponseFromBytes(in []byte) (*CalculationResponse, error) {
	if len(in) != 7 {
		return nil, errors.New("Calc response needs 7 bytes!")
	}

	result := bytesToResult(in[3:7])

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
	out = append(out, resp.Result.toBytes()...)
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
