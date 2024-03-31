package main

import (
	"genetic/intern"
	"math/rand"
	"time"
)

var target = 1<<20 - 1

func main() {
	p := intern.NewPopulation(500, target, 10000, 0.01, 0.9, 10)
	p.SetFitness(fitness)
	p.SetMutation(mutation)
	p.SetHybridization(hybridization)

	// 新建500个 0-700的随机数
	rand.Seed(time.Now().UnixNano())

	numbers := make([]any, 500)
	for i := 0; i < 500; i++ {
		numbers[i] = rand.Intn(100) // Generate a random number between 0 and 700
	}
	p.SetObj(numbers)
	p.Run()
	p.PrintAns()
}

func fitness(o any) int {
	return o.(int)
}

func mutation(o any) any {
	// 生成0-10的随机数i,计算2的i次方
	s := 1 << uint(rand.Intn(20))
	return o.(int) ^ s
}

func hybridization(o1, o2 any) (any, any) {
	a := o1.(int)
	b := o2.(int)
	//取a的前10位和后10位
	af := a & 1023
	al := a >> 10
	bf := b & 1023
	bl := b >> 10

	// 交换a,b 的后五位
	ta := al<<10 + bf
	tb := bl<<10 + af
	return ta, tb
}
