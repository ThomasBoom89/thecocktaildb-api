package main

import (
	"fmt"
	"github.com/ThomasBoom89/thecocktaildb-api/pkg/thecocktaildb"
)

func main() {
	thecocktaildbapi := thecocktaildb.New("")
	err, ingredient := thecocktaildbapi.SearchIngredientByName("tequila")
	if err != nil {
		panic(err)
	}
	fmt.Println(ingredient.Name)
	fmt.Println(".................................................................................")

	err, categories := thecocktaildbapi.GetCategories()
	if err != nil {
		panic(err)
	}
	fmt.Println(categories.Items[0].Name)
	fmt.Println(".................................................................................")

	err, ingredients := thecocktaildbapi.GetIngredients()
	if err != nil {
		panic(err)
	}
	fmt.Println(ingredients.Items[0].Name)
	fmt.Println(".................................................................................")

	err, glasses := thecocktaildbapi.GetGlasses()
	if err != nil {
		panic(err)
	}
	fmt.Println(glasses.Items[0].Name)
	fmt.Println(".................................................................................")
	err, alcohol := thecocktaildbapi.GetAlcoholFilter()
	if err != nil {
		panic(err)
	}
	fmt.Println(alcohol.Items[0].Name)
	fmt.Println(".................................................................................")

	err, drinkList := thecocktaildbapi.SearchByFilter("Light rum", "Alcoholic", "Ordinary Drink", "Highball glass")
	if err != nil {
		panic(err)
	}
	fmt.Println(drinkList.Items[0].Name)
	fmt.Println(".................................................................................")

	err, drink := thecocktaildbapi.GetRandomDrink()
	if err != nil {
		panic(err)
	}

	fmt.Println(drink.Name)
}
