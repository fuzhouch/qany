package query

import (
	"testing"
)

func TestQuery_SmokeTest(t *testing.T) {
	// Semantics:
	// SELECT * FROM test_csv_file LIMITS 1;
	/*
		queryExpr := Select(ALL()).Top(1)
		stmt, err := queryExpr.From(mem)
		assert.Nil(t, err)

		scanner := bufio.NewScanner(stmt.Reader())
		for scanner.Next() {
			text := scanner.Text()
			assert.NotEqual(t, "", text)
		}
	*/
}
