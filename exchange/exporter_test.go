package exchange_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/qor5/x/exchange"
	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {
	initTables()
	records := []*TestExchangeModel{
		{
			ID:    1,
			Name:  "Tom",
			Age:   ptrInt(6),
			Birth: ptrTime(time.Date(1939, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			ID:    2,
			Name:  "Jerry",
			Age:   ptrInt(5),
			Birth: ptrTime(time.Date(1940, 2, 10, 0, 0, 0, 0, time.UTC)),
		},
	}
	err := db.Create(&records).Error
	if err != nil {
		panic(err)
	}

	for _, c := range []struct {
		name             string
		metas            []*exchange.Meta
		whereCondition   string
		expectCSVContent string
		expectError      error
	}{
		{
			name: "normal",
			metas: []*exchange.Meta{
				exchange.NewMeta("ID").PrimaryKey(true),
				exchange.NewMeta("Name").Header("Nameeee"),
				exchange.NewMeta("Age"),
				exchange.NewMeta("Birth"),
			},
			expectCSVContent: `ID,Nameeee,Age,Birth
1,Tom,6,1939-01-01 00:00:00 +0000 UTC
2,Jerry,5,1940-02-10 00:00:00 +0000 UTC
`,
			expectError: nil,
		},

		{
			name: "valuer",
			metas: []*exchange.Meta{
				exchange.NewMeta("ID").PrimaryKey(true),
				exchange.NewMeta("Name").Header("Nameeee"),
				exchange.NewMeta("Age"),
				exchange.NewMeta("Birth").Valuer(func(record interface{}) (string, error) {
					m := record.(*TestExchangeModel)
					b := m.Birth.Format("2006-01-02")
					return b, nil
				}),
			},
			expectCSVContent: `ID,Nameeee,Age,Birth
1,Tom,6,1939-01-01
2,Jerry,5,1940-02-10
`,
			expectError: nil,
		},
	} {
		exporter := exchange.NewExporter(&TestExchangeModel{}).
			Metas(c.metas...)

		buf := bytes.Buffer{}
		w, err := exchange.NewCSVWriter(&buf)
		assert.NoError(t, err, c.name)

		err = exporter.Exec(db, w)
		if err != nil {
			assert.Equal(t, c.expectError, err, c.name)
			continue
		}
		assert.Equal(t, c.expectCSVContent, buf.String(), c.name)
	}
}
