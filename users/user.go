package users

import (
	"fmt"
	"time"

	"github.com/DavidAlf/GoDesde0/modelos"
)

func AltaUsuario() {
	user := new(modelos.User)

	user.AddUser(1, "David", time.Now(), true)

	fmt.Println(user)
}
