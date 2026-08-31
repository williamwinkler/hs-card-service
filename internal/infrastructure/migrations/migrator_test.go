package migrations

import "testing"

func TestDatabaseOperationRemovesQueryText(t *testing.T) {
	if got := databaseOperation("SELECT * FROM cards WHERE name = 'sensitive'"); got != "SELECT" {
		t.Fatalf("databaseOperation() = %q, want SELECT", got)
	}
	if got := databaseOperation("   "); got != "" {
		t.Fatalf("databaseOperation() = %q, want empty", got)
	}
}
