// Code generated by protoc-gen-go.
// source: steammessages_remoteclient.proto
// DO NOT EDIT!

package internal

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import steammessages_base "steammessages_base.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CMsgRemoteClientBroadcastHeader struct {
	ClientId         *uint64 `protobuf:"varint,1,opt,name=client_id" json:"client_id,omitempty"`
	MsgType          *uint32 `protobuf:"varint,2,opt,name=msg_type" json:"msg_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgRemoteClientBroadcastHeader) Reset()         { *m = CMsgRemoteClientBroadcastHeader{} }
func (m *CMsgRemoteClientBroadcastHeader) String() string { return proto.CompactTextString(m) }
func (*CMsgRemoteClientBroadcastHeader) ProtoMessage()    {}

func (m *CMsgRemoteClientBroadcastHeader) GetClientId() uint64 {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return 0
}

func (m *CMsgRemoteClientBroadcastHeader) GetMsgType() uint32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

type CMsgRemoteClientBroadcastStatus struct {
	Version          *int32  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	MinVersion       *int32  `protobuf:"varint,2,opt,name=min_version" json:"min_version,omitempty"`
	ConnectPort      *uint32 `protobuf:"varint,3,opt,name=connect_port" json:"connect_port,omitempty"`
	Hostname         *string `protobuf:"bytes,4,opt,name=hostname" json:"hostname,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgRemoteClientBroadcastStatus) Reset()         { *m = CMsgRemoteClientBroadcastStatus{} }
func (m *CMsgRemoteClientBroadcastStatus) String() string { return proto.CompactTextString(m) }
func (*CMsgRemoteClientBroadcastStatus) ProtoMessage()    {}

func (m *CMsgRemoteClientBroadcastStatus) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *CMsgRemoteClientBroadcastStatus) GetMinVersion() int32 {
	if m != nil && m.MinVersion != nil {
		return *m.MinVersion
	}
	return 0
}

func (m *CMsgRemoteClientBroadcastStatus) GetConnectPort() uint32 {
	if m != nil && m.ConnectPort != nil {
		return *m.ConnectPort
	}
	return 0
}

func (m *CMsgRemoteClientBroadcastStatus) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

type CMsgRemoteClientAppStatus struct {
	StatusUpdates    []*CMsgRemoteClientAppStatus_AppStatus `protobuf:"bytes,1,rep,name=status_updates" json:"status_updates,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *CMsgRemoteClientAppStatus) Reset()         { *m = CMsgRemoteClientAppStatus{} }
func (m *CMsgRemoteClientAppStatus) String() string { return proto.CompactTextString(m) }
func (*CMsgRemoteClientAppStatus) ProtoMessage()    {}

func (m *CMsgRemoteClientAppStatus) GetStatusUpdates() []*CMsgRemoteClientAppStatus_AppStatus {
	if m != nil {
		return m.StatusUpdates
	}
	return nil
}

type CMsgRemoteClientAppStatus_AppStatus struct {
	AppId            *uint32 `protobuf:"varint,1,opt,name=app_id" json:"app_id,omitempty"`
	AppState         *uint32 `protobuf:"varint,2,opt,name=app_state" json:"app_state,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgRemoteClientAppStatus_AppStatus) Reset()         { *m = CMsgRemoteClientAppStatus_AppStatus{} }
func (m *CMsgRemoteClientAppStatus_AppStatus) String() string { return proto.CompactTextString(m) }
func (*CMsgRemoteClientAppStatus_AppStatus) ProtoMessage()    {}

func (m *CMsgRemoteClientAppStatus_AppStatus) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *CMsgRemoteClientAppStatus_AppStatus) GetAppState() uint32 {
	if m != nil && m.AppState != nil {
		return *m.AppState
	}
	return 0
}

type CMsgRemoteClientAuth struct {
	ClientId         *uint64                          `protobuf:"varint,1,opt,name=client_id" json:"client_id,omitempty"`
	Status           *CMsgRemoteClientBroadcastStatus `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *CMsgRemoteClientAuth) Reset()         { *m = CMsgRemoteClientAuth{} }
func (m *CMsgRemoteClientAuth) String() string { return proto.CompactTextString(m) }
func (*CMsgRemoteClientAuth) ProtoMessage()    {}

func (m *CMsgRemoteClientAuth) GetClientId() uint64 {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return 0
}

func (m *CMsgRemoteClientAuth) GetStatus() *CMsgRemoteClientBroadcastStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type CMsgRemoteClientAuthResponse struct {
	Eresult          *int32 `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMsgRemoteClientAuthResponse) Reset()         { *m = CMsgRemoteClientAuthResponse{} }
func (m *CMsgRemoteClientAuthResponse) String() string { return proto.CompactTextString(m) }
func (*CMsgRemoteClientAuthResponse) ProtoMessage()    {}

const Default_CMsgRemoteClientAuthResponse_Eresult int32 = 2

func (m *CMsgRemoteClientAuthResponse) GetEresult() int32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgRemoteClientAuthResponse_Eresult
}

func init() {
}
