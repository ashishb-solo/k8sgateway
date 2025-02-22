// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/extensions/proxylatency/proxylatency.proto

package proxylatency

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *ProxyLatency) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("envoy.config.filter.http.proxylatency.v2.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/extensions/proxylatency.ProxyLatency")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRequest())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMeasureRequestInternally())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetResponse())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetChargeClusterStat()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ChargeClusterStat")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetChargeClusterStat(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ChargeClusterStat")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetChargeListenerStat()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ChargeListenerStat")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetChargeListenerStat(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ChargeListenerStat")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEmitDynamicMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EmitDynamicMetadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEmitDynamicMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EmitDynamicMetadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
