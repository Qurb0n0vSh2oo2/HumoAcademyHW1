package pkg

import "bankCLI/pkg/models"

func FillCities() {
	Cities = append(Cities, models.City{
		City:   "Душанбе",
		Region: "РРП",
	})

	Cities = append(Cities, models.City{
		City:   "Куляб",
		Region: "Хатлон",
	})

	Cities = append(Cities, models.City{
		City:   "Худжанд",
		Region: "Согд",
	})

	Cities = append(Cities, models.City{
		City:   "Памир",
		Region: "ГБАО",
	})

}
