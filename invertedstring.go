package stringinverted

const strShow = `## 题目说明 ##
将字符串反序，但A,a,z,Z的字符串所在位置不变，其他依减前移。`

func invert(array string) string{
    //字符串转为字符数组
    if 0 == len(array) {
        return ""
    }
    runes := []rune(array)
    length := len(runes)
    if length - 1 < 0 {
        return ""
    }
    for from, to := 0, length - 1;from < to;from, to = from + 1, to - 1 {
        from_flag := isMove(runes[from])
        to_flag := isMove(runes[to])
        if true == from_flag {
            from++
        }
        if true == to_flag {
            to--
        }
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
}

func isMove(s rune) bool {
    if 'A' == s || 'a' == s || 'Z' == s || 'z' == s {
        return true
    }
    return false
}
