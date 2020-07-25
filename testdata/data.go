package testdata

var data = []interface{}{
	map[string]interface{}{
		"index":      0,
		"guid":       "11c73dad-9a34-4368-9522-af455bcdeef9",
		"isActive":   true,
		"balance":    "$2,857.19",
		"picture":    "http://placehold.it/32x32",
		"age":        23,
		"eyeColor":   "green",
		"name":       "Rosalind Oconnor",
		"gender":     "female",
		"company":    "ISOSWITCH",
		"email":      "rosalindoconnor@isoswitch.com",
		"phone":      "+1 (999) 459-3700",
		"address":    "341 Hubbard Street, Craig, Connecticut, 3709",
		"registered": "2014-12-07T02:57:57 -06:-30",
		"latitude":   -28.225805,
		"longitude":  33.195572,
		"tags": []string{
			"labore",
			"qui",
			"reprehenderit",
			"cillum",
			"voluptate",
			"laborum",
			"in",
		},
		"friends": []interface{}{
			map[string]interface{}{
				"id":   0,
				"name": "Justine Bird",
			},
			map[string]interface{}{
				"id":   0,
				"name": "Justine Bird",
			},
			map[string]interface{}{
				"id":   1,
				"name": "Marianne Rutledge",
			},
		},
	},
}

func Get() interface{} {
	return data
}

func GetObject() map[string]interface{} {
	return map[string]interface{}{
		"id":   0,
		"name": "Justine Bird",
	}
}

func GetArray() []interface{} {
	return []interface{}{
		map[string]interface{}{
			"id":   0,
			"name": "Justine Bird",
		},
		map[string]interface{}{
			"id":   0,
			"name": "Justine Bird",
		},
		map[string]interface{}{
			"id":   1,
			"name": "Marianne Rutledge",
		},
	}
}
