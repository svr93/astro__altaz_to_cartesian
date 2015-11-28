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
