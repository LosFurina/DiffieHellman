package utils

import "math/big"

func FindPrimitiveRoot(p *big.Int) *big.Int {
	one := big.NewInt(1) // 创建一个 big.Int 值为 1 的变量 one

	// 遍历可能的候选原根值
	for g := big.NewInt(2); g.Cmp(p) < 0; g.Add(g, one) {
		// 从 2 开始遍历可能的原根 g，直到 g 小于 p 为止
		if isPrimitiveRoot(g, p) {
			// 如果 g 是一个适用的原根，调用 isPrimitiveRoot 函数检查
			// 找到了一个适用的原根，返回 g
			return g
		}
	}
	// 如果没有找到原根，则返回 nil
	return nil
}

func isPrimitiveRoot(g, p *big.Int) bool {
	one := big.NewInt(1)       // 创建一个 big.Int 值为 1 的变量 one
	exp := new(big.Int).Set(p) // 创建一个新的 big.Int，其值等于 p
	exp.Sub(exp, one)          // 计算 exp = p - 1

	// 检查 g ^ (p-1) mod p 是否等于 1
	result := new(big.Int).Exp(g, exp, p) // 计算 g 的 exp 次幂对 p 取模的结果
	return result.Cmp(one) == 0           // 检查结果是否等于 1，如果等于 1，则 g 是一个原根
}
