package thecocktaildb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TheCocktailDB struct {
	ApiKey string
}

type Drinks struct {
	Drinks []Drink `json:"drinks"`
}

type Drink struct {
	DateModified             string `json:"dateModified"`
	IdDrink                  string `json:"idDrink"`
	Alcoholic                string `json:"strAlcoholic"`
	Category                 string `json:"strCategory"`
	CreativeCommonsConfirmed string `json:"strCreativeCommonsConfirmed"`
	Name                     string `json:"strDrink"`
	DrinkAlternate           string `json:"strDrinkAlternate"`
	Thumbnail                string `json:"strDrinkThumb"`
	Glass                    string `json:"strGlass"`
	IBA                      string `json:"strIBA"`
	ImageAttribution         string `json:"strImageAttribution"`
	ImageSource              string `json:"strImageSource"`
	Ingredient1              string `json:"strIngredient1"`
	Ingredient2              string `json:"strIngredient2"`
	Ingredient3              string `json:"strIngredient3"`
	Ingredient4              string `json:"strIngredient4"`
	Ingredient5              string `json:"strIngredient5"`
	Ingredient6              string `json:"strIngredient6"`
	Ingredient7              string `json:"strIngredient7"`
	Ingredient8              string `json:"strIngredient8"`
	Ingredient9              string `json:"strIngredient9"`
	Ingredient10             string `json:"strIngredient10"`
	Ingredient11             string `json:"strIngredient11"`
	Ingredient12             string `json:"strIngredient12"`
	Ingredient13             string `json:"strIngredient13"`
	Ingredient14             string `json:"strIngredient14"`
	Ingredient15             string `json:"strIngredient15"`
	Instructions             string `json:"strInstructions"`
	InstructionsDE           string `json:"strInstructionsDE"`
	InstructionsES           string `json:"strInstructionsES"`
	InstructionsFR           string `json:"strInstructionsFR"`
	InstructionsIT           string `json:"strInstructionsIT"`
	InstructionsZHHANS       string `json:"strInstructionsZH-HANS"`
	InstructionsZHHANT       string `json:"strInstructionsZH-HANT"`
	Measure1                 string `json:"strMeasure1"`
	Measure2                 string `json:"strMeasure2"`
	Measure3                 string `json:"strMeasure3"`
	Measure4                 string `json:"strMeasure4"`
	Measure5                 string `json:"strMeasure5"`
	Measure6                 string `json:"strMeasure6"`
	Measure7                 string `json:"strMeasure7"`
	Measure8                 string `json:"strMeasure8"`
	Measure9                 string `json:"strMeasure9"`
	Measure10                string `json:"strMeasure10"`
	Measure11                string `json:"strMeasure11"`
	Measure12                string `json:"strMeasure12"`
	Measure13                string `json:"strMeasure13"`
	Measure14                string `json:"strMeasure14"`
	Measure15                string `json:"strMeasure15"`
	Tags                     string `json:"strTags"`
	Video                    string `json:"strVideo"`
}

type Ingredients struct {
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Id          string `json:"idIngredient"`
	ABV         string `json:"strABV"`
	Alcohol     string `json:"strAlcohol"`
	Description string `json:"strDescription"`
	Name        string `json:"strIngredient"`
	Type        string `json:"strType"`
}

type DrinkList struct {
	Items []DrinkListItem `json:"drinks"`
}

type DrinkListItem struct {
	IdDrink   string `json:"idDrink"`
	Name      string `json:"strDrink"`
	Thumbnail string `json:"strDrinkThumb"`
}

type CategoryListItem struct {
	Name string `json:"strCategory"`
}

type CategoryList struct {
	Items []CategoryListItem `json:"drinks"`
}

type GlassList struct {
	Items []GlassListItem `json:"drinks"`
}

type GlassListItem struct {
	Name string `json:"strGlass"`
}

type IngredientList struct {
	Items []IngredientListItem `json:"drinks"`
}

type IngredientListItem struct {
	Name string `json:"strIngredient1"`
}

type AlcoholFilterList struct {
	Items []AlcoholFilterListItem `json:"drinks"`
}

type AlcoholFilterListItem struct {
	Name string `json:"strAlcoholic"`
}

func New(apiKey string) *TheCocktailDB {
	if apiKey == "" {
		apiKey = "1"
	}

	return &TheCocktailDB{ApiKey: apiKey}
}

func (T *TheCocktailDB) SearchByName(name string) (error, Drinks) {
	url := T.getBaseUrl() + "search.php?s=" + name
	err, drinks := doRequest(url, Drinks{})
	if err != nil {
		return err, Drinks{}
	}

	return nil, drinks
}

func (T *TheCocktailDB) SearchByFirstLetter(letter string) (error, Drinks) {
	url := T.getBaseUrl() + "search.php?f=" + letter
	err, drinks := doRequest(url, Drinks{})
	if err != nil {
		return err, Drinks{}
	}

	return nil, drinks
}

func (T *TheCocktailDB) SearchIngredientByName(ingredient string) (error, Ingredient) {
	url := T.getBaseUrl() + "search.php?i=" + ingredient
	err, ingredients := doRequest(url, Ingredients{})
	if err != nil {
		return err, Ingredient{}
	}

	return nil, ingredients.Ingredients[0]
}

func (T *TheCocktailDB) SearchByCocktailId(id string) (error, Drink) {
	url := T.getBaseUrl() + "lookup.php?i=" + id
	err, drinks := doRequest(url, Drinks{})
	if err != nil {
		return err, Drink{}
	}

	return nil, drinks.Drinks[0]
}

func (T *TheCocktailDB) SearchIngredientById(id string) (error, Ingredient) {
	url := T.getBaseUrl() + "lookup.php?iid=" + id
	err, ingredients := doRequest(url, Ingredients{})
	if err != nil {
		return err, Ingredient{}
	}

	return nil, ingredients.Ingredients[0]
}

func (T *TheCocktailDB) GetRandomDrink() (error, Drink) {
	url := T.getBaseUrl() + "random.php"
	err, drinks := doRequest(url, Drinks{})
	if err != nil {
		return err, Drink{}
	}

	return nil, drinks.Drinks[0]
}

func (T *TheCocktailDB) GetRandomDrinks() (error, Drinks) {
	url := T.getBaseUrl() + "randomselection.php"
	err, drinks := doRequest(url, Drinks{})
	if err != nil {
		return err, Drinks{}
	}

	return nil, drinks
}

func (T *TheCocktailDB) GetPopularDrinks() (error, Drinks) {
	url := T.getBaseUrl() + "popular.php"
	err, drinks := doRequest(url, Drinks{})
	if err != nil {
		return err, Drinks{}
	}

	return nil, drinks
}

func (T *TheCocktailDB) GetLatestDrinks() (error, Drinks) {
	url := T.getBaseUrl() + "popular.php"
	err, drinks := doRequest(url, Drinks{})
	if err != nil {
		return err, Drinks{}
	}

	return nil, drinks
}

func (T *TheCocktailDB) SearchByFilter(ingredient string, alcoholic string, category string, glassType string) (error, DrinkList) {
	url := T.getBaseUrl() + "filter.php?"
	if ingredient != "" {
		url += "i=" + ingredient + "&"
	}

	if alcoholic != "" {
		url += "a=" + alcoholic + "&"
	}

	if category != "" {
		url += "c=" + category + "&"
	}

	if glassType != "" {
		url += "g=" + glassType + "&"
	}

	err, drinkList := doRequest(url, DrinkList{})
	if err != nil {
		return err, DrinkList{}
	}

	return nil, drinkList
}
func (T *TheCocktailDB) SearchByMultiFilter(ingredient []string, alcoholic []string, category []string, glassType []string) (error, DrinkList) {
	url := T.getBaseUrl() + "filter.php?"
	if len(ingredient) > 0 {
		url += "i=" + strings.Join(ingredient, ",") + "&"
	}

	if len(alcoholic) > 0 {
		url += "a=" + strings.Join(alcoholic, ",") + "&"
	}

	if len(category) > 0 {
		url += "c=" + strings.Join(category, ",") + "&"
	}

	if len(glassType) > 0 {
		url += "g=" + strings.Join(glassType, ",") + "&"
	}

	err, drinkList := doRequest(url, DrinkList{})
	if err != nil {
		return err, DrinkList{}
	}

	return nil, drinkList
}

func (T *TheCocktailDB) GetCategories() (error, CategoryList) {
	url := T.getBaseUrl() + "list.php?c=list"
	err, categories := doRequest(url, CategoryList{})
	if err != nil {
		return err, CategoryList{}
	}

	return nil, categories
}

func (T *TheCocktailDB) GetGlasses() (error, GlassList) {
	url := T.getBaseUrl() + "list.php?g=list"
	err, glasses := doRequest(url, GlassList{})
	if err != nil {
		return err, GlassList{}
	}

	return nil, glasses
}

func (T *TheCocktailDB) GetIngredients() (error, IngredientList) {
	url := T.getBaseUrl() + "list.php?i=list"
	err, ingredients := doRequest(url, IngredientList{})
	if err != nil {
		return err, IngredientList{}
	}

	return nil, ingredients
}

func (T *TheCocktailDB) GetAlcoholFilter() (error, AlcoholFilterList) {
	url := T.getBaseUrl() + "list.php?a=list"
	err, alcohol := doRequest(url, AlcoholFilterList{})
	if err != nil {
		return err, AlcoholFilterList{}
	}

	return nil, alcohol
}

func (T *TheCocktailDB) getBaseUrl() string {
	return fmt.Sprintf("https://www.thecocktaildb.com/api/json/v1/%s/", T.ApiKey)
}

func doRequest[T any](url string, resp T) (error, T) {
	response, err := http.Get(url)
	if err != nil {
		return err, resp
	}

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err, resp
	}

	err = json.Unmarshal(result, &resp)
	if err != nil {
		return err, resp
	}
	return nil, resp
}
