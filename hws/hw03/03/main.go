package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func DedupByEmail(users []User) []User {
	seen := make(map[string]bool, len(users))
	result := make([]User, 0, len(users))

	for _, u := range users {
		if seen[u.Email] {
			continue
		}
		seen[u.Email] = true
		result = append(result, u)
	}
	return result
}

func main() {
	users := []User{
		{1, "Alice Johnson", "alice.j@gmail.com"},
		{2, "Bob Smith", "bob.smith@company.com"},
		{3, "Carol White", "carol.w@yahoo.com"},
		{4, "Dave Brown", "dave.b@outlook.com"},
		{5, "Eve Davis", "eve.d@mail.ru"},
		{6, "Frank Miller", "frank.m@gmail.com"},
		{7, "Grace Wilson", "grace.w@company.com"},
		{8, "Hugo Moore", "hugo.m@yahoo.com"},
		{9, "Iris Taylor", "iris.t@outlook.com"},
		{10, "Jack Anderson", "jack.a@mail.ru"},
		{11, "Kate Thomas", "kate.t@gmail.com"},
		{12, "Leo Jackson", "leo.j@company.com"},
		{13, "Mia Harris", "mia.h@yahoo.com"},
		{14, "Noah Martin", "noah.m@outlook.com"},
		{15, "Olivia Garcia", "olivia.g@mail.ru"},
		{16, "Paul Martinez", "paul.m@gmail.com"},
		{17, "Ruby Robinson", "ruby.r@company.com"},
		{18, "Steve Clark", "steve.c@yahoo.com"},
		{19, "Tara Rodriguez", "tara.r@outlook.com"},
		{20, "Uma Lewis", "uma.l@mail.ru"},
		{21, "Victor Lee", "victor.l@gmail.com"},
		{22, "Wendy Walker", "wendy.w@company.com"},
		{23, "Xavier Hall", "xavier.h@yahoo.com"},
		{24, "Yara Allen", "yara.a@outlook.com"},
		{25, "Zach Young", "zach.y@mail.ru"},
		{26, "Amy King", "amy.k@gmail.com"},
		{27, "Brian Wright", "brian.w@company.com"},
		{28, "Chloe Lopez", "chloe.l@yahoo.com"},
		{29, "Dan Hill", "dan.h@outlook.com"},
		{30, "Emma Scott", "emma.s@mail.ru"},
		{31, "Finn Green", "finn.g@gmail.com"},
		{32, "Gina Adams", "gina.a@company.com"},
		{33, "Henry Baker", "henry.b@yahoo.com"},
		{34, "Ivy Nelson", "ivy.n@outlook.com"},
		{35, "Jake Carter", "jake.c@mail.ru"},
		{36, "Alice Smith", "alice.j@gmail.com"},
		{37, "Bob Jones", "bob.smith@company.com"},
		{38, "Carol Lee", "carol.w@yahoo.com"},
		{39, "Dave Wilson", "dave.b@outlook.com"},
		{40, "Eve Brown", "eve.d@mail.ru"},
		{41, "Frank Taylor", "frank.m@gmail.com"},
		{42, "Grace Moore", "grace.w@company.com"},
		{43, "Hugo Anderson", "hugo.m@yahoo.com"},
		{44, "Iris Thomas", "iris.t@outlook.com"},
		{45, "Jack Jackson", "jack.a@mail.ru"},
		{46, "Kate Harris", "kate.t@gmail.com"},
		{47, "Leo Martin", "leo.j@company.com"},
		{48, "Mia Garcia", "mia.h@yahoo.com"},
		{49, "Noah Martinez", "noah.m@outlook.com"},
		{50, "Olivia Robinson", "olivia.g@mail.ru"},
	}

	unique := DedupByEmail(users)

	fmt.Printf("Было записей: %d, осталось после дедупликации: %d\n", len(users), len(unique))
	for _, u := range unique {
		fmt.Printf("%2d  %-16s %s\n", u.ID, u.Name, u.Email)
	}
}
