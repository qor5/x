package gormx_test

import (
	"testing"

	"github.com/qor5/x/v3/gormx"
	"github.com/qor5/x/v3/gormx/postgresx"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gormx.Model
	Name      string
	Age       int
	Addresses []*Address
}

type Address struct {
	gormx.Model
	AddressLine string
	UserID      string
	User        User
}

func resetTables(t *testing.T, db *gorm.DB) {
	require.NoError(t, db.Migrator().DropTable(&User{}, &Address{}))
	require.NoError(t, db.AutoMigrate(&User{}, &Address{}))
}

func TestOmitAssociationsPlugin(t *testing.T) {
	db, err := gorm.Open(postgresx.Open(suite.DSN()))
	require.NoError(t, err)

	require.NoError(t, db.Use(gormx.OmitAssociationsPlugin))

	t.Cleanup(func() {
		require.NoError(t, db.Migrator().DropTable(&User{}, &Address{}))
	})

	t.Run("create without associations", func(t *testing.T) {
		resetTables(t, db)

		user := User{
			Name: "Alice",
			Addresses: []*Address{
				{AddressLine: "123 Street"},
				{AddressLine: "456 Avenue"},
			},
		}
		require.NoError(t, db.Create(&user).Error)
		require.NoError(t, db.Where("name = ?", "Alice").First(&user).Error)
		// Verify associations were not created
		require.ErrorIs(t, db.Where("address_line = ?", "123 Street").First(&Address{}).Error, gorm.ErrRecordNotFound)
	})

	t.Run("update without associations", func(t *testing.T) {
		resetTables(t, db)

		user := User{Name: "Bob"}
		require.NoError(t, db.Create(&user).Error)

		// Create addresses manually
		addresses := []*Address{
			{AddressLine: "123 Street", UserID: user.ID},
			{AddressLine: "456 Avenue", UserID: user.ID},
		}
		require.NoError(t, db.Create(&addresses).Error)

		// Try to update with associations
		user.Addresses = addresses
		user.Addresses[0].AddressLine = "789 Boulevard"
		require.NoError(t, db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user).Error)

		// Verify address was not updated
		var address Address
		require.NoError(t, db.Where("id = ?", user.Addresses[0].ID).First(&address).Error)
		require.Equal(t, "123 Street", address.AddressLine)
	})

	t.Run("delete without associations", func(t *testing.T) {
		resetTables(t, db)

		user := User{Name: "Charlie"}
		require.NoError(t, db.Create(&user).Error)

		// Create addresses manually
		addresses := []*Address{
			{AddressLine: "123 Street", UserID: user.ID},
			{AddressLine: "456 Avenue", UserID: user.ID},
		}
		require.NoError(t, db.Create(&addresses).Error)

		// Try to delete with Select("Addresses")
		require.NoError(t, db.Select("Addresses").Delete(&user).Error)

		// Verify addresses were not deleted
		var count int64
		require.NoError(t, db.Model(&Address{}).Where("user_id = ?", user.ID).Count(&count).Error)
		require.Equal(t, int64(2), count)

		// Try to delete with Select(clause.Associations)
		user2 := User{Name: "Charlie2"}
		require.NoError(t, db.Create(&user2).Error)

		addresses2 := []*Address{
			{AddressLine: "789 Street", UserID: user2.ID},
			{AddressLine: "012 Avenue", UserID: user2.ID},
		}
		require.NoError(t, db.Create(&addresses2).Error)

		// Even with Select(clause.Associations), associations should not be deleted
		require.NoError(t, db.Select(clause.Associations).Delete(&user2).Error)

		// Verify addresses were not deleted
		var count2 int64
		require.NoError(t, db.Model(&Address{}).Where("user_id = ?", user2.ID).Count(&count2).Error)
		require.Equal(t, int64(2), count2)
	})

	t.Run("preload still works", func(t *testing.T) {
		resetTables(t, db)

		user := User{Name: "David"}
		require.NoError(t, db.Create(&user).Error)

		addresses := []*Address{
			{AddressLine: "123 Street", UserID: user.ID},
			{AddressLine: "456 Avenue", UserID: user.ID},
		}
		require.NoError(t, db.Create(&addresses).Error)

		var loadedUser User
		require.NoError(t, db.Preload("Addresses").Where("name = ?", "David").First(&loadedUser).Error)
		require.Len(t, loadedUser.Addresses, 2)
	})
}
