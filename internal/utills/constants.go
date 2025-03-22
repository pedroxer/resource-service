package utills

const (
	PageSize = 15
)

var IdsToConditions = map[int64]string{
	1: "Новый",
	2: "В порядке",
	3: "Сломан",
}

var ConditionsToIds = map[string]int64{
	"Новый":     1,
	"В порядке": 2,
	"Сломан":    3,
}
