package gormx_test

import (
	"context"
	"testing"

	"github.com/qor5/x/v3/gormx/gormxtest"
)

var suite *gormxtest.TestSuite

func TestMain(m *testing.M) {
	ctx := context.Background()
	suite = gormxtest.MustStartTestSuite(ctx)
	defer func() { _ = suite.Stop(ctx) }()
	m.Run()
}
