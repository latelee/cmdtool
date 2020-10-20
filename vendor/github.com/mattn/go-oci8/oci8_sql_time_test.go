package oci8

import (
	"context"
	"database/sql"
	"testing"
	"time"
)

// TestSelectDualNullTime checks nulls
func TestSelectDualNullTime(t *testing.T) {
	if TestDisableDatabase {
		t.SkipNow()
	}

	t.Parallel()

	// TIMESTAMP(9)
	queryResults := testQueryResults{
		query:        "select cast (null as TIMESTAMP(9)) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// TIMESTAMP(9) WITH TIME ZONE
	queryResults = testQueryResults{
		query:        "select cast (null as TIMESTAMP(9) WITH TIME ZONE) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

	// TIMESTAMP(9) WITH LOCAL TIME ZONE
	queryResults = testQueryResults{
		query:        "select cast (null as TIMESTAMP(9) WITH LOCAL TIME ZONE) from dual",
		queryResults: []testQueryResult{{results: [][]interface{}{{nil}}}},
	}
	testRunQueryResults(t, queryResults)

}

// TestSelectDualTime checks select dual for time types
func TestSelectDualTime(t *testing.T) {
	if TestDisableDatabase {
		t.SkipNow()
	}

	t.Parallel()

	queryResults := testQueryResults{}

	queryResultTime := []testQueryResult{
		{
			args:    []interface{}{time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC)},
			results: [][]interface{}{{time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC)}},
		},
	}

	queryResultTimeToTZ := make([]testQueryResult, 0, len(timeLocations))
	for i := 0; i < len(timeLocations); i++ {
		queryResultTimeToTZ = append(queryResultTimeToTZ,
			testQueryResult{
				args:    []interface{}{time.Date(2099, 1, 2, 3, 4, 5, 123456789, timeLocations[i])},
				results: [][]interface{}{{time.Date(2099, 1, 2, 3, 4, 5, 123456789, timeLocations[i])}},
			},
		)
	}

	// https://ss64.com/ora/syntax-datatypes.html

	// TIMESTAMP(9)
	queryResults.query = "select cast (:1 as TIMESTAMP(9)) from dual"
	queryResults.queryResults = queryResultTime
	testRunQueryResults(t, queryResults)

	// TIMESTAMP(9) WITH TIME ZONE
	queryResults.query = "select cast (:1 as TIMESTAMP(9) WITH TIME ZONE) from dual"
	queryResults.queryResults = queryResultTimeToTZ
	testRunQueryResults(t, queryResults)

	// TIMESTAMP(9) WITH LOCAL TIME ZONE
	queryResults.query = "select cast (:1 as TIMESTAMP(9) WITH LOCAL TIME ZONE) from dual"
	queryResults.queryResults = queryResultTimeToTZ
	testRunQueryResults(t, queryResults)

	// https://tour.golang.org/basics/11

	// Go
	queryResults.query = "select :1 from dual"
	queryResults.queryResults = queryResultTimeToTZ
	testRunQueryResults(t, queryResults)

	queryResultTimeYearToMonth := []testQueryResult{
		{
			args:    []interface{}{int64(-2)},
			results: [][]interface{}{{int64(-24)}},
		},
		{
			args:    []interface{}{int64(-1)},
			results: [][]interface{}{{int64(-12)}},
		},
		{
			args:    []interface{}{int64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{int64(1)},
			results: [][]interface{}{{int64(12)}},
		},
		{
			args:    []interface{}{int64(2)},
			results: [][]interface{}{{int64(24)}},
		},
		{
			args:    []interface{}{float64(-2.5)},
			results: [][]interface{}{{int64(-30)}},
		},
		{
			args:    []interface{}{float64(-1.25)},
			results: [][]interface{}{{int64(-15)}},
		},
		{
			args:    []interface{}{float64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{float64(1.25)},
			results: [][]interface{}{{int64(15)}},
		},
		{
			args:    []interface{}{float64(2.5)},
			results: [][]interface{}{{int64(30)}},
		},
	}

	// INTERVAL DAY TO MONTH - YEAR
	queryResults.query = "select NUMTOYMINTERVAL(:1, 'YEAR') from dual"
	queryResults.queryResults = queryResultTimeYearToMonth
	testRunQueryResults(t, queryResults)

	queryResultTimeMonthToMonth := []testQueryResult{
		{
			args:    []interface{}{int64(-2)},
			results: [][]interface{}{{int64(-2)}},
		},
		{
			args:    []interface{}{int64(-1)},
			results: [][]interface{}{{int64(-1)}},
		},
		{
			args:    []interface{}{int64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{int64(1)},
			results: [][]interface{}{{int64(1)}},
		},
		{
			args:    []interface{}{int64(2)},
			results: [][]interface{}{{int64(2)}},
		},
		{
			args:    []interface{}{float64(-2.75)},
			results: [][]interface{}{{int64(-3)}},
		},
		{
			args:    []interface{}{float64(-1.25)},
			results: [][]interface{}{{int64(-1)}},
		},
		{
			args:    []interface{}{float64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{float64(1.25)},
			results: [][]interface{}{{int64(1)}},
		},
		{
			args:    []interface{}{float64(2.75)},
			results: [][]interface{}{{int64(3)}},
		},
	}

	// INTERVAL DAY TO MONTH - MONTH
	queryResults.query = "select NUMTOYMINTERVAL(:1, 'MONTH') from dual"
	queryResults.queryResults = queryResultTimeMonthToMonth
	testRunQueryResults(t, queryResults)

	queryResultTimeDayToSecond := []testQueryResult{
		{
			args:    []interface{}{int64(-2)},
			results: [][]interface{}{{int64(-172800000000000)}},
		},
		{
			args:    []interface{}{int64(-1)},
			results: [][]interface{}{{int64(-86400000000000)}},
		},
		{
			args:    []interface{}{int64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{int64(1)},
			results: [][]interface{}{{int64(86400000000000)}},
		},
		{
			args:    []interface{}{int64(2)},
			results: [][]interface{}{{int64(172800000000000)}},
		},
		{
			args:    []interface{}{float64(-2.5)},
			results: [][]interface{}{{int64(-216000000000000)}},
		},
		{
			args:    []interface{}{float64(-1.25)},
			results: [][]interface{}{{int64(-108000000000000)}},
		},
		{
			args:    []interface{}{float64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{float64(1.25)},
			results: [][]interface{}{{int64(108000000000000)}},
		},
		{
			args:    []interface{}{float64(2.5)},
			results: [][]interface{}{{int64(216000000000000)}},
		},
	}

	// INTERVAL DAY TO SECOND - DAY
	queryResults.query = "select NUMTODSINTERVAL(:1, 'DAY') from dual"
	queryResults.queryResults = queryResultTimeDayToSecond
	testRunQueryResults(t, queryResults)

	queryResultTimeHourToSecond := []testQueryResult{
		{
			args:    []interface{}{int64(-2)},
			results: [][]interface{}{{int64(-7200000000000)}},
		},
		{
			args:    []interface{}{int64(-1)},
			results: [][]interface{}{{int64(-3600000000000)}},
		},
		{
			args:    []interface{}{int64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{int64(1)},
			results: [][]interface{}{{int64(3600000000000)}},
		},
		{
			args:    []interface{}{int64(2)},
			results: [][]interface{}{{int64(7200000000000)}},
		},
		{
			args:    []interface{}{float64(-2.5)},
			results: [][]interface{}{{int64(-9000000000000)}},
		},
		{
			args:    []interface{}{float64(-1.25)},
			results: [][]interface{}{{int64(-4500000000000)}},
		},
		{
			args:    []interface{}{float64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{float64(1.25)},
			results: [][]interface{}{{int64(4500000000000)}},
		},
		{
			args:    []interface{}{float64(2.5)},
			results: [][]interface{}{{int64(9000000000000)}},
		},
	}

	// INTERVAL DAY TO SECOND - HOUR
	queryResults.query = "select NUMTODSINTERVAL(:1, 'HOUR') from dual"
	queryResults.queryResults = queryResultTimeHourToSecond
	testRunQueryResults(t, queryResults)

	queryResultTimeMinuteToSecond := []testQueryResult{
		{
			args:    []interface{}{int64(-2)},
			results: [][]interface{}{{int64(-120000000000)}},
		},
		{
			args:    []interface{}{int64(-1)},
			results: [][]interface{}{{int64(-60000000000)}},
		},
		{
			args:    []interface{}{int64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{int64(1)},
			results: [][]interface{}{{int64(60000000000)}},
		},
		{
			args:    []interface{}{int64(2)},
			results: [][]interface{}{{int64(120000000000)}},
		},
		{
			args:    []interface{}{float64(-2.5)},
			results: [][]interface{}{{int64(-150000000000)}},
		},
		{
			args:    []interface{}{float64(-1.25)},
			results: [][]interface{}{{int64(-75000000000)}},
		},
		{
			args:    []interface{}{float64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{float64(1.25)},
			results: [][]interface{}{{int64(75000000000)}},
		},
		{
			args:    []interface{}{float64(2.5)},
			results: [][]interface{}{{int64(150000000000)}},
		},
	}

	// INTERVAL DAY TO SECOND - MINUTE
	queryResults.query = "select NUMTODSINTERVAL(:1, 'MINUTE') from dual"
	queryResults.queryResults = queryResultTimeMinuteToSecond
	testRunQueryResults(t, queryResults)

	queryResultTimeSecondToSecond := []testQueryResult{
		{
			args:    []interface{}{int64(-2)},
			results: [][]interface{}{{int64(-2000000000)}},
		},
		{
			args:    []interface{}{int64(-1)},
			results: [][]interface{}{{int64(-1000000000)}},
		},
		{
			args:    []interface{}{int64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{int64(1)},
			results: [][]interface{}{{int64(1000000000)}},
		},
		{
			args:    []interface{}{int64(2)},
			results: [][]interface{}{{int64(2000000000)}},
		},
		{
			args:    []interface{}{float64(-2.5)},
			results: [][]interface{}{{int64(-2500000000)}},
		},
		{
			args:    []interface{}{float64(-1.25)},
			results: [][]interface{}{{int64(-1250000000)}},
		},
		{
			args:    []interface{}{float64(0)},
			results: [][]interface{}{{int64(0)}},
		},
		{
			args:    []interface{}{float64(1.25)},
			results: [][]interface{}{{int64(1250000000)}},
		},
		{
			args:    []interface{}{float64(2.5)},
			results: [][]interface{}{{int64(2500000000)}},
		},
	}

	// INTERVAL DAY TO SECOND - SECOND
	queryResults.query = "select NUMTODSINTERVAL(:1, 'SECOND') from dual"
	queryResults.queryResults = queryResultTimeSecondToSecond
	testRunQueryResults(t, queryResults)
}

// TestDestructiveTime checks insert, select, update, and delete of time types
func TestDestructiveTime(t *testing.T) {
	if TestDisableDatabase || TestDisableDestructive {
		t.SkipNow()
	}

	// https://ss64.com/ora/syntax-datatypes.html

	// TIMESTAMP(9)
	tableName := "TIMESTAMP_" + TestTimeString
	err := testExec(t, "create table "+tableName+" ( A int, B TIMESTAMP(9), C TIMESTAMP(9) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	rowsTimestamp := [][]interface{}{
		{
			1,
			time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC),
			time.Date(9876, 5, 4, 3, 2, 1, 987654321, time.UTC),
		},
	}
	resultsTimestamp := [][]interface{}{
		{
			int64(1),
			time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC),
			time.Date(9876, 5, 4, 3, 2, 1, 987654321, time.UTC),
		},
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)", rowsTimestamp)
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults := testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: resultsTimestamp,
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExec(t, "delete from "+tableName+" where A = :1", []interface{}{1})
	if err != nil {
		t.Error("delete error:", err)
	}

	// TIMESTAMP(9) WITH TIME ZONE
	tableName = "TIMESTAMPWTZ_" + TestTimeString
	err = testExec(t, "create table "+tableName+" ( A int, B TIMESTAMP(9) WITH TIME ZONE, C TIMESTAMP(9) WITH TIME ZONE )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	rowsTimestamp = make([][]interface{}, len(timeLocations))
	resultsTimestamp = make([][]interface{}, len(timeLocations))

	for i := 0; i < len(timeLocations); i++ {
		rowsTimestamp[i] = []interface{}{
			i + 1,
			time.Date(2099, 1, 2, 3, 4, 5, 123456789, timeLocations[i]),
			time.Date(9876, 5, 4, 3, 2, 1, 987654321, timeLocations[i]),
		}
		resultsTimestamp[i] = []interface{}{
			int64(i + 1),
			time.Date(2099, 1, 2, 3, 4, 5, 123456789, timeLocations[i]),
			time.Date(9876, 5, 4, 3, 2, 1, 987654321, timeLocations[i]),
		}
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)", rowsTimestamp)
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: resultsTimestamp,
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{4},
			{5},
			{6},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	// TIMESTAMP(9) WITH LOCAL TIME ZONE
	tableName = "TIMESTAMPWLTZ_" + TestTimeString
	err = testExec(t, "create table "+tableName+
		" ( A int, B TIMESTAMP(9) WITH LOCAL TIME ZONE, C TIMESTAMP(9) WITH LOCAL TIME ZONE )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, :2, :3)", rowsTimestamp)
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: resultsTimestamp,
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{4},
			{5},
			{6},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	// INTERVAL YEAR TO MONTH
	tableName = "INTERVALYTM_" + TestTimeString
	err = testExec(t, "create table "+tableName+
		" ( A int, B INTERVAL YEAR TO MONTH, C INTERVAL YEAR TO MONTH )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, NUMTOYMINTERVAL(:2, 'YEAR'), NUMTOYMINTERVAL(:3, 'MONTH'))",
		[][]interface{}{
			{1, -2, -2},
			{2, -1, -1},
			{3, 1, 1},
			{4, 2, 2},
			{5, 1.25, 2.1},
			{6, 1.5, 2.9},
			{7, 2.75, 3},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(1), int64(-24), int64(-2)},
					{int64(2), int64(-12), int64(-1)},
					{int64(3), int64(12), int64(1)},
					{int64(4), int64(24), int64(2)},
					{int64(5), int64(15), int64(2)},
					{int64(6), int64(18), int64(3)},
					{int64(7), int64(33), int64(3)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{5},
			{6},
			{7},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(1), int64(-24), int64(-2)},
					{int64(2), int64(-12), int64(-1)},
					{int64(3), int64(12), int64(1)},
					{int64(4), int64(24), int64(2)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INTERVAL DAY TO SECOND
	tableName = "INTERVALDTS_" + TestTimeString
	err = testExec(t, "create table "+tableName+
		" ( A int, B INTERVAL DAY TO SECOND, C INTERVAL DAY TO SECOND )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, NUMTODSINTERVAL(:2, 'DAY'), NUMTODSINTERVAL(:3, 'HOUR'))",
		[][]interface{}{
			{1, -2, -2},
			{2, -1, -1},
			{3, 1, 1},
			{4, 2, 2},
			{5, 1.25, 1.25},
			{6, 1.5, 1.5},
			{7, 2.75, 2.75},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(1), int64(-172800000000000), int64(-7200000000000)},
					{int64(2), int64(-86400000000000), int64(-3600000000000)},
					{int64(3), int64(86400000000000), int64(3600000000000)},
					{int64(4), int64(172800000000000), int64(7200000000000)},
					{int64(5), int64(108000000000000), int64(4500000000000)},
					{int64(6), int64(129600000000000), int64(5400000000000)},
					{int64(7), int64(237600000000000), int64(9900000000000)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{5},
			{6},
			{7},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(1), int64(-172800000000000), int64(-7200000000000)},
					{int64(2), int64(-86400000000000), int64(-3600000000000)},
					{int64(3), int64(86400000000000), int64(3600000000000)},
					{int64(4), int64(172800000000000), int64(7200000000000)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExec(t, "truncate table "+tableName, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C ) values (:1, NUMTODSINTERVAL(:2, 'MINUTE'), NUMTODSINTERVAL(:3, 'SECOND'))",
		[][]interface{}{
			{1, -2, -2},
			{2, -1, -1},
			{3, 1, 1},
			{4, 2, 2},
			{5, 1.25, 1.25},
			{6, 1.5, 1.5},
			{7, 2.75, 2.75},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(1), int64(-120000000000), int64(-2000000000)},
					{int64(2), int64(-60000000000), int64(-1000000000)},
					{int64(3), int64(60000000000), int64(1000000000)},
					{int64(4), int64(120000000000), int64(2000000000)},
					{int64(5), int64(75000000000), int64(1250000000)},
					{int64(6), int64(90000000000), int64(1500000000)},
					{int64(7), int64(165000000000), int64(2750000000)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from "+tableName+" where A = :1",
		[][]interface{}{
			{5},
			{6},
			{7},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = testQueryResults{
		query: "select A, B, C from " + tableName + " order by A",
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{int64(1), int64(-120000000000), int64(-2000000000)},
					{int64(2), int64(-60000000000), int64(-1000000000)},
					{int64(3), int64(60000000000), int64(1000000000)},
					{int64(4), int64(120000000000), int64(2000000000)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)
}

func TestDestructiveTimeColumnTypes(t *testing.T) {
	if TestDisableDatabase || TestDisableDestructive {
		t.SkipNow()
	}

	tableName := "TIME_TYPES_" + TestTimeString
	err := testExec(t, "create table "+tableName+
		" ( A TIMESTAMP(9), B TIMESTAMP(9) WITH TIME ZONE, C TIMESTAMP(9) WITH LOCAL TIME ZONE, D INTERVAL YEAR TO MONTH, E INTERVAL DAY TO SECOND )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer testDropTable(t, tableName)

	err = testExecRows(t, "insert into "+tableName+" ( A, B, C, D, E ) values (:1, :2, :3, NUMTOYMINTERVAL(:4, 'MONTH'), NUMTODSINTERVAL(:5, 'HOUR'))",
		[][]interface{}{
			{time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC),
				time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC),
				time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC),
				10, 10},
		})
	if err != nil {
		t.Fatal("insert error:", err)
	}

	queryResults := testQueryResults{
		query: "select A, B, C, D, E from " + tableName,
		queryResults: []testQueryResult{
			{
				results: [][]interface{}{
					{time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC),
						time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC),
						time.Date(2099, 1, 2, 3, 4, 5, 123456789, time.UTC),
						int64(10), int64(36000000000000)},
				}},
		},
	}
	testRunQueryResults(t, queryResults)

	ctx, cancel := context.WithTimeout(context.Background(), TestContextTimeout)
	stmt, err := TestDB.PrepareContext(ctx, "select A, B, C, D, E from "+tableName)
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

	if len(columnTypes) != 5 {
		t.Fatal("len columnTypes not equal to 5")
	}

	// A

	columnNum := 0

	if columnTypes[columnNum].DatabaseTypeName() != "SQLT_TIMESTAMP" {
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

	if columnTypes[columnNum].ScanType() != typeTime {
		t.Error("ScanType does not match -", columnTypes[columnNum].ScanType())
	}

	// B

	columnNum = 1

	if columnTypes[columnNum].DatabaseTypeName() != "SQLT_TIMESTAMP_TZ" {
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

	if columnTypes[columnNum].ScanType() != typeTime {
		t.Error("ScanType does not match -", columnTypes[columnNum].ScanType())
	}

	// C

	columnNum = 2

	if columnTypes[columnNum].DatabaseTypeName() != "SQLT_TIMESTAMP_TZ" {
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

	if columnTypes[columnNum].ScanType() != typeTime {
		t.Error("ScanType does not match -", columnTypes[columnNum].ScanType())
	}

	// D

	columnNum = 3

	if columnTypes[columnNum].DatabaseTypeName() != "SQLT_INTERVAL_YM" {
		t.Error("DatabaseTypeName does not match -", columnTypes[columnNum].DatabaseTypeName())
	}

	length, ok = columnTypes[columnNum].Length()
	if length != 8 {
		t.Error("Length does not match -", length)
	}
	if ok != true {
		t.Error("Length ok does not match -", ok)
	}

	if columnTypes[columnNum].Name() != "D" {
		t.Error("Name does not match -", columnTypes[columnNum].Name())
	}

	if columnTypes[columnNum].ScanType() != typeInt64 {
		t.Error("ScanType does not match -", columnTypes[columnNum].ScanType())
	}

	// E

	columnNum = 4

	if columnTypes[columnNum].DatabaseTypeName() != "SQLT_INTERVAL_DS" {
		t.Error("DatabaseTypeName does not match -", columnTypes[columnNum].DatabaseTypeName())
	}

	length, ok = columnTypes[columnNum].Length()
	if length != 8 {
		t.Error("Length does not match -", length)
	}
	if ok != true {
		t.Error("Length ok does not match -", ok)
	}

	if columnTypes[columnNum].Name() != "E" {
		t.Error("Name does not match -", columnTypes[columnNum].Name())
	}

	if columnTypes[columnNum].ScanType() != typeInt64 {
		t.Error("ScanType does not match -", columnTypes[columnNum].ScanType())
	}

}
