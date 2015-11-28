package translators

import (
    "math"
    "github.com/gonum/matrix/mat64"
)

const comprCoeff = 1 / 298.3

const eqEarthRadiusKm = 6378.137

const antennaAngleR = 0

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

func getDirectionCosMatrix(geoPos *GeoPos) *mat64.Dense {

    var val float64
    directionCosMatrix := mat64.NewDense(3, 3, nil)

    val = - math.Sin(geoPos.LngR) * math.Cos(geoPos.LatR) * math.Cos(antennaAngleR) -
            math.Sin(geoPos.LatR) * math.Sin(antennaAngleR)
    directionCosMatrix.Set(0, 0, val)

    val =   math.Cos(geoPos.LngR) * math.Cos(geoPos.LatR)
    directionCosMatrix.Set(0, 1, val)

    val =   math.Sin(geoPos.LngR) * math.Cos(geoPos.LatR) * math.Sin(antennaAngleR) -
            math.Sin(geoPos.LatR) * math.Cos(antennaAngleR)
    directionCosMatrix.Set(0, 2, val)

    val = - math.Sin(geoPos.LngR) * math.Sin(geoPos.LatR) * math.Cos(antennaAngleR) +
            math.Cos(geoPos.LatR) * math.Sin(antennaAngleR)
    directionCosMatrix.Set(1, 0, val)

    val =   math.Cos(geoPos.LngR) * math.Sin(geoPos.LatR)
    directionCosMatrix.Set(1, 1, val)

    val =   math.Sin(geoPos.LngR) * math.Sin(geoPos.LatR) * math.Sin(antennaAngleR) +
            math.Cos(geoPos.LatR) * math.Cos(antennaAngleR)
    directionCosMatrix.Set(1, 2, val)

    val =   math.Cos(geoPos.LngR) * math.Cos(antennaAngleR)
    directionCosMatrix.Set(2, 0, val)

    val =   math.Sin(geoPos.LngR)
    directionCosMatrix.Set(2, 1, val)

    val = - math.Cos(geoPos.LngR) * math.Sin(antennaAngleR)
    directionCosMatrix.Set(2, 2, val)

    return directionCosMatrix
}

func ConvertLocalCartezianToECEF(xKmL, yKmL, zKmL float64, geoPos *GeoPos) (x, y, z float64) {

    directionCosMatrix := getDirectionCosMatrix(geoPos)
    inverseDirectionCosMatrix := mat64.NewDense(3, 3, nil)
    inverseDirectionCosMatrix.Inverse(directionCosMatrix)

    localCartezianMatrix := mat64.NewDense(3, 1, []float64{ xKmL, yKmL, zKmL })

    stationXKm, stationYKm, stationZKm := convertGeoPosToCartesianPos(geoPos)
    stationCartezianMatrix := mat64.NewDense(3, 1, []float64{ stationXKm, stationYKm, stationZKm })

    mulMatrix := mat64.NewDense(3, 1, nil)
    mulMatrix.Mul(inverseDirectionCosMatrix, localCartezianMatrix)
    res := mat64.NewDense(3, 1, nil)
    res.Add(mulMatrix, stationCartezianMatrix)

    x = res.At(0, 0)
    y = res.At(1, 0)
    z = res.At(2, 0)

    return
}
