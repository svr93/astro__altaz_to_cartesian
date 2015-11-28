package translators

import (
    "math"
)

func ConvertAltAzRadiusToLocalCartezian(altR, azR, radius float64) (x, y, z float64) {

    x = radius * math.Sin(azR) * math.Cos(altR)
    y = radius * math.Sin(azR) * math.Sin(altR)
    z = radius * math.Cos(azR)
    return
}

func GetVelocityVector(altR, azR, radius, altRVel, azRVel, radiusVel float64) (x, y, z float64) {

    x = radiusVel * math.Sin(azR) * math.Cos(altR) +
        radius * azRVel * math.Cos(azR) * math.Cos(altR) -
        radius * altRVel * math.Sin(azR) * math.Sin(altR)

    y = radiusVel * math.Sin(azR) * math.Sin(altR) +
        radius * azRVel * math.Cos(azR) * math.Sin(altR) +
        radius * altRVel * math.Sin(azR) * math.Cos(altR)

    z = radiusVel * math.Cos(azR) - radius * azRVel * math.Sin(azR)

    return
}
