1. To set the path of the directory while running
    JSON_FILE_PATH="/var/www/html/filecheck/sample.json" go run main.go
    For compiled version : JSON_FILE_PATH=/Users/satyapraneelholla/www/GO/filecheck/sample.json ./filecheck
    NOTE: Provide the absolute path

2. Create a json file and give the files name
    JSON_FILE_PATH="/var/www/html/filecheck/sample.json"
    Follow the da.json structure example file
    Find more details in DAConfig.go file
    NOTE: Provide the absolute path
    
    
    i. To set Date format for "date_format_for_file" key, please follow the below link:
       https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
       
    ii. Files list will contain only variable string not the full name.  
        EX: "productproperties_EG" instead of "landmark_productproperties_EG_28012022.csv"
        
    iii. "notification_type" can be "email","slack" as an array of values
    
    iv. "email_to" can have multiple values as an array


3.  If you make any changes to the json file then execution has to be stopped.
