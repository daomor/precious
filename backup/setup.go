package backup

import (
	"gopkg.in/ini.v1"
	"time"
	"path/filepath"
	"fmt"
	"os"
	"github.com/jessevdk/go-flags"
	"strings"
	"dao/precious/servers"
	"dao/precious/logger"
	"errors"
)

var flagData			*FlagsData
var args []string


// CLI flag options.
var options struct {

	// Server (-s, --server).
	Server string `short:"s" long:"server" description:"Specify a single server to backup."`

	// Databases (-d, --databases).
	Databases string `short:"d" long:"databases"  description:"Specify databases to backup from."`
}

// Place to store data from flag options..
type FlagsData struct {
	Server		string
	Databases	[]string
}


// Set up the package.
func Setup(c *ini.File, s *servers.ServersData, l *logger.LogService) error {

	// Populate package vars.
	conf = c
	serversData = s
	lg = l

	// Create the backup directory.
	CreateBackupDirectory(conf)

	// Get CLI options.
	err = PrepareFlagData()
	return err
}

// Creates backup directory, possibility with date folders, depending on config.
func CreateBackupDirectory(conf *ini.File) {

	// Get the backup section from config.
	backup := conf.Section("backup")

	// Get destination property.
	backupDestination = backup.Key("destination").String()

	useDateStructure, _ := backup.Key("use_date_structure").Bool()

	// Check if we should be building a date folder structure.
	if(useDateStructure) {
		t := time.Now()
		dateFolder := t.Format("2006/01/02")
		backupDestination = filepath.Clean(backupDestination)
		path := filepath.Join(backupDestination, dateFolder)

		fmt.Println(path)
		err := os.MkdirAll(path, os.ModePerm)
		if(err != nil) {
			lg.Log("info").Message(err.Error())
		} else {
			backupDestination = path
		}
	}
}

// Gets CLI options defined by user.
func PrepareFlagData() error {

	// Parse the args, returning the maintained order, without flags.
	args, err = flags.ParseArgs(&options, os.Args)

	// Start flags data struct.
	flagData = &FlagsData{}

	// Set server.
	if(options.Server != "") {
		flagData.Server = options.Server

		// Check the server exists in the config.
		s := serversData.GetServerByName(options.Server)

		if(s == nil) {
			return errors.New("There is no server in the config with the name '" + options.Server + "'.")
		}
	}

	// Set Databases (only if server was set).
	if(flagData.Server != "" && options.Databases != "") {
		flagData.Databases = strings.Split(options.Databases, ",")
	}

	return nil
}