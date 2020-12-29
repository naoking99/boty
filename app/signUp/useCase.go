package signup

import (
	"fmt"
)

// UseCase acts the use case of sign-up.
func UseCase(param *Param) {
	fmt.Printf("email is %s, password is %s in UseCase", param.Email, param.Password)
}
