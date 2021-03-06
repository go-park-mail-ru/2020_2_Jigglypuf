package tests

//
// import (
//	userRepository "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/repository"
//	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice/delivery"
//	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice/repository"
//	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice/usecase"
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"sync"
//	"testing"
// )
//
// func TestMovieServiceAPISuccessCases(t *testing.T) {
//	mutex := sync.RWMutex{}
//	movieRep := repository.NewMovieRepository(&mutex)
//	movieUC := usecase.NewMovieUseCase(movieRep)
//	authrep := userRepository.NewUserRepository(&mutex)
//	movieHandler := delivery.NewMovieHandler(movieUC, authrep)
//
//	movieList := "[{\"ID\":1,\"Name\":\"Гренландия\",\"Description\":\"Greenland description\",\"Rating\":0,\"PathToAvatar\"" +
//		":\"/media/greenland.jpg\"},{\"ID\":2,\"Name\":\"Антибеллум\",\"Description\":\"Антибеллум description\",\"Rating\"" +
//		":0,\"PathToAvatar\":\"/media/antibellum.jpg\"},{\"ID\":3,\"Name\":\"Довод\",\"Description\"" +
//		":\"Довод description\",\"Rating\":0,\"PathToAvatar\":\"/media/dovod.jpg\"},{\"ID\":4,\"Name\"" +
//		":\"Гнездо\",\"Description\":\"Гнездо description\",\"Rating\":0,\"PathToAvatar\"" +
//		":\"/media/gnezdo.jpg\"},{\"ID\":5,\"Name\":\"Сделано в Италии\",\"Description\":\"Италиан description\",\"Rating\"" +
//		":0,\"PathToAvatar\":\"/media/italian.jpg\"},{\"ID\":6,\"Name\":\"Мулан\",\"Description\":\"Мулан description\",\"Rating\"" +
//		":0,\"PathToAvatar\":\"/media/mulan.jpg\"},{\"ID\":7,\"Name\":\"Никогда всегда всегда никогда\",\"Description\"" +
//		":\"Никогда description\",\"Rating\":0,\"PathToAvatar\":\"/media/nikogda.jpg\"},{\"ID\":8,\"Name\":\"После\",\"Description\"" +
//		":\"После description\",\"Rating\":0,\"PathToAvatar\":\"/media/posle.jpg\"},{\"ID\":9,\"Name\":\"Стрельцов\",\"Description\"" +
//		":\"Стрельцов description\",\"Rating\":0,\"PathToAvatar\":\"/media/strelcov.jpg\"}]"
//	movie := "{\"ID\":1,\"Name\":\"Гренландия\",\"Description\":\"Greenland description\",\"Rating\":0,\"PathToAvatar\":\"/media/greenland.jpg\"}"
//	var testCases = []struct {
//		TestName           string
//		TestRequest        *http.Request
//		TestResponse       http.Response
//		TestResponseWriter *httptest.ResponseRecorder
//		TestHandler        func(w http.ResponseWriter, r *http.Request)
//	}{
//		{
//			"Проверяем что getmovielist возвращает список фильмов",
//			httptest.NewRequest("GET", "/getmovielist/?limit=10&page=1", nil),
//			http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(movieList))},
//			httptest.NewRecorder(),
//			movieHandler.GetMovieList,
//		},
//		{
//			"Проверяем что getmovie возвращает фильм",
//			httptest.NewRequest("GET", "/getmovie/?name=Гренландия", nil),
//			http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(movie))},
//			httptest.NewRecorder(),
//			movieHandler.GetMovie,
//		},
//	}
//
//	for _, val := range testCases {
//		val.TestHandler(val.TestResponseWriter, val.TestRequest)
//		if val.TestResponseWriter.Code != val.TestResponse.StatusCode {
//			t.Fatalf("TEST: "+val.TestName+" "+
//				"handler returned wrong status code: got %v want %v", val.TestResponseWriter.Code, val.TestResponse.StatusCode)
//		}
//		/* cin := new(models.Movie)
//		decoder := json.NewDecoder(val.TestResponse.Body)
//		cin2 := new(models.Movie)
//		decoder.Decode(cin2)
//		if json.Unmarshal(val.TestResponseWriter.Body.Bytes(), cin); cin.Name != "" && cin.Name != cin2.Name{
//			t.Fatalf("TEST: " + val.TestName + " " +
//				"handler returned wrong value: got %v want %v",cin2.Name, cin.Name)
//		} */
//	}
// }
