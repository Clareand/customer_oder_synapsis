package usecase

import (
	"context"
	"fmt"

	"github.com/Clareand/customer_oder_synapsis/libs/models"
	"github.com/Clareand/customer_oder_synapsis/pkg/auth/model"
	"github.com/Clareand/customer_oder_synapsis/pkg/auth/repository"
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/v2"
)

type loginUsecase struct {
	repo repository.LoginRepo
	ctx  context.Context
}

func NewLoginRepo(repo repository.LoginRepo) LoginUsecase {
	return &loginUsecase{repo: repo}
}

func (u *loginUsecase) WithContext(echoCtx echo.Context) LoginUsecase {
	u.ctx = echoCtx.Request().Context()
	u.repo = u.repo.WithContext(echoCtx)
	return u
}

func (u *loginUsecase) NewLoginUser(req model.ReqNewLogin, ipnumber string) <-chan models.Result {
	output := make(chan models.Result)
	request := req
	go func() {
		span, _ := apm.StartSpan(u.ctx, "NewLoginUser", "usecase")
		defer span.End()
		defer close(output)

		// var checkUser bool

		checkUser := u.repo.CheckUser(req.Username)

		fmt.Println("CHECK USER : ", checkUser)

		if checkUser.IsTrue {

			fmt.Println("user :OK")

			userID := checkUser.CustomerID
			// check password di DB
			hashPassword := u.repo.GetPassword(userID)

			fmt.Println("password :", hashPassword)
			fmt.Println("request password :", request.Password)
			if hashPassword != "" {

				// check password antara di DB dan Param

				var matchingPassword = u.repo.CheckPasswordHash(request.Password, hashPassword)
				if matchingPassword {
					fmt.Println("password matching :", matchingPassword)
					// menampilkan data dari user
					userData := u.repo.GetDataUser(userID, request.RememberMe)

					fmt.Println("userdata 1: ", userData)
					if userData.CustomerID != "" {
						var createToken = u.repo.CreateToken(userData)
						fmt.Println("Creating token : ", createToken)
						if createToken != "" {
							loginResult := u.repo.ResultLogin(userData, createToken)

							resp := &models.Response{Code: 200, MessageCode: 0000, Message: "Success", Data: loginResult}

							output <- models.Result{Data: resp}
							return

						}

					} else {
						fmt.Println("1")
						resp := &models.Response{Code: 400, MessageCode: 0000}
						output <- models.Result{Data: resp, Error: &CustomError{message: "Incorrect username or password"}}
						return
					}

				} else {
					fmt.Println("2")
					resp := &models.Response{Code: 400, MessageCode: 0000}
					output <- models.Result{Data: resp, Error: &CustomError{message: "Incorrect username or password"}}
					return
				}

			} else {
				fmt.Println("hashpassword not ok")
				resp := &models.Response{Code: 400, MessageCode: 0000}
				output <- models.Result{Data: resp, Error: &CustomError{message: "Incorrect username or password"}}
				return
			}

		} else {
			fmt.Println("3")
			resp := &models.Response{Code: 400, MessageCode: 0000}
			output <- models.Result{Data: resp, Error: &CustomError{message: "Incorrect username or password"}}
		}

	}()

	return output
}

func (u *loginUsecase) RefreshToken(req model.AccessToken) <-chan models.Result {
	output := make(chan models.Result)

	go func() {
		defer close(output)

		resp := <-u.repo.RefreshToken(req)
		if resp.Error != nil {
			output <- models.Result{Error: resp.Error}
			return
		}

		output <- models.Result{Data: resp.Data}
	}()

	return output
}

func (u *loginUsecase) Logout(req model.AccessToken) <-chan models.Result {
	output := make(chan models.Result)

	go func() {
		span, _ := apm.StartSpan(u.ctx, "Logout", "handler")
		defer close(output)

		resp := <-u.repo.Logout(req)

		if resp.Error != nil {
			output <- models.Result{Error: resp.Error}
			return
		}
		span.End()

		output <- models.Result{Data: resp.Data}
	}()

	return output
}

type CustomError struct {
	message string
}

func (m *CustomError) Error() string {
	return m.message
}
