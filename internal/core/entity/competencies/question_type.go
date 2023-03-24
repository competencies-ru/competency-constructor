package competencies

type QuestionType uint8

const (
	// Test - тип вопроса с выбором ответа одиночным или множественным
	Test Type = iota + 1
	// Sequence - тип вопроса с установкой последовательности ответов
	Sequence
	// Conformity - тип вопроса с сопоставлением левой и правой части
	Conformity

	// Complete - тип вопроса с дополнением
	Complete

	// Essay - тип вопорса свободный ответ
	Essay
)
