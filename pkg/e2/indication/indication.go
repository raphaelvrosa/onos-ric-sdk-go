// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package indication

import "github.com/onosproject/onos-ric-sdk-go/pkg/e2/encoding"

// Indication is an E2 indication
type Indication struct {
	// EncodingType payload encoding type
	EncodingType encoding.Type

	// Payload is the indication payload
	Payload Payload
}

// Payload is an E2 indication payload
type Payload struct {
	Header  []byte
	Message []byte
}
