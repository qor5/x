package errornotifierx

import (
	"strings"
	"testing"

	"github.com/qor5/confx"

	kitlog "github.com/theplant/appkit/log"
)

func TestConfig_Validation(t *testing.T) {
	suite := confx.NewValidationSuite(t)
	logger := kitlog.Default()
	suite.RunTests([]confx.ExpectedValidation{
		{
			Name: "valid log config",
			Config: &Config{
				Kind:   KindLog,
				Logger: &logger,
			},
		},
		{
			Name: "valid airbrake config",
			Config: &Config{
				Kind:   KindAirbrake,
				Logger: &logger,
				Airbrake: AirbrakeConfig{
					ProjectID:   123,
					Token:       strings.Repeat("1", 32), // 32 chars
					Environment: "test",
				},
			},
		},
		{
			Name: "missing kind",
			Config: &Config{
				Logger: &logger,
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "Kind", Tag: "required"},
			},
		},
		{
			Name: "invalid kind",
			Config: &Config{
				Kind:   "invalid",
				Logger: &logger,
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "Kind", Tag: "oneof"},
			},
		},
		{
			Name: "missing airbrake config when kind is airbrake",
			Config: &Config{
				Kind:   KindAirbrake,
				Logger: &logger,
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "Airbrake.ProjectID", Tag: "required"},
				{Path: "Airbrake.Token", Tag: "required"},
				{Path: "Airbrake.Environment", Tag: "required"},
			},
		},
		{
			Name: "invalid token length",
			Config: &Config{
				Kind:   KindAirbrake,
				Logger: &logger,
				Airbrake: AirbrakeConfig{
					ProjectID:   123,
					Token:       "short-token",
					Environment: "test",
				},
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "Airbrake.Token", Tag: "len"},
			},
		},
		{
			Name: "airbrake config not validated when kind is log",
			Config: &Config{
				Kind:     KindLog,
				Logger:   &logger,
				Airbrake: AirbrakeConfig{
					// missing required fields but should be skipped
				},
			},
		},
	})
}
