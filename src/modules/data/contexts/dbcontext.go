package contexts

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/entities"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbContext struct {
	Database *gorm.DB
}

func New(environment string) *DbContext {
	db := connectToDB(environment)
	dbContext := &DbContext{Database: db}
	dbContext.AutoMigrate()

	return dbContext
}

func connectToDB(environment string) *gorm.DB {

	// logLevel := utils.GetLogLevel()
	gormLogLevel := logger.Silent
	// if logLevel == core.LogLevels.Verbose {
	// 	gormLogLevel = logger.Info
	// } else if logLevel == core.LogLevels.Warn {
	// 	gormLogLevel = logger.Warn
	// } else if logLevel == core.LogLevels.Error || logLevel == core.LogLevels.Fatal {
	// 	gormLogLevel = logger.Error
	// }

	db, err := gorm.Open(sqlite.Open(utils.GetDBLocation(environment)), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	db.Set("gorm:json_type", "json1")

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func (s *DbContext) AutoMigrate() {
	// entities := []interface{}{
	// 	&entities.Workspace{},
	// }
	// s.define(entities)
	s.define(
		&entities.Workspace{},
		&entities.Group{},
		&entities.Settings{},
		&entities.Project{},
		&entities.Template{},
		&entities.Task{},
		&entities.Resource{},
		&entities.GenerationHistory{},
	)
}

func (s *DbContext) define(dst ...interface{}) {
	err := s.Database.AutoMigrate(dst...)
	if err != nil {
		panic(err)
	}
}
