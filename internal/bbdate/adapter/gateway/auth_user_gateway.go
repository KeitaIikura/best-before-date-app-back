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

type AuthUserOnRDB struct {
	Reader *sql.DB
	Writer *sql.DB
}

func NewAuthUserGateway(m db.IMySQL) repository.AuthUserRepository {
	return &AuthUserOnRDB{
		Reader: m.GetReaderConn(),
		Writer: m.GetWriterConn(),
	}
}

func (m *AuthUserOnRDB) GetByID(ctx context.Context, xrid string, id int64) (*model.AuthUser, error) {

	logging.Info(xrid, "AuthUserGateway.GetByID()")

	var mods []qm.QueryMod
	mods = append(mods, dbmodels.AuthUserWhere.ID.EQ(id))
	authUser, err := dbmodels.AuthUsers(mods...).One(ctx, m.Reader)
	if err != nil {
		return nil, fmt.Errorf("err: AuthUserGateway.GetByID: %w", err)
	}
	return m.convertToModel(authUser), err
}

func (m *AuthUserOnRDB) GetByEmailAddress(ctx context.Context, xrid string, email_address string) (*model.AuthUser, error) {

	logging.Info(xrid, "AuthUserGateway.GetByEmailAddress()")

	var mods []qm.QueryMod
	mods = append(mods, dbmodels.AuthUserWhere.EmailAddress.EQ(email_address))
	authUser, err := dbmodels.AuthUsers(mods...).One(ctx, m.Reader)
	if err != nil {
		if err == sql.ErrNoRows {
			// NoRowsのときはエラーにしない
			return nil, nil
		}
		return nil, fmt.Errorf("err: AuthUserGateway.GetByEmailAddress: %w", err)
	}
	return m.convertToModel(authUser), err
}

func (m *AuthUserOnRDB) Create(ctx context.Context, xrid string, mu model.AuthUser) error {
	logging.Info(xrid, "AuthUserGateway.Create()")

	orm := dbmodels.AuthUser{
		Name:         mu.UserName,
		EmailAddress: mu.EmailAddress,
		Password:     mu.Password,
	}

	if err := orm.Insert(ctx, m.Writer, boil.Infer()); err != nil {
		return fmt.Errorf(" AuthUserGateway.Create: %w", err)
	}

	return nil
}

func (m *AuthUserOnRDB) Update(ctx context.Context, xrid string, mu model.AuthUser) error {
	logging.Info(xrid, "AuthUserGateway.Update()")

	orm := m.convertToDbmodel(&mu)

	_, err := orm.Update(ctx, m.Writer, boil.Blacklist(dbmodels.AuthUserColumns.CreatedAt))
	if err != nil {
		return fmt.Errorf("err: AuthUserGateway.Update: %w", err)
	}

	return nil
}

// dbmodels ⇔ model 変換用メソッド
func (m *AuthUserOnRDB) convertToModel(dm *dbmodels.AuthUser) *model.AuthUser {
	return &model.AuthUser{
		ID:           dm.ID,
		UserName:     dm.Name,
		EmailAddress: dm.EmailAddress,
		Password:     dm.Password,
	}
}

func (m *AuthUserOnRDB) convertToDbmodel(mu *model.AuthUser) dbmodels.AuthUser {
	dbm := dbmodels.AuthUser{
		Name:         mu.UserName,
		EmailAddress: mu.EmailAddress,
		Password:     mu.Password,
	}

	if mu.ID != 0 {
		dbm.ID = mu.ID
	}

	return dbm
}
