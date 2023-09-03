package testdata

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"testing"
)

func TestSexing(t *testing.T) {
	input := "UserNameID"
	snakeCase := strcase.ToSnake(input)
	fmt.Println(snakeCase) // 输出：name_id
}
