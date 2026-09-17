package funcs

func Swap(x, y string) (string, string) {
	return y, x
}

func Swaps(x *int, y *int) {

	var temp int
	temp = *x /* 保存 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
