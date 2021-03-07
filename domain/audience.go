package domain

import "errors"

type Qualification int

const (
	// 一般
	QualificationGeneral Qualification = iota + 1
	// 中高生
	QualificationSecondarySchoolStudent
	// 学生(大・専門)
	QualificationStudent
	// 障害者
	QualificationHandicapped
	// 障害者(高校生以下)
	QualificationHandicappedStudent
	// シネマシチズン
	QualificationCinemaCitizen
	// シネマシチズン(60歳以上)
	QualificationCinemaCitizenOver60
	// 50夫婦割対象
	QualificationCouple50

	//以下、年齢で自動判定できるもの
	//シニア
	QualificationSenior
	//小学生以下
	QualificationPrimaryStudentOrBelow
)

var QualificationNameMap = map[Qualification]string{
	QualificationGeneral:                "一般",
	QualificationSenior:                 "シニア(70歳以上)",
	QualificationStudent:                "学生(大・専)",
	QualificationSecondarySchoolStudent: "中・高校生",
	QualificationPrimaryStudentOrBelow:  "幼児(3歳以上）・小学生",
	QualificationHandicapped:            "障がい者(学生以上)",
	QualificationHandicappedStudent:     "障がい者(高校生以下)",
	QualificationCinemaCitizen:          "シネマシチズン",
	QualificationCinemaCitizenOver60:    "シネマシチズン(60歳以上)",
	QualificationCouple50:               "夫婦50割",
}

type Audience struct {
	// 年齢
	Age int
	// 資格(高校生かつ障害者がありうるのでスライス)
	Qualifications []Qualification
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

	q := make([]Qualification, 0)
	//シニア
	if isSenior(age) {
		q = append(q, QualificationSenior)
	}
	//小学生以下
	if isPrimaryStudentOrBelow(age) {
		q = append(q, QualificationPrimaryStudentOrBelow)
	}
	//中高生
	if isSecondarySchoolStudent {
		q = append(q, QualificationSecondarySchoolStudent)
	}
	//大学生
	if isStudent {
		q = append(q, QualificationStudent)
	}

	//障害者
	if isHandicapped {
		// 高校生以下
		if isPrimaryStudentOrBelow(age) || isSecondarySchoolStudent {
			q = append(q, QualificationHandicappedStudent)
		} else {
			q = append(q, QualificationHandicapped)
		}
	}
	//シネマシチズン
	if isCinemaCitizen {
		//60歳以上
		if age >= 60 {
			q = append(q, QualificationCinemaCitizenOver60)
		} else {
			q = append(q, QualificationCinemaCitizen)
		}
	}
	if isCouple50 {
		q = append(q, QualificationCouple50)
	}

	return &Audience{
		Age:            age,
		Qualifications: q,
	}, nil
}
