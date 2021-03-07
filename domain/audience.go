package domain

import "errors"

type TicketType int

const (
	// 一般
	TicketTypeGeneral TicketType = iota + 1
	// 中高生
	TicketTypeSecondarySchoolStudent
	// 学生(大・専門)
	TicketTypeStudent
	// 障害者
	TicketTypeHandicapped
	// 障害者(高校生以下)
	TicketTypeHandicappedStudent
	// シネマシチズン
	TicketTypeCinemaCitizen
	// シネマシチズン(60歳以上)
	TicketTypeCinemaCitizenOver60
	// 50夫婦割対象
	TicketTypeCouple50

	//以下、年齢で自動判定できるもの
	//シニア
	TicketTypeSenior
	//小学生以下
	TicketTypePrimaryStudentOrBelow
)

var TicketTypeNameMap = map[TicketType]string{
	TicketTypeGeneral:                "一般",
	TicketTypeSenior:                 "シニア(70歳以上)",
	TicketTypeStudent:                "学生(大・専)",
	TicketTypeSecondarySchoolStudent: "中・高校生",
	TicketTypePrimaryStudentOrBelow:  "幼児(3歳以上）・小学生",
	TicketTypeHandicapped:            "障がい者(学生以上)",
	TicketTypeHandicappedStudent:     "障がい者(高校生以下)",
	TicketTypeCinemaCitizen:          "シネマシチズン",
	TicketTypeCinemaCitizenOver60:    "シネマシチズン(60歳以上)",
	TicketTypeCouple50:               "夫婦50割",
}

type Audience struct {
	// 年齢
	Age int
	// 購入可能なチケット種別
	AvailableTicketTypes []TicketType
}

func isSenior(age int) bool {
	return age >= 70
}

func isPrimaryStudentOrBelow(age int) bool {
	return age <= 12
}

func NewAudience(age int, isSecondarySchoolStudent, isStudent, isHandicapped, isCinemaCitizen, isCouple50 bool) (*Audience, error) {
	//バリデーション
	if isSecondarySchoolStudent && isStudent {
		return nil, errors.New("中高生と大学生の資格は同時に満たせません")
	}

	t := make([]TicketType, 0)
	//シニア
	if isSenior(age) {
		t = append(t, TicketTypeSenior)
	}
	//小学生以下
	if isPrimaryStudentOrBelow(age) {
		t = append(t, TicketTypePrimaryStudentOrBelow)
	}
	//中高生
	if isSecondarySchoolStudent {
		t = append(t, TicketTypeSecondarySchoolStudent)
	}
	//大学生
	if isStudent {
		t = append(t, TicketTypeStudent)
	}

	//障害者
	if isHandicapped {
		// 高校生以下
		if isPrimaryStudentOrBelow(age) || isSecondarySchoolStudent {
			t = append(t, TicketTypeHandicappedStudent)
		} else {
			t = append(t, TicketTypeHandicapped)
		}
	}
	//シネマシチズン
	if isCinemaCitizen {
		//60歳以上
		if age >= 60 {
			t = append(t, TicketTypeCinemaCitizenOver60)
		} else {
			t = append(t, TicketTypeCinemaCitizen)
		}
	}
	//夫婦50割
	if isCouple50 {
		t = append(t, TicketTypeCouple50)
	}

	return &Audience{
		Age:                  age,
		AvailableTicketTypes: t,
	}, nil
}
