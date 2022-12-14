deadline := flag.String("deadline", "", "The deadline for the countdown timer in RFC3339 format (e.g. 2019-12-25T15:00:00+01:00)")
	// flag.Parse()
	// // fmt.Println(deadline)

	// // deadline := timerGo("2022-12-01T55:00:00+01:00")

	// if *deadline == "" {
	// 	fmt.Println(time.Parse(time.RFC3339, *deadline))
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	// v, err := time.Parse(time.RFC3339, *deadline)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(time.Parse(time.RFC3339, *deadline))
	// fmt.Println(v)