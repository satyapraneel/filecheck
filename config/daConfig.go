package config

//da.json file mapper
type DAConfig struct {
	StartTime         string      `json:"start_time"`           // 9:30PM (21:30)
	EndTime           string      `json:"end_time"`             // 02:40 AM (02:40)
	FolderPath        string      `json:"folder_path"`          // /home/lmsone/DynamicAction/output/archive/
	FileNamePrefix    string      `json:"file_name_prefix"`     //landmark_
	MinFileSize       int64       `json:"min_file_size_in_kb"`  //10KB
	EmailTo           []string    `json:"email_to"`             //["xyz@gmail.com","pqr@gmail.com"]
	FileExtension     string      `json:"file_extension"`       //.csv
	NotificationType  []string    `json:"notification_type"`    //[email, slack]
	DateFormatForFile string      `json:"date_format_for_file"` //28012021
	Schedule          []Scheduler `json:"schedule"`
	EnvFile           string      `json:"env_file"`
	RootPath          string      `json:"root_path"`
}

type Scheduler struct {
	CronTime  string   `json:"cron_timing"`
	FilesList []string `json:"files_list"` //["productproperties_EG", "productattributes_UAE"]
}
