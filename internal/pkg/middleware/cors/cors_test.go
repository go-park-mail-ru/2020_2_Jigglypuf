package cors

import(
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestCORSSuccessGet(t *testing.T){
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		return
	})

	testFunc := MiddlewareCORS(testHandler)
	testReq := httptest.NewRequest(http.MethodGet, "/somepath/", nil)
	testRec := httptest.NewRecorder()
	testFunc.ServeHTTP(testRec, testReq)

	assert.Equal(t, http.StatusOK, testRec.Code)
}

func TestCORSSuccessOption(t *testing.T){
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		return
	})

	testFunc := MiddlewareCORS(testHandler)
	testReq := httptest.NewRequest(http.MethodOptions, "/somepath/", nil)
	testRec := httptest.NewRecorder()
	testFunc.ServeHTTP(testRec, testReq)

	assert.Equal(t, http.StatusOK, testRec.Code)
}