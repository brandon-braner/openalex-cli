package database

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CreateTestSuite struct {
	suite.Suite
	numberOfTables  int
	numberOfIndexes int
	db              *sql.DB
}

func (suite *CreateTestSuite) SetupTest() {
	suite.db = SetupTestDb()
	// 33 tables in the schema
	suite.numberOfTables = 33
	suite.numberOfIndexes = 6
}

func (suite *CreateTestSuite) TearDownTest() {
	TeardownTestDb()
}

func (suite *CreateTestSuite) TestCreateTable() {
	CreateTables()
	var numberOfTables int
	err := suite.db.QueryRow("SELECT count(*) FROM information_schema.tables WHERE table_schema = 'openalex'").Scan(&numberOfTables)
	suite.Nil(err)
	suite.Equal(numberOfTables, suite.numberOfTables)
}

func (suite *CreateTestSuite) TestCreateIndexes() {

	CreateTables()
	CreateIndexes()

	var numberOfIndexes int
	err := suite.db.QueryRow("select count(*) from pg_indexes where schemaname = 'openalex';").Scan(&numberOfIndexes)
	suite.Nil(err)
	suite.Equal(numberOfIndexes, suite.numberOfIndexes)
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
