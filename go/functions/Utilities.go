package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func DateStringToIntSlice(dateString string) ([]int, error) {
	dateParts := strings.Split(dateString, "-")

	var intSlice []int
	for _, part := range dateParts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		intSlice = append(intSlice, num)
	}

	return intSlice, nil
}

func DateStringToStringSlice(dateString string) ([]string, error) {
	if dateString[0] == '*' {
		dateString = dateString[1:]
	}
	dateParts := strings.Split(dateString, "-")

	var stringSlice []string
	for _, part := range dateParts {
		stringSlice = append(stringSlice, part)
	}

	return stringSlice, nil
}

func IsNumberinSlice(number int, slice []int) bool {
	for _, value := range slice {
		if value == number {
			return true
		}
	}
	return false
}

func IsStringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func parseDate(dateStr string) string {
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return ""
	}
	return parts[2] + "-" + parts[1] + "-" + parts[0]
}

func DateStringToYear(date string) (float64, error) {
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid date format: %s", date)
	}

	year, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse year: %s", parts[2])
	}

	return year, nil
}

func TranslateEnglishCity(city string) string {
	cities := map[string]string{
		"CAROLINE DU NORD":       "NORTH CAROLINA",
		"GÉORGIE":                "GEORGIA",
		"LOS ANGELES":            "LOS ANGELES",
		"SAITAMA":                "SAITAMA",
		"OSAKA":                  "OSAKA",
		"NAGOYA":                 "NAGOYA",
		"PENROSE":                "PENROSE",
		"DUNEDIN":                "DUNEDIN",
		"PLAYA DEL CARMEN":       "PLAYA DEL CARMEN",
		"PAPEETE":                "PAPEETE",
		"NOUMÉA":                 "NOUMEA",
		"LONDRES":                "LONDON",
		"LAUSANNE":               "LAUSANNE",
		"LYON":                   "LYON",
		"VICTORIA":               "VICTORIA",
		"NOUVELLE-GALLES DU SUD": "NEW SOUTH WALES",
		"QUEENSLAND":             "QUEENSLAND",
		"AUCKLAND":               "AUCKLAND",
		"YOGYAKARTA":             "YOGYAKARTA",
		"BRATISLAVA":             "BRATISLAVA",
		"BUDAPEST":               "BUDAPEST",
		"MINSK":                  "MINSK",
		"CALIFORNIE":             "CALIFORNIA",
		"NEVADA":                 "NEVADA",
		"SAO PAULO":              "SAO PAULO",
		"SAN ISIDRO":             "SAN ISIDRO",
		"ARIZONA":                "ARIZONA",
		"TEXAS":                  "TEXAS",
		"STOCKHOLM":              "STOCKHOLM",
		"WERCHTER":               "WERCHTER",
		"LISBONNE":               "LISBON",
		"BILBAO":                 "BILBAO",
		"BOGOTA":                 "BOGOTA",
		"NEW YORK":               "NEW YORK",
		"DÜSSELDORF":             "DUSSELDORF",
		"AARHUS":                 "AARHUS",
		"MANCHESTER":             "MANCHESTER",
		"FRANCFORT":              "FRANKFURT",
		"BERLIN":                 "BERLIN",
		"COPENHAGUE":             "COPENHAGEN",
		"DOHA":                   "DOHA",
		"MINNESOTA":              "MINNESOTA",
		"ILLINOIS":               "ILLINOIS",
		"MUMBAI":                 "MUMBAI",
		"ABOU DHABI":             "ABU DHABI",
		"PENNSYLVANIE":           "PENNSYLVANIA",
		"WESTCLIFF-ON-SEA":       "WESTCLIFF ON SEA",
		"MERKERS":                "MERKERS",
		"MAINE":                  "MAINE",
		"GOTHENBURG":             "GOTHENBURG",
		"FLORIDE":                "FLORIDA",
		"CAROLINE DU SUD":        "SOUTH CAROLINA",
		"PAGNEY-DERRIÈRE-BARINE": "PAGNEY DERRIERE BARINE",
		"HAMBURG":                "HAMBURG",
		"BOULOGNE-BILLANCOURT":   "BOULOGNE BILLANCOURT",
		"NÎMES":                  "NIMES",
		"SION":                   "SION",
		"OSTRAVA":                "OSTRAVA",
		"KLAGENFURT":             "KLAGENFURT",
		"FREYMING-MERLEBACH":     "FREYMING MERLEBACH",
		"ZARAGOZA":               "ZARAGOZA",
		"MADRID":                 "MADRID",
		"BARCELONE":              "BARCELONA",
		"RIO DE JANEIRO":         "RIO DE JANEIRO",
		"RECIFE":                 "RECIFE",
		"LEIPZIG":                "LEIPZIG",
		"SALEM":                  "SALEM",
		"MONCHENGLADBACH":        "MONCHENGLADBACH",
		"CUXHAVEN":               "CUXHAVEN",
		"SKANDERBORG":            "SKANDERBORG",
		"AMSTERDAM":              "AMSTERDAM",
		"BURRIANA":               "BURRIANA",
		"OULU":                   "OULU",
		"NAPOCA":                 "NAPOCA",
		"RIYAD":                  "RIYADH",
		"CANTON":                 "CANTON",
		"QUÉBEC":                 "QUEBEC",
		"LAS VEGAS":              "LAS VEGAS",
		"MEXICO":                 "MEXICO CITY",
		"MONTERREY":              "MONTERREY",
		"DEL MAR":                "DEL MAR",
		"WASHINGTON":             "WASHINGTON",
		"WEST MELBOURNE":         "WEST MELBOURNE",
		"PARIS":                  "PARIS",
		"MISSOURI":               "MISSOURI",
		"CHICAGO":                "CHICAGO",
		"BIRMINGHAM":             "BIRMINGHAM",
		"SYDNEY":                 "SYDNEY",
		"MADISON":                "MADISON",
		"TORONTO":                "TORONTO",
		"CLEVELAND":              "CLEVELAND",
		"BOSTON":                 "BOSTON",
		"UTAH":                   "UTAH",
		"GLASGOW":                "GLASGOW",
		"DUBLIN":                 "DUBLIN",
		"CARDIFF":                "CARDIFF",
		"ABERDEEN":               "ABERDEEN",
		"VARSOVIE":               "WARSAW",
		"MILAN":                  "MILAN",
		"MICHIGAN":               "MICHIGAN",
		"NEW HAMPSHIRE":          "NEW HAMPSHIRE",
		"SOCHAUX":                "SOCHAUX",
		"EINDHOVEN":              "EINDHOVEN",
		"OSLO":                   "OSLO",
		"COLORADO":               "COLORADO",
		"JAKARTA":                "JAKARTA",
		"HUIZHOU":                "HUIZHOU",
		"CHANGZHOU":              "CHANGZHOU",
		"HONG KONG":              "HONG KONG",
		"SANYA":                  "SANYA",
		"AALBORG":                "AALBORG",
		"SEATTLE":                "SEATTLE",
		"OMAHA":                  "OMAHA",
		"KANSAS CITY":            "KANSAS CITY",
		"ST LOUIS":               "ST LOUIS",
		"INDIANAPOLIS":           "INDIANAPOLIS",
		"ROSEMONT":               "ROSEMONT",
		"GRAND RAPIDS":           "GRAND RAPIDS",
		"MONTRÉAL":               "MONTREAL",
		"NEWARK":                 "NEWARK",
		"UNIONDALE":              "UNIONDALE",
		"PHILADELPHIE":           "PHILADELPHIA",
		"HERSHEY":                "HERSHEY",
		"PITTSBURGH":             "PITTSBURGH",
		"COLUMBIA":               "COLUMBIA",
		"SANTIAGO":               "SANTIAGO",
		"HOUSTON":                "HOUSTON",
		"ATLANTA":                "ATLANTA",
		"NOUVELLE-ORLÉANS":       "NEW ORLEANS",
		"FRAUENFELD":             "FRAUENFELD",
		"TURKU":                  "TURKU",
		"BROOKLYN":               "BROOKLYN",
		"IMOLA":                  "IMOLA",
		"VIENNE":                 "VIENNA",
		"CRACOVIE":               "KRAKOW",
		"ZURICH":                 "ZURICH",
		"AMITYVILLE":             "AMITYVILLE",
		"MINNEAPOLIS":            "MINNEAPOLIS",
		"DÉTROIT":                "DETROIT",
		"OAKLAND":                "OAKLAND",
		"CHARLOTTE":              "CHARLOTTE",
		"INGLEWOOD":              "INGLEWOOD",
		"WINDSOR":                "WINDSOR",
		"CINCINNATI":             "CINCINNATI",
		"ANAHEIM":                "ANAHEIM",
		"MANILLE":                "MANILA",
		"BRISBANE":               "BRISBANE",
		"MELBOURNE":              "MELBOURNE",
		"LIMA":                   "LIMA",
		"GRONINGEN":              "GRONINGEN",
		"ANVERS":                 "ANTWERP",
		"PICO RIVERA":            "PICO RIVERA",
		"BERWYN":                 "BERWYN",
		"DALLAS":                 "DALLAS",
		"BRIXTON":                "BRIXTON",
		"ROTSELAAR":              "ROTSELAAR",
		"ALABAMA":                "ALABAMA",
		"MASSACHUSETTS":          "MASSACHUSETTS",
		"ATHÈNES":                "ATHENS",
		"FLORENCE":               "FLORENCE",
		"LANDGRAAF":              "LANDGRAAF",
		"BURSWOOD":               "BURSWOOD",
		"WELLINGTON":             "WELLINGTON",
		"SÉVILLE":                "SEVILLE",
		"BANGKOK":                "BANGKOK",
		"TAÏPEI":                 "TAIPEI",
		"SÉOUL":                  "SEOUL",
		"MUNICH":                 "MUNICH",
		"MANNHEIM":               "MANNHEIM",
		"SAN FRANCISCO":          "SAN FRANCISCO",
		"BUENOS AIRES":           "BUENOS AIRES",
		"PORTO ALEGRE":           "PORTO ALEGRE",
		"BELO HORIZONTE":         "BELO HORIZONTE",
		"LA PLATA":               "LA PLATA",
		"DUBAÏ":                  "DUBAI",
		"WILLEMSTAD":             "WILLEMSTAD",
		"BRASILIA":               "BRASILIA",
		"OKLAHOMA":               "OKLAHOMA",
		"SCHEESSEL":              "SCHEESSEL",
		"SAINT-GALL":             "ST GALLEN",
		"GDYNIA":                 "GDYNIA",
		"ARRAS":                  "ARRAS",
		"SAN JOSÉ":               "SAN JOSE",
		"NICKELSDORF":            "NICKELSDORF",
		"OREGON":                 "OREGON",
		"VANCOUVER":              "VANCOUVER",
		"PRAGUE":                 "PRAGUE",
	}

	// Convertit le nom de la ville entré en majuscules pour correspondre aux clés de la map
	if translatedCity, ok := cities[city]; ok {
		return translatedCity
	}
	// Renvoie la ville non trouvée telle quelle ou un message spécifique
	return city
}

func TranslateFrenchCountries(country string) string {
	invertedCountryTranslations := map[string]string{
		"GERMANY":              "ALLEMAGNE",
		"SAUDI ARABIA":         "ARABIE SAOUDITE",
		"NETHERLANDS ANTILLES": "ANTILLES NÉERLANDAISES",
		"ARGENTINA":            "ARGENTINE",
		"AUSTRALIA":            "AUSTRALIE",
		"AUSTRIA":              "AUTRICHE",
		"BELGIUM":              "BELGIQUE",
		"BELARUS":              "BIÉLORUSSIE",
		"BRAZIL":               "BRÉSIL",
		"CANADA":               "CANADA",
		"CHILE":                "CHILI",
		"CHINA":                "CHINE",
		"COLOMBIA":             "COLOMBIE",
		"SOUTH KOREA":          "CORÉE DU SUD",
		"COSTA RICA":           "COSTA RICA",
		"DENMARK":              "DANEMARK",
		"UNITED ARAB EMIRATES": "ÉMIRATS ARABES UNIS",
		"USA":                  "ÉTATS-UNIS",
		"SPAIN":                "ESPAGNE",
		"FINLAND":              "FINLANDE",
		"FRANCE":               "FRANCE",
		"GREECE":               "GRÈCE",
		"HUNGARY":              "HONGRIE",
		"INDIA":                "INDE",
		"INDONESIA":            "INDONÉSIE",
		"IRELAND":              "IRLANDE",
		"ITALY":                "ITALIE",
		"JAPAN":                "JAPON",
		"MEXICO":               "MEXIQUE",
		"NORWAY":               "NORVÈGE",
		"NEW CALEDONIA":        "NOUVELLE-CALÉDONIE",
		"NEW ZEALAND":          "NOUVELLE-ZÉLANDE",
		"NETHERLANDS":          "PAYS-BAS",
		"PERU":                 "PÉROU",
		"PHILIPPINES":          "PHILIPPINES",
		"POLAND":               "POLOGNE",
		"FRENCH POLYNESIA":     "POLYNÉSIE FRANÇAISE",
		"PORTUGAL":             "PORTUGAL",
		"QATAR":                "QATAR",
		"ROMANIA":              "ROUMANIE",
		"UK":                   "ROYAUME-UNI",
		"SLOVAKIA":             "SLOVAQUIE",
		"SWEDEN":               "SUÈDE",
		"SWITZERLAND":          "SUISSE",
		"TAIWAN":               "TAÏWAN",
		"CZECHIA":              "TCHÉQUIE",
		"THAILAND":             "THAÏLANDE",
	}

	// Convertit le nom du pays entré en majuscules pour correspondre aux clés de la map
	if translatedCountry, ok := invertedCountryTranslations[country]; ok {
		return translatedCountry
	}
	// Renvoie le pays non trouvé tel quel ou un message spécifique
	return country
}
