package build_pattern

//### 难易度
type Level int

const (
	Easy = iota
	Mid
	Hard
)

type LevelObj interface {
	TellYouLevel() string
}

type EasyLevel struct {
}

func (level *EasyLevel) TellYouLevel() string {
	return "easy"
}

type MidLevel struct {
}

func (level *MidLevel) TellYouLevel() string {
	return "mid"
}

type HardLevel struct {
}

func (level *HardLevel) TellYouLevel() string {
	return "hard"
}

//### 科目
type Subject int

const (
	Math = iota
	English
	Chinese
)

type SubjectObj interface {
	TellYouSubject() string
}

type MathSubject struct {
}

func (subject *MathSubject) TellYouSubject() string {
	return "math"
}

type EnglishSubject struct {
}

func (subject *EnglishSubject) TellYouSubject() string {
	return "english"
}

type ChineseSubject struct {
}

func (subject *ChineseSubject) TellYouSubject() string {
	return "chinese"
}

type AbstractFactoryAction interface {
	GetLevel(level Level) LevelObj
	GetSubject(subject Subject) SubjectObj
}

type AbstractFactory struct {
}

func (factory *AbstractFactory) GetLevel(level Level) LevelObj {
	switch level {
	case Easy:
		return &EasyLevel{}
	case Mid:
		return &MidLevel{}
	case Hard:
		return &HardLevel{}
	default:
		panic("unsupport level")
	}
}

func (factory *AbstractFactory) GetSubject(subject Subject) SubjectObj {
	switch subject {
	case Math:
		return &MathSubject{}
	case Chinese:
		return &ChineseSubject{}
	case English:
		return &EnglishSubject{}
	default:
		panic("unsupport subject")
	}
}
