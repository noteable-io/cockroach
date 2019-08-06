// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package exec

import (
	"fmt"
	"unsafe"

	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
)

const (
	sizeOfBool    = int(unsafe.Sizeof(true))
	sizeOfInt8    = int(unsafe.Sizeof(int8(0)))
	sizeOfInt16   = int(unsafe.Sizeof(int16(0)))
	sizeOfInt32   = int(unsafe.Sizeof(int32(0)))
	sizeOfInt64   = int(unsafe.Sizeof(int64(0)))
	sizeOfFloat32 = int(unsafe.Sizeof(float32(0)))
	sizeOfFloat64 = int(unsafe.Sizeof(float64(0)))
)

// EstimateBatchSizeBytes returns an estimated amount of bytes needed to
// store a batch in memory that has column types vecTypes.
// WARNING: This only is correct for fixed width types, and returns an
// estimate for non fixed width types. In future it might be possible to
// remove the need for estimation by specifying batch sizes in terms of bytes.
func EstimateBatchSizeBytes(vecTypes []types.T, batchLength int) int {
	// acc represents the number of bytes to represent a row in the batch.
	acc := 0
	for _, t := range vecTypes {
		switch t {
		case types.Bool:
			acc += sizeOfBool
		case types.Bytes:
			// We don't know without looking at the data in a batch to see how
			// much space each byte array takes up. Use some default value as a
			// heuristic right now.
			acc += 100
		case types.Int8:
			acc += sizeOfInt8
		case types.Int16:
			acc += sizeOfInt16
		case types.Int32:
			acc += sizeOfInt32
		case types.Int64:
			acc += sizeOfInt64
		case types.Float32:
			acc += sizeOfFloat32
		case types.Float64:
			acc += sizeOfFloat64
		case types.Decimal:
			// Similar to byte arrays, we can't tell how much space is used
			// to hold the arbitrary precision decimal objects.
			acc += 50
		default:
			panic(fmt.Sprintf("unhandled type %s", t))
		}
	}
	return acc * batchLength
}