package delivery

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/mock"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/magiconair/properties/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ProfileTesting struct {
	handler          *ProfileHandler
	useCaseMock      *mock.MockManagerInterface
	goMockController *gomock.Controller
}

var (
	testingStruct *ProfileTesting = nil
)

func setup(t *testing.T) {
	testingStruct = new(ProfileTesting)
	testingStruct.goMockController = gomock.NewController(t)

	testingStruct.useCaseMock = mock.NewMockManagerInterface(testingStruct.goMockController)
	testingStruct.handler = NewProfileHandler(testingStruct.useCaseMock)
}

func teardown() {
	testingStruct.goMockController.Finish()
}

func TestGetProfileSuccess(t *testing.T) {
	setup(t)
	testReq := httptest.NewRequest(http.MethodGet, "/somepath/", nil)
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, cookieService.ContextIsAuthName, true)
	ctx = context.WithValue(ctx, cookieService.ContextUserIDName, uint64(0))
	testingStruct.useCaseMock.EXPECT().GetProfileByID(gomock.Any(),gomock.Any()).Return(&profileService.Profile{}, nil)

	testRec := httptest.NewRecorder()

	testingStruct.handler.GetProfile(testRec, testReq.WithContext(ctx), httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusOK)

	teardown()
}
func TestGetProfileUCError(t *testing.T) {
	setup(t)
	testReq := httptest.NewRequest(http.MethodGet, "/somepath/", nil)
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, cookieService.ContextIsAuthName, true)
	ctx = context.WithValue(ctx, cookieService.ContextUserIDName, uint64(0))
	testingStruct.useCaseMock.EXPECT().GetProfileByID(gomock.Any(),gomock.Any()).Return(nil, errors.New("someerror"))

	testRec := httptest.NewRecorder()

	testingStruct.handler.GetProfile(testRec, testReq.WithContext(ctx), httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusBadRequest)

	teardown()
}

func TestGetProfileInvalidMethod(t *testing.T) {
	setup(t)

	testReq := httptest.NewRequest(http.MethodPost, "/somepath", nil)
	testRec := httptest.NewRecorder()

	testingStruct.handler.GetProfile(testRec, testReq, httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusMethodNotAllowed)

	teardown()
}

func TestGetProfileNoAuth(t *testing.T) {
	setup(t)
	testReq := httptest.NewRequest(http.MethodGet, "/somepath/", nil)

	testRec := httptest.NewRecorder()

	testingStruct.handler.GetProfile(testRec, testReq, httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusUnauthorized)

	teardown()
}

func TestUpdateProfileFail(t *testing.T) {
	setup(t)
	testReq := httptest.NewRequest(http.MethodPost, "/somepath/", nil)
	testRec := httptest.NewRecorder()

	testingStruct.handler.UpdateProfile(testRec, testReq, httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusMethodNotAllowed)
	teardown()
}

func TestUpdateProfileFailNoForm(t *testing.T) {
	setup(t)
	testReq := httptest.NewRequest(http.MethodPut, "/somepath/", nil)
	testRec := httptest.NewRecorder()

	testingStruct.handler.UpdateProfile(testRec, testReq, httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusBadRequest)
	teardown()
}

func TestUpdateProfileFailBody(t *testing.T) {
	setup(t)
	testReq := httptest.NewRequest(http.MethodPut, "/somepath/", nil)
	testRec := httptest.NewRecorder()
	testReq.MultipartForm = new(multipart.Form)

	testingStruct.handler.UpdateProfile(testRec, testReq, httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusUnauthorized)
	teardown()
}

func TestUpdateProfileSuccess(t *testing.T) {
	setup(t)
	testReq := httptest.NewRequest(http.MethodPut, "/somepath/", nil)
	testRec := httptest.NewRecorder()
	testReq.MultipartForm = new(multipart.Form)
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, cookieService.ContextIsAuthName, true)
	ctx = context.WithValue(ctx, cookieService.ContextUserIDName, uint64(1))

	testingStruct.useCaseMock.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(nil,nil)
	testingStruct.handler.UpdateProfile(testRec, testReq.WithContext(ctx), httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusOK)
	teardown()
}

func TestUpdateProfileUCFail(t *testing.T) {
	setup(t)
	testReq := httptest.NewRequest(http.MethodPut, "/somepath/", nil)
	testRec := httptest.NewRecorder()
	testReq.MultipartForm = new(multipart.Form)
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, cookieService.ContextIsAuthName, true)
	ctx = context.WithValue(ctx, cookieService.ContextUserIDName, uint64(1))

	testingStruct.useCaseMock.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(nil,errors.New("some error"))
	testingStruct.handler.UpdateProfile(testRec, testReq.WithContext(ctx), httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusBadRequest)
	teardown()
}

func TestUpdateProfileSaveAvatarError(t *testing.T) {
	setup(t)
	testReq := httptest.NewRequest(http.MethodPut, "/somepath/", nil)
	testRec := httptest.NewRecorder()
	testReq.MultipartForm = new(multipart.Form)
	testReq.MultipartForm.File = map[string][]*multipart.FileHeader{
		profile.AvatarFormName: {
			&multipart.FileHeader{
				Filename: "some file",
			},
		},
	}
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, cookieService.ContextIsAuthName, true)
	ctx = context.WithValue(ctx, cookieService.ContextUserIDName, uint64(1))

	testingStruct.useCaseMock.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(nil,nil)
	testingStruct.handler.UpdateProfile(testRec, testReq.WithContext(ctx), httprouter.Params{})
	assert.Equal(t, testRec.Code, http.StatusOK)
	teardown()
}
