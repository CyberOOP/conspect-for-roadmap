package homework

import (
	"github.com/k0kubun/pp"
)

func Mass() {
	arr := [5]string{
		"первый",
		"второй",
		"третий",
		"четвёртый",
		"пятый",
	}

	pp.Println(arr)

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	pp.Println(arr)
}
