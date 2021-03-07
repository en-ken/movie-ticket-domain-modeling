package domain

type QualificationPriceMap map[Qualification]int

var PriceMap = map[TimeCategory]QualificationPriceMap{
	TimeCategoryWeekDayGeneral: {QualificationCinemaCitizen: 1000, QualificationCinemaCitizenOver60: 1000, QualificationGeneral: 1800, QualificationSenior: 1100, QualificationStudent: 1500, QualificationSecondarySchoolStudent: 1000, QualificationPrimaryStudentOrBelow: 1000, QualificationHandicapped: 1000, QualificationHandicappedStudent: 900, QualificationCouple50: 1100},
	TimeCategoryWeekDayLate:    {QualificationCinemaCitizen: 1000, QualificationCinemaCitizenOver60: 1000, QualificationGeneral: 1300, QualificationSenior: 1100, QualificationStudent: 1300, QualificationSecondarySchoolStudent: 1000, QualificationPrimaryStudentOrBelow: 1000, QualificationHandicapped: 1000, QualificationHandicappedStudent: 900, QualificationCouple50: 1100},
	TimeCategoryWeekendGeneral: {QualificationCinemaCitizen: 1300, QualificationCinemaCitizenOver60: 1000, QualificationGeneral: 1800, QualificationSenior: 1100, QualificationStudent: 1500, QualificationSecondarySchoolStudent: 1000, QualificationPrimaryStudentOrBelow: 1000, QualificationHandicapped: 1000, QualificationHandicappedStudent: 900, QualificationCouple50: 1100},
	TimeCategoryWeekendLate:    {QualificationCinemaCitizen: 1000, QualificationCinemaCitizenOver60: 1000, QualificationGeneral: 1300, QualificationSenior: 1100, QualificationStudent: 1300, QualificationSecondarySchoolStudent: 1000, QualificationPrimaryStudentOrBelow: 1000, QualificationHandicapped: 1000, QualificationHandicappedStudent: 900, QualificationCouple50: 1100},
	TimeCategoryMovieDay:       {QualificationCinemaCitizen: 1100, QualificationCinemaCitizenOver60: 1000, QualificationGeneral: 1100, QualificationSenior: 1100, QualificationStudent: 1100, QualificationSecondarySchoolStudent: 1000, QualificationPrimaryStudentOrBelow: 1000, QualificationHandicapped: 1000, QualificationHandicappedStudent: 900, QualificationCouple50: 1100},
}

type Ticket struct {
	TicketTypeName string
	Price          int
}

type PurchaseInfo struct {
	//合計料金
	TotalPrice int
	Tickets    []*Ticket
}

func CalculateTicketPrice(movie *Movie, audiences ...*Audience) *PurchaseInfo {
	prices := PriceMap[movie.TimeCategory]

	total := 0
	tickets := make([]*Ticket, len(audiences))
	for i, aud := range audiences {
		ticket := &Ticket{
			TicketTypeName: QualificationNameMap[QualificationGeneral],
			Price:          prices[QualificationGeneral],
		}
		for _, q := range aud.Qualifications {
			p := prices[q]
			//最低料金を適用する
			if ticket.Price >= p {
				ticket = &Ticket{
					TicketTypeName: QualificationNameMap[q],
					Price:          p,
				}
			}
		}
		tickets[i] = ticket
		total += ticket.Price
	}

	return &PurchaseInfo{
		TotalPrice: total,
		Tickets:    tickets,
	}
}
