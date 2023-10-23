func isMatch(s string, p string) bool {
    strLen := len(s) + 1
    regLen := len(p) + 1

    isMatch := make([][]bool, strLen)
    isMatch[0] = make([]bool, regLen)

    for j := 0; j < regLen; j++ {
        if j == 0 {
            isMatch[0][j] = true
        } else if j > 1 {
            if p[j-1] == '*' && isMatch[0][j-2] {
                isMatch[0][j] = true
            }
        }
    }

    for i := 1; i < strLen; i++ {
        isMatch[i] = make([]bool, regLen)
        for j := 1; j < regLen; j++ {
            if s[i-1] == p[j-1] || p[j-1] == '.' {
                isMatch[i][j] = isMatch[i-1][j-1]
            } else if p[j-1] == '*' {
                lastMatch := s[i-1] == p[j-2] || p[j-2] == '.'
                isMatch[i][j] = isMatch[i][j-2] || (lastMatch && isMatch[i-1][j])
            }
        }
    }
 
    return isMatch[strLen-1][regLen-1]
}
