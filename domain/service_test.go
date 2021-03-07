package domain_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/en-ken/movie-ticket-domain-modeling/domain"
	"github.com/google/go-cmp/cmp"
)

func TestCalculateTicketPrice(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	time.Local = loc

	type args struct {
		movie     *domain.Movie
		audiences []*domain.Audience
	}

	//シネマシチズン
	citizen, _ := domain.NewAudience(30, false, false, false, true, false)
	//シネマシチズン(60歳以上)
	citizenSenior, _ := domain.NewAudience(60, false, false, false, true, false)
	//一般
	general, _ := domain.NewAudience(30, false, false, false, false, false)
	//シニア
	senior, _ := domain.NewAudience(70, false, false, false, false, false)
	//学生
	student, _ := domain.NewAudience(20, false, true, false, false, false)
	//中高生
	ssStudent, _ := domain.NewAudience(17, true, false, false, false, false)
	//小学生
	psStudent, _ := domain.NewAudience(10, false, false, false, false, false)
	//障がい者
	handicapped, _ := domain.NewAudience(25, false, false, true, false, false)
	//障がい者(中高生)
	handicappedStudent, _ := domain.NewAudience(13, true, false, true, false, false)
	//夫婦50割
	couple1, _ := domain.NewAudience(50, false, false, false, false, true)
	couple2, _ := domain.NewAudience(45, false, false, false, false, true)

	tests := []struct {
		name string
		args args
		want *domain.PurchaseInfo
	}{
		{
			name: "通常料金: 一般1枚、シニア1枚、学生1枚、中高生1枚、小学生1枚、障がい者1枚、障がい者(中高生以下)1枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 2, 12, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					general,
					senior,
					student,
					ssStudent,
					psStudent,
					handicapped,
					handicappedStudent,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 8300,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "一般",
						Price:          1800,
					},
					{
						TicketTypeName: "シニア(70歳以上)",
						Price:          1100,
					},
					{
						TicketTypeName: "学生(大・専)",
						Price:          1500,
					},
					{
						TicketTypeName: "中・高校生",
						Price:          1000,
					},
					{
						TicketTypeName: "幼児(3歳以上）・小学生",
						Price:          1000,
					},
					{
						TicketTypeName: "障がい者(学生以上)",
						Price:          1000,
					},
					{
						TicketTypeName: "障がい者(高校生以下)",
						Price:          900,
					},
				},
			},
		},
		{
			name: "通常料金: シネマシチズン1枚, シネマシチズン(60歳以上)1枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 2, 12, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					citizen,
					citizenSenior,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 2000,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "シネマシチズン",
						Price:          1000,
					},
					{
						TicketTypeName: "シネマシチズン(60歳以上)",
						Price:          1000,
					},
				},
			},
		},
		{
			name: "通常料金: 夫婦50割2枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 2, 12, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					couple1,
					couple2,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 2200,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "夫婦50割",
						Price:          1100,
					},
					{
						TicketTypeName: "夫婦50割",
						Price:          1100,
					},
				},
			},
		},

		{
			name: "平日レイトショー料金: 一般1枚、シニア1枚、学生1枚、中高生1枚、小学生1枚、障がい者1枚、障がい者(中高生以下)1枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 2, 20, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					general,
					senior,
					student,
					ssStudent,
					psStudent,
					handicapped,
					handicappedStudent,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 7600,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "一般",
						Price:          1300,
					},
					{
						TicketTypeName: "シニア(70歳以上)",
						Price:          1100,
					},
					{
						TicketTypeName: "学生(大・専)",
						Price:          1300,
					},
					{
						TicketTypeName: "中・高校生",
						Price:          1000,
					},
					{
						TicketTypeName: "幼児(3歳以上）・小学生",
						Price:          1000,
					},
					{
						TicketTypeName: "障がい者(学生以上)",
						Price:          1000,
					},
					{
						TicketTypeName: "障がい者(高校生以下)",
						Price:          900,
					},
				},
			},
		},
		{
			name: "平日レイトショー料金: シネマシチズン1枚, シネマシチズン(60歳以上)1枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 2, 20, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					citizen,
					citizenSenior,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 2000,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "シネマシチズン",
						Price:          1000,
					},
					{
						TicketTypeName: "シネマシチズン(60歳以上)",
						Price:          1000,
					},
				},
			},
		},
		{
			name: "平日レイトショー料金: 夫婦50割2枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 2, 20, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					couple1,
					couple2,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 2200,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "夫婦50割",
						Price:          1100,
					},
					{
						TicketTypeName: "夫婦50割",
						Price:          1100,
					},
				},
			},
		},

		{
			name: "休日料金: 一般1枚、シニア1枚、学生1枚、中高生1枚、小学生1枚、障がい者1枚、障がい者(中高生以下)1枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 2, 11, 12, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					general,
					senior,
					student,
					ssStudent,
					psStudent,
					handicapped,
					handicappedStudent,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 8300,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "一般",
						Price:          1800,
					},
					{
						TicketTypeName: "シニア(70歳以上)",
						Price:          1100,
					},
					{
						TicketTypeName: "学生(大・専)",
						Price:          1500,
					},
					{
						TicketTypeName: "中・高校生",
						Price:          1000,
					},
					{
						TicketTypeName: "幼児(3歳以上）・小学生",
						Price:          1000,
					},
					{
						TicketTypeName: "障がい者(学生以上)",
						Price:          1000,
					},
					{
						TicketTypeName: "障がい者(高校生以下)",
						Price:          900,
					},
				},
			},
		},
		{
			name: "休日料金: シネマシチズン1枚, シネマシチズン(60歳以上)1枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 2, 11, 12, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					citizen,
					citizenSenior,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 2300,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "シネマシチズン",
						Price:          1300,
					},
					{
						TicketTypeName: "シネマシチズン(60歳以上)",
						Price:          1000,
					},
				},
			},
		},
		{
			name: "休日料金: 夫婦50割2枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 2, 12, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					couple1,
					couple2,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 2200,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "夫婦50割",
						Price:          1100,
					},
					{
						TicketTypeName: "夫婦50割",
						Price:          1100,
					},
				},
			},
		},

		{
			name: "映画の日: 一般1枚、シニア1枚、学生1枚、中高生1枚、小学生1枚、障がい者1枚、障がい者(中高生以下)1枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 1, 12, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					general,
					senior,
					student,
					ssStudent,
					psStudent,
					handicapped,
					handicappedStudent,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 7200,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "一般",
						Price:          1100,
					},
					{
						TicketTypeName: "シニア(70歳以上)",
						Price:          1100,
					},
					{
						TicketTypeName: "学生(大・専)",
						Price:          1100,
					},
					{
						TicketTypeName: "中・高校生",
						Price:          1000,
					},
					{
						TicketTypeName: "幼児(3歳以上）・小学生",
						Price:          1000,
					},
					{
						TicketTypeName: "障がい者(学生以上)",
						Price:          1000,
					},
					{
						TicketTypeName: "障がい者(高校生以下)",
						Price:          900,
					},
				},
			},
		},
		{
			name: "映画の日: シネマシチズン1枚, シネマシチズン(60歳以上)1枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 1, 12, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					citizen,
					citizenSenior,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 2100,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "シネマシチズン",
						Price:          1100,
					},
					{
						TicketTypeName: "シネマシチズン(60歳以上)",
						Price:          1000,
					},
				},
			},
		},
		{
			name: "映画の日: 夫婦50割2枚",
			args: args{
				movie: domain.NewMovie(time.Date(2021, 3, 1, 12, 0, 0, 0, time.Local)),
				audiences: []*domain.Audience{
					couple1,
					couple2,
				},
			},
			want: &domain.PurchaseInfo{
				TotalPrice: 2200,
				Tickets: []*domain.Ticket{
					{
						TicketTypeName: "夫婦50割",
						Price:          1100,
					},
					{
						TicketTypeName: "夫婦50割",
						Price:          1100,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := domain.CalculateTicketPrice(tt.args.movie, tt.args.audiences...); !reflect.DeepEqual(got, tt.want) {
				if len(tt.args.audiences) >= 3 {
					t.Logf("%v", tt.args.audiences[2])
				}
				t.Errorf("CalculateTicketPrice() = %v, want %v, diff %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
