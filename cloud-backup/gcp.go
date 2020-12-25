package main

import (
	"fmt",
	"flag",
	"os",
	"cloud.google.com/go",
	"google.golang.org/appengine/cloudsql",
	"github.com/go-sql-driver/mysql",
	"database/sql"
)

// function to parse command line arguments
func commandLineArgParser()(*string, *string, *string, *string){
	gcpProjectId := flag.String("projecid")
	gcpProjectName := flag.String("projectname")
	gcpInstanceName := flag.String("instancename")
	gcpCloudUser := flag.String("cloudusername")
	databaseType := flag.String("databasetype")
	databaseName := flag.String("databasename")
}

// google cloud sql connector
func googleCloudSqlConnector(gcpProjectId *string, gcpProjectName *string, gcpInstanceName *string, databaseType *string, databaseName *string, gcpCloudUser *string){
	db, err := sql.Open("%s", "%s@cloudsql(%s:%s)/%s", *databaseType, *gcpCloudUser, *gcpProjectId, *gcpInstanceName, *databaseName)
}

