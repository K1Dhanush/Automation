package eventhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "GoProject/model"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func SetDB(database *gorm.DB) {
	db = database
}

func GetAllReports(w http.ResponseWriter, r *http.Request) {
	var response []m.Report
	if err := db.Find(&response).Error; err != nil {
		fmt.Fprint(w, err)
		return
	}
	var count int64
	if err := db.Model(&m.Report{}).Count(&count).Error; err != nil {
		fmt.Fprint(w, "Unable to count the toatal no. of rows")
		return
	}
	data := map[string]interface{}{
		"count":   count,
		"reports": response,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println(w, "Unable to get the data")
		return
	}
}

func CreateNamespace(w http.ResponseWriter, r *http.Request) {
	var namespace m.Namespace
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&namespace)
	if err != nil {
		fmt.Println(w, "Data is not in JSON Format")
		return
	}

	if namespace.Region == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Please, Enter the region to create Namespace."})
		return
	}

	if namespace.Instance == "" {
		namespace.Instance = "Global"
	}
	var existingNamespace m.Namespace
	result := db.Where("instance = ? AND region = ?", namespace.Instance, namespace.Region).First(&existingNamespace)
	if result.RowsAffected != 0 {
		json.NewEncoder(w).Encode(map[string]string{"error": "Namespace with this instance is already present."})
		return
	}

	//If Global -- is not client.
	// if namespace.Instance != "Global" {
	// 	var count1 int64
	// 	var exitNamespace m.Namespace
	// 	_ = db.Where("instance = ? AND region = ?", "Global", namespace.Region).First(&exitNamespace).Count(&count1).Error
	// 	if count1 == 1 {
	// 		db.Delete(&exitNamespace)
	// 	}
	// }

	var count int64
	if err := db.Model(&m.Namespace{}).Where("region = ?", namespace.Region).Count(&count).Error; err != nil {
		panic("failed to count region")
	}
	generateNamespaceid(&namespace, count)

	if err := db.Create(&namespace).Error; err != nil {
		fmt.Println("failed to create/add the namespace")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"namespaceId": namespace.Id})
}
func generateNamespaceid(namespace *m.Namespace, count int64) {
	numericPart := strconv.FormatInt(count+1, 10)
	numericPart = fmt.Sprintf("%03s", numericPart)

	namespace.Id = namespace.Region + numericPart
}
