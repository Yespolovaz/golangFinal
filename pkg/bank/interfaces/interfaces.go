package interfaces

type ResponseAccount struct {
    ID uint
    Name string
    Balance int
}

type ResponseUser struct {
    ID uint
    Username string
    Email string
    Accounts []ResponseAccount
}