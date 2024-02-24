package goFoundation

import "fmt"

type Vehicle interface {
	isRunning() bool
	getSpeed() string
}

/* Car */

type Car struct {
	name         string
	currentSpeed int32
	price        int32
}

func (c Car) isRunning() bool {
	return c.currentSpeed != 0
}

func (c Car) getSpeed() string {
	if c.isRunning() {
		return "James is driving his car in " + fmt.Sprint(c.currentSpeed) + " km/h"
	}

	return "James's car is parking"
}

/* Motorbike */

type Motorbike struct {
	name         string
	currentSpeed int32
	color        string
}

func (m Motorbike) isRunning() bool {
	return m.currentSpeed != 0
}

func (m Motorbike) getSpeed() string {
	if m.isRunning() {
		return "Brew is racing his motorbike in " + fmt.Sprint(m.currentSpeed) + " km/h"
	}

	return "Brew's motorbike is parking"
}

/* Get driving info */

func getDrivingInfo(v Vehicle) {
	fmt.Printf("getSpeed of %T return: %s\n", v, v.getSpeed())
}

func PlayWithInterface() {
	mercedes := Car{
		name:  "Mercedes",
		currentSpeed: 0,
		price: 120000,
	}

	s1000rr := Motorbike{
		name:  "BMW S1000RR",
		currentSpeed: 180,
		color: "blue",
	}

	getDrivingInfo(mercedes)
	getDrivingInfo(s1000rr)
}