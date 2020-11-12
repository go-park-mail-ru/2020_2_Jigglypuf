package delivery

import (
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
)

type ProfileHandler struct {
	useCase profile.UseCase
}

type SavingError struct{}

func (t SavingError) Error() string {
	return "Cannot save the file!"
}

func NewProfileHandler(useCase profile.UseCase) *ProfileHandler {
	return &ProfileHandler{
		useCase: useCase,
	}
}

func SaveAvatarImage(image multipart.File, handler *multipart.FileHeader, fileErr error) (string, error) {
	returnPath := profile.MediaPath
	if fileErr != nil {
		return "", SavingError{}
	}

	buff := make([]byte, 512)
	_, err := image.Read(buff)
	if err != nil{
		return "",SavingError{}
	}
	filetype := http.DetectContentType(buff)
	if matched, regexErr := regexp.Match("image/.*", []byte(filetype)); !matched || regexErr != nil{
		return "",SavingError{}
	}
	defer image.Close()
	uniqueName := models.RandStringRunes(10)
	fileName := uniqueName + handler.Filename
	f, saveErr := os.OpenFile(profile.SavingPath+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if saveErr != nil {
		return "", SavingError{}
	}

	defer f.Close()
	_, copyingError := io.Copy(f, image)
	if copyingError != nil {
		return "", SavingError{}
	}
	returnPath += fileName
	return returnPath, nil
}

// Profile godoc
// @Summary GetProfile
// @Description Get Profile
// @ID profile-id
// @Param Cookie_info header string true "Cookie information"
// @Success 200 {object} models.Profile
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 401 {object} models.ServerResponse "No authorization"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Router /api/profile/ [get]
func (t *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	profileUserID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || profileUserID == nil {
		log.Println(isAuth, profileUserID)
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	requiredProfile, profileError := t.useCase.GetProfileViaID(profileUserID.(uint64))

	if profileError != nil {
		models.BadBodyHTTPResponse(&w, profileError)
		return
	}

	w.WriteHeader(http.StatusOK)
	responseProfile, _ := json.Marshal(requiredProfile)

	_, _ = w.Write(responseProfile)
}

// Profile godoc
// @Summary GetProfile
// @Description Get Profile
// @ID profile-update-id
// @Param UpdateProfileInfo formData models.ProfileFormData true "Profile update information"
// @Success 200
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 401 {object} models.ServerResponse "No authorization"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Router /api/profile/ [put]
func (t *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		models.BadMethodHTTPResponse(&w)
		return
	}

	translationError := r.ParseMultipartForm(32 << 20)

	if translationError != nil {
		models.BadBodyHTTPResponse(&w, translationError)
		return
	}

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	profileUserID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || profileUserID == nil {
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	avatarPath, savingErr := SaveAvatarImage(r.FormFile(profile.AvatarFormName))
	if savingErr != nil {
		avatarPath = ""
	}

	err := t.useCase.UpdateProfile(profileUserID.(uint64), r.FormValue(profile.NameFormName), r.FormValue(profile.SurnameFormName), avatarPath)

	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	// http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
