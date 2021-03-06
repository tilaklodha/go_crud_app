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
	suite.repository = NewUserRepository()
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

func (suite *UserRepositoryTestSuite) TestGetAllUsers() {
	userOne := domain.NewUser("abc", "xyz", "city")
	userTwo := domain.NewUser("def", "uvw", "city")
	errOne := suite.repository.InsertUser(userOne)
	errTwo := suite.repository.InsertUser(userTwo)
	assert.Nil(suite.T(), errOne)
	assert.Nil(suite.T(), errTwo)
	expectedUser := []domain.User{*userOne, *userTwo}
	observedUser, _ := suite.repository.GetAllUser()
	assert.Equal(suite.T(), expectedUser, observedUser)
	suite.cleanUpDB()
}

func (suite *UserRepositoryTestSuite) TestDeletesUser() {
	suite.repository.db.MustExec("ALTER SEQUENCE users_id_seq RESTART 1;")
	userOne := domain.NewUser("abc", "xyz", "city")
	err := suite.repository.InsertUser(userOne)
	assert.Nil(suite.T(), err)
	err = suite.repository.DeleteUser(1)
	assert.Nil(suite.T(), err)
	_, err = suite.repository.GetUser(1)
	assert.NotNil(suite.T(), err)
	suite.cleanUpDB()
}

func (suite *UserRepositoryTestSuite) TestUpdateUser() {
	suite.repository.db.MustExec("ALTER SEQUENCE users_id_seq RESTART 1;")
	userOne := domain.NewUser("abc", "xyz", "city")
	err := suite.repository.InsertUser(userOne)
	assert.Nil(suite.T(), err)
	userUpdated := domain.NewUser("abcd", "vxyz", "city")
	updatedUser, err := suite.repository.UpdateUser(userUpdated, 1)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), updatedUser, userUpdated)
	suite.cleanUpDB()
}

func (suite *UserRepositoryTestSuite) TestFailsToUpdateUser() {
	suite.repository.db.MustExec("ALTER SEQUENCE users_id_seq RESTART 1;")
	userOne := domain.NewUser("abc", "xyz", "city")
	err := suite.repository.InsertUser(userOne)
	assert.Nil(suite.T(), err)
	userUpdated := domain.NewUser("abcd", "vxyz", "city")
	updatedUser, err := suite.repository.UpdateUser(userUpdated, 2)
	assert.NotNil(suite.T(), err)
	assert.NotEqual(suite.T(), updatedUser, userUpdated)
	suite.cleanUpDB()
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) cleanUpDB() {
	db := appcontext.GetDB()
	db.MustExec("DELETE from users")
}
