package pkg

func tpfn1[T any]() T { var zero T; return zero }

func tpfn() {
	if tpfn1[bool]() == true { //@ diag(`simplified to tpfn1[bool]()`)
	}
	if tpfn1[any]() == true {
	}
}

func tpfn2[T bool](x T) {
	if x == true { //@ diag(`omit comparison to bool constant`)
	}
}

func tpfn3[T ~bool](x T) {
	if x == true { //@ diag(`omit comparison to bool constant`)
	}
}

type MyBool bool

func tpfn4[T bool | MyBool](x T) {
	if x == true { //@ diag(`omit comparison to bool constant`)
	}
}
