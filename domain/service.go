package domain

type QualificationPriceMap map[TicketType]int

var PriceMap = map[TimeCategory]QualificationPriceMap{
	TimeCategoryWeekDayGeneral: {TicketTypeCinemaCitizen: 1000, TicketTypeCinemaCitizenOver60: 1000, TicketTypeGeneral: 1800, TicketTypeSenior: 1100, TicketTypeStudent: 1500, TicketTypeSecondarySchoolStudent: 1000, TicketTypePrimaryStudentOrBelow: 1000, TicketTypeHandicapped: 1000, TicketTypeHandicappedStudent: 900, TicketTypeCouple50: 1100},
	TimeCategoryWeekDayLate:    {TicketTypeCinemaCitizen: 1000, TicketTypeCinemaCitizenOver60: 1000, TicketTypeGeneral: 1300, TicketTypeSenior: 1100, TicketTypeStudent: 1300, TicketTypeSecondarySchoolStudent: 1000, TicketTypePrimaryStudentOrBelow: 1000, TicketTypeHandicapped: 1000, TicketTypeHandicappedStudent: 900, TicketTypeCouple50: 1100},
	TimeCategoryWeekendGeneral: {TicketTypeCinemaCitizen: 1300, TicketTypeCinemaCitizenOver60: 1000, TicketTypeGeneral: 1800, TicketTypeSenior: 1100, TicketTypeStudent: 1500, TicketTypeSecondarySchoolStudent: 1000, TicketTypePrimaryStudentOrBelow: 1000, TicketTypeHandicapped: 1000, TicketTypeHandicappedStudent: 900, TicketTypeCouple50: 1100},
	TimeCategoryWeekendLate:    {TicketTypeCinemaCitizen: 1000, TicketTypeCinemaCitizenOver60: 1000, TicketTypeGeneral: 1300, TicketTypeSenior: 1100, TicketTypeStudent: 1300, TicketTypeSecondarySchoolStudent: 1000, TicketTypePrimaryStudentOrBelow: 1000, TicketTypeHandicapped: 1000, TicketTypeHandicappedStudent: 900, TicketTypeCouple50: 1100},
	TimeCategoryMovieDay:       {TicketTypeCinemaCitizen: 1100, TicketTypeCinemaCitizenOver60: 1000, TicketTypeGeneral: 1100, TicketTypeSenior: 1100, TicketTypeStudent: 1100, TicketTypeSecondarySchoolStudent: 1000, TicketTypePrimaryStudentOrBelow: 1000, TicketTypeHandicapped: 1000, TicketTypeHandicappedStudent: 900, TicketTypeCouple50: 1100},
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
			TicketTypeName: TicketTypeNameMap[TicketTypeGeneral],
			Price:          prices[TicketTypeGeneral],
		}
		for _, q := range aud.AvailableTicketTypes {
			p := prices[q]
			//購入可能チケットから最低料金を適用する
			if ticket.Price >= p {
				ticket = &Ticket{
					TicketTypeName: TicketTypeNameMap[q],
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
