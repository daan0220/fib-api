package fib

import (
	"net/http"

	echoContext "github.com/daan0220/fib-api/context"
	"github.com/daan0220/fib-api/functions/calculate"
	"github.com/daan0220/fib-api/presentation/common"
	"github.com/go-playground/validator/v10"
)

type Parameter struct {
	N uint `query:"n" validate:"required"`
}

// Get Fibonacci number
// @Summary Calculate Fibonacci
// @Description Calculate Fibonacci number based on the provided input
// @Tags Fibonacci
// @Accept json
// @Produce json
// @Param n query int true "Number to calculate Fibonacci (must be a positive integer)" format(int32) minimum(1)
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /fibonacci [get]
func Get(c echoContext.Context) error {
	var p Parameter

	// リクエストのパラメータのバインドに失敗した場合、無効な入力値としてエラーメッセージが含まれたJSONレスポンスが返される。
	if err := c.Bind(&p); err != nil {
		return common.ErrorWithJsonItem(c, http.StatusBadRequest, "Invalid input value.")
	}

	validator := validator.New()
	if err := validator.Struct(p); err != nil {
		// バリデーションエラーが発生した場合、エラーメッセージを生成してJSONレスポンスとして返す
		return common.ErrorWithJsonItem(c, http.StatusBadRequest, "Required input value.")
	}

	fib := calculate.CalculateFibonacci(p.N)

	// 正常な場合にはステータスコード200と共にフィボナッチ数列の計算結果が含まれたJSONレスポンスが返される
	return common.OKWithJsonItem(c, http.StatusOK, NewResponse(fib.String()))
}
