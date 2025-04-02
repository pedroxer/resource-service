package utills

const (
	PageSize = 15
)

var IdsToConditions = map[int64]string{
	1: "Хорошее",
	2: "Плохое",
}

var ConditionsToIds = map[string]int64{
	"Хорошее": 1,
	"Плохое":  2,
}
