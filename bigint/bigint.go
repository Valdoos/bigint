package bigint

import "fmt"

type BigInt struct {
	Value    []int8
	Positive bool
}

func NewIntFromString(s string) *BigInt {
	positive := true
	b := []byte(s)
	if b[0] == '-' {
		positive = false
		b = b[1:]
	} else if b[0] == '+' {
		b = b[1:]
	}
	value := make([]int8, len(b))
	hasNotZero := false
	for i := range value {
		value[i] = int8(b[i]) - 48
		if value[i] != 0 {
			hasNotZero = true
		}
	}
	if !hasNotZero {
		return &BigInt{
			Positive: false,
			Value:    []int8{0},
		}
	}
	return &BigInt{
		Positive: positive,
		Value:    value,
	}
}

func NewIntFromInt32(i int32) *BigInt {
	s := fmt.Sprintf("%d", i)
	return NewIntFromString(s)
}

func NewIntFromInt64(i int64) *BigInt {
	s := fmt.Sprintf("%d", i)
	return NewIntFromString(s)
}

func (this *BigInt) Equal(big *BigInt) bool {
	if this.Positive != big.Positive {
		return false
	}
	if len(this.Value) != len(big.Value) {
		return false
	}
	for i, k := range this.Value {
		if k != big.Value[i] {
			return false
		}
	}
	return true
}

func (this *BigInt) Abs() *BigInt {
	return &BigInt{
		Positive: true,
		Value:    this.Value,
	}
}

func (this *BigInt) ToString() string {
	ans := make([]byte, len(this.Value))
	for i, k := range this.Value {
		ans[i] = byte(byte(k)) + '0'
	}
	if !this.Positive && !(len(ans) == 1 && ans[0] == '0') {
		return "-" + string(ans)
	}
	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (this *BigInt) Add(big *BigInt) *BigInt {
	positive := true
	if this.Positive != big.Positive {
		if this.Positive {
			return this.Sub(big.Abs())
		} else {
			return big.Sub(this.Abs())
		}
	} else {
		positive = this.Positive
	}

	n1 := &this.Value
	n2 := &big.Value
	if len(*n1) < len(*n2) {
		n1, n2 = n2, n1
	}
	l1 := len(*n1)
	l2 := len(*n2)
	l := max(l1, l2) + 1
	ans := make([]int8, l)
	var supple int8
	l1, l2, l = l1-1, l2-1, l-1
	for l2 >= 0 {
		ans[l] = (*n1)[l1] + (*n2)[l2] + supple
		if ans[l] > 9 {
			ans[l] -= 10
			supple = 1
		} else {
			supple = 0
		}
		l1--
		l2--
		l--
	}
	for l1 >= 0 {
		ans[l] = (*n1)[l1] + supple
		if ans[l] > 9 {
			ans[l] -= 10
			supple = 1
		} else {
			supple = 0
		}
		l1--
		l--
	}
	if supple == 0 {
		ans = ans[1:]
	} else {
		ans[0] = supple
	}
	for len(ans) > 1 && ans[0] == 0 {
		ans = ans[1:]
	}
	return &BigInt{
		Value:    ans,
		Positive: positive,
	}
}

func (this *BigInt) Sub(big *BigInt) *BigInt {

	if this.Positive != big.Positive {
		if this.Positive {
			return this.Add(big.Abs())
		} else {
			nSub := big.Add(this.Abs())
			nSub.Positive = false
			return nSub
		}
	}
	ba := this.Abs()
	oa := big.Abs()
	if ba.GreaterAbs(oa) {
		tSub := ba.SubAbs(oa)
		if !this.Positive {
			tSub.Positive = false
		}
		return tSub
	} else {
		tSub := oa.SubAbs(ba)
		if this.Positive {
			tSub.Positive = false
		}
		return tSub
	}
}

func (this *BigInt) SubAbs(big *BigInt) *BigInt {
	n1 := &this.Value
	n2 := &big.Value
	if len(*n1) < len(*n2) {
		n1, n2 = n2, n1
	}
	l1 := len(*n1)
	l2 := len(*n2)
	l := max(l1, l2) + 1
	ans := make([]int8, l)
	var supple int8
	l1, l2, l = l1-1, l2-1, l-1
	for l2 >= 0 {
		ans[l] = (*n1)[l1] - (*n2)[l2] - supple
		if ans[l] < 0 {
			ans[l] += 10
			supple = 1
		} else {
			supple = 0
		}
		l1--
		l2--
		l--
	}
	for l1 >= 0 {
		ans[l] = (*n1)[l1] - supple
		if ans[l] < 0 {
			ans[l] += 10
			supple = 1
		} else {
			supple = 0
		}
		l1--
		l--
	}
	if supple == 0 {
		ans = ans[1:]
	} else {
		ans[0] -= supple
	}
	for len(ans) > 1 && ans[0] == 0 {
		ans = ans[1:]
	}
	return &BigInt{
		Value:    ans,
		Positive: true,
	}
}

func (this *BigInt) Multi(big *BigInt) *BigInt {
	positive := true
	if this.Positive != big.Positive {
		positive = false
	}

	num1 := &this.Value
	num2 := &big.Value
	l1, l2 := len(*num1), len(*num2)
	res := make([]int8, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			val := (*num1)[i] * (*num2)[j]
			res[i+j+1] += val
			if res[i+j+1] >= 10 {
				res[i+j] += (res[i+j+1] / 10)
				res[i+j+1] %= 10
			}
		}
	}
	for len(res) > 1 && res[0] == 0 {
		res = res[1:]
	}
	return &BigInt{
		Value:    res,
		Positive: positive,
	}
}

func (this *BigInt) DivInt(d int) *BigInt {
	if d == 0 {
		panic("Cannot divide by 0!")
	}
	positive := true
	if d > 0 {
		positive = this.Positive
	} else {
		d *= -1
		if this.Positive {
			positive = false
		} else {
			positive = true
		}
	}
	num := this.Value
	var ans string
	i := 0
	temp := int(num[i])
	for i < len(num)-1 && temp < d {
		i++
		temp = temp*10 + int(num[i])
	}
	i++
	for ; i < len(num); i++ {
		ans += string(byte(temp/d) + '0')
		temp = (temp%d)*10 + int(num[i])
	}
	ans += string(byte(temp/d) + '0')
	if len(ans) == 0 {
		return NewIntFromString("0")
	}
	if !positive {
		ans = "-" + ans
	}

	return NewIntFromString(ans)
}

func (a *BigInt) LessAbs(b *BigInt) bool {
	aLen := len(a.Value)
	bLen := len(b.Value)
	if aLen > bLen {
		return false
	} else if aLen < bLen {
		return true
	} else {
		for i := aLen - 1; i >= 0; i-- {
			if a.Value[i] < b.Value[i] {
				return true
			} else if a.Value[i] > b.Value[i] {
				return false
			}
		}
	}
	return false
}

func (a *BigInt) Greater(b *BigInt) bool {
	if a.GreaterAbs(b) {
		if a.Positive {
			return true
		} else {
			return false
		}
	} else {
		if b.Positive {
			return false
		} else {
			return true
		}
	}
}

func (a *BigInt) GreaterEqual(b *BigInt) bool {
	return a.Greater(b) || a.Equal(b)
}

func (a *BigInt) Less(b *BigInt) bool {
	if a.LessAbs(b) {
		if b.Positive {
			return true
		} else {
			return false
		}
	} else {
		if a.Positive {
			return false
		} else {
			return true
		}
	}
}

func (a *BigInt) LessEqual(b *BigInt) bool {
	return a.Less(b) || a.Equal(b)
}

func (a *BigInt) NotEqual(b *BigInt) bool {
	return !a.Equal(b)
}

func (this *BigInt) GreaterAbs(big *BigInt) bool {
	if len(this.Value) > len(big.Value) {
		return true
	} else if len(this.Value) < len(big.Value) {
		return false
	}
	for i, k := range this.Value {
		if k > big.Value[i] {
			return true
		} else if k < big.Value[i] {
			return false
		}
	}
	return false
}
