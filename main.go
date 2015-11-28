package main

import (
    "fmt"
    "math"
    "./translators"
)

func main() {

    stationPos := translators.GeoPos{ LatR: 0, LngR: math.Pi / 2, HKm: 0 }
    xKmL, yKmL, zKmL := translators.ConvertAltAzRadiusToLocalCartezian(0, 0, 0)
    fmt.Println(translators.ConvertLocalCartezianToECEF(xKmL, yKmL, zKmL, &stationPos))
}
