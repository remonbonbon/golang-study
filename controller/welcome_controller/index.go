package welcome_controller

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
<h1>Hello, world!</h1>
<ul>
<li><a href="./users/1">users</a></li>
</ul>
`))
}
