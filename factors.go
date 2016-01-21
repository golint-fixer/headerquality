package headerquality

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Factors - return quality factors from give Request headers
func Factors(factor string, request *http.Request) (output Parameters) {
	output = make(Parameters, 0)
	if request == nil || request.Header == nil || factor == "" {
		return output
	}

	if header := request.Header.Get(factor); header != "" {
		factorHeaderValues := strings.Split(header, ",")
		output = make(Parameters, len(factorHeaderValues))

		for i, factorRange := range factorHeaderValues {
			// Check if a given range is qualified or not
			if qualifiedRange := strings.Split(factorRange, ";q="); len(qualifiedRange) == 2 {
				quality, error := strconv.ParseFloat(trim(qualifiedRange[1]), 32)
				if error != nil {
					// When the quality is unparseable, assume it's 1
					output[i] = Parameter{trim(qualifiedRange[0]), 1}
				} else {
					output[i] = Parameter{trim(qualifiedRange[0]), float32(quality)}
				}
			} else {
				output[i] = Parameter{trim(factorRange), 1}
			}
		}

		sort.Sort(output)
		return output
	}
	return output
}

func trim(factor string) string {
	return strings.Trim(factor, " ")
}
