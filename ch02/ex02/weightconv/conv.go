package weightconv

type Pound float64
type Kilogram float64

func PToK(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}

func KToP(k Kilogram) Pound {
	return Pound(k * 2.2046)
}
