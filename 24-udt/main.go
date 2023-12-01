package main

import "fmt"

func main() {
	var e1 Empty
	var p1 Person
	p1.Id = 101
	p1.Name = "JP"
	p1.Email = "JP@gmail.com"
	fmt.Println(e1)
	fmt.Println(p1)

	emp1 := Employee{Id: 101, Name: "Jiten", Email: "Jitenp@outlook.com", Address: Address{City: "Bangalore", Pincode: "560096"}}
	emp1.Address.Pincode = "560086"
	fmt.Println(emp1)

	//cmp1 := Company{Id: 101, Name: "ABC Solutions", Email: "abc@outlook.com", Address: Address{City: "Bangalore", Pincode: "560096"}}

	var cmp2 Company
	cmp2.Id = 12312
	cmp2.Name = "ABC Sol"
	cmp2.Email = "Abc@AbcSol.com"
	cmp2.Status = "Active" // status of Company
	// since address is a promoted field, can call direcly with cmp2 object

	cmp2.City = "Bangalore"
	cmp2.Pincode = "560011"
	cmp2.Address.Status = "Active" // since there is a status field in both, call specifically
	fmt.Println(cmp2)

	var cc1 ColorCode
	cc1.int = 10001
	cc1.string = "grey"
	fmt.Println(cc1)
	println(cc1.int, cc1.string)

	stu1 := Student{Id: 101, Name: "Jiten", Address: struct{ City, Country, Pincode string }{City: "Bangalore", Pincode: "560096", Country: "India"}}
	fmt.Println(stu1)

	addr1 := struct{ City, Country, Pincode string }{City: "Bangalore", Pincode: "560096", Country: "India"} // struct on the fly

	fmt.Println(addr1)

	var socialMediaa struct{ LinkedIn, Twitter, Youtube string } // a variable with the structure but no struct name, only a variable name
	socialMediaa.LinkedIn = "linkedin.com/jpalaparti"
	socialMediaa.Twitter = "Jiten_1981"
	socialMediaa.Youtube = "Jiten.Palaparthi@Gmail.Com"
	fmt.Println(socialMediaa)

}

type Empty struct{} // empty struct .. no fields
type Person struct {
	Id          int
	Name, Email string
}

type Address struct {
	City, Pincode, Status string
}

type Employee struct {
	Id          int
	Name, Email string
	Address     Address // composition
}

type Company struct {
	Id                  int
	Name, Email, Status string
	Address             // promoted field. Dont give name of the object instead just give the type Name
}

type ColorCode struct { // struct with anonymous fieds
	int
	string
}

type Student struct {
	Id      int
	Name    string
	Address struct {
		City, Country, Pincode string
	}
}

// there is no class
// there is no inheritance in Go
// empty struct
// embedded struct
// composition of stuct object
// struct with anonymous fields
// struct with promoted field
