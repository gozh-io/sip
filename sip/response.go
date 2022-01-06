package sip

import "github.com/gozh-io/sip/sip/primitive"

type Response interface {
	Message
	StatusCode() primitive.StatusCode
	SetStatusCode(code primitive.StatusCode)
	Reason() string
	SetReason(reason string)
	// Previous returns previous provisional responses
	Previous() []Response
	SetPrevious(responses []Response)
	/* Common helpers */
	IsProvisional() bool
	IsSuccess() bool
	IsRedirection() bool
	IsClientError() bool
	IsServerError() bool
	IsGlobalError() bool
}
