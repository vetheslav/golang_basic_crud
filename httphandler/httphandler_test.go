package httphandler_test

// func TestInsert(t *testing.T) {
// 	tt := []struct {
// 		name   string
// 		method string
// 		body   []byte
// 		status int
// 		err    string
// 	}{
// 		{name: "post request", method: "POST", body: nil, status: http.StatusCreated},
// 	}
// 	// addedObject which will used in request body.
// 	addedObject := &network.PersonalData{DocumentID: "", Name: "firstName", LastName: "secondName", Phone: "", Email: "", YearOfBirth: 1980}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			var err error
// 			tc.body, err = json.Marshal(addedObject)
// 			assert.NoError(t, err, fmt.Sprintf("couldn't marshal request body: %v", err))

// 			req, err := http.NewRequest(tc.method, "localhost:8000/", bytes.NewReader(tc.body))
// 			assert.NoError(t, err, fmt.Sprintf("couldn't create requset: %v", err))

// 			rec := httptest.NewRecorder()
// 			network.Insert(rec, req)

// 			res := rec.Result()
// 			defer res.Body.Close()

// 			if tc.err != "" {
// 				assert.Equal(t, http.StatusBadRequest, res.StatusCode, fmt.Sprintf("expected status Bad Request; got: %v", res.StatusCode))
// 				return
// 			}
// 			assert.Equal(t, tc.status, res.StatusCode, fmt.Sprintf("expected status %v; got %v", tc.status, res.StatusCode))
// 		})
// 	}
// }

// func TestList(t *testing.T) {
// 	var expectedObject *network.PersonalData

// 	tt := []struct {
// 		name   string
// 		method string
// 		body   []byte
// 		status int
// 		err    string
// 	}{
// 		{name: "get request", method: "GET", body: nil, status: http.StatusOK},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			req, err := http.NewRequest(tc.method, "localhost:8000/", bytes.NewReader(tc.body))
// 			assert.NoError(t, err, fmt.Sprintf("couldn't create requset: %v", err))

// 			rec := httptest.NewRecorder()
// 			network.List(rec, req)

// 			res := rec.Result()
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			assert.NoError(t, err, fmt.Sprintf("couldn't read responce body: %v", err))

// 			if tc.err != "" {
// 				assert.Equal(t, http.StatusBadRequest, res.StatusCode, fmt.Sprintf("expected status Bad Request; got: %v", res.StatusCode))
// 				assert.Equal(t, tc.err, string(bytes.TrimSpace(body)), fmt.Sprintf("expected message %q; got %q", tc.err, string(bytes.TrimSpace(body))))
// 				return
// 			}
// 			assert.Equal(t, tc.status, res.StatusCode, fmt.Sprintf("expected status %v; got %v", tc.status, res.StatusCode))
// 			assert.NoError(t, json.Unmarshal(body, &expectedObject), err.Error())
// 		})
// 	}
// }

// func TestShowListByID(t *testing.T) {
// 	var expectedObject *network.PersonalData
// 	tt := []struct {
// 		name   string
// 		method string
// 		param  string
// 		value  string
// 		body   []byte
// 		status int
// 		err    string
// 	}{
// 		{name: "get request by id", method: "GET", param: "id", value: "1", body: nil, status: http.StatusOK},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			req, err := http.NewRequest(tc.method, "localhost:8000/", bytes.NewReader(tc.body))
// 			assert.NoError(t, err, fmt.Sprintf("couldn't create requset: %v", err))

// 			// add url params to request.
// 			q := req.URL.Query()
// 			q.Add(tc.param, tc.value)
// 			req.URL.RawQuery = q.Encode()

// 			rec := httptest.NewRecorder()
// 			network.ByID(rec, req)

// 			res := rec.Result()
// 			defer res.Body.Close()

// 			body, err := ioutil.ReadAll(res.Body)
// 			assert.NoError(t, err, fmt.Sprintf("couldn't read responce body: %v", err))

// 			if tc.err != "" {
// 				assert.Equal(t, http.StatusBadRequest, res.StatusCode, fmt.Sprintf("expected status Bad Request; got: %v", res.StatusCode))
// 				assert.Equal(t, tc.err, string(bytes.TrimSpace(body)), fmt.Sprintf("expected message %q; got %q", tc.err, string(bytes.TrimSpace(body))))
// 				return
// 			}

// 			assert.Equal(t, tc.status, res.StatusCode, fmt.Sprintf("expected status %v; got %v", tc.status, res.StatusCode))
// 			assert.NoError(t, json.Unmarshal(body, &expectedObject), err.Error())

// 		})
// 	}
// }

// func TestUpdate(t *testing.T) {
// 	tt := []struct {
// 		name   string
// 		method string
// 		body   []byte
// 		status int
// 		err    string
// 	}{
// 		{name: "update request", method: "PUT", body: nil, status: http.StatusOK},
// 	}

// 	// updateO which will used in request body.
// 	updateO := &network.PersonalData{DocumentID: "", Name: "firstName", LastName: "secondName", Phone: "12345678", Email: "test@test.com", YearOfBirth: 1980}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			var err error
// 			tc.body, err = json.Marshal(updateO)
// 			assert.NoError(t, err, fmt.Sprintf("couldn't marshal request body: %v", err))

// 			req, err := http.NewRequest(tc.method, "localhost:8000/", bytes.NewReader(tc.body))
// 			assert.NoError(t, err, fmt.Sprintf("couldn't create requset: %v", err))

// 			rec := httptest.NewRecorder()
// 			network.Insert(rec, req)

// 			res := rec.Result()
// 			defer res.Body.Close()

// 			if tc.err != "" {
// 				assert.Equal(t, http.StatusBadRequest, res.StatusCode, fmt.Sprintf("expected status Bad Request; got: %v", res.StatusCode))
// 				return
// 			}
// 			assert.Equal(t, tc.status, res.StatusCode, fmt.Sprintf("expected status %v; got %v", tc.status, res.StatusCode))
// 		})
// 	}
// }

// func TestRemove(t *testing.T) {
// 	tt := []struct {
// 		name   string
// 		method string
// 		param  string
// 		value  string
// 		body   []byte
// 		status int
// 		err    string
// 	}{
// 		{name: "delete data by id", method: "DELETE", param: "id", value: "1", body: nil, status: http.StatusOK},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			req, err := http.NewRequest(tc.method, "localhost:8000/", bytes.NewReader(tc.body))
// 			assert.NoError(t, err, fmt.Sprintf("couldn't create requset: %v", err))

// 			// add url params to request.
// 			q := req.URL.Query()
// 			q.Add(tc.param, tc.value)
// 			req.URL.RawQuery = q.Encode()

// 			rec := httptest.NewRecorder()
// 			network.Remove(rec, req)

// 			res := rec.Result()
// 			defer res.Body.Close()
// 			if tc.err != "" {
// 				assert.Equal(t, http.StatusBadRequest, res.StatusCode, fmt.Sprintf("expected status Bad Request; got: %v", res.StatusCode))
// 				return
// 			}
// 			assert.Equal(t, tc.status, res.StatusCode, fmt.Sprintf("expected status %v; got %v", tc.status, res.StatusCode))
// 		})
// 	}
// }
