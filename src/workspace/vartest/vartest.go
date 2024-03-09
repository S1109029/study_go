package vartest

func Vartest() (int,int,[]byte) {
    a := []int{}
    bufa := make([]byte, 0, 1024)
    for i := 0; i < 10; i++ {
        a = append(a, i)
    }
    return len(a),cap(a),bufa
}