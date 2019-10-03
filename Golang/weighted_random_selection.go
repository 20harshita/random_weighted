package main

import (
    "errors"
     "fmt"
     "math/rand"
     "time"
)

type Game struct {
    Name   string
    Weight int
}

func main() {
    games := []Game{
        Game{Name: "Mario", Weight: 1},
        Game{Name: "Sonic", Weight: 2},
        Game{Name: "CoD", Weight: 4},
        Game{Name: "Link", Weight: 8},
        Game{Name: "Fantasy", Weight: 10},
        Game{Name: "Destiny", Weight: 7},
    }
    
    var totalWeight int
        for _, g := range games {
        totalWeight += g.Weight
    }
    
    results := map[string]int{}
    for i := 0; i < 10000; i++ {
        g, err := RandomWeightedSelect(games, totalWeight)
        if err != nil {
            panic(err)
        }
        if _, ok := results[g.Name]; ok {
            results[g.Name]++
        } else {
            results[g.Name] = 1
        }
    }
    fmt.Println(results)
}

func RandomWeightedSelect(games []Game, totalWeight int) (Game, error) {
    rand.Seed(time.Now().UnixNano())
    r := rand.Intn(totalWeight)
    for _, g := range games {
        r -= g.Weight
        if r <= 0 {
            return g, nil
        }
    }
    return Game{}, errors.New("No game selected")
}