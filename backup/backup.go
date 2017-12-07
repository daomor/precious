package backup

import (
	"dao/precious/servers"
	"gopkg.in/ini.v1"
	"github.com/keighl/barkup"
	_ "github.com/go-sql-driver/mysql"
	"dao/precious/logger"
	"dao/utils"
)

var serversData 		*servers.ServersData
var conf 				*ini.File
var backupDestination 	string
var lg					*logger.LogService
var err					error


func Run() {

	var backupCount int

	stringUtils := utils.GetStrings()

	// Loop the servers.
	for _, s := range serversData.Servers {

		// If flag is set and this isn't the server we want, skip it.
		if(flagData.Server != "" && flagData.Server != s.Name) {
			continue
		}

		// Create backup struct.
		backup := &barkup.MySQL{
			Host: s.Host,
			Port: s.Port,
			User: s.User,
			Password: s.Password,
		}

		// Check the connection details can connect.
		conn, err := s.CheckConnection()

		// Check for connection errors.
		if(err != nil) {

			// Log the error
			lg.Log("info").Message("Server '" + s.Name + "': " + err.Error())
			continue
		}

		// If there is no connection, then skip.
		if(!conn) {

			// Log the error
			lg.Log("error").Message("Could not connect to the database for reasons unknown.")
			continue
		}

		var exported []string

		// Loop each database and override user/pass if needed.
		for _, database := range s.Databases {

			// If flag is set and this isn't the server, skip it.
			if(flagData.Server != "" && len(flagData.Databases) > 0 && !stringUtils.StringInSlice(database, flagData.Databases)) {
				continue
			}

			// Already exported that one....
			if(stringUtils.StringInSlice(database, exported)) {

				lg.Log("info").Message("Server '" + s.Name + "', database '" + database + "': already exported.")
				continue
			}

			// Add the database to the backup struct.
			backup.DB = database

			// Export the database.
			exportServiceError := exportDatabase(backup)

			if(exportServiceError != nil) {

				// Log the error
				lg.Log("error").Message("Server '" + s.Name + "', database '" + database + "': could not export.")
			}

			// Log backup.
			lg.Log("info").Message("Backing up '" + backup.DB + "(db)' from " + s.Name + "(server).")

			// Add to the exported list.
			exported = append(exported, database)

			backupCount++
		}
	}

	if(backupCount == 0) {
		lg.Log("info").Message("No backups were made. Please test your server details.")
	}
}

func exportDatabase(backup *barkup.MySQL) *barkup.Error {

	// Run the export.
	err := backup.Export().To(backupDestination, nil)

	// Check if something went wrong with the export.
	if(err != nil) {

		/*
			For some reason, the barkup package creates an sql file, even if the
			export fails. This means that the file needs clearing up on a failed
			export.
		 */

		// Clean up the file that was created and return error
		//os.Remove(export.Path + export.Filename())
		return err
	}

	return nil
}