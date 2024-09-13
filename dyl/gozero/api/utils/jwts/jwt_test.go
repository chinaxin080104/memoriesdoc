package jwts

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken(JwtPayLoad{
		UserID:   1,
		Username: "天堂龙井",
		Role:     1,
	}, "home", 36000)
	fmt.Println(token, err)
}

func TestParseToken(t *testing.T) {
	fmt.Println("测试jwt-引入了testing 包。")
}
