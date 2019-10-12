package utils

func SaltPasswordWithSecret(firstMd5 string) string {
	return Md5(firstMd5 + Secret)
}

func SaltPassword(password string) string {
	firstMd5 := Md5(password)
	return SaltPasswordWithSecret(firstMd5)
}