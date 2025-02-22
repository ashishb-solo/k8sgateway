// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/load_balancer.proto

package v1

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
func (m *LoadBalancerConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.LoadBalancerConfig")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHealthyPanicThreshold()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HealthyPanicThreshold")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHealthyPanicThreshold(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HealthyPanicThreshold")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUpdateMergeWindow()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UpdateMergeWindow")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpdateMergeWindow(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UpdateMergeWindow")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUseHostnameForHashing()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UseHostnameForHashing")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUseHostnameForHashing(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UseHostnameForHashing")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetCloseConnectionsOnHostSetChange())
	if err != nil {
		return 0, err
	}

	switch m.Type.(type) {

	case *LoadBalancerConfig_RoundRobin_:

		if h, ok := interface{}(m.GetRoundRobin()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RoundRobin")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRoundRobin(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RoundRobin")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *LoadBalancerConfig_LeastRequest_:

		if h, ok := interface{}(m.GetLeastRequest()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("LeastRequest")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetLeastRequest(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("LeastRequest")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *LoadBalancerConfig_Random_:

		if h, ok := interface{}(m.GetRandom()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Random")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRandom(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Random")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *LoadBalancerConfig_RingHash_:

		if h, ok := interface{}(m.GetRingHash()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RingHash")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRingHash(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RingHash")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *LoadBalancerConfig_Maglev_:

		if h, ok := interface{}(m.GetMaglev()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Maglev")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMaglev(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Maglev")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.LocalityConfig.(type) {

	case *LoadBalancerConfig_LocalityWeightedLbConfig:

		if h, ok := interface{}(m.GetLocalityWeightedLbConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("LocalityWeightedLbConfig")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetLocalityWeightedLbConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("LocalityWeightedLbConfig")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *LoadBalancerConfig_RoundRobin) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.LoadBalancerConfig_RoundRobin")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSlowStartConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SlowStartConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSlowStartConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SlowStartConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *LoadBalancerConfig_LeastRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.LoadBalancerConfig_LeastRequest")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetChoiceCount())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSlowStartConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SlowStartConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSlowStartConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SlowStartConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *LoadBalancerConfig_Random) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.LoadBalancerConfig_Random")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *LoadBalancerConfig_RingHashConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.LoadBalancerConfig_RingHashConfig")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMinimumRingSize())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMaximumRingSize())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *LoadBalancerConfig_RingHash) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.LoadBalancerConfig_RingHash")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRingHashConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RingHashConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRingHashConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RingHashConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *LoadBalancerConfig_Maglev) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.LoadBalancerConfig_Maglev")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *LoadBalancerConfig_SlowStartConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.LoadBalancerConfig_SlowStartConfig")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetSlowStartWindow()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SlowStartWindow")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSlowStartWindow(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SlowStartWindow")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAggression()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Aggression")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAggression(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Aggression")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMinWeightPercent()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MinWeightPercent")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMinWeightPercent(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MinWeightPercent")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
