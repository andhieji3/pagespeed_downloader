package googleDriveLibrary

import (
	"fmt"
	"google.golang.org/api/drive/v3"
	"pagespeed_reader/pkg/configs"
	"pagespeed_reader/pkg/googleauth"
	"pagespeed_reader/pkg/googlesheet"
	"strings"
	"time"
)

func GetFileId() string {
	directoryId, err := getMainDirectoryID()
	currentTime := time.Now()
	templateFileName := "0000000 - Template"
	templateId := ""
	fileName := currentTime.Format("200601") + " - Pagespeed Monitoring Report (Cron)"
	fileExist := 0
	fileId := ""

	if err != nil {
		fmt.Println(err)
	}

	childrenFileList, err := googleAuthLibrary.DriveService.Files.List().PageSize(1000).Q("'" + directoryId + "' in parents").Fields("nextPageToken, files(id, name)").Do()

	if err != nil {
		return "Unable to retrieve directory"
	}

	if len(childrenFileList.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, index := range childrenFileList.Files {
			if strings.Compare(index.Name, fileName) == 0 {
				fileId = index.Id
				fileExist++
			}

			if strings.Compare(index.Name, templateFileName) == 0 {
				templateId = index.Id
			}
		}
	}

	if fileExist == 0 {
		fileId = copyFileTemplate(templateId, fileName)
		googleSheetLibrary.ConfigureNewFile(fileId)
		shareFile(fileId)
	}

	return fileId
}

func getMainDirectoryID() (string, error) {
	directoryId := ""
	configuration := configs.GetConfiguration()
	topFileList, err := googleAuthLibrary.DriveService.Files.List().PageSize(1000).Q("mimeType = 'application/vnd.google-apps.folder'").Fields("nextPageToken, files(id, name)").Do()

	if err != nil {
		return "Unable to retrieve directory", err
	}

	if len(topFileList.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, index := range topFileList.Files {
			if strings.Compare(index.Name, configuration.GoogleDrive_Main_Directory) == 0 {
				directoryId = index.Id
			}
		}
	}

	return directoryId, err
}

func copyFileTemplate(templateId string, fileName string) string {
	newFile := &drive.File{Name: fileName}
	driveFile, err := googleAuthLibrary.DriveService.Files.Copy(templateId, newFile).Do()

	if err != nil {
		fmt.Println(err)
	}

	return driveFile.Id
}

func shareFile(fileId string) {
	configuration := configs.GetConfiguration()
	permission := &drive.Permission{Type: "user", Role: "writer", EmailAddress: configuration.Shared_Email}
	googleAuthLibrary.DriveService.Permissions.Create(fileId, permission).Do()
	fmt.Println("Shared file " + fileId + " to " + configuration.Shared_Email)
}
