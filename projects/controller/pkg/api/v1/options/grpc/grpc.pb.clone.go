// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/grpc/grpc.proto

package grpc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/transformation"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *ServiceSpec) Clone() proto.Message {
	var target *ServiceSpec
	if m == nil {
		return target
	}
	target = &ServiceSpec{}

	if m.GetDescriptors() != nil {
		target.Descriptors = make([]byte, len(m.GetDescriptors()))
		copy(target.Descriptors, m.GetDescriptors())
	}

	if m.GetGrpcServices() != nil {
		target.GrpcServices = make([]*ServiceSpec_GrpcService, len(m.GetGrpcServices()))
		for idx, v := range m.GetGrpcServices() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.GrpcServices[idx] = h.Clone().(*ServiceSpec_GrpcService)
			} else {
				target.GrpcServices[idx] = proto.Clone(v).(*ServiceSpec_GrpcService)
			}

		}
	}

	return target
}

// Clone function
func (m *DestinationSpec) Clone() proto.Message {
	var target *DestinationSpec
	if m == nil {
		return target
	}
	target = &DestinationSpec{}

	target.Package = m.GetPackage()

	target.Service = m.GetService()

	target.Function = m.GetFunction()

	if h, ok := interface{}(m.GetParameters()).(clone.Cloner); ok {
		target.Parameters = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Parameters)
	} else {
		target.Parameters = proto.Clone(m.GetParameters()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Parameters)
	}

	return target
}

// Clone function
func (m *ServiceSpec_GrpcService) Clone() proto.Message {
	var target *ServiceSpec_GrpcService
	if m == nil {
		return target
	}
	target = &ServiceSpec_GrpcService{}

	target.PackageName = m.GetPackageName()

	target.ServiceName = m.GetServiceName()

	if m.GetFunctionNames() != nil {
		target.FunctionNames = make([]string, len(m.GetFunctionNames()))
		for idx, v := range m.GetFunctionNames() {

			target.FunctionNames[idx] = v

		}
	}

	return target
}
