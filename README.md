# TheCocktailDB API

![Go](https://img.shields.io/github/go-mod/go-version/thomasboom89/thecocktaildb-api/main)
![License](https://img.shields.io/badge/license-MIT-green?style=plastic)

A wrapper for https://www.thecocktaildb.com API

### Attention!

I am not owner or maintainer of the API (https://www.thecocktaildb.com). This is only a wrapper for it.
If you want to use it professional you need to purchase an api-key -> https://www.thecocktaildb.com/api.php

## Installation

Include this module into your existing go project

```zsh
go get github.com/thomasboom89/thecocktaildb-api
```

## Update

```zsh
go get -u github.com/thomasboom89/thecocktaildb-api
```

## Usage

To use the wrapper you can look into the example file example/main.go
Here is also a little text for you

```go
// create instance and provide an api key, defaults to development key if none is provided
thecocktaildbapi := thecocktaildb.New("your-api-key")
// now you can call the specific endpoint
err, drink := thecocktaildbapi.GetRandomDrink()
if err != nil {
	panic(err)
}
fmt.Println(drink.Name)
```

## License

TheCocktailDB API
Copyright (C) 2022 ThomasBoom89. MIT license.