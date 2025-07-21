package api

import (
	"net/http"

	"github.com/josevitorrodriguess/goCloud/internal/jsonutils"
)

func (api *Api) updateAvatarHandler(w http.ResponseWriter, r *http.Request) {
	// Exemplo: espera receber {"user_id": 123, "avatar_url": "url"} no body
	type reqBody struct {
		UserID    uint   `json:"user_id"`
		AvatarURL string `json:"avatar_url"`
	}
	body, err := jsonutils.DecodeJson[reqBody](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
		return
	}
	if err := api.UserUsecase.UpdateAvatar(body.UserID, body.AvatarURL); err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]string{"message": "Avatar atualizado com sucesso"})
}

func (api *Api) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	type reqBody struct {
		UserID uint `json:"user_id"`
	}
	var body reqBody
	body, err := jsonutils.DecodeJson[reqBody](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
		return
	}
	if err := api.UserUsecase.DeleteUser(body.UserID); err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]string{"message": "Usuário deletado com sucesso"})
}
