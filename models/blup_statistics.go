package models

type Statistics struct {
	MinGeneralValue    float64 // Минимальная Общая оценка по EBV
	MinEbvMilk         float64 // Минимальная Оценка удоя по EBV
	MinEbvFat          float64 // Минимальная Оценка жира по EBV
	MinEbvProtein      float64 // Минимальная Оценка белка по EBV
	MinEbvInsemenation float64 // Минимальная Оценка кратности осеменения по EBV
	MinEvbService      float64 // Минимальная Оценка длительности сервисного периода по EBV

	MaxGeneralValue    float64 // Максимальная Общая оценка по EBV
	MaxEbvMilk         float64 // Максимальная Оценка удоя по EBV
	MaxEbvFat          float64 // Максимальная Оценка жира по EBV
	MaxEbvProtein      float64 // Максимальная Оценка белка по EBV
	MaxEbvInsemenation float64 // Максимальная Оценка кратности осеменения по EBV
	MaxEvbService      float64 // Максимальная Оценка длительности сервисного периода по EBV
}

type RegionalStatistics struct {
	ID       uint `gorm:"primaryKey"`
	RegionID uint
	Region   Region
	Statistics
}

type HozStatistics struct {
	ID    uint `gorm:"primaryKey"`
	HozID uint
	Hoz   Farm
	Statistics
}
