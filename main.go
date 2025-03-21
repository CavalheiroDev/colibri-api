package main

import (
	"colibri-api/src/city"

	"github.com/colibri-project-io/colibri-sdk-go"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/cacheDB"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

func init() {
	colibri.InitializeApp()
	sqlDB.Initialize()
	cacheDB.Initialize()
}

func main() {
	restserver.AddRoutes(city.NewCityController().Routes())

	restserver.ListenAndServe()
}
