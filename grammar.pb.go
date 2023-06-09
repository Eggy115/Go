// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/grammar.proto

/*
Package grammar is a generated protocol buffer package.

It is generated from these files:
	proto/grammar.proto

It has these top-level messages:
	Rule
	Library
*/
package grammar

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Rule struct {
	Name                string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ScopeName           string           `protobuf:"bytes,2,opt,name=scopeName" json:"scopeName,omitempty"`
	ContentName         string           `protobuf:"bytes,3,opt,name=contentName" json:"contentName,omitempty"`
	Match               string           `protobuf:"bytes,4,opt,name=match" json:"match,omitempty"`
	Begin               string           `protobuf:"bytes,5,opt,name=begin" json:"begin,omitempty"`
	While               string           `protobuf:"bytes,6,opt,name=while" json:"while,omitempty"`
	End                 string           `protobuf:"bytes,7,opt,name=end" json:"end,omitempty"`
	Include             string           `protobuf:"bytes,8,opt,name=include" json:"include,omitempty"`
	Patterns            []*Rule          `protobuf:"bytes,9,rep,name=patterns" json:"patterns,omitempty"`
	Captures            map[string]*Rule `protobuf:"bytes,10,rep,name=captures" json:"captures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BeginCaptures       map[string]*Rule `protobuf:"bytes,11,rep,name=beginCaptures" json:"beginCaptures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	WhileCaptures       map[string]*Rule `protobuf:"bytes,12,rep,name=whileCaptures" json:"whileCaptures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	EndCaptures         map[string]*Rule `protobuf:"bytes,13,rep,name=endCaptures" json:"endCaptures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Repository          map[string]*Rule `protobuf:"bytes,14,rep,name=repository" json:"repository,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Injections          map[string]*Rule `protobuf:"bytes,15,rep,name=injections" json:"injections,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Disabled            bool             `protobuf:"varint,16,opt,name=disabled" json:"disabled,omitempty"`
	ApplyEndPatternLast bool             `protobuf:"varint,17,opt,name=applyEndPatternLast" json:"applyEndPatternLast,omitempty"`
	IncludeResetBase    bool             `protobuf:"varint,18,opt,name=includeResetBase" json:"includeResetBase,omitempty"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (m *Rule) String() string            { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetScopeName() string {
	if m != nil {
		return m.ScopeName
	}
	return ""
}

func (m *Rule) GetContentName() string {
	if m != nil {
		return m.ContentName
	}
	return ""
}

func (m *Rule) GetMatch() string {
	if m != nil {
		return m.Match
	}
	return ""
}

func (m *Rule) GetBegin() string {
	if m != nil {
		return m.Begin
	}
	return ""
}

func (m *Rule) GetWhile() string {
	if m != nil {
		return m.While
	}
	return ""
}

func (m *Rule) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *Rule) GetInclude() string {
	if m != nil {
		return m.Include
	}
	return ""
}

func (m *Rule) GetPatterns() []*Rule {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func (m *Rule) GetCaptures() map[string]*Rule {
	if m != nil {
		return m.Captures
	}
	return nil
}

func (m *Rule) GetBeginCaptures() map[string]*Rule {
	if m != nil {
		return m.BeginCaptures
	}
	return nil
}

func (m *Rule) GetWhileCaptures() map[string]*Rule {
	if m != nil {
		return m.WhileCaptures
	}
	return nil
}

func (m *Rule) GetEndCaptures() map[string]*Rule {
	if m != nil {
		return m.EndCaptures
	}
	return nil
}

func (m *Rule) GetRepository() map[string]*Rule {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *Rule) GetInjections() map[string]*Rule {
	if m != nil {
		return m.Injections
	}
	return nil
}

func (m *Rule) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Rule) GetApplyEndPatternLast() bool {
	if m != nil {
		return m.ApplyEndPatternLast
	}
	return false
}

func (m *Rule) GetIncludeResetBase() bool {
	if m != nil {
		return m.IncludeResetBase
	}
	return false
}

type Library struct {
	Grammars map[string]*Rule `protobuf:"bytes,1,rep,name=grammars" json:"grammars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Library) Reset()                    { *m = Library{} }
func (m *Library) String() string            { return proto.CompactTextString(m) }
func (*Library) ProtoMessage()               {}
func (*Library) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Library) GetGrammars() map[string]*Rule {
	if m != nil {
		return m.Grammars
	}
	return nil
}

func init() {
	proto.RegisterType((*Rule)(nil), "grammar.Rule")
	proto.RegisterType((*Library)(nil), "grammar.Library")
}

func init() { proto.RegisterFile("proto/grammar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x55, 0x66, 0xda, 0x69, 0x7a, 0x4b, 0x99, 0x72, 0x87, 0x85, 0x55, 0x1e, 0x8a, 0x86, 0x4d,
	0x61, 0x51, 0x10, 0x2c, 0x40, 0x23, 0x21, 0xa1, 0x41, 0x05, 0x81, 0xca, 0x43, 0xd9, 0xb0, 0x76,
	0x13, 0x6b, 0x26, 0x90, 0x3a, 0x91, 0xed, 0x82, 0xf2, 0x19, 0x7c, 0x19, 0xbf, 0x84, 0x7c, 0xed,
	0xa6, 0x49, 0xdb, 0x5d, 0x76, 0xbe, 0xe7, 0x25, 0x3b, 0x3e, 0x0e, 0x5c, 0x94, 0xaa, 0x30, 0xc5,
	0xf3, 0x1b, 0xc5, 0xd7, 0x6b, 0xae, 0xe6, 0x34, 0xe1, 0xc0, 0x8f, 0x97, 0xff, 0x86, 0xd0, 0x8b,
	0x37, 0xb9, 0x40, 0x84, 0x9e, 0xe4, 0x6b, 0xc1, 0x82, 0x28, 0x98, 0x0d, 0x63, 0x5a, 0xe3, 0x43,
	0x18, 0xea, 0xa4, 0x28, 0xc5, 0x57, 0x4b, 0x9c, 0x10, 0xb1, 0x03, 0x30, 0x82, 0x51, 0x52, 0x48,
	0x23, 0xa4, 0x21, 0xfe, 0x94, 0xf8, 0x26, 0x84, 0xf7, 0xa1, 0xbf, 0xe6, 0x26, 0xb9, 0x65, 0x3d,
	0xe2, 0xdc, 0x60, 0xd1, 0x95, 0xb8, 0xc9, 0x24, 0xeb, 0x3b, 0x94, 0x06, 0x8b, 0xfe, 0xb9, 0xcd,
	0x72, 0xc1, 0xce, 0x1c, 0x4a, 0x03, 0x4e, 0xe0, 0x54, 0xc8, 0x94, 0x0d, 0x08, 0xb3, 0x4b, 0x64,
	0x30, 0xc8, 0x64, 0x92, 0x6f, 0x52, 0xc1, 0x42, 0x42, 0xb7, 0x23, 0x3e, 0x85, 0xb0, 0xe4, 0xc6,
	0x08, 0x25, 0x35, 0x1b, 0x46, 0xa7, 0xb3, 0xd1, 0xcb, 0xf1, 0x7c, 0x7b, 0x6a, 0x7b, 0xc4, 0xb8,
	0xa6, 0xf1, 0x35, 0x84, 0x09, 0x2f, 0xcd, 0x46, 0x09, 0xcd, 0x80, 0xa4, 0x0f, 0x5a, 0xd2, 0xf9,
	0x7b, 0xcf, 0x2e, 0xa4, 0x51, 0x55, 0x5c, 0x8b, 0xf1, 0x03, 0x8c, 0x69, 0xbb, 0x5b, 0x9e, 0x8d,
	0xc8, 0x1d, 0xb5, 0xdd, 0xd7, 0x4d, 0x89, 0x8b, 0x68, 0xdb, 0x6c, 0x0e, 0x1d, 0xb0, 0xce, 0xb9,
	0x73, 0x2c, 0xe7, 0x47, 0x53, 0xe2, 0x73, 0x5a, 0x36, 0x7c, 0x07, 0x23, 0x21, 0xd3, 0x3a, 0x65,
	0x4c, 0x29, 0x8f, 0xdb, 0x29, 0x8b, 0x9d, 0xc0, 0x65, 0x34, 0x2d, 0xf8, 0x16, 0x40, 0x89, 0xb2,
	0xd0, 0x99, 0x29, 0x54, 0xc5, 0xee, 0x52, 0xc0, 0xa3, 0x76, 0x40, 0x5c, 0xf3, 0xce, 0xdf, 0x30,
	0x58, 0x7b, 0x26, 0x7f, 0x8a, 0xc4, 0x64, 0x85, 0xd4, 0xec, 0xfc, 0x98, 0xfd, 0x53, 0xcd, 0x7b,
	0xfb, 0xce, 0x80, 0x53, 0x08, 0xd3, 0x4c, 0xf3, 0x55, 0x2e, 0x52, 0x36, 0x89, 0x82, 0x59, 0x18,
	0xd7, 0x33, 0xbe, 0x80, 0x0b, 0x5e, 0x96, 0x79, 0xb5, 0x90, 0xe9, 0x77, 0x77, 0x71, 0x4b, 0xae,
	0x0d, 0xbb, 0x47, 0xb2, 0x63, 0x14, 0x3e, 0x83, 0x89, 0x2f, 0x43, 0x2c, 0xb4, 0x30, 0xd7, 0x5c,
	0x0b, 0x86, 0x24, 0x3f, 0xc0, 0xa7, 0x9f, 0x61, 0xdc, 0xfa, 0x2a, 0xb6, 0x6a, 0xbf, 0x44, 0xe5,
	0xfb, 0x6f, 0x97, 0xf8, 0x04, 0xfa, 0xbf, 0x79, 0xbe, 0x71, 0xd5, 0x3f, 0x68, 0x93, 0xe3, 0xae,
	0x4e, 0xde, 0x04, 0xd3, 0x6f, 0x80, 0x87, 0x57, 0xde, 0x31, 0xf0, 0xf0, 0xee, 0xbb, 0x04, 0x7e,
	0x81, 0xc9, 0x7e, 0x0d, 0xba, 0xc4, 0x2d, 0xe1, 0x7c, 0xaf, 0x14, 0x1d, 0xd3, 0xf6, 0x3a, 0xd2,
	0x21, 0xed, 0xf2, 0x6f, 0x00, 0x83, 0x65, 0xb6, 0x52, 0x5c, 0x55, 0x78, 0x05, 0xa1, 0x97, 0x69,
	0x16, 0xec, 0xbd, 0x0d, 0xaf, 0x99, 0x7f, 0xf4, 0x02, 0xff, 0xd4, 0xb7, 0x7a, 0x5b, 0x90, 0x16,
	0xd5, 0x61, 0x4f, 0xab, 0x33, 0xfa, 0xeb, 0xbe, 0xfa, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x2e,
	0xec, 0x55, 0x8c, 0x05, 0x00, 0x00,
}
