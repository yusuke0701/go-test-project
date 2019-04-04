package main

// Player は、操作キャラクター
type Player struct {
	ID      int
	HP      int // 体力
	Satiety int // 満腹度
}

// checkHP は、体力が0以下になっていないかチェックする
func (p Player) checkHP() bool {
	return p.HP <= 0
}

// checkSatiety は、満腹度が0以下になっていないかチェックする
func (p Player) checkSatiety() bool {
	return p.Satiety <= 0
}

// CheckHealth は、健康状態をチェックする。trueは健康。falseは死亡を表す
func (p Player) CheckHealth() bool {
	return (!p.checkHP() && !p.checkSatiety())
}
