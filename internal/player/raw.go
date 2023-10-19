// Этот модуль содержит энтити игрока,
// все подробности которые мы хотим экспортировать
// из страницы
package player

// Raw содержит чистые строки, как были получены из HTML
type Raw struct {
	FirstName      string
	LastName       string
	Nationality    string
	DateOfBirth    string
	Age            string
	CountryOfBirth string
	Position       string
	Height         string
	Foot           string
}
