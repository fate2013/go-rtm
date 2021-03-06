// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rtmcommon

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type FriendPair struct {
	UidA int64 `thrift:"uidA,1" json:"uidA"`
	UidB int64 `thrift:"uidB,2" json:"uidB"`
}

func NewFriendPair() *FriendPair {
	return &FriendPair{}
}

func (p *FriendPair) GetUidA() int64 {
	return p.UidA
}

func (p *FriendPair) GetUidB() int64 {
	return p.UidB
}
func (p *FriendPair) Read(iprot thrift.TProtocol) error {
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

func (p *FriendPair) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.UidA = v
	}
	return nil
}

func (p *FriendPair) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.UidB = v
	}
	return nil
}

func (p *FriendPair) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("friendPair"); err != nil {
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

func (p *FriendPair) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("uidA", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:uidA: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.UidA)); err != nil {
		return fmt.Errorf("%T.uidA (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:uidA: %s", p, err)
	}
	return err
}

func (p *FriendPair) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("uidB", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:uidB: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.UidB)); err != nil {
		return fmt.Errorf("%T.uidB (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:uidB: %s", p, err)
	}
	return err
}

func (p *FriendPair) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FriendPair(%+v)", *p)
}

type RTMException struct {
	Code    int32  `thrift:"code,1" json:"code"`
	Message string `thrift:"message,2" json:"message"`
}

func NewRTMException() *RTMException {
	return &RTMException{}
}

func (p *RTMException) GetCode() int32 {
	return p.Code
}

func (p *RTMException) GetMessage() string {
	return p.Message
}
func (p *RTMException) Read(iprot thrift.TProtocol) error {
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

func (p *RTMException) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Code = v
	}
	return nil
}

func (p *RTMException) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *RTMException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RTMException"); err != nil {
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

func (p *RTMException) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *RTMException) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:message: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return fmt.Errorf("%T.message (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:message: %s", p, err)
	}
	return err
}

func (p *RTMException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RTMException(%+v)", *p)
}

func (p *RTMException) Error() string {
	return p.String()
}

type MsgNum struct {
	MsgType int8  `thrift:"msg_type,1" json:"msg_type"`
	FromXid int64 `thrift:"from_xid,2" json:"from_xid"`
	Num     int32 `thrift:"num,3" json:"num"`
}

func NewMsgNum() *MsgNum {
	return &MsgNum{}
}

func (p *MsgNum) GetMsgType() int8 {
	return p.MsgType
}

func (p *MsgNum) GetFromXid() int64 {
	return p.FromXid
}

func (p *MsgNum) GetNum() int32 {
	return p.Num
}
func (p *MsgNum) Read(iprot thrift.TProtocol) error {
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
		case 3:
			if err := p.ReadField3(iprot); err != nil {
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

func (p *MsgNum) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadByte(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		temp := int8(v)
		p.MsgType = temp
	}
	return nil
}

func (p *MsgNum) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.FromXid = v
	}
	return nil
}

func (p *MsgNum) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Num = v
	}
	return nil
}

func (p *MsgNum) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("MsgNum"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
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

func (p *MsgNum) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msg_type", thrift.BYTE, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:msg_type: %s", p, err)
	}
	if err := oprot.WriteByte(byte(p.MsgType)); err != nil {
		return fmt.Errorf("%T.msg_type (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:msg_type: %s", p, err)
	}
	return err
}

func (p *MsgNum) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("from_xid", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:from_xid: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.FromXid)); err != nil {
		return fmt.Errorf("%T.from_xid (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:from_xid: %s", p, err)
	}
	return err
}

func (p *MsgNum) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:num: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num)); err != nil {
		return fmt.Errorf("%T.num (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:num: %s", p, err)
	}
	return err
}

func (p *MsgNum) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MsgNum(%+v)", *p)
}

type MsgParam struct {
	MsgType int8  `thrift:"msg_type,1" json:"msg_type"`
	FromXid int64 `thrift:"from_xid,2" json:"from_xid"`
	Num     int32 `thrift:"num,3" json:"num"`
	Offset  int64 `thrift:"offset,4" json:"offset"`
}

func NewMsgParam() *MsgParam {
	return &MsgParam{}
}

func (p *MsgParam) GetMsgType() int8 {
	return p.MsgType
}

func (p *MsgParam) GetFromXid() int64 {
	return p.FromXid
}

func (p *MsgParam) GetNum() int32 {
	return p.Num
}

func (p *MsgParam) GetOffset() int64 {
	return p.Offset
}
func (p *MsgParam) Read(iprot thrift.TProtocol) error {
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
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
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

func (p *MsgParam) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadByte(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		temp := int8(v)
		p.MsgType = temp
	}
	return nil
}

func (p *MsgParam) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.FromXid = v
	}
	return nil
}

func (p *MsgParam) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Num = v
	}
	return nil
}

func (p *MsgParam) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Offset = v
	}
	return nil
}

func (p *MsgParam) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("MsgParam"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
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

func (p *MsgParam) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msg_type", thrift.BYTE, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:msg_type: %s", p, err)
	}
	if err := oprot.WriteByte(byte(p.MsgType)); err != nil {
		return fmt.Errorf("%T.msg_type (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:msg_type: %s", p, err)
	}
	return err
}

func (p *MsgParam) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("from_xid", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:from_xid: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.FromXid)); err != nil {
		return fmt.Errorf("%T.from_xid (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:from_xid: %s", p, err)
	}
	return err
}

func (p *MsgParam) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:num: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num)); err != nil {
		return fmt.Errorf("%T.num (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:num: %s", p, err)
	}
	return err
}

func (p *MsgParam) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("offset", thrift.I64, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:offset: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Offset)); err != nil {
		return fmt.Errorf("%T.offset (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:offset: %s", p, err)
	}
	return err
}

func (p *MsgParam) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MsgParam(%+v)", *p)
}

type MsgContent struct {
	Body  string `thrift:"body,1" json:"body"`
	Mtime int32  `thrift:"mtime,2" json:"mtime"`
	Mtype int8   `thrift:"mtype,3" json:"mtype"`
	Uid   int64  `thrift:"uid,4" json:"uid"`
}

func NewMsgContent() *MsgContent {
	return &MsgContent{}
}

func (p *MsgContent) GetBody() string {
	return p.Body
}

func (p *MsgContent) GetMtime() int32 {
	return p.Mtime
}

func (p *MsgContent) GetMtype() int8 {
	return p.Mtype
}

func (p *MsgContent) GetUid() int64 {
	return p.Uid
}
func (p *MsgContent) Read(iprot thrift.TProtocol) error {
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
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
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

func (p *MsgContent) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Body = v
	}
	return nil
}

func (p *MsgContent) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Mtime = v
	}
	return nil
}

func (p *MsgContent) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadByte(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		temp := int8(v)
		p.Mtype = temp
	}
	return nil
}

func (p *MsgContent) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Uid = v
	}
	return nil
}

func (p *MsgContent) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("MsgContent"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
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

func (p *MsgContent) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("body", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:body: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Body)); err != nil {
		return fmt.Errorf("%T.body (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:body: %s", p, err)
	}
	return err
}

func (p *MsgContent) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mtime", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:mtime: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Mtime)); err != nil {
		return fmt.Errorf("%T.mtime (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:mtime: %s", p, err)
	}
	return err
}

func (p *MsgContent) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mtype", thrift.BYTE, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:mtype: %s", p, err)
	}
	if err := oprot.WriteByte(byte(p.Mtype)); err != nil {
		return fmt.Errorf("%T.mtype (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:mtype: %s", p, err)
	}
	return err
}

func (p *MsgContent) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("uid", thrift.I64, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:uid: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Uid)); err != nil {
		return fmt.Errorf("%T.uid (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:uid: %s", p, err)
	}
	return err
}

func (p *MsgContent) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MsgContent(%+v)", *p)
}

type MsgResult_ struct {
	MsgType int8          `thrift:"msg_type,1" json:"msg_type"`
	FromXid int64         `thrift:"from_xid,2" json:"from_xid"`
	Num     int32         `thrift:"num,3" json:"num"`
	Offset  int64         `thrift:"offset,4" json:"offset"`
	Msgs    []*MsgContent `thrift:"msgs,5" json:"msgs"`
}

func NewMsgResult_() *MsgResult_ {
	return &MsgResult_{}
}

func (p *MsgResult_) GetMsgType() int8 {
	return p.MsgType
}

func (p *MsgResult_) GetFromXid() int64 {
	return p.FromXid
}

func (p *MsgResult_) GetNum() int32 {
	return p.Num
}

func (p *MsgResult_) GetOffset() int64 {
	return p.Offset
}

func (p *MsgResult_) GetMsgs() []*MsgContent {
	return p.Msgs
}
func (p *MsgResult_) Read(iprot thrift.TProtocol) error {
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
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
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

func (p *MsgResult_) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadByte(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		temp := int8(v)
		p.MsgType = temp
	}
	return nil
}

func (p *MsgResult_) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.FromXid = v
	}
	return nil
}

func (p *MsgResult_) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Num = v
	}
	return nil
}

func (p *MsgResult_) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Offset = v
	}
	return nil
}

func (p *MsgResult_) ReadField5(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*MsgContent, 0, size)
	p.Msgs = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &MsgContent{}
		if err := _elem0.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem0, err)
		}
		p.Msgs = append(p.Msgs, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *MsgResult_) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("MsgResult"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
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

func (p *MsgResult_) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msg_type", thrift.BYTE, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:msg_type: %s", p, err)
	}
	if err := oprot.WriteByte(byte(p.MsgType)); err != nil {
		return fmt.Errorf("%T.msg_type (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:msg_type: %s", p, err)
	}
	return err
}

func (p *MsgResult_) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("from_xid", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:from_xid: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.FromXid)); err != nil {
		return fmt.Errorf("%T.from_xid (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:from_xid: %s", p, err)
	}
	return err
}

func (p *MsgResult_) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:num: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num)); err != nil {
		return fmt.Errorf("%T.num (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:num: %s", p, err)
	}
	return err
}

func (p *MsgResult_) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("offset", thrift.I64, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:offset: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Offset)); err != nil {
		return fmt.Errorf("%T.offset (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:offset: %s", p, err)
	}
	return err
}

func (p *MsgResult_) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msgs", thrift.LIST, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:msgs: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Msgs)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.Msgs {
		if err := v.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", v, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:msgs: %s", p, err)
	}
	return err
}

func (p *MsgResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MsgResult_(%+v)", *p)
}
