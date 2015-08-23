Faker
=====
This package  is a port of Perl's Data::Faker library that generates fake data such as names, addresses, and phone numbers, internet domains.

=Install

$go get github.com/jalil/faker


=Use

import "fmt"
import "faker/name"

func main() {
  fmt.Println(name.FirstName()) // "Sam"
  fmt.Println(name.LastName()) // "Kitson"
}




