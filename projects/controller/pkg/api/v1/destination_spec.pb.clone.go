// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/destination_spec.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/aws"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_azure "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/azure"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/grpc"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_rest "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/rest"
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
func (m *DestinationSpec) Clone() proto.Message {
	var target *DestinationSpec
	if m == nil {
		return target
	}
	target = &DestinationSpec{}

	switch m.DestinationType.(type) {

	case *DestinationSpec_Aws:

		if h, ok := interface{}(m.GetAws()).(clone.Cloner); ok {
			target.DestinationType = &DestinationSpec_Aws{
				Aws: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws.DestinationSpec),
			}
		} else {
			target.DestinationType = &DestinationSpec_Aws{
				Aws: proto.Clone(m.GetAws()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws.DestinationSpec),
			}
		}

	case *DestinationSpec_Azure:

		if h, ok := interface{}(m.GetAzure()).(clone.Cloner); ok {
			target.DestinationType = &DestinationSpec_Azure{
				Azure: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_azure.DestinationSpec),
			}
		} else {
			target.DestinationType = &DestinationSpec_Azure{
				Azure: proto.Clone(m.GetAzure()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_azure.DestinationSpec),
			}
		}

	case *DestinationSpec_Rest:

		if h, ok := interface{}(m.GetRest()).(clone.Cloner); ok {
			target.DestinationType = &DestinationSpec_Rest{
				Rest: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_rest.DestinationSpec),
			}
		} else {
			target.DestinationType = &DestinationSpec_Rest{
				Rest: proto.Clone(m.GetRest()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_rest.DestinationSpec),
			}
		}

	case *DestinationSpec_Grpc:

		if h, ok := interface{}(m.GetGrpc()).(clone.Cloner); ok {
			target.DestinationType = &DestinationSpec_Grpc{
				Grpc: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc.DestinationSpec),
			}
		} else {
			target.DestinationType = &DestinationSpec_Grpc{
				Grpc: proto.Clone(m.GetGrpc()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc.DestinationSpec),
			}
		}

	}

	return target
}
