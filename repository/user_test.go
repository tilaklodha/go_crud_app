package repository

import (
	"go_crud_app/appcontext"
	"go_crud_app/config"
	"go_crud_app/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	repository *userRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	config.Load()
	appcontext.Initiate()
	suite.repository = &userRepository{db: appcontext.GetDB()}
}

func (suite *UserRepositoryTestSuite) TestInsertUser() {
	userOne := domain.NewUser("abc", "xyz", "city")
	err := suite.repository.InsertUser(userOne)
	assert.Nil(suite.T(), err)
	suite.cleanUpDB()
}

func (suite *UserRepositoryTestSuite) TestGetUserWhenUserExists() {
	suite.repository.db.MustExec("ALTER SEQUENCE users_id_seq RESTART 1;")
	expectedUser := domain.NewUser("abc", "xyz", "city")
	err := suite.repository.InsertUser(expectedUser)
	assert.Nil(suite.T(), err)
	observedUser, err := suite.repository.GetUser(1)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedUser, observedUser)
	suite.cleanUpDB()
}

func (suite *UserRepositoryTestSuite) TestFailsGetUserWhenUserDoesNotExist() {
	_, err := suite.repository.GetUser(1)
	assert.NotNil(suite.T(), err)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) cleanUpDB() {
	db := appcontext.GetDB()
	db.MustExec("DELETE from users")
}
