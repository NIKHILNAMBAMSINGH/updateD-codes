package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Employee struct {
	ID    int64  `json:"id"`
	NAME  string `json:"name"`
	PHONE string `json:"phone"`
}
type MyError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
	Status     string `json:"status"`
}

func main() {
	http.HandleFunc("/hello", func(response http.ResponseWriter, Request *http.Request) {
		empId, err := strconv.ParseInt(Request.URL.Query().Get("id"), 10, 64)
		//fmt.Println(empId)
		//fmt.Fprintf(w, "employee ID:%s\n", empId)
		/*if err != nil {
			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte("emp_id must be a number"))
			return
		}*/
		employee := map[int64]*Employee{
			100: {
				ID: 2, NAME: "Nikhil Nambam Singh", PHONE: "9856884132",
			},
			101: {
				ID: 3, NAME: "Niraj Nambam Singh", PHONE: "8974548849",
			},
		}

		if err != nil {
			userError := &MyError{
				Message:    "emp_id must be a number",
				StatusCode: http.StatusBadGateway,
				Status:     "bad request",
			}
			jsonValue, _ := json.Marshal(userError)
			response.Write(jsonValue)
			return
		}

		if newValue := employee[empId]; newValue != nil {
			jsonValue, _ := json.Marshal(newValue)
			response.Write(jsonValue)
			return
		}
		{
			NotFOUND := &MyError{
				Message:    "emp_not_found",
				StatusCode: http.StatusNotFound,
				Status:     "data not found",
			}
			jsonvalue, _ := json.Marshal(NotFOUND)
			response.Write(jsonvalue)
		}

	})
	http.ListenAndServe(":3000", nil)
}
