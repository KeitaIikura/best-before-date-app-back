package gateway

import (
	"bbdate/internal/bbdate/domain/model"
	"bbdate/internal/bbdate/domain/repository"
	"bbdate/pkg/db"
	"bbdate/pkg/dbmodels"
	"bbdate/pkg/logging"
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserOnRDB struct {
	Reader *sql.DB
	Writer *sql.DB
}

func NewUserGateway(m db.IMySQL) repository.UserRepository {
	return &UserOnRDB{
		Reader: m.GetReaderConn(),
		Writer: m.GetWriterConn(),
	}
}

func (m *UserOnRDB) GetByID(ctx context.Context, xrid string, id int64) (*model.User, error) {

	logging.Info(xrid, "UserGateway.GetByID()")

	var mods []qm.QueryMod
	mods = append(mods, dbmodels.UserWhere.ID.EQ(id))
	user, err := dbmodels.Users(mods...).One(ctx, m.Reader)
	if err != nil {
		return nil, fmt.Errorf("err: UserGateway.GetByID: %w", err)
	}
	return m.convertToModel(user), err
}

func (m *UserOnRDB) GetByEmailAddress(ctx context.Context, xrid string, email_address string) (*model.User, error) {

	logging.Info(xrid, "UserGateway.GetByEmailAddress()")

	var mods []qm.QueryMod
	mods = append(mods, dbmodels.UserWhere.EmailAddress.EQ(email_address))
	user, err := dbmodels.Users(mods...).One(ctx, m.Reader)
	if err != nil {
		if err == sql.ErrNoRows {
			// NoRowsのときはエラーにしない
			return nil, nil
		}
		return nil, fmt.Errorf("err: UserGateway.GetByEmailAddress: %w", err)
	}
	return m.convertToModel(user), err
}

func (m *UserOnRDB) Create(ctx context.Context, xrid string, mu model.User) error {
	logging.Info(xrid, "UserGateway.Create()")

	orm := dbmodels.User{
		UserName:     mu.UserName,
		EmailAddress: mu.EmailAddress,
		Password:     mu.Password,
	}

	if err := orm.Insert(ctx, m.Writer, boil.Infer()); err != nil {
		return fmt.Errorf(" UserGateway.Create: %w", err)
	}

	return nil
}

func (m *UserOnRDB) Update(ctx context.Context, xrid string, mu model.User) error {
	logging.Info(xrid, "UserGateway.Update()")

	orm := m.convertToDbmodel(&mu)

	_, err := orm.Update(ctx, m.Writer, boil.Blacklist(dbmodels.UserColumns.CreatedAt))
	if err != nil {
		return fmt.Errorf("err: UserGateway.Update: %w", err)
	}

	return nil
}

// dbmodels ⇔ model 変換用メソッド
func (m *UserOnRDB) convertToModel(dm *dbmodels.User) *model.User {
	return &model.User{
		ID:           dm.ID,
		UserName:     dm.UserName,
		EmailAddress: dm.EmailAddress,
		Password:     dm.Password,
	}
}

func (m *UserOnRDB) convertToDbmodel(mu *model.User) dbmodels.User {
	dbm := dbmodels.User{
		UserName:     mu.UserName,
		EmailAddress: mu.EmailAddress,
		Password:     mu.Password,
	}

	if mu.ID != 0 {
		dbm.ID = mu.ID
	}

	return dbm
}
