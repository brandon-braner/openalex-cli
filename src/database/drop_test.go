package database

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DropTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (suite *DropTestSuite) SetupTest() {
	suite.db = SetupTestDb()
	CreateTables()
	CreateIndexes()

}

func (suite *DropTestSuite) TearDownTest() {
	TeardownTestDb()
}

func (suite *DropTestSuite) TestCreateTable() {
	var numberOfTables int
	DropTables()
	err := suite.db.QueryRow("SELECT count(*) FROM information_schema.tables WHERE table_schema = 'openalex'").Scan(&numberOfTables)
	suite.Nil(err)
	suite.Equal(numberOfTables, 0)
}

func (suite *DropTestSuite) TestCreateIndexes() {
	DropIndexes()
	var numberOfIndexes int
	err := suite.db.QueryRow("select count(*) from pg_indexes where schemaname = 'openalex';").Scan(&numberOfIndexes)
	suite.Nil(err)
	suite.Equal(numberOfIndexes, 0)
}

func TestDropTestSuite(t *testing.T) {
	suite.Run(t, new(DropTestSuite))
}
