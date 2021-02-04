package reflection

// START OMIT
import "net/url"

func sample(url string) {
	// ... 処理 ...
	u, _ := url.Parse(target) // url.Parse undefined // HL
// END OMIT
}
