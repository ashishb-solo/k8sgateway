// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/secret.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extauth/v1"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
func (m *Secret) Clone() proto.Message {
	var target *Secret
	if m == nil {
		return target
	}
	target = &Secret{}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	switch m.Kind.(type) {

	case *Secret_Aws:

		if h, ok := interface{}(m.GetAws()).(clone.Cloner); ok {
			target.Kind = &Secret_Aws{
				Aws: h.Clone().(*AwsSecret),
			}
		} else {
			target.Kind = &Secret_Aws{
				Aws: proto.Clone(m.GetAws()).(*AwsSecret),
			}
		}

	case *Secret_Azure:

		if h, ok := interface{}(m.GetAzure()).(clone.Cloner); ok {
			target.Kind = &Secret_Azure{
				Azure: h.Clone().(*AzureSecret),
			}
		} else {
			target.Kind = &Secret_Azure{
				Azure: proto.Clone(m.GetAzure()).(*AzureSecret),
			}
		}

	case *Secret_Tls:

		if h, ok := interface{}(m.GetTls()).(clone.Cloner); ok {
			target.Kind = &Secret_Tls{
				Tls: h.Clone().(*TlsSecret),
			}
		} else {
			target.Kind = &Secret_Tls{
				Tls: proto.Clone(m.GetTls()).(*TlsSecret),
			}
		}

	case *Secret_Oauth:

		if h, ok := interface{}(m.GetOauth()).(clone.Cloner); ok {
			target.Kind = &Secret_Oauth{
				Oauth: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.OauthSecret),
			}
		} else {
			target.Kind = &Secret_Oauth{
				Oauth: proto.Clone(m.GetOauth()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.OauthSecret),
			}
		}

	case *Secret_ApiKey:

		if h, ok := interface{}(m.GetApiKey()).(clone.Cloner); ok {
			target.Kind = &Secret_ApiKey{
				ApiKey: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ApiKey),
			}
		} else {
			target.Kind = &Secret_ApiKey{
				ApiKey: proto.Clone(m.GetApiKey()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ApiKey),
			}
		}

	case *Secret_Header:

		if h, ok := interface{}(m.GetHeader()).(clone.Cloner); ok {
			target.Kind = &Secret_Header{
				Header: h.Clone().(*HeaderSecret),
			}
		} else {
			target.Kind = &Secret_Header{
				Header: proto.Clone(m.GetHeader()).(*HeaderSecret),
			}
		}

	case *Secret_Credentials:

		if h, ok := interface{}(m.GetCredentials()).(clone.Cloner); ok {
			target.Kind = &Secret_Credentials{
				Credentials: h.Clone().(*AccountCredentialsSecret),
			}
		} else {
			target.Kind = &Secret_Credentials{
				Credentials: proto.Clone(m.GetCredentials()).(*AccountCredentialsSecret),
			}
		}

	case *Secret_Encryption:

		if h, ok := interface{}(m.GetEncryption()).(clone.Cloner); ok {
			target.Kind = &Secret_Encryption{
				Encryption: h.Clone().(*EncryptionKeySecret),
			}
		} else {
			target.Kind = &Secret_Encryption{
				Encryption: proto.Clone(m.GetEncryption()).(*EncryptionKeySecret),
			}
		}

	case *Secret_Extensions:

		if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
			target.Kind = &Secret_Extensions{
				Extensions: h.Clone().(*Extensions),
			}
		} else {
			target.Kind = &Secret_Extensions{
				Extensions: proto.Clone(m.GetExtensions()).(*Extensions),
			}
		}

	}

	return target
}

// Clone function
func (m *AwsSecret) Clone() proto.Message {
	var target *AwsSecret
	if m == nil {
		return target
	}
	target = &AwsSecret{}

	target.AccessKey = m.GetAccessKey()

	target.SecretKey = m.GetSecretKey()

	target.SessionToken = m.GetSessionToken()

	return target
}

// Clone function
func (m *AzureSecret) Clone() proto.Message {
	var target *AzureSecret
	if m == nil {
		return target
	}
	target = &AzureSecret{}

	if m.GetApiKeys() != nil {
		target.ApiKeys = make(map[string]string, len(m.GetApiKeys()))
		for k, v := range m.GetApiKeys() {

			target.ApiKeys[k] = v

		}
	}

	return target
}

// Clone function
func (m *TlsSecret) Clone() proto.Message {
	var target *TlsSecret
	if m == nil {
		return target
	}
	target = &TlsSecret{}

	target.CertChain = m.GetCertChain()

	target.PrivateKey = m.GetPrivateKey()

	target.RootCa = m.GetRootCa()

	if m.GetOcspStaple() != nil {
		target.OcspStaple = make([]byte, len(m.GetOcspStaple()))
		copy(target.OcspStaple, m.GetOcspStaple())
	}

	return target
}

// Clone function
func (m *HeaderSecret) Clone() proto.Message {
	var target *HeaderSecret
	if m == nil {
		return target
	}
	target = &HeaderSecret{}

	if m.GetHeaders() != nil {
		target.Headers = make(map[string]string, len(m.GetHeaders()))
		for k, v := range m.GetHeaders() {

			target.Headers[k] = v

		}
	}

	return target
}

// Clone function
func (m *AccountCredentialsSecret) Clone() proto.Message {
	var target *AccountCredentialsSecret
	if m == nil {
		return target
	}
	target = &AccountCredentialsSecret{}

	target.Username = m.GetUsername()

	target.Password = m.GetPassword()

	return target
}

// Clone function
func (m *EncryptionKeySecret) Clone() proto.Message {
	var target *EncryptionKeySecret
	if m == nil {
		return target
	}
	target = &EncryptionKeySecret{}

	target.Key = m.GetKey()

	return target
}
