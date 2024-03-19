package controllers

import (
	"DebTour/database"
	"time"
)

func ApproveAgency(adminUsername string, agencyUsername string) error {
	agency, err := database.GetAgencyByUsername(agencyUsername, database.MainDB)
	if err != nil {
		return err
	}
	agency.AuthorizeStatus = "Approved"
	agency.AuthorizeAdminUsername = adminUsername
	tim := time.Now()
	agency.ApproveTime = &tim
	err = database.UpdateAgencyByUsername(agencyUsername, agency, database.MainDB)
	if err != nil {
		return err
	}
	return nil
}
