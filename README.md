# cmpe273-lab2

This is a simple “Hello World” REST API in Go to understand how handler and Http router work, and how to parse the request and send the response.
GET /hello/xxxx
http://localhost:8080/hello/foo
        Response:
                Hello, xxxx!
                hello, foo!
                
POST /hello
Request:
{
   “name”: “foo”
}
Response:
{
   “greeting” : “Hello, foo!”
}

Post is run through POSTER using Mozilla Firefox
and also through curl command : 
C:\>curl -H "Content-Type: application/json" -X POST -d "{\"name\":\"Foo Bar\"}"
 http://127.0.0.1:8080/hello
{"Greeting":"Hello,Foo Bar!"}


