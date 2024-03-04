# core

Core package for Golang.

## Install

```shell
go get github.com/innotechdevops/core
```

## How to use

- Compare Password

```go
got := gobcrypt.ComparePassword(hashedPwd, plainPwd)
```

- Hash Password

```go
got, err := gobcrypt.HashPassword(password)
```