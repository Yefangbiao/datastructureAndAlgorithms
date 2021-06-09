package chapter17

import "testing"

func TestTable_Insert(t *testing.T) {
	table := Table{}
	table.Insert(1)
	if table.size != 1 {
		t.Errorf("error: want: %v, got: %v", 1, table.size)
	}

	table.Insert(2)
	table.Insert(3)
	table.Insert(4)
	table.Insert(5)
	if table.size != 8 {
		t.Errorf("error: want: %v, got: %v", 8, table.size)
	}
}

func TestTable_Delete(t *testing.T) {
	table := Table{}
	table.Insert(1)
	table.Delete()
	if table.size != 0 {
		t.Errorf("error: want: %v, got: %v", 0, table.size)
	}

	table.Insert(1)
	table.Insert(1)
	table.Delete()
	table.Delete()
	if table.size != 0 {
		t.Errorf("error: want: %v, got: %v", 0, table.size)
	}

	table.Insert(1)
	table.Insert(1)
	table.Insert(1)
	table.Insert(1)
	table.Delete()
	table.Delete()
	table.Delete()
	if table.size != 1 {
		t.Errorf("error: want: %v, got: %v", 1, table.size)
	}
}
