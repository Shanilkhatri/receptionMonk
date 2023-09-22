package controllers

import (
	"reakgo/models"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type Env struct {
	authentication interface {
		GetUserByEmail(email string) (models.Authentication, error)
		ForgotPassword(id int32) (string, error)
		TokenVerify(token string, newPassword string) (bool, error)
	}
	data interface {
		All() ([]models.Data, error)
	}
	orders interface {
		GetOrders(models.OrderDataCondition, *sqlx.Tx) ([]models.Orders, error)
		GetParamsForFilterOrderData(models.OrderDataCondition) models.OrderDataCondition
	}
	users interface {
		PutUser(tablename string, structure models.Users) error
		PostUser(tablename string, structure models.Users) error
	}
}

var Db *Env

func init() {
	// Initialize DB
	Db = &Env{
		authentication: models.AuthenticationModel{DB: utility.Db},
		data:           models.DataModel{DB: utility.Db},
		orders:         models.OrdersModel{DB: utility.Db},
		users:          models.UserModel{DB: utility.Db},
	}
}
