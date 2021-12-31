package header

import (
	"bytes"
	"fmt"

	"github.com/gozh-io/sip/sip/primitive"
	"github.com/gozh-io/sip/util"
)

type Uri interface {
	Equals(other interface{}) bool
	String() string
	Clone() Uri
	IsEncrypted() bool
	SetEncrypted(flag bool)
	GetUser() primitive.Stringer
	SetUser(user primitive.Stringer)
	GetPassword() primitive.Stringer
	SetPassword(pwd primitive.Stringer)
	GetHost() string
	SetHost(host string)
	GetPort() *primitive.Port
	SetPort(port *primitive.Port)
	GetUriParams() Params
	SetUriParams(params Params)
	GetHeaders() Params
	SetHeaders(params Params)
	IsWildCard() bool
}

type ContactUri interface {
	Uri
}

type SipUri struct {
	Encrypted bool
	User      primitive.Stringer
	Password  primitive.Stringer
	Host      string
	Port      *primitive.Port
	UriParams Params
	Headers   Params
}

func NewSipUri(tls bool, user, pwd primitive.Stringer, host string, port *primitive.Port, params, header Params) *SipUri {
	return &SipUri{
		Encrypted: tls,
		User:      user,
		Password:  pwd,
		Host:      host,
		Port:      port,
		UriParams: params,
		Headers:   header,
	}
}

func (s *SipUri) Equals(other interface{}) bool {
	otherPtr, ok := other.(*SipUri)
	if !ok {
		return false
	}

	if s == otherPtr {
		return true
	}

	if s == nil && otherPtr != nil || s != nil && otherPtr == nil {
		return false
	}

	val := *otherPtr

	result := s.Encrypted == val.Encrypted &&
		s.User == val.User &&
		s.Password == val.Password &&
		s.Host == val.Host &&
		util.Uint16PtrEq((*uint16)(s.Port), (*uint16)(val.Port))

	if !result {
		return false
	}

	if s.UriParams != otherPtr.UriParams {
		if s.UriParams == nil {
			result = result && otherPtr.UriParams != nil
		} else {
			result = result && s.UriParams.Equals(otherPtr.UriParams)
		}
	}

	if s.Headers != otherPtr.Headers {
		if s.Headers == nil {
			result = result && otherPtr.Headers != nil
		} else {
			result = result && s.Headers.Equals(otherPtr.Headers)
		}
	}

	return result
}

func (s *SipUri) String() string {
	var buffer bytes.Buffer

	if s.Encrypted {
		buffer.WriteString("sips")
	} else {
		buffer.WriteString("sip")
	}
	buffer.WriteString(":")

	if user, ok := s.User.(*primitive.String); ok && user.String() != "" {
		buffer.WriteString(s.User.String())
		if pwd, ok := s.Password.(*primitive.String); ok && pwd.String() != "" {
			buffer.WriteString(":")
			buffer.WriteString(s.Password.String())
		}
		buffer.WriteString("@")
	}

	// hostname
	buffer.WriteString(s.Host)

	// optional port number
	if s.Port != nil {
		buffer.WriteString(fmt.Sprintf(":%d", s.Port))
	}

	if s.UriParams != nil && s.UriParams.Length() > 0 {
		buffer.WriteString(";")
		buffer.WriteString(s.UriParams.String())
	}

	if s.Headers != nil && s.Headers.Length() > 0 {
		buffer.WriteString("?")
		buffer.WriteString(s.Headers.String())
	}

	return buffer.String()
}

func (s *SipUri) Clone() Uri {
	var newUri *SipUri

	if s == nil {
		return newUri
	}

	newUri = &SipUri{
		Encrypted: s.Encrypted,
		User:      s.User,
		Password:  s.Password,
		Host:      s.Host,
		UriParams: cloneWithNil(s.UriParams),
		Headers:   cloneWithNil(s.Headers),
	}

	if s.Port != nil {
		newUri.Port = s.Port
	}

	return newUri
}

func (s *SipUri) IsEncrypted() bool {
	return s.Encrypted
}

func (s *SipUri) SetEncrypted(flag bool) {
	s.Encrypted = flag
}

func (s *SipUri) GetUser() primitive.Stringer {
	return s.User
}

func (s *SipUri) SetUser(user primitive.Stringer) {
	s.User = user
}

func (s *SipUri) GetPassword() primitive.Stringer {
	return s.Password
}

func (s *SipUri) SetPassword(pwd primitive.Stringer) {
	s.Password = pwd
}

func (s *SipUri) GetHost() string {
	return s.Host
}

func (s *SipUri) SetHost(host string) {
	s.Host = host
}

func (s *SipUri) GetPort() *primitive.Port {
	return s.Port
}

func (s *SipUri) SetPort(port *primitive.Port) {
	s.Port = port
}

func (s *SipUri) GetUriParams() Params {
	return s.UriParams
}

func (s *SipUri) SetUriParams(params Params) {
	s.UriParams = params
}

func (s *SipUri) GetHeaders() Params {
	return s.Headers
}

func (s *SipUri) SetHeaders(params Params) {
	s.Headers = params
}

func (s *SipUri) IsWildCard() bool {
	return false
}
