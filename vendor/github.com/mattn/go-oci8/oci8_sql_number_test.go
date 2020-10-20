package oci8

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"sync"
	"testing"
)

// TestSelectDualNullNumber checks nulls
func TestSelectDualNullNumber(t *testing.T) {
	if TestDisableDatabase {
		t.SkipNow()
	}

	t.Parallel()

	// null
	queryResults := testQueryResults{
		query:        "select null from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// INTEGER
	queryResults = testQueryResults{
		query:        "select cast (null as INTEGER) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// INT(
	queryResults = testQueryResults{
		query:        "select cast (null as INT) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// SMALLINT
	queryResults = testQueryResults{
		query:        "select cast (null as SMALLINT) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// NUMBER(38,10)
	queryResults = testQueryResults{
		query:        "select cast (null as NUMBER(38,10)) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// DEC(38,10)
	queryResults = testQueryResults{
		query:        "select cast (null as DEC(38,10)) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// DECIMAL(38,10)
	queryResults = testQueryResults{
		query:        "select cast (null as DECIMAL(38,10)) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// NUMERIC(38,10)
	queryResults = testQueryResults{
		query:        "select cast (null as NUMERIC(38,10)) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// FLOAT
	queryResults = testQueryResults{
		query:        "select cast (null as FLOAT) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// REAL
	queryResults = testQueryResults{
		query:        "select cast (null as REAL) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_FLOAT
	queryResults = testQueryResults{
		query:        "select cast (null as BINARY_FLOAT) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_DOUBLE
	queryResults = testQueryResults{
		query:        "select cast (null as BINARY_DOUBLE) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)
}

// TestSelectDualNumber checks select dual for number types
func TestSelectDualNumber(t *testing.T) {
	if TestDisableDatabase {
		t.SkipNow()
	}

	t.Parallel()

	queryResults := testQueryResults{}

	// bool
	queryResultBoolToInt64 := []testQueryResult{
		{
			args:    []interface{}{false},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{true},
			results: [][]interface{}{{int64(1)}},
		},
	}
	// int8: -128 to 127
	queryResultInt8ToInt64 := []testQueryResult{
		{
			args:    []interface{}{int8(-128)},
			results: [][]interface{}{{int64(-128)}},
		},
		{
			args:    []interface{}{int8(-1)},
			results: [][]interface{}{{int64(-1)}},
		},
		{
			args:    []interface{}{int8(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{int8(1)},
			results: [][]interface{}{{int64(1)}},
		},
		{
			args:    []interface{}{int8(127)},
			results: [][]interface{}{{int64(127)}},
		},
	}
	// int16: -32768 to 32767
	queryResultInt16ToInt64 := []testQueryResult{
		{
			args:    []interface{}{int16(-32768)},
			results: [][]interface{}{{int64(-32768)}},
		},
		{
			args:    []interface{}{int16(-128)},
			results: [][]interface{}{{int64(-128)}},
		},
		{
			args:    []interface{}{int16(127)},
			results: [][]interface{}{{int64(127)}},
		},
		{
			args:    []interface{}{int16(32767)},
			results: [][]interface{}{{int64(32767)}},
		},
	}
	// int32: -2147483648 to 2147483647
	queryResultInt32ToInt64 := []testQueryResult{
		{
			args:    []interface{}{int32(-2147483648)},
			results: [][]interface{}{{int64(-2147483648)}},
		},
		{
			args:    []interface{}{int32(-32768)},
			results: [][]interface{}{{int64(-32768)}},
		},
		{
			args:    []interface{}{int32(32767)},
			results: [][]interface{}{{int64(32767)}},
		},
		{
			args:    []interface{}{int32(2147483647)},
			results: [][]interface{}{{int64(2147483647)}},
		},
	}
	// int64: -9223372036854775808 to 9223372036854775807
	queryResultInt64ToInt64 := []testQueryResult{
		{
			args:    []interface{}{int64(-9223372036854775808)},
			results: [][]interface{}{{int64(-9223372036854775808)}},
		},
		{
			args:    []interface{}{int64(-2147483648)},
			results: [][]interface{}{{int64(-2147483648)}},
		},
		{
			args:    []interface{}{int64(2147483647)},
			results: [][]interface{}{{int64(2147483647)}},
		},
		{
			args:    []interface{}{int64(9223372036854775807)},
			results: [][]interface{}{{int64(9223372036854775807)}},
		},
	}
	// uint8: 0 to 255
	queryResultUint8ToInt64 := []testQueryResult{
		{
			args:    []interface{}{uint8(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{uint8(1)},
			results: [][]interface{}{{int64(1)}},
		},
		{
			args:    []interface{}{uint8(127)},
			results: [][]interface{}{{int64(127)}},
		},
		{
			args:    []interface{}{uint8(128)},
			results: [][]interface{}{{int64(128)}},
		},
		{
			args:    []interface{}{uint8(255)},
			results: [][]interface{}{{int64(255)}},
		},
	}
	// uint16: 0 to 65535
	queryResultUint16ToInt64 := []testQueryResult{
		{
			args:    []interface{}{uint16(255)},
			results: [][]interface{}{{int64(255)}},
		},
		{
			args:    []interface{}{uint16(65535)},
			results: [][]interface{}{{int64(65535)}},
		},
	}
	// uint32: 0 to 4294967295
	queryResultUint32ToInt64 := []testQueryResult{
		{
			args:    []interface{}{uint32(65535)},
			results: [][]interface{}{{int64(65535)}},
		},
		{
			args:    []interface{}{uint32(4294967295)},
			results: [][]interface{}{{int64(4294967295)}},
		},
	}
	// uint64: 0 to 18446744073709551615
	// for 18446744073709551615 get: get rows error: query error: sql: converting argument $1 type: uint64 values with high bit set are not supported
	queryResultUint64ToInt64 := []testQueryResult{
		{
			args:    []interface{}{uint64(4294967295)},
			results: [][]interface{}{{int64(4294967295)}},
		},
		{
			args:    []interface{}{uint64(9223372036854775807)},
			results: [][]interface{}{{int64(9223372036854775807)}},
		},
	}
	// bool
	queryResultBoolToFloat64 := []testQueryResult{
		{
			args:    []interface{}{false},
			results: [][]interface{}{{float64(0)}},
		},
		{
			args:    []interface{}{true},
			results: [][]interface{}{{float64(1)}},
		},
	}
	// int8: -128 to 127
	queryResultInt8ToFloat64 := []testQueryResult{
		{
			args:    []interface{}{int8(-128)},
			results: [][]interface{}{{float64(-128)}},
		},
		{
			args:    []interface{}{int8(-1)},
			results: [][]interface{}{{float64(-1)}},
		},
		{
			args:    []interface{}{int8(0)},
			results: [][]interface{}{{float64(0)}},
		},
		{
			args:    []interface{}{int8(1)},
			results: [][]interface{}{{float64(1)}},
		},
		{
			args:    []interface{}{int8(127)},
			results: [][]interface{}{{float64(127)}},
		},
	}
	// int16: -32768 to 32767
	queryResultInt16ToFloat64 := []testQueryResult{
		{
			args:    []interface{}{int16(-32768)},
			results: [][]interface{}{{float64(-32768)}},
		},
		{
			args:    []interface{}{int16(-128)},
			results: [][]interface{}{{float64(-128)}},
		},
		{
			args:    []interface{}{int16(127)},
			results: [][]interface{}{{float64(127)}},
		},
		{
			args:    []interface{}{int16(32767)},
			results: [][]interface{}{{float64(32767)}},
		},
	}
	// int32: -2147483648 to 2147483647
	queryResultInt32ToFloat64 := []testQueryResult{
		{
			args:    []interface{}{int32(-2147483648)},
			results: [][]interface{}{{float64(-2147483648)}},
		},
		{
			args:    []interface{}{int32(-32768)},
			results: [][]interface{}{{float64(-32768)}},
		},
		{
			args:    []interface{}{int32(32767)},
			results: [][]interface{}{{float64(32767)}},
		},
		{
			args:    []interface{}{int32(2147483647)},
			results: [][]interface{}{{float64(2147483647)}},
		},
	}
	// int64: -9223372036854775808 to 9223372036854775807
	queryResultInt64ToFloat64 := []testQueryResult{
		{
			args:    []interface{}{int64(-9223372036854775808)},
			results: [][]interface{}{{float64(-9223372036854775808)}},
		},
		{
			args:    []interface{}{int64(-2147483648)},
			results: [][]interface{}{{float64(-2147483648)}},
		},
		{
			args:    []interface{}{int64(2147483647)},
			results: [][]interface{}{{float64(2147483647)}},
		},
		{
			args:    []interface{}{int64(9223372036854775807)},
			results: [][]interface{}{{float64(9223372036854775807)}},
		},
	}
	// uint8: 0 to 255
	queryResultUint8ToFloat64 := []testQueryResult{
		{
			args:    []interface{}{uint8(0)},
			results: [][]interface{}{{float64(0)}},
		},
		{
			args:    []interface{}{uint8(1)},
			results: [][]interface{}{{float64(1)}},
		},
		{
			args:    []interface{}{uint8(127)},
			results: [][]interface{}{{float64(127)}},
		},
		{
			args:    []interface{}{uint8(128)},
			results: [][]interface{}{{float64(128)}},
		},
		{
			args:    []interface{}{uint8(255)},
			results: [][]interface{}{{float64(255)}},
		},
	}
	// uint16: 0 to 65535
	queryResultUint16ToFloat64 := []testQueryResult{
		{
			args:    []interface{}{uint16(255)},
			results: [][]interface{}{{float64(255)}},
		},
		{
			args:    []interface{}{uint16(65535)},
			results: [][]interface{}{{float64(65535)}},
		},
	}
	// uint32: 0 to 4294967295
	queryResultUint32ToFloat64 := []testQueryResult{
		{
			args:    []interface{}{uint32(65535)},
			results: [][]interface{}{{float64(65535)}},
		},
		{
			args:    []interface{}{uint32(4294967295)},
			results: [][]interface{}{{float64(4294967295)}},
		},
	}
	// uint64: 0 to 18446744073709551615
	// for 18446744073709551615 get: get rows error: query error: sql: converting argument $1 type: uint64 values with high bit set are not supported
	queryResultUint64ToFloat64 := []testQueryResult{
		{
			args:    []interface{}{uint64(4294967295)},
			results: [][]interface{}{{float64(4294967295)}},
		},
		{
			args:    []interface{}{uint64(9223372036854775807)},
			results: [][]interface{}{{float64(9223372036854775807)}},
		},
	}
	// float32: sign 1 bit, exponent 8 bits, mantissa 23 bits
	queryResultFloat32ToFloat64 := []testQueryResult{
		{ // 0 00000000 00000000000000000000000
			args:    []interface{}{math.Float32frombits(0x00000000)},
			results: [][]interface{}{{float64(math.Float32frombits(0x00000000))}},
		},
		{ // 1 00000000 00000000000000000000000
			args:    []interface{}{math.Float32frombits(0x80000000)},
			results: [][]interface{}{{float64(math.Float32frombits(0x80000000))}},
		},
	}
	// TODO: look at being able to test at greater bit range, seem to be losing precision
	for x := 0; x < 4; x++ { // positive/negtive number and positive/negtive exponent loop
		for i := uint32(23); i < 26; i++ { // exponent[30:24] starts at 23 to have bits start at 0
			for j := uint32(24); j > 18; j-- { // mantissa [23:0] starts at 24 to have bits start at 0
				bits := uint32(0)
				if x == 1 || x == 3 {
					// negtive number has bit 31 set
					bits |= uint32(0x80000000)
				}
				if x < 2 {
					// positive exponent has bit 30 set
					bits |= uint32(0x40000000)
					for k := uint32(24); k <= i; k++ {
						// exponent starts at bit 24 and goes up
						bits |= 1 << k
					}
				} else {
					// negtive exponent starts with bits 29 to 24 set for 0
					// remove bits starting at 24
					for k := uint32(29); k > i; k-- {
						bits |= 1 << k
					}
				}
				for l := uint32(23); l >= j; l-- {
					// mantissa starts at bit 23 and goes down
					bits |= 1 << l
				}
				float := math.Float32frombits(bits)
				queryResultFloat32ToFloat64 = append(queryResultFloat32ToFloat64,
					testQueryResult{
						args:    []interface{}{float},
						results: [][]interface{}{{float64(float)}},
					},
				)
			}
		}
	}

	// TODO: added float64

	// https://ss64.com/ora/syntax-datatypes.html

	// INTEGER
	queryResults.query = "select cast (:1 as INTEGER) from dual"
	queryResults.queryResults = queryResultBoolToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToInt64
	testRunQueryResults(t, queryResults)

	// INT
	queryResults.query = "select cast (:1 as INT) from dual"
	queryResults.queryResults = queryResultBoolToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToInt64
	testRunQueryResults(t, queryResults)

	// SMALLINT
	queryResults.query = "select cast (:1 as SMALLINT) from dual"
	queryResults.queryResults = queryResultBoolToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToInt64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToInt64
	testRunQueryResults(t, queryResults)

	// NUMBER(38,10)
	queryResults.query = "select cast (:1 as NUMBER(38,10)) from dual"
	queryResults.queryResults = queryResultBoolToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultFloat32ToFloat64
	testRunQueryResults(t, queryResults)

	// DEC(38,10)
	queryResults.query = "select cast (:1 as DEC(38,10)) from dual"
	queryResults.queryResults = queryResultBoolToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultFloat32ToFloat64
	testRunQueryResults(t, queryResults)

	// DECIMAL(38,10)
	queryResults.query = "select cast (:1 as DECIMAL(38,10)) from dual"
	queryResults.queryResults = queryResultBoolToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultFloat32ToFloat64
	testRunQueryResults(t, queryResults)

	// NUMERIC(38,10)
	queryResults.query = "select cast (:1 as NUMERIC(38,10)) from dual"
	queryResults.queryResults = queryResultBoolToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultFloat32ToFloat64
	testRunQueryResults(t, queryResults)

	// FLOAT
	queryResults.query = "select cast (:1 as FLOAT) from dual"
	queryResults.queryResults = queryResultBoolToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultFloat32ToFloat64
	testRunQueryResults(t, queryResults)

	// REAL
	queryResults.query = "select cast (:1 as REAL) from dual"
	queryResults.queryResults = queryResultBoolToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultFloat32ToFloat64
	testRunQueryResults(t, queryResults)

	// BINARY_FLOAT
	queryResults.query = "select cast (:1 as BINARY_FLOAT) from dual"
	queryResults.queryResults = queryResultBoolToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultFloat32ToFloat64
	testRunQueryResults(t, queryResults)

	// BINARY_DOUBLE
	queryResults.query = "select cast (:1 as BINARY_DOUBLE) from dual"
	queryResults.queryResults = queryResultBoolToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultFloat32ToFloat64
	testRunQueryResults(t, queryResults)

	// https://tour.golang.org/basics/11

	// Go
	queryResults.query = "select :1 from dual"
	queryResults.queryResults = queryResultBoolToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultInt64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint8ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint16ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint32ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultUint64ToFloat64
	testRunQueryResults(t, queryResults)
	queryResults.queryResults = queryResultFloat32ToFloat64
	testRunQueryResults(t, queryResults)

	// sum
	queryResults.query = "select sum(A) from (select :1 as A from dual union select :2 as A from dual)"
	queryResults.queryResults = []testQueryResult{
		{
			args:    []interface{}{int64(1), int64(2)},
			results: [][]interface{}{{float64(3)}},
		},
		{
			args:    []interface{}{int64(1), float64(2.25)},
			results: [][]interface{}{{float64(3.25)}},
		},
		{
			args:    []interface{}{float64(1.5), float64(2.25)},
			results: [][]interface{}{{float64(3.75)}},
		},
	}
	testRunQueryResults(t, queryResults)
}

// TestDestructiveNumber checks insert, select, update, and delete of number types
func TestDestructiveNumber(t *testing.T) {
	if TestDisableDatabase || TestDisableDestructive {
		t.SkipNow()
	}

	// NUMBER negative
	tableName := "NUMBER_" + TestTimeString
	err := testExec(t, "create table "+tableName+" ( A NUMBER(10,2), B NUMBER(20,4), C NUMBER(38,8) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults := testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// test truncate
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName,
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// NUMBER positive
	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0), float64(0), float64(0)},
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1), float64(1), float64(1)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{0},
			{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// DEC negative
	tableName = "DEC_" + TestTimeString
	err = testExec(t, "create table "+tableName+
		" ( A DEC(10,2), B DEC(20,4), C DEC(38,8) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// DEC positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0), float64(0), float64(0)},
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1), float64(1), float64(1)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{0},
			{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// DECIMAL negative
	tableName = "DECIMAL_" + TestTimeString
	err = testExec(t, "create table "+tableName+" ( A DECIMAL(10,2), B DECIMAL(20,4), C DECIMAL(38,8) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// DECIMAL positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0), float64(0), float64(0)},
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1), float64(1), float64(1)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{0},
			{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// NUMERIC negative
	tableName = "NUMERIC_" + TestTimeString
	err = testExec(t, "create table "+tableName+" ( A NUMERIC(10,2), B NUMERIC(20,4), C NUMERIC(38,8) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// NUMERIC positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0), float64(0), float64(0)},
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1), float64(1), float64(1)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{0},
			{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// FLOAT negative
	tableName = "FLOAT_" + TestTimeString
	err = testExec(t, "create table "+tableName+
		" ( A FLOAT(28), B FLOAT(32), C FLOAT(38) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// FLOAT positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0), float64(0), float64(0)},
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1), float64(1), float64(1)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{0},
			{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INTEGER negative
	tableName = "INTEGER_" + TestTimeString
	err = testExec(t, "create table "+tableName+" ( A INTEGER, B INTEGER, C INTEGER )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(-21474836), int64(-2147483648), int64(-2147483648)},
					{int64(-10000000), int64(-1000000000000000), int64(-1000000000000000)},
					{int64(-1234567), int64(-123456792), int64(-123456792)},
					{int64(-2), int64(-2), int64(-2)},
					{int64(-1), int64(-1), int64(-1)},
					{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-10000000},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(-1234567), int64(-123456792), int64(-123456792)},
					{int64(-2), int64(-2), int64(-2)},
					{int64(-1), int64(-1), int64(-1)},
					{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INTEGER positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(0), int64(0), int64(0)},
					{int64(1), int64(1), int64(1)},
					{int64(1), int64(1), int64(1)},
					{int64(2), int64(2), int64(2)},
					{int64(10000000), int64(1000000000000000), int64(1000000000000000)},
					{int64(12345679), int64(123456792), int64(123456792)},
					{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{10000000},
			{12345679},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(0), int64(0), int64(0)},
					{int64(1), int64(1), int64(1)},
					{int64(1), int64(1), int64(1)},
					{int64(2), int64(2), int64(2)},
					{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INT negative
	tableName = "INT_" + TestTimeString
	err = testExec(t, "create table "+tableName+" ( A INT, B INT, C INT )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(-21474836), int64(-2147483648), int64(-2147483648)},
					{int64(-10000000), int64(-1000000000000000), int64(-1000000000000000)},
					{int64(-1234567), int64(-123456792), int64(-123456792)},
					{int64(-2), int64(-2), int64(-2)},
					{int64(-1), int64(-1), int64(-1)},
					{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-10000000},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(-1234567), int64(-123456792), int64(-123456792)},
					{int64(-2), int64(-2), int64(-2)},
					{int64(-1), int64(-1), int64(-1)},
					{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INT positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(0), int64(0), int64(0)},
					{int64(1), int64(1), int64(1)},
					{int64(1), int64(1), int64(1)},
					{int64(2), int64(2), int64(2)},
					{int64(10000000), int64(1000000000000000), int64(1000000000000000)},
					{int64(12345679), int64(123456792), int64(123456792)},
					{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{10000000},
			{12345679},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(0), int64(0), int64(0)},
					{int64(1), int64(1), int64(1)},
					{int64(1), int64(1), int64(1)},
					{int64(2), int64(2), int64(2)},
					{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// SMALLINT negative
	tableName = "SMALLINT_" + TestTimeString
	err = testExec(t, "create table "+tableName+" ( A SMALLINT, B SMALLINT, C SMALLINT )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(-21474836), int64(-2147483648), int64(-2147483648)},
					{int64(-10000000), int64(-1000000000000000), int64(-1000000000000000)},
					{int64(-1234567), int64(-123456792), int64(-123456792)},
					{int64(-2), int64(-2), int64(-2)},
					{int64(-1), int64(-1), int64(-1)},
					{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-10000000},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(-1234567), int64(-123456792), int64(-123456792)},
					{int64(-2), int64(-2), int64(-2)},
					{int64(-1), int64(-1), int64(-1)},
					{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// SMALLINT positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(0), int64(0), int64(0)},
					{int64(1), int64(1), int64(1)},
					{int64(1), int64(1), int64(1)},
					{int64(2), int64(2), int64(2)},
					{int64(10000000), int64(1000000000000000), int64(1000000000000000)},
					{int64(12345679), int64(123456792), int64(123456792)},
					{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{10000000},
			{12345679},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(0), int64(0), int64(0)},
					{int64(1), int64(1), int64(1)},
					{int64(1), int64(1), int64(1)},
					{int64(2), int64(2), int64(2)},
					{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// REAL negative
	tableName = "REAL_" + TestTimeString
	err = testExec(t, "create table "+tableName+
		" ( A REAL, B REAL, C REAL )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// REAL positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0), float64(0), float64(0)},
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1), float64(1), float64(1)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{0},
			{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_FLOAT negative
	tableName = "BINARY_FLOAT_" + TestTimeString
	err = testExec(t, "create table "+tableName+
		" ( A BINARY_FLOAT, B BINARY_FLOAT, C BINARY_FLOAT )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{float64(-288230381928101358902502915674136903680), float64(-288230381928101358902502915674136903680), float64(-288230381928101358902502915674136903680)},
			{-2147483648, -2147483648, -2147483648},
			{-123456792, -123456792, -123456792},
			{-1.99999988079071044921875, -1.99999988079071044921875, -1.99999988079071044921875},
			{-1, -1, -1},
			{-0.00415134616196155548095703125, -0.00415134616196155548095703125, -0.00415134616196155548095703125},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-288230381928101358902502915674136903680), float64(-288230381928101358902502915674136903680), float64(-288230381928101358902502915674136903680)},
					{float64(-2147483648), float64(-2147483648), float64(-2147483648)},
					{float64(-123456792), float64(-123456792), float64(-123456792)},
					{float64(-1.99999988079071044921875), float64(-1.99999988079071044921875), float64(-1.99999988079071044921875)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.00415134616196155548095703125), float64(-0.00415134616196155548095703125), float64(-0.00415134616196155548095703125)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-2147483648},
			{float64(-288230381928101358902502915674136903680)},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-123456792), float64(-123456792), float64(-123456792)},
					{float64(-1.99999988079071044921875), float64(-1.99999988079071044921875), float64(-1.99999988079071044921875)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.00415134616196155548095703125), float64(-0.00415134616196155548095703125), float64(-0.00415134616196155548095703125)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_FLOAT positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.00415134616196155548095703125, 0.00415134616196155548095703125, 0.00415134616196155548095703125},
			{1, 1, 1},
			{1.99999988079071044921875, 1.99999988079071044921875, 1.99999988079071044921875},
			{123456792, 123456792, 123456792},
			{2147483648, 2147483648, 2147483648},
			{float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680)},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0), float64(0), float64(0)},
					{float64(0.00415134616196155548095703125), float64(0.00415134616196155548095703125), float64(0.00415134616196155548095703125)},
					{float64(1), float64(1), float64(1)},
					{float64(1.99999988079071044921875), float64(1.99999988079071044921875), float64(1.99999988079071044921875)},
					{float64(123456792), float64(123456792), float64(123456792)},
					{float64(2147483648), float64(2147483648), float64(2147483648)},
					{float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{0},
			{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0.00415134616196155548095703125), float64(0.00415134616196155548095703125), float64(0.00415134616196155548095703125)},
					{float64(1.99999988079071044921875), float64(1.99999988079071044921875), float64(1.99999988079071044921875)},
					{float64(123456792), float64(123456792), float64(123456792)},
					{float64(2147483648), float64(2147483648), float64(2147483648)},
					{float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_DOUBLE negative
	tableName = "BINARY_DOUBLE_" + TestTimeString
	err = testExec(t, "create table "+tableName+" ( A BINARY_DOUBLE, B BINARY_DOUBLE, C BINARY_DOUBLE )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			{-21474836, -2147483648, -2147483648},
			{-1234567, -123456792, -123456792},
			{-1.98, -1.9873, -1.98730468},
			{-1, -1, -1},
			{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{-21474836},
			{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(-1234567), float64(-123456792), float64(-123456792)},
					{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					{float64(-1), float64(-1), float64(-1)},
					{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_DOUBLE positive
	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{0, 0, 0},
			{0.76, 0.7617, 0.76171875},
			{1, 1, 1},
			{1.98, 1.9873, 1.98730468},
			{12345679, 123456792, 123456792},
			{21474836, 2147483647, 2147483647},
			{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0), float64(0), float64(0)},
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1), float64(1), float64(1)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{0},
			{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(0.76), float64(0.7617), float64(0.76171875)},
					{float64(1.98), float64(1.9873), float64(1.98730468)},
					{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					{float64(12345679), float64(123456792), float64(123456792)},
					{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)
}

func TestFunctionCallNumber(t *testing.T) {
	if TestDisableDatabase {
		t.SkipNow()
	}

	t.Parallel()

	// https://ss64.com/ora/syntax-datatypes.html

	var execResults testExecResults

	// bool
	execResultBool := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: true, In: true}},
			results: map[string]interface{}{"num1": false},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: true, In: true}},
			results: map[string]interface{}{"num1": false},
		},
	}

	// bool with INTEGER
	execResults.query = `
declare
	function GET_NUMBER(p_number INTEGER) return INTEGER as
	begin
		if p_number = 0 then
			return 1;
		end if;
		return 0;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultBool
	testRunExecResults(t, execResults)

	// bool
	execResultBool = []testExecResult{
		{
			args:    map[string]sql.Out{"bool1": {Dest: true, In: true}},
			results: map[string]interface{}{"bool1": false},
		},
	}

	execResults.query = `begin :bool1 := 0; end;`
	execResults.execResults = execResultBool
	testRunExecResults(t, execResults)

	// bool
	execResultBool = []testExecResult{
		{
			args:    map[string]sql.Out{"bool1": {Dest: false, In: true}},
			results: map[string]interface{}{"bool1": true},
		},
	}

	execResults.query = `begin :bool1 := 1; end;`
	execResults.execResults = execResultBool
	testRunExecResults(t, execResults)

	// bool
	execResultBool = []testExecResult{
		{
			args:    map[string]sql.Out{"bool1": {Dest: false, In: true}},
			results: map[string]interface{}{"bool1": false},
		},
		{
			args:    map[string]sql.Out{"bool1": {Dest: true, In: true}},
			results: map[string]interface{}{"bool1": true},
		},
	}

	execResults.query = `begin :bool1 := :bool1; end;`
	execResults.execResults = execResultBool
	testRunExecResults(t, execResults)

	// int8: -128 to 127
	execResultInt8 := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: int8(-128), In: true}},
			results: map[string]interface{}{"num1": int8(-128)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int8(-1), In: true}},
			results: map[string]interface{}{"num1": int8(-1)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int8(0), In: true}},
			results: map[string]interface{}{"num1": int8(0)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int8(1), In: true}},
			results: map[string]interface{}{"num1": int8(1)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int8(127), In: true}},
			results: map[string]interface{}{"num1": int8(127)},
		},
	}
	// int16: -32768 to 32767
	execResultInt16 := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: int16(-32768), In: true}},
			results: map[string]interface{}{"num1": int16(-32768)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int16(-128), In: true}},
			results: map[string]interface{}{"num1": int16(-128)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int16(127), In: true}},
			results: map[string]interface{}{"num1": int16(127)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int16(32767), In: true}},
			results: map[string]interface{}{"num1": int16(32767)},
		},
	}
	// int: -2147483648 to 2147483647
	execResultInt := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: int(-2147483648), In: true}},
			results: map[string]interface{}{"num1": int(-2147483648)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int(-32768), In: true}},
			results: map[string]interface{}{"num1": int(-32768)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int(32767), In: true}},
			results: map[string]interface{}{"num1": int(32767)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int(2147483647), In: true}},
			results: map[string]interface{}{"num1": int(2147483647)},
		},
	}
	// int32: -2147483648 to 2147483647
	execResultInt32 := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: int32(-2147483648), In: true}},
			results: map[string]interface{}{"num1": int32(-2147483648)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int32(-32768), In: true}},
			results: map[string]interface{}{"num1": int32(-32768)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int32(32767), In: true}},
			results: map[string]interface{}{"num1": int32(32767)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int32(2147483647), In: true}},
			results: map[string]interface{}{"num1": int32(2147483647)},
		},
	}
	// int64: -9223372036854775808 to 9223372036854775807
	execResultInt64 := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(-9223372036854775808), In: true}},
			results: map[string]interface{}{"num1": int64(-9223372036854775808)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(-2147483648), In: true}},
			results: map[string]interface{}{"num1": int64(-2147483648)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(2147483647), In: true}},
			results: map[string]interface{}{"num1": int64(2147483647)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(9223372036854775807), In: true}},
			results: map[string]interface{}{"num1": int64(9223372036854775807)},
		},
	}
	// uint8: 0 to 255
	execResultUint8 := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: uint8(0), In: true}},
			results: map[string]interface{}{"num1": uint8(0)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: uint8(1), In: true}},
			results: map[string]interface{}{"num1": uint8(1)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: uint8(127), In: true}},
			results: map[string]interface{}{"num1": uint8(127)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: uint8(128), In: true}},
			results: map[string]interface{}{"num1": uint8(128)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: uint8(255), In: true}},
			results: map[string]interface{}{"num1": uint8(255)},
		},
	}
	// uint16: 0 to 65535
	execResultUint16 := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: uint16(255), In: true}},
			results: map[string]interface{}{"num1": uint16(255)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: uint16(65535), In: true}},
			results: map[string]interface{}{"num1": uint16(65535)},
		},
	}
	// uint32: 0 to 4294967295
	execResultUint32 := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: uint32(65535), In: true}},
			results: map[string]interface{}{"num1": uint32(65535)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: uint32(4294967295), In: true}},
			results: map[string]interface{}{"num1": uint32(4294967295)},
		},
	}
	// uint64: 0 to 18446744073709551615
	// for 18446744073709551615 get: get rows error: query error: sql: converting argument $1 type: uint64 values with high bit set are not supported
	execResultUint64 := []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: uint64(4294967295), In: true}},
			results: map[string]interface{}{"num1": uint64(4294967295)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: uint64(9223372036854775807), In: true}},
			results: map[string]interface{}{"num1": uint64(9223372036854775807)},
		},
	}
	// float32: sign 1 bit, exponent 8 bits, mantissa 23 bits
	execResultFloat32 := []testExecResult{
		{ // 0 00000000 00000000000000000000000
			args:    map[string]sql.Out{"num1": {Dest: math.Float32frombits(0x00000000), In: true}},
			results: map[string]interface{}{"num1": math.Float32frombits(0x00000000)},
		},
		{ // 1 00000000 00000000000000000000000
			args:    map[string]sql.Out{"num1": {Dest: math.Float32frombits(0x80000000), In: true}},
			results: map[string]interface{}{"num1": math.Float32frombits(0x80000000)},
		},
	}
	// TODO: look at being able to test at greater bit range, seem to be losing precision
	for x := 0; x < 4; x++ { // positive/negtive number and positive/negtive exponent loop
		for i := uint32(23); i < 26; i++ { // exponent[30:24] starts at 23 to have bits start at 0
			for j := uint32(24); j > 18; j-- { // mantissa [23:0] starts at 24 to have bits start at 0
				bits := uint32(0)
				if x == 1 || x == 3 {
					// negtive number has bit 31 set
					bits |= uint32(0x80000000)
				}
				if x < 2 {
					// positive exponent has bit 30 set
					bits |= uint32(0x40000000)
					for k := uint32(24); k <= i; k++ {
						// exponent starts at bit 24 and goes up
						bits |= 1 << k
					}
				} else {
					// negtive exponent starts with bits 29 to 24 set for 0
					// remove bits starting at 24
					for k := uint32(29); k > i; k-- {
						bits |= 1 << k
					}
				}
				for l := uint32(23); l >= j; l-- {
					// mantissa starts at bit 23 and goes down
					bits |= 1 << l
				}
				float := math.Float32frombits(bits)
				execResultFloat32 = append(execResultFloat32,
					testExecResult{
						args:    map[string]sql.Out{"num1": {Dest: float, In: true}},
						results: map[string]interface{}{"num1": float},
					},
				)
			}
		}
	}

	// TODO: added float64

	// INTEGER
	execResults.query = `
declare
	function GET_NUMBER(p_number INTEGER) return INTEGER as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint64
	testRunExecResults(t, execResults)
	// TODO: add float

	// INT
	execResults.query = `
declare
	function GET_NUMBER(p_number INT) return INT as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint64
	testRunExecResults(t, execResults)
	// TODO: add float

	// SMALLINT
	execResults.query = `
declare
	function GET_NUMBER(p_number SMALLINT) return SMALLINT as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint64
	testRunExecResults(t, execResults)
	// TODO: add float

	// PLS_INTEGER
	execResults.query = `
declare
	function GET_NUMBER(p_number PLS_INTEGER) return PLS_INTEGER as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	// execResults.execResults = execResultInt64
	// testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	// execResults.execResults = execResultUint32
	// testRunExecResults(t, execResults)
	// execResults.execResults = execResultUint64
	// testRunExecResults(t, execResults)
	// TODO: add float

	// NUMBER
	execResults.query = `
declare
	function GET_NUMBER(p_number NUMBER) return NUMBER as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultFloat32
	testRunExecResults(t, execResults)

	// DEC
	execResults.query = `
declare
	function GET_NUMBER(p_number DEC) return DEC as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint64
	testRunExecResults(t, execResults)
	// TODO: add float

	// DECIMAL
	execResults.query = `
declare
	function GET_NUMBER(p_number DECIMAL) return DECIMAL as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint64
	testRunExecResults(t, execResults)
	// TODO: add float

	// NUMERIC
	execResults.query = `
declare
	function GET_NUMBER(p_number NUMERIC) return NUMERIC as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint64
	testRunExecResults(t, execResults)
	// TODO: add float

	// FLOAT
	execResults.query = `
declare
	function GET_NUMBER(p_number FLOAT) return FLOAT as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultFloat32
	testRunExecResults(t, execResults)

	// REAL
	execResults.query = `
declare
	function GET_NUMBER(p_number REAL) return REAL as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint64
	testRunExecResults(t, execResults)
	execResults.execResults = execResultFloat32
	testRunExecResults(t, execResults)

	// BINARY_FLOAT
	execResults.query = `
declare
	function GET_NUMBER(p_number BINARY_FLOAT) return BINARY_FLOAT as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	// execResults.execResults = execResultInt32
	// testRunExecResults(t, execResults)
	// execResults.execResults = execResultInt64
	// testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	// execResults.execResults = execResultUint32
	// testRunExecResults(t, execResults)
	// execResults.execResults = execResultUint64
	// testRunExecResults(t, execResults)
	execResults.execResults = execResultFloat32
	testRunExecResults(t, execResults)

	// BINARY_DOUBLE
	execResults.query = `
declare
	function GET_NUMBER(p_number BINARY_DOUBLE) return BINARY_DOUBLE as
	begin
		return p_number;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`
	execResults.execResults = execResultInt8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultInt32
	testRunExecResults(t, execResults)
	// execResults.execResults = execResultInt64
	// testRunExecResults(t, execResults)
	execResults.execResults = execResultUint8
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint16
	testRunExecResults(t, execResults)
	execResults.execResults = execResultUint32
	testRunExecResults(t, execResults)
	// execResults.execResults = execResultUint64
	// testRunExecResults(t, execResults)
	execResults.execResults = execResultFloat32
	testRunExecResults(t, execResults)
}

// TestDestructiveNumberSequence checks insert sequence
func TestDestructiveNumberSequence(t *testing.T) {
	if TestDisableDatabase || TestDisableDestructive {
		t.SkipNow()
	}

	// test sequence from dual

	sequenceName := "S_A_" + TestTimeString
	err := testExec(t, "create sequence "+sequenceName+" increment by 1 start with 2147483646", nil)
	if err != nil {
		t.Fatal("create sequence error:", err)
	}

	defer testExecQuery(t, "drop sequence "+sequenceName, nil)

	queryResults := testQueryResults{
		query: "select " + sequenceName + ".NEXTVAL from dual",
		queryResults: []testQueryResult{
			{results: [][]interface{}{{float64(2147483646)}}},
			{results: [][]interface{}{{float64(2147483647)}}},
			{results: [][]interface{}{{float64(2147483648)}}},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExec(t, "alter sequence "+sequenceName+" increment by 4294967294", nil)
	if err != nil {
		t.Fatal("alter sequence error:", err)
	}

	queryResults = testQueryResults{
		query: "select " + sequenceName + ".NEXTVAL from dual",
		queryResults: []testQueryResult{
			{results: [][]interface{}{{float64(6442450942)}}},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExec(t, "alter sequence "+sequenceName+" increment by 1", nil)
	if err != nil {
		t.Fatal("alter sequence error:", err)
	}

	queryResults = testQueryResults{
		query: "select " + sequenceName + ".NEXTVAL from dual",
		queryResults: []testQueryResult{
			{results: [][]interface{}{{float64(6442450943)}}},
			{results: [][]interface{}{{float64(6442450944)}}},
			{results: [][]interface{}{{float64(6442450945)}}},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExec(t, "alter sequence "+sequenceName+" increment by 8589934588", nil)
	if err != nil {
		t.Fatal("alter sequence error:", err)
	}

	queryResults = testQueryResults{
		query: "select " + sequenceName + ".NEXTVAL from dual",
		queryResults: []testQueryResult{
			{results: [][]interface{}{{float64(15032385533)}}},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExec(t, "alter sequence "+sequenceName+" increment by 1", nil)
	if err != nil {
		t.Fatal("alter sequence error:", err)
	}

	queryResults = testQueryResults{
		query: "select " + sequenceName + ".NEXTVAL from dual",
		queryResults: []testQueryResult{
			{results: [][]interface{}{{float64(15032385534)}}},
			{results: [][]interface{}{{float64(15032385535)}}},
			{results: [][]interface{}{{float64(15032385536)}}},
		},
	}
	testRunQueryResults(t, queryResults)

	// test sequence insert into table

	sequenceName = "S_B_" + TestTimeString
	err = testExec(t, "create sequence "+sequenceName+" increment by 1 start with 2147483646", nil)
	if err != nil {
		t.Fatal("create sequence error:", err)
	}

	defer testExecQuery(t, "drop sequence "+sequenceName, nil)

	tableName := "sequence_" + TestTimeString
	err = testExec(t, "create table "+tableName+" ( A INTEGER )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	execResults := testExecResults{
		query: "insert into " + tableName + " ( A ) values (" + sequenceName + ".NEXTVAL) returning A into :num1",
		execResults: []testExecResult{
			{
				args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
				results: map[string]interface{}{"num1": int64(2147483646)},
			},
			{
				args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
				results: map[string]interface{}{"num1": int64(2147483647)},
			},
			{
				args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
				results: map[string]interface{}{"num1": int64(2147483648)},
			},
		},
	}
	testRunExecResults(t, execResults)

	err = testExec(t, "alter sequence "+sequenceName+" increment by 4294967294", nil)
	if err != nil {
		t.Fatal("alter sequence error:", err)
	}

	execResults.execResults = []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
			results: map[string]interface{}{"num1": int64(6442450942)},
		},
	}

	err = testExec(t, "alter sequence "+sequenceName+" increment by 1", nil)
	if err != nil {
		t.Fatal("alter sequence error:", err)
	}

	execResults.execResults = []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
			results: map[string]interface{}{"num1": int64(6442450943)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
			results: map[string]interface{}{"num1": int64(6442450944)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
			results: map[string]interface{}{"num1": int64(6442450945)},
		},
	}

	err = testExec(t, "alter sequence "+sequenceName+" increment by 8589934588", nil)
	if err != nil {
		t.Fatal("alter sequence error:", err)
	}

	execResults.execResults = []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
			results: map[string]interface{}{"num1": int64(15032385533)},
		},
	}

	err = testExec(t, "alter sequence "+sequenceName+" increment by 1", nil)
	if err != nil {
		t.Fatal("alter sequence error:", err)
	}

	execResults.execResults = []testExecResult{
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
			results: map[string]interface{}{"num1": int64(15032385534)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
			results: map[string]interface{}{"num1": int64(15032385535)},
		},
		{
			args:    map[string]sql.Out{"num1": {Dest: int64(0)}},
			results: map[string]interface{}{"num1": int64(15032385536)},
		},
	}

}

// TestDestructiveNumberRowsAffected checks insert RowsAffected
// Also checks count
func TestDestructiveNumberRowsAffected(t *testing.T) {
	if TestDisableDatabase || TestDisableDestructive {
		t.SkipNow()
	}

	tableName := "rows_affected_" + TestTimeString
	err := testExec(t, "create table "+tableName+" ( A INTEGER )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	ctx, cancel := context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err := TestDB.PrepareContext(ctx, "insert into "+tableName+" ( A ) values (:1)")
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	var result sql.Result
	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	result, err = stmt.ExecContext(ctx, 1)
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}

	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}

	var count int64
	count, err = result.RowsAffected()
	if err != nil {
		t.Fatal("rows affected error:", err)
	}
	expected := int64(1)
	if count != expected {
		t.Fatalf("rows affected: received: %v - expected: %v", count, expected)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err = TestDB.PrepareContext(ctx, "insert into "+tableName+" ( A ) select :1 from dual union all select :2 from dual")
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	result, err = stmt.ExecContext(ctx, 2, 3)
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}

	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}

	count, err = result.RowsAffected()
	if err != nil {
		t.Fatal("rows affected error:", err)
	}
	expected = int64(2)
	if count != expected {
		t.Fatalf("rows affected: received: %v - expected: %v", count, expected)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err = TestDB.PrepareContext(ctx, "insert into "+tableName+" ( A ) select :1 from dual union all select :2 from dual union all select :3 from dual")
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	result, err = stmt.ExecContext(ctx, 4, 5, 6)
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}

	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}

	count, err = result.RowsAffected()
	if err != nil {
		t.Fatal("rows affected error:", err)
	}
	expected = int64(3)
	if count != expected {
		t.Fatalf("rows affected: received: %v - expected: %v", count, expected)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err = TestDB.PrepareContext(ctx, "update "+tableName+" set A = :1 where A = :2")
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	result, err = stmt.ExecContext(ctx, 2, 1)
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}

	count, err = result.RowsAffected()
	if err != nil {
		t.Fatal("rows affected error:", err)
	}
	expected = int64(1)
	if count != expected {
		t.Fatalf("rows affected: received: %v - expected: %v", count, expected)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	result, err = stmt.ExecContext(ctx, 3, 2)
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}

	count, err = result.RowsAffected()
	if err != nil {
		t.Fatal("rows affected error:", err)
	}
	expected = int64(2)
	if count != expected {
		t.Fatalf("rows affected: received: %v - expected: %v", count, expected)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	result, err = stmt.ExecContext(ctx, 4, 3)
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}

	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}

	count, err = result.RowsAffected()
	if err != nil {
		t.Fatal("rows affected error:", err)
	}
	expected = int64(3)
	if count != expected {
		t.Fatalf("rows affected: received: %v - expected: %v", count, expected)
	}

	// count test
	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err = TestDB.PrepareContext(ctx, "select count(1) from "+tableName)
	if err != nil {
		cancel()
		t.Fatal("prepare error:", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		cancel()
		stmt.Close()
		t.Fatal("query error:", err)
	}

	if !rows.Next() {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("expected row")
	}

	err = rows.Scan(&count)
	if err != nil {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("scan error:", err)
	}

	if rows.Next() {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("unexpected row")
	}

	err = rows.Err()
	if err != nil {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("rows error:", err)
	}

	cancel()

	err = rows.Close()
	if err != nil {
		stmt.Close()
		t.Fatal("rows close error", err)
	}

	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}

	expected = int64(6)
	if count != expected {
		t.Fatalf("count: received: %v - expected: %v", count, expected)
	}

}

// TestNullNumber tests NullFloat64 and NullInt64
func TestNullNumber(t *testing.T) {
	if TestDisableDatabase {
		t.SkipNow()
	}

	t.Parallel()

	query := `
declare
	function GET_NUMBER(p_number1 NUMERIC) return NUMERIC as
	begin
		if p_number1 is not null then
			return p_number1;
		end if;
		return 11;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`

	ctx, cancel := context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err := TestDB.PrepareContext(ctx, query)
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	var nullFloat1 sql.NullFloat64

	nullFloat1.Float64 = 1
	nullFloat1.Valid = false

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	_, err = stmt.ExecContext(ctx, sql.Named("num1", sql.Out{Dest: &nullFloat1, In: true}))
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}
	if !nullFloat1.Valid {
		t.Fatal("nullFloat1 not Valid")
	}
	if nullFloat1.Float64 != 11 {
		t.Fatal("nullFloat1 not equal to 11")
	}

	nullFloat1.Float64 = 2

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	_, err = stmt.ExecContext(ctx, sql.Named("num1", sql.Out{Dest: &nullFloat1, In: true}))
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}
	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}
	if !nullFloat1.Valid {
		t.Fatal("nullFloat1 not Valid")
	}
	if nullFloat1.Float64 != 2 {
		t.Fatal("nullFloat1 not equal to 2")
	}

	query = `
declare
	function GET_NUMBER(p_number1 NUMERIC) return NUMERIC as
	begin
		return null;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err = TestDB.PrepareContext(ctx, query)
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	nullFloat1.Float64 = 3
	nullFloat1.Valid = true

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	_, err = stmt.ExecContext(ctx, sql.Named("num1", sql.Out{Dest: &nullFloat1, In: true}))
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}
	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}
	if nullFloat1.Valid {
		t.Fatal("nullFloat1 is Valid")
	}
	if nullFloat1.Float64 != 0 {
		t.Fatal("nullFloat1 not equal to 0")
	}

	query = `
declare
	function GET_NUMBER(p_number1 NUMERIC) return NUMERIC as
	begin
		if p_number1 is not null then
			return p_number1;
		end if;
		return 11;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err = TestDB.PrepareContext(ctx, query)
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	var nullInt1 sql.NullInt64

	nullInt1.Int64 = 1
	nullInt1.Valid = false

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	_, err = stmt.ExecContext(ctx, sql.Named("num1", sql.Out{Dest: &nullInt1, In: true}))
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}
	if !nullInt1.Valid {
		t.Fatal("nullInt1 not Valid")
	}
	if nullInt1.Int64 != 11 {
		t.Fatal("nullInt1 not equal to 11")
	}

	nullInt1.Int64 = 2

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	_, err = stmt.ExecContext(ctx, sql.Named("num1", sql.Out{Dest: &nullInt1, In: true}))
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}
	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}
	if !nullInt1.Valid {
		t.Fatal("nullInt1 not Valid")
	}
	if nullInt1.Int64 != 2 {
		t.Fatal("nullInt1 not equal to 2")
	}

	query = `
declare
	function GET_NUMBER(p_number1 NUMERIC) return NUMERIC as
	begin
		return null;
	end GET_NUMBER;
begin
	:num1 := GET_NUMBER(:num1);
end;`

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err = TestDB.PrepareContext(ctx, query)
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	nullInt1.Int64 = 3
	nullInt1.Valid = true

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	_, err = stmt.ExecContext(ctx, sql.Named("num1", sql.Out{Dest: &nullInt1, In: true}))
	cancel()
	if err != nil {
		stmt.Close()
		t.Fatal("exec error:", err)
	}
	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}
	if nullInt1.Valid {
		t.Fatal("nullInt1 is Valid")
	}
	if nullInt1.Int64 != 0 {
		t.Fatal("nullInt1 not equal to 0")
	}
}

// TestSelectDualNumberScan testing select numbers from dual with scan
func TestSelectDualNumberScan(t *testing.T) {
	if TestDisableDatabase {
		t.SkipNow()
	}

	t.Parallel()

	// float64 to float64
	ctx, cancel := context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err := TestDB.PrepareContext(ctx, "select :1 from dual")
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	var rows *sql.Rows
	var float float64
	floats := []float64{-4294967295, -65535, -255, -128, -1, 0, 1, 128, 255, 65535, 4294967295}
	for _, data := range floats {
		ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
		rows, err = stmt.QueryContext(ctx, data)
		if err != nil {
			cancel()
			stmt.Close()
			t.Fatal("query error:", err)
		}

		if !rows.Next() {
			cancel()
			rows.Close()
			stmt.Close()
			t.Fatal("expected row")
		}

		err = rows.Scan(&float)
		if err != nil {
			cancel()
			rows.Close()
			stmt.Close()
			t.Fatal("scan error:", err)
		}

		cancel()

		err = rows.Close()
		if err != nil {
			stmt.Close()
			t.Fatal("rows close error:", err)
		}

		if float != data {
			stmt.Close()
			t.Fatalf("float64 to float64 - received: %v - wanted: %v", float, data)
		}
	}

	// float64 to int64
	// note that Go database/sql convertAssignRows does not handle float64 to int64 over 6 digits as of Go ~1.11
	// Go database/sql source code: https://golang.org/src/database/sql/convert.go
	floats = []float64{-999999, -65535, -255, -128, -1, 0, 1, 128, 255, 65535, 999999}
	var aint64 int64
	for _, data := range floats {
		ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
		rows, err = stmt.QueryContext(ctx, data)
		if err != nil {
			cancel()
			stmt.Close()
			t.Fatal("query error:", err)
		}

		if !rows.Next() {
			cancel()
			rows.Close()
			stmt.Close()
			t.Fatal("expected row")
		}

		err = rows.Scan(&aint64)
		if err != nil {
			cancel()
			rows.Close()
			stmt.Close()
			t.Fatal("scan error:", err)
		}

		cancel()

		err = rows.Close()
		if err != nil {
			stmt.Close()
			t.Fatal("rows close error:", err)
		}

		if aint64 != int64(data) {
			stmt.Close()
			t.Fatalf("int64 to float64 - received: %v - wanted: %v", aint64, int64(data))
		}
	}

	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}

	// count int64
	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err = TestDB.PrepareContext(ctx, "select count(1) from dual")
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		cancel()
		stmt.Close()
		t.Fatal("query error:", err)
	}

	if !rows.Next() {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("expected row")
	}

	err = rows.Scan(&aint64)
	if err != nil {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("scan error:", err)
	}

	if rows.Next() {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("unexpected row")
	}

	cancel()

	err = rows.Close()
	if err != nil {
		stmt.Close()
		t.Fatal("rows close error:", err)
	}

	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}

	if aint64 != 1 {
		t.Fatal("aint64 not equal to 1")
	}

}

func TestSelectCountLarge(t *testing.T) {
	// skip test because it takes too long to be run all the time
	t.SkipNow()

	if TestDisableDatabase || TestDisableDestructive {
		t.SkipNow()
	}

	tableName := "COUNT_TABLE_" + TestTimeString
	err := testExec(t, "create table "+tableName+" ( A INTEGER )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	// insert into table
	query := "insert into " + tableName + ` ( A )
select :1 from dual
union all select :2 from dual
union all select :3 from dual
union all select :4 from dual
union all select :5 from dual
union all select :6 from dual
union all select :7 from dual
union all select :8 from dual
union all select :9 from dual
union all select :10 from dual
union all select :11 from dual
union all select :12 from dual
union all select :13 from dual
union all select :14 from dual
union all select :15 from dual
union all select :16 from dual
union all select :17 from dual
union all select :18 from dual
union all select :19 from dual
union all select :20 from dual`
	ctx, cancel := context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err := TestDB.PrepareContext(ctx, query)
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	var hasError bool
	var waitGroup sync.WaitGroup
	chanGoLimit := make(chan struct{}, 100)
	waitGroup.Add(50001)
	for i := 0; i < 1000020; i += 20 {
		chanGoLimit <- struct{}{}
		go func(num int) {
			for j := 0; j < 100; j++ {
				ctx, cancel := context.WithTimeout(context.Background(), TestContextTimeout)
				_, err := stmt.ExecContext(ctx, num, num+1, num+2, num+3, num+4, num+5, num+6, num+7, num+8, num+9,
					num+10, num+11, num+12, num+13, num+14, num+15, num+16, num+17, num+18, num+19)
				cancel()
				if err == nil {
					break
				}
				if len(err.Error()) < 9 {
					break
				}
				if err.Error()[0:9] != "ORA-12519" {
					break
				}
			}
			<-chanGoLimit
			waitGroup.Done()
			if err != nil {
				hasError = true
				t.Fatal("exec error:", err)
			}
		}(i)
		if hasError {
			fmt.Println("has error at", i)
			break
		}
	}

	if !hasError {
		fmt.Println("waiting")
		waitGroup.Wait()
	}

	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}

	// count float64
	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err = TestDB.PrepareContext(ctx, "select count(1) from "+tableName)
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		cancel()
		stmt.Close()
		t.Fatal("query error:", err)
	}

	if !rows.Next() {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("expected row")
	}

	var float float64
	err = rows.Scan(&float)
	if err != nil {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("scan error:", err)
	}

	if rows.Next() {
		cancel()
		rows.Close()
		stmt.Close()
		t.Fatal("unexpected row")
	}

	cancel()

	err = rows.Close()
	if err != nil {
		stmt.Close()
		t.Fatal("rows close error:", err)
	}

	err = stmt.Close()
	if err != nil {
		t.Fatal("stmt close error:", err)
	}

	if float != 1000020 {
		t.Fatal("count not equaal to  -", float)
	}

}

func TestDestructiveNumberColumnTypes(t *testing.T) {
	if TestDisableDatabase || TestDisableDestructive {
		t.SkipNow()
	}

	tableName := "NUMBER_TYPES_" + TestTimeString
	err := testExec(t, "create table "+tableName+" ( A NUMBER(10,2), B FLOAT(20), C INTEGER )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			{12345, 234.5, 3456789},
		})
	if err != nil {
		t.Fatal("insert error:", err)
	}

	queryResults := testQueryResults{
		query: "select A, B, C from " + tableName,
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{float64(12345), float64(234.5), int64(3456789)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	ctx, cancel := context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err := TestDB.PrepareContext(ctx, "select A, B, C from "+tableName)
	cancel()
	if err != nil {
		t.Fatal("prepare error:", err)
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			t.Error("stmt close error:", err)
		}
	}()

	ctx, cancel = context.WithTimeout(context.Background(), TestContextTimeout)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		cancel()
		t.Fatal("query error", err)
	}

	defer func() {
		cancel()
		err = rows.Close()
		if err != nil {
			t.Error("rows close error", err)
		}
	}()

	var columnTypes []*sql.ColumnType
	columnTypes, err = rows.ColumnTypes()

	if len(columnTypes) != 3 {
		t.Fatal("len columnTypes not equal to 3")
	}

	// A

	columnNum := 0

	if columnTypes[columnNum].DatabaseTypeName() != "SQLT_BDOUBLE" {
		t.Error("DatabaseTypeName does not match -", columnTypes[columnNum].DatabaseTypeName())
	}

	length, ok := columnTypes[columnNum].Length()
	if length != 8 {
		t.Error("Length does not match -", length)
	}
	if ok != true {
		t.Error("Length ok does not match -", ok)
	}

	if columnTypes[columnNum].Name() != "A" {
		t.Error("Name does not match -", columnTypes[columnNum].Name())
	}

	if columnTypes[columnNum].ScanType() != typeFloat64 {
		t.Error("ScanType does not match -", columnTypes[columnNum].ScanType())
	}

	// B

	columnNum = 1

	if columnTypes[columnNum].DatabaseTypeName() != "SQLT_BDOUBLE" {
		t.Error("DatabaseTypeName does not match -", columnTypes[columnNum].DatabaseTypeName())
	}

	length, ok = columnTypes[columnNum].Length()
	if length != 8 {
		t.Error("Length does not match -", length)
	}
	if ok != true {
		t.Error("Length ok does not match -", ok)
	}

	if columnTypes[columnNum].Name() != "B" {
		t.Error("Name does not match -", columnTypes[columnNum].Name())
	}

	if columnTypes[columnNum].ScanType() != typeFloat64 {
		t.Error("ScanType does not match -", columnTypes[columnNum].ScanType())
	}

	// C

	columnNum = 2

	if columnTypes[columnNum].DatabaseTypeName() != "SQLT_INT" {
		t.Error("DatabaseTypeName does not match -", columnTypes[columnNum].DatabaseTypeName())
	}

	length, ok = columnTypes[columnNum].Length()
	if length != 8 {
		t.Error("Length does not match -", length)
	}
	if ok != true {
		t.Error("Length ok does not match -", ok)
	}

	if columnTypes[columnNum].Name() != "C" {
		t.Error("Name does not match -", columnTypes[columnNum].Name())
	}

	if columnTypes[columnNum].ScanType() != typeInt64 {
		t.Error("ScanType does not match -", columnTypes[columnNum].ScanType())
	}

}
