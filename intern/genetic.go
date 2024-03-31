package intern

import (
	"fmt"
	"math/rand"
	"sort"
)

// 种群
type population struct {
	obj                 []any                     // 个体集合
	maxSize             int                       // 个体最大长度
	target              int                       // 目标
	maxIterations       int                       // 最大迭代次数
	mutationProbability float64                   // 变异概率
	reserved            int                       // 保留个数，每次迭代保留前n各优质个体基因，之间进入下一代
	minHybridization    float64                   // 最低杂交概率
	fitness             func(any) int             // 适应度函数 any为个体 返回目标个体的适应度
	mutation            func(any) any             // 变异
	hybridization       func(any, any) (any, any) // 杂交,杂交后产生两各子代
}

// 设置适应度函数
func (p *population) SetFitness(f func(any) int) {
	p.fitness = f
}

// 设置变异函数
func (p *population) SetMutation(m func(any) any) {
	p.mutation = m
}

// 设置杂交函数
func (p *population) SetHybridization(h func(any, any) (any, any)) {
	p.hybridization = h
}

// 设置种群初始个体
func (p *population) SetObj(obj []any) {
	p.obj = obj
}

// 新建种群
func NewPopulation(maxSize int, target int, maxIterations int, mutationProbability, minHybridization float64, reserved int) *population {
	return &population{
		maxSize:             maxSize,
		target:              target,
		maxIterations:       maxIterations,
		mutationProbability: mutationProbability,
		minHybridization:    minHybridization,
		reserved:            reserved,
	}
}

// 开始迭代
func (p *population) Run() any {
	for len(p.obj) < p.maxSize {
		p.obj = append(p.obj, p.obj...)
	}
	p.obj = p.obj[:p.maxSize]
	for i := 0; i < p.maxIterations; i++ {
		fmt.Printf("正在进行第%d轮进化，当前种群的最优解是：", i)
		fmt.Println(p.obj[0])
		if p.fitness(p.obj[0]) >= p.target {
			return p.obj[0]
		}
		save := p.obj[:p.reserved]
		// 两两杂交
		for j := 0; j < len(p.obj)/2; j++ {
			// 从0-1取随机数
			g := rand.Float64()
			if g < p.minHybridization {
				p1, p2 := p.hybridization(p.obj[j], p.obj[j+1])
				b1 := rand.Float64()
				if b1 < p.mutationProbability {
					p1 = p.mutation(p1)
				}
				b2 := rand.Float64()
				if b2 < p.mutationProbability {
					p2 = p.mutation(p2)
				}
				p.obj = append(p.obj, p1, p2)
			}
		}
		p.obj = append(p.obj, save...)
		// 使用sort包对p.obj做排序,其中大小为p.fitness(any)
		sort.Slice(p.obj, func(i, j int) bool { return p.fitness(p.obj[i]) > p.fitness(p.obj[j]) })
		p.obj = p.obj[:p.maxSize]
	}
	return nil
}
func (p *population) PrintAns() {
	fmt.Println(p.obj[0])
}
