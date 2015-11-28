package translators

import (
    "math"
)

const comprCoeff = 1 / 298.3

const eqEarthRadiusKm = 6378.137

type GeoPos struct {

    LatR float64
    LngR float64
    HKm float64
}

func getComprInPos(geoPos *GeoPos) (comprInPos float64) {

    k := comprCoeff * (2 - comprCoeff)
    comprInPos = math.Sqrt(1 - k * math.Pow(math.Sin(geoPos.LngR), 2))
    return
}

func convertGeoPosToCartesianPos(geoPos *GeoPos) (xKm, yKm, zKm float64) {

    comprInPos := getComprInPos(geoPos)
    posEarthRadiusKm := eqEarthRadiusKm / comprInPos

    xKm = (posEarthRadiusKm + geoPos.HKm) * math.Cos(geoPos.LngR) * math.Cos(geoPos.LatR)

    yKm = (posEarthRadiusKm + geoPos.HKm) * math.Cos(geoPos.LngR) * math.Sin(geoPos.LatR)

    zKm = (posEarthRadiusKm + geoPos.HKm) * math.Sin(geoPos.LngR)

    return
}
