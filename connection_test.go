package dbtconf

import "testing"

func TestConnection(t *testing.T) {
	expect := "user=user password=pass dbname=dbname sslmode=disable"
	if conn, err := Connection("config.toml.example"); conn != expect {
		if err != nil {
			t.Fatal(err)
		}
		t.Errorf("Expected %q, got %q", expect, conn)
	}
}
