/*
Package enum validates your enums.

Usage

Create your enum containing struct:
	type PayloadUpdateUser struct {
		UserName string `json:"username"`
		Role 	 string `json:"role" enum:"customer;partner;employee;admin"`
	}

Than validate your request body:
	func Update(req *http.Request) {
		body := new(PayloadUpdateUser)

		if err := json.NewDecoder(req.Body).Decode(body); err != nil {
			c.Response.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := enum.Validate(body); err != nil {
			c.Response.WriteHeader(http.StatusBadRequest)
			return
		}

	...

*/
package enum
