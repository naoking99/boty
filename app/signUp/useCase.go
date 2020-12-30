package signup

import (
	"fmt"
)

// UseCase acts the use case of sign-up.
func UseCase(p *Param) {
	fmt.Printf("email is %s, password is %s in UseCase", p.email, p.password)
}
