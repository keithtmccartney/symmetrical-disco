# symmetrical-disco
[Cory Lanou] Making a RESTful JSON API in Go

Grab the article at [https://thenewstack.io/make-a-restful-json-api-go/](https://thenewstack.io/make-a-restful-json-api-go/)

## Tips/Run

* At CLI run "go run main.go" to serve at localhost:8080
* I'm still a little confused by the closing of the function not being required, including it results in "exported function Index should have comment or be unexported" (will investigate and learn)
* There's an alternative router in the form of [mux](http://www.gorillatoolkit.org/pkg/mux) from the [Gorilla Web Toolkit](http://www.gorillatoolkit.org/)
* There's 'ToDo' routing in place, accessible via "go run main.go", followed by the example localhost:8080 URLs: [http://localhost:8080](http://localhost:8080); [http://localhost:8080/todos](http://localhost:8080/todos); [http://localhost:8080/todos/1](http://localhost:8080/todos/1); [http://localhost:8080/todos/2](http://localhost:8080/todos/2); [http://localhost:8080/todos/3](http://localhost:8080/todos/3);
* Learned about trying to instantiate a new mux router via 'NewRouter' today when using "go run", turns out I can't so a build is needed ("go build"): https://github.com/corylanou/tns-restful-json-api/issues/2; (Y)
* Run the following curl command in a Bash CLI (this will create a new Todo entry with a POST to the '/todos' URL): curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos; I resulted in also trying Postman but to no avail was there any success, problems were had and installing Delve ([Windows](https://github.com/derekparker/delve/blob/master/Documentation/installation/windows/install.md)) for debugging was required alongside configuring a [launch.json](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code)

## What's it all about?

* ..."we will not only cover how to use Go to create a RESTful JSON API, but we will also talk about good RESTful design"...

## Thanks

Thanks goes out to Cory Lanou for the FREE stuff!

* [Abc](https://abc.com) ...Abc...
