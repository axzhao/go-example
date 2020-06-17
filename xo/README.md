# 1. xo pgsql

mkdir models
xo 'pgsql://postgres:password@localhost:5432/demo?sslmode=disable' -o models


```go

func Example() {

	db, err := sql.Open("postgres", "postgres://postgres:password@127.0.0.1/demo?sslmode=disable")
	if err != nil {
		panic(err)
	}

	a := Article{
		ID:    2,
		Title: "aaaa",
	}
	if err := a.Insert(db); err != nil {
		panic(err)
	}

	aa, err := ArticleByID(db, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(aa)

	// Output:
}

```

# 2. templates

`cp "$GOPATH/src/github.com/xo/xo/templates/*" templates/`
vi templates/postgres.type.tpl.go
add GetMostRecent func
xo 'pgsql://postgres:password@localhost:5432/demo?sslmode=disable' -o models --template-path templates

