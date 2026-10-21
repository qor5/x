// Package gormxtest starts PostgreSQL in a container for tests: OpenContainer
// and SetupContainer for one container, TemplatePool for a private database per
// test cloned from one migrated template, and TestSuite for a container plus a
// gormx connection under a lifecycle.
//
// It was moved out of gormx so that production code importing gormx no longer
// links testcontainers, the moby Docker client and the testing package. Those
// are only ever used by tests, but while they sat in gormx's non-test files
// every binary that opened a database carried them, and every advisory against
// them counted as one against that binary.
//
// Moving over is an import change: gormx.OpenContainer, gormx.NewTemplatePool,
// gormx.MustStartTestSuite and the rest of the container, template-pool and
// test-suite API are here under the same names, with the same behaviour. Keep
// this package out of production code — importing it brings those packages
// back into the binary.
package gormxtest
