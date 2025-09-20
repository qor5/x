package gormx_test

import (
	"context"
	"testing"

	"github.com/qor5/x/v3/gormx"
)

var suite *gormx.TestSuite

func TestMain(m *testing.M) {
	suite = gormx.MustStartTestSuite(context.Background())
	defer suite.Stop(context.Background())
	m.Run()
}
