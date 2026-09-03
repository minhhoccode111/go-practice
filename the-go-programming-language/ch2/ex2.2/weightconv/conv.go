package weightconv

// KgToLb converts a Kilogram weight to Pound.
func KgToLb(k Kilogram) Pound { return Pound(k * 2.2046) }

// LbToKg converts a Pound weight to Kilogram.
func LbToKg(p Pound) Kilogram { return Kilogram(p / 2.2046) }
