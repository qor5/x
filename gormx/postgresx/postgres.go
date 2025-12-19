package postgresx

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Dialector is an enhanced PostgreSQL dialector that preserves original error details
// alongside GORM error translations. It embeds the standard PostgreSQL dialector,
// inheriting all its capabilities while adding enhanced error handling.
type Dialector struct {
	*postgres.Dialector
}

// Open creates a new enhanced PostgreSQL dialector using a DSN string.
// This is a drop-in replacement for postgres.Open() that provides enhanced error handling.
func Open(dsn string) gorm.Dialector {
	return New(postgres.Config{DSN: dsn})
}

// New creates a new enhanced PostgreSQL dialector using a postgres.Config.
// This is a drop-in replacement for postgres.New() that provides enhanced error handling.
func New(conf postgres.Config) gorm.Dialector {
	return &Dialector{
		Dialector: &postgres.Dialector{Config: &conf},
	}
}

func (d *Dialector) Translate(cause error) error {
	translatedErr := d.Dialector.Translate(cause)
	if translatedErr == cause {
		return translatedErr
	}
	return errors.Join(translatedErr, cause)
}

// ConfigureTimezone returns stdlib.OptionOpenDB options that configure the connection
// to use the timezone specified in conf.RuntimeParams["timezone"]. This ensures that
// timestamp values are correctly converted to the specified timezone when scanning.
// The timezone is automatically populated by pgx.ParseConfig when parsing DSN.
func ConfigureTimezone(conf *pgx.ConnConfig) []stdlib.OptionOpenDB {
	tz := conf.RuntimeParams["timezone"]
	if tz == "" {
		return nil
	}

	return []stdlib.OptionOpenDB{
		stdlib.OptionAfterConnect(func(ctx context.Context, conn *pgx.Conn) error {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				return err
			}
			conn.TypeMap().RegisterType(&pgtype.Type{
				Name:  "timestamp",
				OID:   pgtype.TimestampOID,
				Codec: &pgtype.TimestampCodec{ScanLocation: loc},
			})
			return nil
		}),
	}
}
