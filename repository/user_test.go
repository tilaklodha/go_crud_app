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
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
