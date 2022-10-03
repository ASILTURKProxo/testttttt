package db

import (
	"e-vet/globals"
	m "e-vet/models"
)

type MigrateParams struct {
	model     interface{}
	field     string
	joinTable interface{}
}

var (
	migrateRelationList = []MigrateParams{
		//* Contracts

		//* Platform
		// {&m.Platform{}, "Hotels", &m.PlatformsHotels{}},
		// {&m.PersonnalRole{}, "Ability", &m.PlatformsHotelsRooms{}},

		// // //* Market
		// {&m.Market{}, "Markups", &m.MarketMarkups{}},
	}
	migrateModelList = []interface{}{
		&m.Ability{}, &m.Role{}, &m.User{}, &m.Module{}}

	seederModelList = []globals.Seeder{
		// &m.AirportCodes{},
		// &m.Term{},
	}
)
