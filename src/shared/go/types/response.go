package types 

import (
	"errors"

	"shared/go/log"
)

type Result int32

type CalculationResponse struct {
	TML byte
	RequestID byte
	ErrorCode byte
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
		TML: in[0],
		RequestID: in[1],
		ErrorCode: in[2],
		Result: result,
	}, nil
}

func (req *CalculationResponse) ToBytes() []byte {
	out := make([]byte, 0, int(req.TML))
	out = append(out, req.TML)
	out = append(out, req.RequestID)
	out = append(out, req.ErrorCode)
	out = append(out, result_toBytes(req.Result)...)
	log.Trace("Request bytes: %+v", out)
	return out
}

func BuildResponse(reqID byte, result Result) *CalculationResponse {
	return &CalculationResponse{
		TML: 7,
		RequestID: reqID,
		ErrorCode: 0,
		Result: result,
	}	
}
