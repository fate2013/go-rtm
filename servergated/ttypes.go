// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package servergated

import (
	"bytes"
	"fmt"
	"fate2013/go-rtm/fp1225"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fate2013/go-rtm/rtmcommon"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = fp1225.GoUnusedProtection__
var _ = rtmcommon.GoUnusedProtection__
var GoUnusedProtection__ int

type ServerGatedException struct {
	Code   int32  `thrift:"code,1" json:"code"`
	Reason string `thrift:"reason,2" json:"reason"`
}

func NewServerGatedException() *ServerGatedException {
	return &ServerGatedException{}
}

func (p *ServerGatedException) GetCode() int32 {
	return p.Code
}

func (p *ServerGatedException) GetReason() string {
	return p.Reason
}
func (p *ServerGatedException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ServerGatedException) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Code = v
	}
	return nil
}

func (p *ServerGatedException) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Reason = v
	}
	return nil
}

func (p *ServerGatedException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("serverGatedException"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ServerGatedException) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:code: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Code)); err != nil {
		return fmt.Errorf("%T.code (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:code: %s", p, err)
	}
	return err
}

func (p *ServerGatedException) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("reason", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:reason: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Reason)); err != nil {
		return fmt.Errorf("%T.reason (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:reason: %s", p, err)
	}
	return err
}

func (p *ServerGatedException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ServerGatedException(%+v)", *p)
}

func (p *ServerGatedException) Error() string {
	return p.String()
}

type Token struct {
	ProjectId int32  `thrift:"project_id,1" json:"project_id"`
	AuthToken string `thrift:"auth_token,2" json:"auth_token"`
}

func NewToken() *Token {
	return &Token{}
}

func (p *Token) GetProjectId() int32 {
	return p.ProjectId
}

func (p *Token) GetAuthToken() string {
	return p.AuthToken
}
func (p *Token) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Token) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.ProjectId = v
	}
	return nil
}

func (p *Token) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.AuthToken = v
	}
	return nil
}

func (p *Token) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Token"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Token) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("project_id", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:project_id: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.ProjectId)); err != nil {
		return fmt.Errorf("%T.project_id (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:project_id: %s", p, err)
	}
	return err
}

func (p *Token) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("auth_token", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:auth_token: %s", p, err)
	}
	if err := oprot.WriteString(string(p.AuthToken)); err != nil {
		return fmt.Errorf("%T.auth_token (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:auth_token: %s", p, err)
	}
	return err
}

func (p *Token) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Token(%+v)", *p)
}