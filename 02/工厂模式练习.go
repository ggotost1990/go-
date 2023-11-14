package main

import "fmt"

// 工厂抽象层
type AbstructFactory interface {
	CreateGraphics() AbstructGraphics
	CreateMemory() AbstructMemory
	CreateCpu() AbstructCpu
}

// 产品抽象层
type AbstructGraphics interface {
	Display()
}

type AbstructMemory interface {
	Storage()
}

type AbstructCpu interface {
	Calculate()
}

// 实现层

// intel产品族
type IntelGraphics struct {
}

func (ig *IntelGraphics) Display() {
	fmt.Println("英特尔显卡")
}

type IntelMemory struct {
}

func (im *IntelMemory) Storage() {
	fmt.Println("英特尔内存")
}

type IntelCpu struct {
}

func (ic *IntelCpu) Calculate() {
	fmt.Println("英特尔cpu")
}

type IntelFactory struct {
}

func (i *IntelFactory) CreateGraphics() AbstructGraphics {
	var intelGraphics AbstructGraphics
	intelGraphics = new(IntelGraphics)
	return intelGraphics
}

func (i *IntelFactory) CreateMemory() AbstructMemory {
	var intelMemory AbstructMemory
	intelMemory = new(IntelMemory)
	return intelMemory
}

func (i *IntelFactory) CreateCpu() AbstructCpu {
	var intelCpu AbstructCpu
	intelCpu = new(IntelCpu)
	return intelCpu
}

// nvidia产品族
type NvidiaGraphics struct {
}

func (ng *NvidiaGraphics) Display() {
	fmt.Println("英伟达显卡")
}

type NvidiaMemory struct {
}

func (nm *NvidiaMemory) Storage() {
	fmt.Println("英伟达内存")
}

type NvidiaCpu struct {
}

func (nc *NvidiaCpu) Calculate() {
	fmt.Println("英伟达cpu")
}

type NvidiaFactory struct {
}

func (nf *NvidiaFactory) CreateGraphics() AbstructGraphics {
	var nvidiaGraphics AbstructGraphics
	nvidiaGraphics = new(NvidiaGraphics)
	return nvidiaGraphics
}

func (nf *NvidiaFactory) CreateMemory() AbstructMemory {
	var nvidiaMemory AbstructMemory
	nvidiaMemory = new(NvidiaMemory)
	return nvidiaMemory
}

func (nf *NvidiaFactory) CreateCpu() AbstructCpu {
	var nvidiaCpu AbstructCpu
	nvidiaCpu = new(NvidiaCpu)
	return nvidiaCpu
}

// Kingston产品族
type KingstonGraphics struct {
}

func (kg *KingstonGraphics) Display() {
	fmt.Println("金士顿显卡")
}

type KingstonMemory struct {
}

func (km *KingstonMemory) Storage() {
	fmt.Println("金士顿内存")
}

type KingstonCpu struct {
}

func (kc *KingstonCpu) Calculate() {
	fmt.Println("金士顿cpu")
}

type KingstonFactory struct {
}

func (kf *KingstonFactory) CreateGraphics() AbstructGraphics {
	var kingstonGraphics AbstructGraphics
	kingstonGraphics = new(KingstonGraphics)
	return kingstonGraphics
}

func (kf *KingstonFactory) CreateMemory() AbstructMemory {
	var kingstonMemory AbstructMemory
	kingstonMemory = new(KingstonMemory)
	return kingstonMemory
}

func (kf *KingstonFactory) CreateCpu() AbstructCpu {
	var kingstonCpu AbstructCpu
	kingstonCpu = new(KingstonCpu)
	return kingstonCpu
}

// 业务层

type Computer struct {
	CPU    AbstructCpu
	GPU    AbstructGraphics
	Memory AbstructMemory
}

func (c *Computer) BuildComputer() {
	fmt.Println("computer start..")
	c.GPU.Display()
	c.Memory.Storage()
	c.CPU.Calculate()
}

func main() {
	intelFactory := new(IntelFactory)
	nvidiaFactory := new(NvidiaFactory)
	kingstonFactory := new(KingstonFactory)
	computer1 := Computer{
		GPU:    intelFactory.CreateGraphics(),
		CPU:    intelFactory.CreateCpu(),
		Memory: intelFactory.CreateMemory(),
	}
	computer1.BuildComputer()

	computer2 := Computer{
		GPU:    nvidiaFactory.CreateGraphics(),
		CPU:    intelFactory.CreateCpu(),
		Memory: kingstonFactory.CreateMemory(),
	}

	computer2.BuildComputer()

}
