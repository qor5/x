package gormx_test

import (
	"context"
	"testing"

	"github.com/qor5/x/v3/gormx"
)

var suite *gormx.TestSuite

func TestMain(m *testing.M) {
	ctx := context.Background()
	suite = gormx.MustStartTestSuite(ctx)
	defer func() { _ = suite.Stop(ctx) }()
	m.Run()
}
