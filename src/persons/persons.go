package persons

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// here we define the enum type where discribes the profile type of the agent
// we map the const to print out the actual name for a more "human-readable" info
type Profile int

const (
	User Profile = iota + 1
	Viewer
	Analist
	Admin
)

var profileType = map[Profile]string{
	User:    "user",
	Viewer:  "viewer",
	Analist: "analist",
	Admin:   "admin",
}

// here we define the enum type where discribes the type of customer created
// we map the const to print out the actual name for a more "human-readable" info
type Entry int

const (
	endUser Entry = iota + 1
	organization
	departament
)

var entryType = map[Entry]string{
	endUser:      "end-user",
	organization: "organization",
	departament:  "departament",
}

// here we define the contructors for the customers and agents
// type Person struct {
// 	cpfCnpj  string
// 	fullName string
// 	username string
// 	password string
// 	email    string
// 	contact  string
// }

// TODO
// teams	  Teams
// type Agent struct {
// 	agent     Person
// 	agentType Profile
// }

// type Customer struct {
// 	customer     Person
// 	customerType Entry
// }

func (pp Profile) String() string {
	return profileType[pp]
}

func (ee Entry) String() string {
	return entryType[ee]
}

func Persons() {

	var choice int8

	db, err := sql.Open("mysql", "openHive:openHive@tcp(127.0.0.1:3306)/openhive")
	if err != nil {
		fmt.Println("Error")
		panic(err.Error())
	}

	// customers := []Customer{}
	// agents := []Agent{}

	fmt.Println("What do you want to do?")
	fmt.Println("1 - Create new record")
	fmt.Println("2 - Update a record")
	fmt.Println("3 - Delete a record")
	fmt.Println("4 - List all Records")
	fmt.Println("5 - Exit")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		{
			fmt.Println("1 for Customer, 2 to Agent")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				newCustomer(db)

			case 2:
				newAgent(db)

			default:
				fmt.Println("Undefined")
			}
		}

	}
}

func newCustomer(db *sql.DB) {
	var cpfCnpj, fullName, username, password, email, contact string
	var customerType Entry

	fmt.Print("Full name: ")
	fmt.Scan(&fullName)
	fmt.Print("CPF or CNPJ: ")
	fmt.Scan(&cpfCnpj)
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Contact: ")
	fmt.Scan(&contact)
	fmt.Print("Customer type? 1 for End-user, 2 for Company or 3 for Departament: ")
	fmt.Scan(&customerType)

	// newCustomer := &Customer{
	// 	customer: Person{
	// 		fullName: fullName,
	// 		username: username,
	// 		cpfCnpj:  cpfCnpj,
	// 		password: password,
	// 		email:    email,
	// 		contact:  contact,
	// 	},
	// 	customerType: customerType,
	// }
	_, err := db.Exec(`INSERT INTO users 	
								(full_name, cpf_cnpj, username, password, email, 
								contact, entry_type)
							VALUES		
								(?, ?, ?, ?, ?, ?, ?)`,
		fullName, cpfCnpj, username, password, email,
		contact, customerType)
	if err != nil {
		fmt.Println("Error")
		panic(err.Error())
	}
	fmt.Println("Customer added!")
}

func newAgent(db *sql.DB) {
	var cpfCnpj, fullName, username, password, email, contact string
	var agentType Profile

	fmt.Print("Full name: ")
	fmt.Scanln(&fullName)
	fmt.Print("CPF or CNPJ:")
	fmt.Scanln(&cpfCnpj)
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	fmt.Print("Contact: ")
	fmt.Scanln(&contact)
	fmt.Print("Customer type? 1 for User, 2 for Viewer, 3 for Analist or 4 for Admin")
	fmt.Scanln(&agentType)

	// newAgent := &Agent{
	// 	agent: Person{
	// 		fullName: fullName,
	// 		cpfCnpj:  cpfCnpj,
	// 		username: username,
	// 		password: password,
	// 		email:    email,
	// 		contact:  contact,
	// 	},
	// 	agentType: agentType,
	// }
	_, err := db.Exec(`INSERT INTO users 	
								(full_name, cpf_cnpj, username, password, email, 
								contact, profile_type)
							VALUES		
								(?, ?, ?, ?, ?, ?, ?)`,
		fullName, cpfCnpj, username, password, email,
		contact, agentType)
	if err != nil {
		fmt.Println("Error")
		panic(err.Error())
	}
	fmt.Println("Agent added!")
}
