package competencies

import (
	"sort"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrQuestionIDEmpty               = errors.New("question: id cannot be empty")
	ErrQuestionRankSegment           = errors.New("question: rank is not included in segment")
	ErrQuestionDescriptionEmpty      = errors.New("question: description cannot be empty")
	ErrQuestionIndicatorIDEmpty      = errors.New("question: indicator id cannot be empty")
	ErrQuestionTypeInvalid           = errors.New("question: type invalid")
	ErrQuestionCompleteAnswerIsEmpty = errors.New("question: complete answer is empty")
	ErrQuestionTaskSequenceIsEmpty   = errors.New("question: test sequence is empty or has one element")
	ErrQuestionTaskConformityIsEmpty = errors.New("question: test conformity is empty or has one element")
)

const (
	minRank = 1
	maxRank = 5
)

type (
	Question struct {
		id           string
		description  string
		questionType QuestionType
		rank         int
		data         QuestionData
		indicatorID  string
	}

	QuestionData struct {
		taskPoint      *TaskPoint
		taskSequence   map[int]*TaskSequence
		taskConformity []*TaskConformity
		completeAnswer string
	}
)

type (
	QuestionParams struct {
		ID           string
		Description  string
		QuestionType QuestionType
		Rank         int
		IndicatorID  string

		TaskSequence   []TaskSequenceParams
		TaskConformity []TaskConformityParams
		TaskPoint      TaskPointParams
		CompleteAnswer string
	}
)

func NewQuestion(params QuestionParams) (*Question, error) {
	if tmp := strings.TrimSpace(params.ID); tmp == "" {
		return nil, ErrQuestionIDEmpty
	}

	if tmp := strings.TrimSpace(params.Description); tmp == "" {
		return nil, ErrQuestionDescriptionEmpty
	}

	if tmp := strings.TrimSpace(params.IndicatorID); tmp == "" {
		return nil, ErrQuestionIndicatorIDEmpty
	}

	if !params.QuestionType.IsValid() {
		return nil, ErrQuestionTypeInvalid
	}

	if params.Rank < minRank || params.Rank > maxRank {
		return nil, ErrQuestionRankSegment
	}

	data, err := createQuestionData(params)
	if err != nil {
		return nil, err
	}

	return &Question{
		id:           params.ID,
		description:  params.Description,
		questionType: params.QuestionType,
		rank:         params.Rank,
		data:         data,
		indicatorID:  params.IndicatorID,
	}, nil
}

func (q *Question) ID() string {
	return q.id
}

func (q *Question) Description() string {
	return q.description
}

func (q *Question) QuestionType() QuestionType {
	return q.questionType
}

func (q *Question) Rank() int {
	return q.rank
}

func (q *Question) Data() QuestionData {
	return q.data
}

func (q *Question) IndicatorID() string {
	return q.indicatorID
}

func (q *Question) CompleteAnswer() string {
	return q.data.completeAnswer
}

func (q *Question) TaskPoint() TaskPoint {
	if q.data.taskPoint == nil {
		return TaskPoint{}
	}

	return *q.data.taskPoint
}

func (q *Question) TaskSequence() []TaskSequence {
	buff := make([]TaskSequence, 0, len(q.data.taskSequence))

	for _, sequence := range q.data.taskSequence {
		tmp := sequence

		buff = append(buff, *tmp)
	}

	return buff
}

func (q *Question) TaskConformity() []TaskConformity {
	buff := make([]TaskConformity, 0, len(q.data.taskConformity))

	for i := 0; i < len(q.data.taskConformity); i++ {
		buff = append(buff, *q.data.taskConformity[i])
	}

	return buff
}

func createQuestionData(params QuestionParams) (QuestionData, error) {
	data := QuestionData{}

	switch params.QuestionType {
	case Sequence:
		mapping, err := taskSequenceMapping(params.TaskSequence)
		if err != nil {
			return QuestionData{}, err
		}

		data.taskSequence = mapping
	case Test:

		point, err := NewTaskPoint(params.TaskPoint)
		if err != nil {
			return QuestionData{}, err
		}

		data.taskPoint = point

	case Conformity:
		mapping, err := taskConformityMapping(params.TaskConformity)
		if err != nil {
			return QuestionData{}, err
		}

		data.taskConformity = mapping

	case Complete:
		if tmp := strings.TrimSpace(params.CompleteAnswer); tmp == "" {
			return QuestionData{}, ErrQuestionCompleteAnswerIsEmpty
		}

		data.completeAnswer = params.CompleteAnswer

	case Essay:
	}

	return data, nil
}

func taskSequenceMapping(params []TaskSequenceParams) (map[int]*TaskSequence, error) {
	if len(params) <= 1 {
		return nil, ErrQuestionTaskSequenceIsEmpty
	}

	testSequence := make(map[int]*TaskSequence, len(params))

	sort.Slice(params, func(i, j int) bool {
		return params[i].Sequence <= params[j].Sequence
	})

	if params[0].Sequence != 1 {
		return nil, ErrTaskSequenceIncorrect
	}

	for i := 0; i < len(params); i++ {
		if len(params)-1 > i && params[i+1].Sequence != (params[i].Sequence+1) {
			return nil, ErrTaskSequenceIncorrect
		}

		sequence, err := NewTaskSequence(params[i])
		if err != nil {
			return nil, err
		}

		testSequence[sequence.Sequence()] = sequence
	}

	return testSequence, nil
}

func taskConformityMapping(params []TaskConformityParams) ([]*TaskConformity, error) {
	if len(params) <= 1 {
		return nil, ErrQuestionTaskConformityIsEmpty
	}

	buff := make([]*TaskConformity, 0, len(params))

	for i := 0; i < len(params); i++ {

		conformity, err := NewTaskConformity(params[i])
		if err != nil {
			return nil, err
		}

		buff = append(buff, conformity)
	}

	return buff, nil
}
