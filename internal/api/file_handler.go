package api

import (
	"io"
	"net/http"

	"github.com/josevitorrodriguess/goCloud/internal/jsonutils"
)

func (api *Api) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao ler arquivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Erro ao ler bytes do arquivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	email, ok := r.Context().Value("user_email").(string)
	if !ok || email == "" {
		http.Error(w, "Usuário não autenticado", http.StatusUnauthorized)
		return
	}

	err = api.FileUsecase.UploadFileUC(fileBytes, handler.Filename, email)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]string{"error": "Erro ao fazer upload: " + err.Error()})
		return
	}
	jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]string{"message": "Arquivo enviado com sucesso!"})
}

func (api *Api) DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value("user_email").(string)
	if !ok || email == "" {
		http.Error(w, "Usuário não autenticado", http.StatusUnauthorized)
		return
	}

	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Nome do arquivo não fornecido", http.StatusBadRequest)
		return
	}

	fileBytes, err := api.FileUsecase.GetFileUC(email, filename)
	if err != nil {
		http.Error(w, "Erro ao baixar arquivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}

func (api *Api) ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value("user_email").(string)
	if !ok || email == "" {
		jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]string{"error": "Usuário não autenticado"})
		return
	}

	files, err := api.FileUsecase.GetAllFilesUC(email)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]string{"error": "Erro ao listar arquivos: " + err.Error()})
		return
	}
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]interface{}{"files": files})
}

func (api *Api) DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value("user_email").(string)
	if !ok || email == "" {
		jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]string{"error": "Usuário não autenticado"})
		return
	}

	filename := r.URL.Query().Get("filename")
	if filename == "" {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{"error": "Nome do arquivo não fornecido"})
		return
	}

	err := api.FileUsecase.DeleteFileUC(email, filename)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]string{"error": "Erro ao deletar arquivo: " + err.Error()})
		return
	}
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]string{"message": "Arquivo deletado com sucesso!"})
}
