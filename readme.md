# Precious

This package was built in Go as a command line tool to backup databases from one or more servers.

## Getting Started

These instructions will help you get the project setup on your machine/server.

To install the latest stable version, run the following command:

```sh
curl -OL https://raw.githubusercontent.com/daomor/precious/v0.1-alpha/install.sh | sh
```

## Settings

The format for the server credentials is YAML. This is the structure to use:
```yaml
servers:
- name: "docker"
  host: "127.0.0.1"
  user: "root"
  pass: "password"
  port: 3307
  databases:
  - "test"
  - "test3"
```
The `name` property is just an alias for the server. Is is for output purposes and is required. 


## In The Works

These are the things I am looking to add in the near future:
* Email notifications including the backup status.
* Backup to external locations (Google Drive, another server).
* The ability to add server and database details via command.
* Removal of backups after X amount of time.
