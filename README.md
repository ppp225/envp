# envp [![GoDoc](https://godoc.org/github.com/ppp225/envp?status.svg)](https://godoc.org/github.com/ppp225/envp) 
envp adds simple env var helpers with debug logging. Uses [GoDotEnv](https://github.com/joho/godotenv) internally.

# Why

* Reduce env var related boilerplate, thus making code more readable,
* Basic validation and fallback
* Log read values in human readible format
* Use `.env` file [convention](https://github.com/joho/godotenv#precedence--conventions)

# Why not?

There are multiple great pkgs that unmarshal env vars right away into structs, like [this](https://github.com/sethvargo/go-envconfig) or [this](https://github.com/Netflix/go-env).

# Usage

```bash
go get -u github.com/ppp225/envp
```

```go
import (
  "github.com/ppp225/envp"
)

func main() {
  env := envp.GetEnvStringFrom("APP_ENV", "development", []string{"development", "demo", "production"}) // key, default, allowedValues
  envp.LoadEnvFromEnvFiles(env)

  ip := envp.GetEnvString("IP_ADDR", "localhost:12345") // key, default
  num := envp.GetEnvFloat("NUM", 100.3, 0.1, 200.5) // key, default, MIN, MAX
  nuInt := envp.GetEnvInt("NUM2", 100, 0, 200) // key, default, MIN, MAX
  booo := envp.GetEnvBool("B", false) // key, default
}
```
