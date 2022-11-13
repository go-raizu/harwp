# Harwp

Harwp provides a proxy type for `http.ResponseWriter` of the Go standard library. 

The intended use of the proxy type is for middlewares to inspect the `http.ResponseWriter` after the encapsulated `http.Handler` has finished their execution.
