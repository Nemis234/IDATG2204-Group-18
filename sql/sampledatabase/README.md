# Populating the database with sample data
You can either use the golang script to create and populate the database or do it manually.
Instructions are written below.

## Prerequisite for both methods
First create a mysql database, make sure it runs locally. Im using xxamp to run apache and sql locally.

### Manually instructions
This method will use a .sql provided in this folder to import a fully populated database.
If you're using phpmyadmin interface simply:
-> Go to your newly created database
-> Click on the tab "Import"
-> Under "File to import" choose the sql-file provided in the folder: sampleData.sql
-> Click on Import

### Golang script instructions
populatedb.go will populate the database with sample data.
It can also be used as a reset to revert back to the original sample data. (Some data is randomized)

Then you need to change the dsn and which database it should use.
In populatedb.go change the dsn variable in the main func to use your newly created database:
```Go
//Change this:
dsn := "root:@tcp(127.0.0.1:3306)/test_project"

//To this:
dsn := "root:@tcp(127.0.0.1:3306)/NAME_OF_YOUR_NEWLY_CREATED_DATABASE"
```

After that you are ready to run the script:
```bash
go run populatedb.go
```