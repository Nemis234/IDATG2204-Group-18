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

First you need to change the content inside of the .txt file provided in folder : DATABASE_NAME.txt
This file should only contain the name of your newly created database: Replace "idatg2204" with the name of your database without the "".
After that you are ready to run the script:
```bash
go run populatedb.go
```

You can set the name of the database manually in the code if you get error reading the name from file, look for:
```Go
//Change this:
databasename = "idatg2204" //<---- REPLACE THIS MANUALLY IF YOU SEE THE ERROR OVER WHEN YOU RUN THIS CODE.

//To this:
databasename = "NAME_OF_YOUR_NEWLY_CREATED_DATABASE" //<---- REPLACE THIS MANUALLY IF YOU SEE THE ERROR OVER WHEN YOU RUN THIS CODE.
```
