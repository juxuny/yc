package cos

import "testing"

func TestParse(t *testing.T) {
	update := map[string]string{
		"db.host": "127.0.0.1",
		"db.port": "3306",
		"db.name": "name",
	}
	type TestOutput struct {
		DbHost string `cos:"db.host"`
		DbPort int    `cos:"db.port"`
		DbName string `cos:"db.name"`
	}
	var data TestOutput
	err := Parse(update, &data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data.DbHost)
	t.Log(data.DbPort)
	t.Log(data.DbName)
}
