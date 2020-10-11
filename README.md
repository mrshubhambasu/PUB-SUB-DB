# PUB-SUB-DB
First of all create hotel table in DB and run the sql scripts in export directory to populate the schema 

Now make sure you set the config.yml in both pub and sub directory with the required fields

Run the publisher by using "go run pub.go" command inside pub directory, it will publish the JSON file

Run the subscriber by using "go run ." command inside pub directory, it will recieve the JSON file and insert it into DB
