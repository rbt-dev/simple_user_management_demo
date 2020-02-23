# Simple User Management Demo

An exercise to use Go and Vue.js


## Built With

* [MySQL](https://www.mysql.com) - The database used
* [Go](https://golang.org) - Programming language used for back-end
* [Vue.js](https://vuejs.org) - Front-end JavaScript framework


### Setup database

Run *db.sql* script from *database* folder


### Run back-end

``` bash
# run server (default localhost:8081)
go run main.go handlers.go tools.go models.go
```

Configuration in: *backend_go/config*


### Run front-end

``` bash
# install dependencies
npm install

# serve with hot reload (default localhost:8082)
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```
Configuration in: *frontend_vue/config/dev.env.js*


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
