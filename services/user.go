package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kimtaek/gamora/pkg/respond"
	"movie.api/ent"
	"movie.api/ent/user"
)

// @Tags users
// @Summary 사용자 리스트
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var source = respond.Source{}
	source.Data, _ = ent.Connection().User.Query().All(c)
	respond.Data(c, &source)
}

// @Tags users
// @Summary 사용자 리스트
// @Router /users/{id}/paginate [get]
// @Param   id path string true "*"
func GetUsersPaginate(c *gin.Context) {
}

// @Tags users
// @Summary 사용자 차 리스트
// @Router /users/{id}/cars [get]
// @Param   id path string true "*"
func GetUserActors(c *gin.Context) {

}

// @Tags users
// @Summary 사용자 추가
// @Router /users [post]
// @Param   movie body ent.User true "User info"
func PostUser(c *gin.Context) {
	ford, err := ent.Connection().Car.
		Create().
		SetModel("tesla").
		Save(c)

	if err != nil {

	}

	u, err := ent.Connection().User.
		Create().
		SetEmail("google@gmail.com").
		SetBirthday("1970-01-01").
		SetNickname("Google").
		SetUsername("google").
		SetPassword("hash").
		AddCars(ford).
		Save(c)
	if err != nil {

	}

	var source = respond.Source{}
	source.Data = u
	respond.Data(c, &source)
}

// @Tags users
// @Summary 사용자 가져 오기
// @Router /users/{id} [get]
// @Param   id path int true "User.Id"
func GetUser(c *gin.Context) {
	u, err := ent.Connection().User.
		Query().
		Where(user.UsernameContains("go")).
		Only(c)
	if err != nil {

	}

	var source = respond.Source{}
	source.Data = u
	respond.Data(c, &source)
}

// @Tags users
// @Summary 단일 사용자 수정
// @Router /users/{id} [put]
// @Param   id path int true "User.Id"
func PutUser(c *gin.Context) {

}

// @Tags users
// @Summary 사용자 삭제
// @Router /users/{id} [delete]
// @Param   id path int true "User.Id"
func DeleteUser(c *gin.Context) {

}
