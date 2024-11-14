package common

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Sairam-04/blog-app/backend/internal/service"
	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/Sairam-04/blog-app/backend/utils"
)

type CommonHandler struct {
	commonService *service.UploadImage
}

func NewCommonHandler(commonService *service.UploadImage) *CommonHandler {
	return &CommonHandler{commonService: commonService}
}

// Upload handler to handle file uploads
func (h *CommonHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()
	dir, err := os.Getwd()
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	// get file nmae
	filename := handler.Filename
	filesDir := filepath.Join(dir, "files")
	if _, err := os.Stat(filesDir); os.IsNotExist(err) {
		err := os.MkdirAll(filesDir, 0755) // Create 'files' folder with appropriate permissions
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}
	// save temporary image in this files folder
	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// send data to module
	ctx := context.Background()
	filePathUrl, err := h.commonService.UploadService(ctx, fileLocation)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	// delete temporary email
	e := os.Remove(fileLocation)
	if e != nil {
		log.Fatal("error in deleting the file", e)
	}
	//return if done
	utils.RespondWithJSON(w, http.StatusCreated, &types.FileUploadResponse{
		GeneralResponse: types.GeneralResponse{
			Success: true,
			Message: "File Uploaded Successfully",
			Error:   "",
		},
		FilePath: filePathUrl,
	})
}
