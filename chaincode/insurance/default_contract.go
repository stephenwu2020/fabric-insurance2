package main

func defaultContractTypes() []ContractTypeWithUUID {
	config := []ContractTypeWithUUID{
		{"63ef076a-33a1-41d2-a9bc-2777505b014f", &ContractType{
			ShopType:        "B",
			FormulaPerDay:   "price * 0.01 + 0.05",
			MaxSumInsured:   4300.00,
			TheftInsured:    true,
			Description:     "Contract for Sea Boat",
			Conditions:      "Contract Terms here",
			Active:          true,
			MinDurationDays: 1,
			MaxDurationDays: 7,
		}},
		{"1d640cf7-9808-4c78-b7f0-55aaad02e9e5", &ContractType{
			ShopType:        "B",
			FormulaPerDay:   "price * 0.02",
			MaxSumInsured:   3500.00,
			TheftInsured:    false,
			Description:     "Insure Your Boat",
			Conditions:      "Simple contract terms.",
			Active:          true,
			MinDurationDays: 3,
			MaxDurationDays: 10,
		}},
		{"17210a72-f505-42bf-a238-65c8898477e1", &ContractType{
			ShopType:        "C",
			FormulaPerDay:   "price * 0.001 + 5.00",
			MaxSumInsured:   1500.00,
			TheftInsured:    true,
			Description:     "Car Insurance Contract",
			Conditions:      "Exemplary contract terms here.",
			Active:          true,
			MinDurationDays: 5,
			MaxDurationDays: 10,
		}},
		{"17d773dc-2624-4c22-a478-87544dd0a17f", &ContractType{
			ShopType:        "C",
			FormulaPerDay:   "price * 0.005 + 10.00",
			MaxSumInsured:   2500.00,
			TheftInsured:    true,
			Description:     "Premium Car Insurance",
			Conditions:      "Only for premium phone owners.",
			Active:          true,
			MinDurationDays: 10,
			MaxDurationDays: 20,
		}},
		{"d804f730-8c77-4583-9247-ec9e753643db", &ContractType{
			ShopType:        "M",
			FormulaPerDay:   "25.00",
			MaxSumInsured:   5000.00,
			TheftInsured:    false,
			Description:     "Short-Term Moto Insurance",
			Conditions:      "Simple contract terms here.",
			Active:          true,
			MinDurationDays: 3,
			MaxDurationDays: 25,
		}},
		{"dcee27d7-bf3c-4995-a272-8a306a35e51f", &ContractType{
			ShopType:        "M",
			FormulaPerDay:   "price * 0.001 + 10.00",
			MaxSumInsured:   3000.00,
			TheftInsured:    true,
			Description:     "Insure Ur Moto",
			Conditions:      "Just do it.",
			Active:          true,
			MinDurationDays: 1,
			MaxDurationDays: 15,
		}},
		{"c06f95d6-9b90-4d24-b8cb-f347d1b33ddf", &ContractType{
			ShopType:        "BCM",
			FormulaPerDay:   "50",
			MaxSumInsured:   3000.00,
			TheftInsured:    false,
			Description:     "Universal Insurance Contract",
			Conditions:      "Universal Contract Terms here. For all types of goods.",
			Active:          true,
			MinDurationDays: 1,
			MaxDurationDays: 10,
		}},
	}
	return config
}
