package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	ProfileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type ProfileHandler struct {
	profile ProfileService.ProfileServiceClient
}

type SavingError struct{}

func (t SavingError) Error() string {
	return "Cannot save the file!"
}

func NewProfileHandler(profile ProfileService.ProfileServiceClient) *ProfileHandler {
	return &ProfileHandler{
		profile: profile,
	}
}

func SaveAvatarImage(image multipart.File, handler *multipart.FileHeader, fileErr error) (string, error) {
	returnPath := profile.MediaPath
	if fileErr != nil {
		return "", SavingError{}
	}

	buff := make([]byte, 512)
	_, err := image.Read(buff)
	_, _ = image.Seek(int64(0), 0)
	if err != nil {
		return "", SavingError{}
	}
	filetype := http.DetectContentType(buff)
	if matched, regexErr := regexp.Match("image/.*", []byte(filetype)); !matched || regexErr != nil {
		return "", SavingError{}
	}
	defer image.Close()
	uniqueName := models.RandStringRunes(25)
	fileName := uniqueName + "." + strings.Split(filetype, "/")[1]
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

	requiredProfile, profileError := t.profile.GetProfileByID(r.Context(), &ProfileService.GetProfileByUserIDRequest{UserID: profileUserID.(uint64)})

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
	_, err := t.profile.UpdateProfile(r.Context(), &ProfileService.UpdateProfileRequest{
		Profile: &ProfileService.Profile{
			UserCredentials: &ProfileService.UserProfile{
				UserID: profileUserID.(uint64),
			},
			Name:       r.FormValue(profile.NameFormName),
			Surname:    r.FormValue(profile.SurnameFormName),
			AvatarPath: avatarPath,
		},
	})

	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	// http.SetCookie(w, session)
	w.WriteHeader(http.StatusOK)
}
