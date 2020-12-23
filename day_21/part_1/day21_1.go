// Álvaro Castellano Vela 2020/12/21
package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"log"
	"os"
	"regexp"
	"strings"
)

type Food struct {
	Ingredients map[string]bool
	Allergens   map[string]bool
}

func processFile(filename string) []Food {

	foods := make([]Food, 0)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	foodsOneAllergenRe := regexp.MustCompile(`^([^(]+) \(contains ([a-z]+)\)$`)
	foodsTwoOrMoreAllergensRe := regexp.MustCompile(`^([^(]+) \(contains ([a-z ,]+)\)$`)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// read Tiles
	for scanner.Scan() {

		foodLine := scanner.Text()

		food := Food{make(map[string]bool), make(map[string]bool)}

		var ingredients []string
		var allergens []string

		matchFoodsOneAllergen := foodsOneAllergenRe.FindAllStringSubmatch(foodLine, -1)
		matchFoodsTwoOrMoreAllergens := foodsTwoOrMoreAllergensRe.FindAllStringSubmatch(foodLine, -1)
		if len(matchFoodsOneAllergen) > 0 {
			ingredientsString := matchFoodsOneAllergen[0][1]
			allergentsString := matchFoodsOneAllergen[0][2]

			ingredientsSplitedString := strings.Split(ingredientsString, " ")
			ingredients = ingredientsSplitedString

			allergens = append(allergens, allergentsString)

		} else {
			ingredientsString := matchFoodsTwoOrMoreAllergens[0][1]
			allergentsString := matchFoodsTwoOrMoreAllergens[0][2]

			ingredientsSplitedString := strings.Split(ingredientsString, " ")
			ingredients = ingredientsSplitedString

			allergentsSplitString := strings.Split(allergentsString, ", ")
			allergens = allergentsSplitString
		}
		for _, ingredient := range ingredients {
			food.Ingredients[ingredient] = true
		}
		for _, allergen := range allergens {
			food.Allergens[allergen] = true
		}
		foods = append(foods, food)
	}

	return foods
}

func calculateAllergensSets(foods []Food) int {

	var result int = 0

	allergensMapSet := make(map[string]mapset.Set)
	allIngredients := make(map[string]bool)
	leftIngredients := make(map[string]bool)

	for _, food := range foods {
		for ingredient, _ := range food.Ingredients {
			allIngredients[ingredient] = true
		}
		for allergen, _ := range food.Allergens {
			if _, ok := allergensMapSet[allergen]; !ok {
				allergensMapSet[allergen] = mapset.NewSet()
				for ingredient, _ := range food.Ingredients {
					allergensMapSet[allergen].Add(ingredient)
				}
			} else {
				intermediateSet := mapset.NewSet()
				for ingredient, _ := range food.Ingredients {
					intermediateSet.Add(ingredient)
				}
				allergensMapSet[allergen] = allergensMapSet[allergen].Intersect(intermediateSet)
			}
		}
	}
	for _, set := range allergensMapSet {
		for ingredient, _ := range allIngredients {
			if set.Contains(ingredient) {
				allIngredients[ingredient] = false
			}
		}
	}

	for ingredient, _ := range allIngredients {
		if allIngredients[ingredient] {
			leftIngredients[ingredient] = true
		}
	}

	for _, food := range foods {
		for ingredient, _ := range food.Ingredients {
			if _, ok := leftIngredients[ingredient]; ok {
				result++
			}
		}
	}
	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("You must supply a file to process.")
	}
	filename := args[0]

	foods := processFile(filename)

	fmt.Println("Result:", calculateAllergensSets(foods))
}
