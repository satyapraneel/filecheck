package services

import (
	"filecheck/jsons"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type ValidateFileStruct struct {
	FileNotFound          string
	FileDirectory         string
	FileSizeLess          string
	FileInvalidInterval   string
	MinimumFileSize       int
	SendErrorNotification bool
}

func (app App) ValidateFiles() ValidateFileStruct {
	fileSizeLessThanThreshold := ""
	fileNotFound := ""
	fileInvalidInterval := ""
	sendNotification := false
	data := app.DAStruct
	dt := time.Now() //.AddDate(0, 0, -1)
	dateFormat := dt.Format(data.DateFormatForFile)
	startTime := getStartTime(data.StartTime)
	endTime := getEndTime(data.EndTime)
	for _, fileList := range data.FilesList {
		fileName := getFileName(data, fileList, dateFormat)
		filePath := data.FolderPath + fileName
		// CreateFile(filePath) //run this if we want to test it in local
		stats, error := fileExists(filePath)
		//check file exists
		if error != nil {
			sendNotification = true
			fileNotFound += fileName + "\n\n\n"
			continue
		}
		//check files minimum threshold
		if stats.Size() < data.MinFileSize {
			sendNotification = true
			fileSize := strconv.Itoa(int(stats.Size() / 1000))
			fileSizeLessThanThreshold += fileName + " " + fileSize + "KB\n\n\n"
			continue
		}
		//grater than endtime/smaller than start time window
		if stats.ModTime().Unix() > endTime || stats.ModTime().Unix() < startTime {
			sendNotification = true
			createdAt := stats.ModTime().Format("2006-01-02 15:04:05")
			fileInvalidInterval += fileName + " created_at:" + createdAt + "\n\n\n"
		}

	}
	return ValidateFileStruct{
		FileNotFound:          fileNotFound,
		FileSizeLess:          fileSizeLessThanThreshold,
		FileInvalidInterval:   fileInvalidInterval,
		FileDirectory:         data.FolderPath,
		MinimumFileSize:       int(data.MinFileSize / 1000),
		SendErrorNotification: sendNotification,
	}
}

func getStartTime(startTime string) int64 {
	dt := time.Now().AddDate(0, 0, -1)
	return parseTime(dt, startTime)
}

func getEndTime(endTime string) int64 {
	dt := time.Now()
	return parseTime(dt, endTime)
}

func parseTime(dt time.Time, scheduledTime string) int64 {
	d := dt.Format("2006-01-02")
	d = d + " " + scheduledTime
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, d)
	if err != nil {
		panic(err)
	}
	return t.Unix()
}

func getFileName(data jsons.DAStruct, fileList string, dateFormat string) string {
	return data.FileNamePrefix + "_" +
		fileList + "_" + dateFormat +
		"." + data.FileExtension
}

// exists returns whether the given file or directory exists
func fileExists(path string) (fs.FileInfo, error) {
	stat, err := os.Stat(path)
	if err == nil {
		return stat, nil
	}
	return nil, err
}

//This method only used for testing, to create files
func CreateFile(filePath string) {
	// to create files for testing uncomment the below code
	err := ioutil.WriteFile(filePath, []byte("Hello world"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
