# httputil
golang httputil lib

### Example:

```golang

http.HandleFunc("/user/current/role", httputil.JSON(GetUserRole))

....

func GetUserRole(w http.ResponseWriter, r *http.Request) (i interface{}, err error) {

	if r.URL.Query().Get("user") == "" {
		return nil, httputil.NewClientError("invalid user")   // will be translated to http 400 
		// in case of any other error, will be translated to 500
		// otherwise, you can write the specific error code to the ResponseWriter
	}

	// do some logic to retrieve user's role

	return struct {				// will be written to http.ResponseWriter as JSON
		Username string
		Role     string
	}{
		Username: "Ali",
		Role:     "Admin",
	}, nil
}



```
