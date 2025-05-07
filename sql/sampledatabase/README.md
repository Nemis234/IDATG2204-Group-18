# Populating the database with sample data
populatedb.go will populate the database with sample data.
It can also be used as a reset to revert back to the original sample data. (Some data is randomized)

## Instructions
First create a mysql database, make sure it runs locally. Im using xxamp to run apache and sql locally.
Use the .sql file provided in the folder to create the tables.

In populatedb.go change the dsn variable in the main func to use your newly created database:
```Go
//Change this:
dsn := "root:@tcp(127.0.0.1:3306)/test_project"

//To this:
dsn := "root:@tcp(127.0.0.1:3306)/NAME_OF_YOUR_NEWLY_CREATED_DATABASE"
```

After that you are ready to run the code:
```bash
go run populatedb.go
```