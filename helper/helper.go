package helper

import "fmt"

func gcd(a, b int) int {
	if a == 0 {
		return b
	} else {
		return gcd(b%a, a)
	}
}
func power(m, e, n int) int {
	if e == 0 {
		return 1
	}
	if e == 1 {
		return m
	} else {
		ans := m
		for i := 1; i < e; i++ {
			ans = (ans * m) % n
		}
		return ans
	}

}
func Encryption(msg, e, n int) int {
	pp := power(msg, e, n)
	return pp
}
func Decryption(cipher, d, n int) int {
	Decrypted_msg := power(cipher, d, n)
	return Decrypted_msg
}

func StringDecryption(cipher, d, n int) byte {
	Decrypted_msg := power(cipher, d, n)
	return byte(Decrypted_msg)
}
func RSA(msg string) string {
	p, q := 9001, 13
	n := p * q
	phi := (p - 1) * (q - 1)

	var e int
	for e = 2; e < phi; e++ {
		if gcd(e, phi) == 1 {
			break
		}
	}
	dd := 0
	for i := 0; i <= 9; i++ {
		x := 1 + (i * phi)
		if x%e == 0 {
			dd = x / e
			break
		}
	}
	// m := 24
	// cipher := Encryption(m, e, n)
	// Decrypted_msg := Decryption(cipher, dd, n)
	// fmt.Println("the original msg is", m)
	// fmt.Println("the decrypted msg is ", Decrypted_msg)
	// // msg := "hello world"
	// fmt.Println("The original message is ", msg)
	data := []byte(msg)
	var encrypted []int
	for i := 0; i < len(data); i++ {
		encrypted = append(encrypted, Encryption(int(data[i]), e, n))
	}
	// fmt.Println("The encrypted message is", s
	var dycrypted []byte
	for i := 0; i < len(encrypted); i++ {
		dycrypted = append(dycrypted, StringDecryption(int(encrypted[i]), dd, n))
	}
	fmt.Println("the dycrypted message inn byte is", dycrypted)
	ans := string(dycrypted)
	// fmt.Println("the decrypted message is", string(dycrypted))
	return ans
}
