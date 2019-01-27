# Simple Web App

This is the simple web app that demonstrate how :-

- Create simple web server
- Decode json request from the web
- Encode json from `struct` as a response
- Best practise/principle for performance purpose on how to decode/encode json
- Use `http.HandleFunc` to create routes.

## Usage

Just run `go run server.go`

Then visit `http://localhost:8080/api/v1` and `http://localhost:8080/api/v1/products` to get see the resuts.
To test `POST` method for json decoding use `curl`, and try `curl http://localhost:8080/api/v1/send -d '{"name":"Juma"}'` you should see a response.
