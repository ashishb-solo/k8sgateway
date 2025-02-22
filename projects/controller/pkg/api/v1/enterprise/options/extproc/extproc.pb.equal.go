// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/enterprise/options/extproc/extproc.proto

package extproc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *Settings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGrpcService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGrpcService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGrpcService(), target.GetGrpcService()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFilterStage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFilterStage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFilterStage(), target.GetFilterStage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFailureModeAllow()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFailureModeAllow()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFailureModeAllow(), target.GetFailureModeAllow()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetProcessingMode()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProcessingMode()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProcessingMode(), target.GetProcessingMode()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAsyncMode()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAsyncMode()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAsyncMode(), target.GetAsyncMode()) {
			return false
		}
	}

	if len(m.GetRequestAttributes()) != len(target.GetRequestAttributes()) {
		return false
	}
	for idx, v := range m.GetRequestAttributes() {

		if strings.Compare(v, target.GetRequestAttributes()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetResponseAttributes()) != len(target.GetResponseAttributes()) {
		return false
	}
	for idx, v := range m.GetResponseAttributes() {

		if strings.Compare(v, target.GetResponseAttributes()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetMessageTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMessageTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMessageTimeout(), target.GetMessageTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatPrefix()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatPrefix()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatPrefix(), target.GetStatPrefix()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMutationRules()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMutationRules()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMutationRules(), target.GetMutationRules()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMaxMessageTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxMessageTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxMessageTimeout(), target.GetMaxMessageTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDisableClearRouteCache()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDisableClearRouteCache()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDisableClearRouteCache(), target.GetDisableClearRouteCache()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetForwardRules()).(equality.Equalizer); ok {
		if !h.Equal(target.GetForwardRules()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetForwardRules(), target.GetForwardRules()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFilterMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFilterMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFilterMetadata(), target.GetFilterMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAllowModeOverride()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAllowModeOverride()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAllowModeOverride(), target.GetAllowModeOverride()) {
			return false
		}
	}

	if len(m.GetMetadataContextNamespaces()) != len(target.GetMetadataContextNamespaces()) {
		return false
	}
	for idx, v := range m.GetMetadataContextNamespaces() {

		if strings.Compare(v, target.GetMetadataContextNamespaces()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetTypedMetadataContextNamespaces()) != len(target.GetTypedMetadataContextNamespaces()) {
		return false
	}
	for idx, v := range m.GetTypedMetadataContextNamespaces() {

		if strings.Compare(v, target.GetTypedMetadataContextNamespaces()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *RouteSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteSettings)
	if !ok {
		that2, ok := that.(RouteSettings)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Override.(type) {

	case *RouteSettings_Disabled:
		if _, ok := target.Override.(*RouteSettings_Disabled); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDisabled()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDisabled()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDisabled(), target.GetDisabled()) {
				return false
			}
		}

	case *RouteSettings_Overrides:
		if _, ok := target.Override.(*RouteSettings_Overrides); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOverrides()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOverrides()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOverrides(), target.GetOverrides()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Override != target.Override {
			return false
		}
	}

	return true
}

// Equal function
func (m *GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService)
	if !ok {
		that2, ok := that.(GrpcService)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetExtProcServerRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtProcServerRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtProcServerRef(), target.GetExtProcServerRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAuthority()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAuthority()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAuthority(), target.GetAuthority()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRetryPolicy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRetryPolicy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRetryPolicy(), target.GetRetryPolicy()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTimeout(), target.GetTimeout()) {
			return false
		}
	}

	if len(m.GetInitialMetadata()) != len(target.GetInitialMetadata()) {
		return false
	}
	for idx, v := range m.GetInitialMetadata() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetInitialMetadata()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetInitialMetadata()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Overrides) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Overrides)
	if !ok {
		that2, ok := that.(Overrides)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetProcessingMode()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProcessingMode()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProcessingMode(), target.GetProcessingMode()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAsyncMode()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAsyncMode()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAsyncMode(), target.GetAsyncMode()) {
			return false
		}
	}

	if len(m.GetRequestAttributes()) != len(target.GetRequestAttributes()) {
		return false
	}
	for idx, v := range m.GetRequestAttributes() {

		if strings.Compare(v, target.GetRequestAttributes()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetResponseAttributes()) != len(target.GetResponseAttributes()) {
		return false
	}
	for idx, v := range m.GetResponseAttributes() {

		if strings.Compare(v, target.GetResponseAttributes()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetGrpcService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGrpcService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGrpcService(), target.GetGrpcService()) {
			return false
		}
	}

	if len(m.GetMetadataContextNamespaces()) != len(target.GetMetadataContextNamespaces()) {
		return false
	}
	for idx, v := range m.GetMetadataContextNamespaces() {

		if strings.Compare(v, target.GetMetadataContextNamespaces()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetTypedMetadataContextNamespaces()) != len(target.GetTypedMetadataContextNamespaces()) {
		return false
	}
	for idx, v := range m.GetTypedMetadataContextNamespaces() {

		if strings.Compare(v, target.GetTypedMetadataContextNamespaces()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *HeaderForwardingRules) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderForwardingRules)
	if !ok {
		that2, ok := that.(HeaderForwardingRules)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetAllowedHeaders()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAllowedHeaders()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAllowedHeaders(), target.GetAllowedHeaders()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDisallowedHeaders()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDisallowedHeaders()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDisallowedHeaders(), target.GetDisallowedHeaders()) {
			return false
		}
	}

	return true
}
