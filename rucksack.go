package main

import (
    "flag"
    "fmt"
    "os"
    "time"
    "strconv"
    "path"
    "os/exec"
)

const (
    backupDir = "Rucksack"
    shell = "bash"
    sql_extension = ".sql"
    backslask = "/"
)

// function to parse the command line arguments
func commandLineArgParser()(*string, *string, *string){
    // using flag to get command line arguments
    databaseName := flag.String("dbname", "demodb", "Name of the database")
    databaseType := flag.String("dbtype", "mysql", "Type of the database")
    gcsBucketName := flag.String("bucketname", "demobucket", "Type the name of your gcs bucketname")
    flag.Parse()
    return databaseName, databaseType, gcsBucketName
}

// function to get current year and month
func getYearAndMonth()(string, string){
    currentDateTime := time.Now()
    // parsing current year and month and converted that to string
    year := strconv.Itoa(currentDateTime.Year())
    month := currentDateTime.Month().String()
    return year, month
}

// function to get backup path 
func getBackupPath(year string, month string) (string, string, string){
    homePath := os.Getenv("HOME")
    backupRootDir := path.Join(homePath, backupDir)
    backupYearDirPath := path.Join(backupRootDir, year)
    backupMonthDirPath := path.Join(backupYearDirPath, month)
    return backupRootDir, backupYearDirPath, backupMonthDirPath
}

// function to create backup directories
func createBackupDirs(rootDir,yearDir,monthDir string){
    os.MkdirAll(rootDir, os.ModePerm)
    fmt.Println("Ruck sack root directory created .......")
    os.MkdirAll(yearDir, os.ModePerm)
    fmt.Println("Year directory created .......")
    os.MkdirAll(monthDir, os.ModePerm)
    fmt.Println("Month directory created .......")
}

func executeShellCommand(command string){
    backupCommand := exec.Command(shell, "-c", string(command))
    stdout, err := backupCommand.Output()
    if err != nil {
        fmt.Println(err.Error())
        return
        }
        fmt.Print(string(stdout))
    }

func GcsStorageAutomate(filepath string, bucketName *string)(string){
    gcs_push_command := fmt.Sprintf("gsutil cp %s %s", filepath, *bucketName)
    return string(gcs_push_command)
}    

func main() {

    databaseName, databaseType, gcsBucketName := commandLineArgParser()
    year, month := getYearAndMonth()
    rootDir,yearDir,monthDir := getBackupPath(year,month)
    createBackupDirs(rootDir,yearDir,monthDir)
    
    if *databaseType == "mysql" {
        mysqlbackupFileName := monthDir + backslask + *databaseName + sql_extension
        mysqlBackupCommand := fmt.Sprintf("mysqldump -u root -p %s > %s", *databaseName, mysqlbackupFileName)
        executeShellCommand(string(mysqlBackupCommand))
        gcsPushCommand := GcsStorageAutomate(mysqlbackupFileName, gcsBucketName)
        executeShellCommand(string(gcsPushCommand))
        fmt.Println(gcsPushCommand)
    } else if *databaseType == "mongodb" {
        mongobackupFileName := monthDir + backslask + *databaseName
        mongoBackupCommand := fmt.Sprintf("mongodump --db %s --out %s", *databaseName, mongobackupFileName)
        executeShellCommand(string(mongoBackupCommand))
    } 
    
    outPutString := fmt.Sprintf("Database %s is backpacked !", *databaseName)
    fmt.Println(outPutString)
    
    }