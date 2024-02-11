package tahap2

func GenerateNik(gender string, year int, number_generated int) string {
	var g string
	if gender == "male" {
		g = "N"
	} else if gender == "female" {
		g = "T"
	}
	nik := "AR" + g
	return nik
}
