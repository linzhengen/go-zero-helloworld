### 1. N/A

1. route definition

- Url: /sayHello
- Method: GET
- Request: `request`
- Response: `response`

2. request definition



```golang
type Request struct {
	Name string `form:"name"`
}
```


3. response definition



```golang
type Response struct {
	Message string `json:"name"`
}
```

