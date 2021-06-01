package main

// type row struct {
// 	name string
// author string
// }

// const (
// 	N_FIELDS = 10000000
// 	N_ROWS   = 10000000
// )

const (
	N_FIELDS = 2
	N_ROWS   = 2
)

func main() {
	// Prepare test
	// t := make([]db.Row, N_ROWS)
	// b := new(bytes.Buffer)
	// r := make([]db.String, N_FIELDS)

	// for i := 0; i < N_ROWS; i++ {
	// 	t[i] = make(db.Row, N_FIELDS)
	// 	for j := 0; j < N_FIELDS; j++ {
	// 		field := db.String("Field")
	// 		t[i][j] = &field
	// 	}
	// }
	// fmt.Println(t)

	// Write Test
	// start := time.Now()
	// for i := 0; i < N_FIELDS; i++ {
	// 	t[i].WriteTo(b)
	// }
	// wTime := time.Since(start)

	// Read Test
	// start = time.Now()
	// for i := 0; i < N_FIELDS; i++ {
	// 	r[i].ReadFrom(b)
	// }
	// rTime := time.Since(start)

	// Results
	// fmt.Println("Writing took ", wTime.Seconds(), "seconds")
	// fmt.Println("Reading took ", rTime.Seconds(), "seconds")
}
