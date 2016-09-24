// Lab 1 - Group 11
// 
// The types package contains custom structs for that encapsulate the packets
// to be sent, including the various fields in said packets.
//
// response.go defines the response type, which encapsulate the packets sent
// by the server containing the result of a requested calculation. 
package types 

import (
	"fmt"

	"shared/go/log"
)

const (
	ErrorCodeError ErrorCode = 127
	ErrorCodeNoError ErrorCode = 0
	ResponseTML	TML = 7
)

// CalculationResponse holds the data for the packet sent from the server
// containing the response to a given request.
type CalculationResponse struct {
	TML       TML       // Total Message Length (in bytes).
	RequestID RequestID // Unique identifier for this request.
	ErrorCode ErrorCode // Indicates whether an error occurred.
	Result 	  Result    // The numerical result of the requested calculation.
}

// CalcResponseFromBytes takes a byte array and returns a pointer to the
// calculation response object stored in those bytes.
func CalcResponseFromBytes(in []byte) (*CalculationResponse, error) {
	if len(in) != int(ResponseTML) {
		return nil, fmt.Errorf("Calc response needs %d bytes!", int(ResponseTML))
	}

	result := bytesToResult(in[3:7])

	return &CalculationResponse{
		TML: TML(in[0]),
		RequestID: RequestID(in[1]),
		ErrorCode: ErrorCode(in[2]),
		Result: result,
	}, nil
}

// ToBytes produces the network byte order representation of the given
// calculation response object.
func (resp *CalculationResponse) ToBytes() []byte {
	out := make([]byte, 0, int(resp.TML))
	out = append(out, byte(resp.TML))
	out = append(out, byte(resp.RequestID))
	out = append(out, byte(resp.ErrorCode))
	out = append(out, resp.Result.toBytes()...)
	log.Trace("Request bytes: %+v", out)
	return out
}

// BuildResponse takes a request id and the results of a calculation and returns
// the appropriate calculation response.
func BuildResponse(reqID RequestID, result Result) *CalculationResponse {
	return &CalculationResponse{
		TML: ResponseTML,
		RequestID: reqID,
		ErrorCode: ErrorCodeNoError,
		Result: result,
	}	
}

// ErrResponse takes a request id and returns the a calculation response that
// indicates the given request was unsuccessful.
func ErrResponse(reqID RequestID) *CalculationResponse {
	return &CalculationResponse {
		TML: ResponseTML,
		RequestID: reqID,
		ErrorCode: ErrorCodeError,
	}
}
