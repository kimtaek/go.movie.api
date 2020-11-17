package services

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kimtaek/gamora/pkg/db"
	"github.com/kimtaek/gamora/pkg/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func engine() *gin.Engine {
	os.Setenv("MYSQL_DATABASE", "homestead")
	os.Setenv("MYSQL_USERNAME", "root")
	os.Setenv("MYSQL_PASSWORD", "root")

	db.Setup()

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	movies := r.Group("movies")
	{
		movies.GET("", GetMovies)
		movies.GET(":id/paginate", GetMoviesPaginate)
		movies.GET(":id", GetMovie)
		movies.POST("", PostMovie)
		movies.PUT(":id", PutMovie)
		movies.DELETE(":id", DeleteMovie)
	}
	return r
}

func TestGetMovies(t *testing.T) {
	ts := httptest.NewServer(engine())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/movies", ts.URL))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetMoviesPaginate(t *testing.T) {
	ts := httptest.NewServer(engine())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/movies/*/paginate", ts.URL))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	response, _ := ioutil.ReadAll(resp.Body)
	helper.Info(string(response))
}

func TestPostMovie(t *testing.T) {
	ts := httptest.NewServer(engine())
	defer ts.Close()

	var json = []byte(`{"title":"아수라", "description": "강력계 형사 한도경(정우성)은 이권과 성공을 위해 각종 범죄를 저지르는 악덕시장 박성배(황정민)의 뒷일을 처리해주는 대가로 돈을 받는다. 악에 계속 노출되는 사이, 말기 암 환자인 아내의 병원비를 핑계로 돈 되는 건 뭐든 하는 악인의 길로 들어서게 된 한도경. 그의 약점을 쥔 독종 검사 김차인(곽도원)과 검찰수사관 도창학(정만식)은 그를 협박하고 이용해 박성배의 비리와 범죄 혐의를 캐려 한다. 각자의 이익과 목적을 위해 한도경의 목을 짓누르는 검찰과 박성배. 그 사이 태풍의 눈처럼 되어 버린 한도경은, 자신을 친형처럼 따르는 후배 형사 문선모(주지훈)를 박성배의 수하로 들여보내고, 살아남기 위해 혈안이 된 나쁜 놈들 사이에서 서로 물지 않으면 물리는 지옥도가 펼쳐진다."}`)
	resp, err := http.Post(fmt.Sprintf("%s/movies", ts.URL), "application/json", bytes.NewBuffer(json))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetMovie(t *testing.T) {
	ts := httptest.NewServer(engine())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/movies/1", ts.URL))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestPutMovie(t *testing.T) {
	ts := httptest.NewServer(engine())
	defer ts.Close()

	client := http.Client{}

	var json = []byte(`{"title":"New World, 2012", "description": "세 남자가 가고 싶었던 서로 다른 신세계"}`)
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/movies/1", ts.URL), bytes.NewBuffer(json))
	request.Header.Add("content-type", "application/json")
	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteMovie(t *testing.T) {
	ts := httptest.NewServer(engine())
	defer ts.Close()

	client := http.Client{}

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/movies/1", ts.URL), nil)
	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
