package query

var Bank = struct {
	Select string
}{
	Select: `SELECT b.*, COUNT(*) OVER () AS count FROM banks b`,
}
