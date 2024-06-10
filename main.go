package main

import (
	"encoding/json"
	"net/http"
)

type ExaminationData struct {
	Filename            string `json:"filename"`
	DocumentType        string `json:"document_type"`
	Employee            string `json:"employee"`
	ExaminationType     string `json:"examination_type"`
	ExaminationDate     string `json:"examination_date"`
	ResultOfExamination string `json:"result_of_examination"`
	Company             string `json:"company"`
	UploadedBy          string `json:"uploaded_by"`
}

func getExaminationData(w http.ResponseWriter, r *http.Request) {
	data := []ExaminationData{
		{
			Filename:            "file1.pdf",
			DocumentType:        "PDF",
			Employee:            "John Doe",
			ExaminationType:     "Annual",
			ExaminationDate:     "2023-06-10",
			ResultOfExamination: "Passed",
			Company:             "ABC Corp",
			UploadedBy:          "Jane Smith",
		},
		{
			Filename:            "file2.docx",
			DocumentType:        "DOCX",
			Employee:            "Alice Brown",
			ExaminationType:     "Quarterly",
			ExaminationDate:     "2023-03-15",
			ResultOfExamination: "Failed",
			Company:             "XYZ Ltd",
			UploadedBy:          "Bob Johnson",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/data", getExaminationData)

	http.ListenAndServe(":8080", nil)
}
