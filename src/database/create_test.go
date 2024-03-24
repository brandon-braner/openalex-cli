package database

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CreateTestSuite struct {
	suite.Suite
	numberOfTables int
	db             *sql.DB
}

func (suite *CreateTestSuite) SetupSuite() {
	suite.db = SetupTestDb()
	// 33 tables in the schema
	suite.numberOfTables = 33
}

func (suite *CreateTestSuite) TearDownSuite() {
	TeardownTestDb()
}

func (suite *CreateTestSuite) TestExample() {
	CreateTables()
	var numberOfTables int
	err := suite.db.QueryRow("SELECT count(*) FROM information_schema.tables WHERE table_schema = 'openalex'").Scan(&numberOfTables)
	suite.Nil(err)
	suite.Equal(numberOfTables, suite.numberOfTables)
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
