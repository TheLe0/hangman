package ui

func drawActor() [7]string{
	var stages [7]string

	stages[0] = `
			+---+
			|   |
			|   0
			|  /|\
			|  / \
			|
		=========
	`

	stages[1] = `
			+---+
			|   |
			|   0
			|  /|\
			|  /
			|
		=========
	`

	stages[2] = `
			+---+
			|   |
			|   0
			|  /|\
			|
			|
		=========
	`

	stages[3] = `
			+---+
			|   |
			|   0
			|  /|
			|
			|
		=========
	`

	stages[4] = `
			+---+
			|   |
			|   0
			|  /
			|
			|
		=========
	`

	stages[5] = `
			+---+
			|   |
			|   0
			|
			|
			|
		=========
	`

	stages[6] = `
			+---+
			|   |
			|    
			|
			|
			|
		=========
	`

	return stages;
}

func RemainLifes(lifes int) string {

	stages := drawActor();

	return stages[lifes];
}