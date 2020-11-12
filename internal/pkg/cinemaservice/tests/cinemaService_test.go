package tests

//
// import (
//	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice/delivery"
//	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice/repository"
//	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice/usecase"
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"sync"
//	"testing"
// )
//
// func TestCinemaServiceAPISuccessCases(t *testing.T) {
//	mutex := sync.RWMutex{}
//	cinemarep := repository.NewCinemaRepository(&mutex)
//	cinemaUC := usecase.NewCinemaUseCase(cinemarep)
//	cinemadelivery := delivery.NewCinemaHandler(cinemaUC)
//
//	cinemaList := "[{\"ID\":1,\"Name\":\"CinemaScope1\",\"Address\":\"Москва, Первая улица, д.1\"},{\"ID\":2,\"Name\":\"CinemaScope2\",\"Address\":\"Москва, Первая улица, д.2\"},{\"ID\":3,\"Name\":\"CinemaScope3\",\"Address\":\"Москва, Первая улица, д.3\"},{\"ID\":4,\"Name\":\"CinemaScope4\",\"Address\":\"Москва, Первая улица, д.4\"}]"
//	cinema := "{\"ID\":1,\"Name\":\"CinemaScope1\",\"Address\":\"Москва, Первая улица, д.1\"}"
//	var testCases = []struct {
//		TestName           string
//		TestRequest        *http.Request
//		TestResponse       http.Response
//		TestResponseWriter *httptest.ResponseRecorder
//		TestHandler        func(w http.ResponseWriter, r *http.Request)
//	}{
//		{
//			"Проверяем что getcinemalist отдаст порядок кинотеатров",
//			httptest.NewRequest("GET", "/getcinemalist/?limit=10&page=1", nil),
//			http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(cinemaList))},
//			httptest.NewRecorder(),
//			cinemadelivery.GetCinemaList,
//		},
//		{
//			"Проверяем что getcinema отдаст кинотеатр",
//			httptest.NewRequest("GET", "/getcinema/?name=CinemaScope1", nil),
//			http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(cinema))},
//			httptest.NewRecorder(),
//			cinemadelivery.GetCinema,
//		},
//	}
//
//	for _, val := range testCases {
//		val.TestHandler(val.TestResponseWriter, val.TestRequest)
//		if val.TestResponseWriter.Code != val.TestResponse.StatusCode {
//			t.Fatalf("TEST: "+val.TestName+" "+
//				"handler returned wrong status code: got %v want %v", val.TestResponseWriter.Code, val.TestResponse.StatusCode)
//		}
//		/* cin := new(models.Cinema)
//		decoder := json.NewDecoder(val.TestResponse.Body)
//		cin2 := new(models.Cinema)
//		decoder.Decode(cin2)
//		if json.Unmarshal(val.TestResponseWriter.Body.Bytes(), cin); cin.Name != "" && cin.Name != cin2.Name{
//			t.Fatalf("TEST: " + val.TestName + " " +
//				"handler returned wrong value: got %v want %v",cin2.Name, cin.Name)
//		} */
//	}
// }
//
// func TestCinemaServiceAPIFAILCases(t *testing.T) {
//	mutex := sync.RWMutex{}
//	cinemarep := repository.NewCinemaRepository(&mutex)
//	cinemaUC := usecase.NewCinemaUseCase(cinemarep)
//	cinemadelivery := delivery.NewCinemaHandler(cinemaUC)
//
//	var testCases = []struct {
//		TestName           string
//		TestRequest        *http.Request
//		TestResponse       http.Response
//		TestResponseWriter *httptest.ResponseRecorder
//		TestHandler        func(w http.ResponseWriter, r *http.Request)
//	}{
//		{
//			"Проверяем что выкинет ошибку при неправильном page",
//			httptest.NewRequest("GET", "/getcinemalist/?limit=10&page=2", nil),
//			http.Response{StatusCode: http.StatusBadRequest},
//			httptest.NewRecorder(),
//			cinemadelivery.GetCinemaList,
//		},
//		{
//			"Проверяем что getcinema выкинет ошибку при неправильном имени",
//			httptest.NewRequest("GET", "/getcinema/?name=CinemaScope12", nil),
//			http.Response{StatusCode: http.StatusBadRequest},
//			httptest.NewRecorder(),
//			cinemadelivery.GetCinema,
//		},
//		{
//			"Проверяем что выкинет ошибку при отсутствии get параметров",
//			httptest.NewRequest("GET", "/getcinemalist/", nil),
//			http.Response{StatusCode: http.StatusBadRequest},
//			httptest.NewRecorder(),
//			cinemadelivery.GetCinemaList,
//		},
//		{
//			"Проверяем что выкинет ошибку при неправильном формате гет запросов",
//			httptest.NewRequest("GET", "/getcinemalist/?limit=askdasd&page=2", nil),
//			http.Response{StatusCode: http.StatusBadRequest},
//			httptest.NewRecorder(),
//			cinemadelivery.GetCinemaList,
//		},
//		{
//			"Проверяем что выкинет ошибку при неправильном методе",
//			httptest.NewRequest("POST", "/getcinemalist/?limit=askdasd&page=2", nil),
//			http.Response{StatusCode: http.StatusMethodNotAllowed},
//			httptest.NewRecorder(),
//			cinemadelivery.GetCinemaList,
//		},
//		{
//			"Проверяем что getcinema выкинет ошибку при неправильном имени",
//			httptest.NewRequest("POST", "/getcinema/?name=CinemaScope12", nil),
//			http.Response{StatusCode: http.StatusMethodNotAllowed},
//			httptest.NewRecorder(),
//			cinemadelivery.GetCinema,
//		},
//	}
//
//	for _, val := range testCases {
//		val.TestHandler(val.TestResponseWriter, val.TestRequest)
//		if val.TestResponseWriter.Code != val.TestResponse.StatusCode {
//			t.Fatalf("TEST: "+val.TestName+" "+
//				"handler returned wrong status code: got %v want %v", val.TestResponseWriter.Code, val.TestResponse.StatusCode)
//		}
//	}
// }
