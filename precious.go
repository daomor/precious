package main

import (
	"fmt"
	"os"
	"dao/precious/backup"
	"dao/precious/servers"
	"gopkg.in/ini.v1"
	"github.com/Unknwon/log"
	"dao/precious/logger"
	"dao/utils"
)

var serversData *servers.ServersData
var conf *ini.File
var err error
var lg *logger.LogService

/*
	Things to do
	------------
	- Put into repo.
	- Write documentation.
	- Make easy to use installation script.


	Nice to haves
	-------------
	- CLI way of adding servers and databases to yaml.
		* precious add -d "server:database"
		* precious add -s "username:password@tcp(127.0.0.1:3307)/dbname" -n "newServer"


	Commands
	--------
	- precious backup (performs all backups in config)
	- precious backup --server=docker performs all backups for a server
	- precious backup --server=docker --database=test1,test2
 */

/**
 * Install options available for command line.
 */

func main() {

	// Check an argument exists.
	if(len(os.Args) > 1 && os.Args != nil) {

		// We have a command, lets see if it exists.
		switch os.Args[1] {

			// Create a .shadow file.
			case "backup":
				fmt.Println("Backup: Starting to backup!")

				// Load configuration.
				bootstrap()

				err = backup.Setup(conf, serversData, lg)
				if(err != nil) {
					lg.Log("info").Message(err.Error())
					return
				}
				backup.Run()
				break

			default:
				fmt.Println("Backup command not found.")
				break
		}

	} else {

		// We don't have
		fmt.Println("Backup requires a command.")
	}

	defer lg.Close()
}

func bootstrap() {

	/**
	 * Configs (.ini).
	 */

	// Load config.
	conf, err = ini.Load("conf/conf.ini")
	if(err != nil) {
		log.Fatal("Could not load config file.")
		os.Exit(1)
	}

	// Conf path.
	s, err := conf.GetSection("servers")
	if(err != nil) {
		log.Fatal("Your config needs a servers section.")
		os.Exit(1)
	}
	serverDataPath, err := s.GetKey("data")
	if(err != nil) {
		log.Fatal("Your config needs a data path for your server details.")
		os.Exit(1)
	}


	/**
	 * Logs.
	 */

	lg := logger.NewLogger()

	// Create (default) log.
	lg.Create("/var/log/precious/access.log")

	// Create custom log.
	lg.CreateCustom("error", "/var/log/precious/error.log")


	/**
	 * Servers data (.yaml).
	 */

	// Get new server data struct.
	serversData = &servers.ServersData{}

	// load YAML into the struct.
	err = utils.GetYaml().LoadYAMLToStruct(serverDataPath.String(), serversData)

	if(err != nil) {
		lg.Log("info").Message(err.Error())
		os.Exit(1)
	}
}
