package main

import (
	"errors"
)

//################ House Definition

type House struct {
	longitude  int //West is negative, East is positive
	latitude   int //North is positive, South is negative
	deliveries uint
}

func (g *House) Initialize(long int, lat int) bool {
	g.longitude = long
	g.latitude = lat
	return true
}

func (g *House) HitTest(long int, lat int) bool {
	if g.latitude != lat {
		return false
	}
	if g.longitude != long {
		return false
	}
	return true
}

func (g *House) DeliverPizza() error {
	g.deliveries++
	if g.deliveries > 200 {
		return errors.New("Obesity Warning - clinical trial exceeding consumption threshold")
	}
	return nil
}

func (g *House) DeliveryCount() uint {
	return g.deliveries
}

func NewHouse(long int, lat int) *House {
	house := &House{}
	house.Initialize(long, lat)
	return house
}

//################ Delivery History

type Deliveries struct {
	houses []*House
}

func (g *Deliveries) FindHouse(long int, lat int) *House {
	for _, house := range g.houses {
		if !house.HitTest(long, lat) {
			continue
		}
		return house
	}
	return nil
}

func (g *Deliveries) AddHouse(long int, lat int) *House {
	house := NewHouse(long, lat)
	g.houses = append(g.houses, house)
	return house
}

func (g *Deliveries) NumHouses() int {
	return len(g.houses)
}

func NewDeliveries() *Deliveries {
	return &Deliveries{houses: make([]*House, 0)}
}

//################ Delivery Agent Definition

type Deliverer struct {
	Longitude int
	Latitude  int
}

func (g *Deliverer) GoEast() {
	g.Longitude++
}

func (g *Deliverer) GoWest() {
	g.Longitude--
}

func (g *Deliverer) GoNorth() {
	g.Latitude++
}

func (g *Deliverer) GoSouth() {
	g.Latitude--
}

func (g *Deliverer) HitTest(long int, lat int) bool {
	if g.Latitude != lat {
		return false
	}
	if g.Longitude != long {
		return false
	}
	return true
}

func NewDeliverer() *Deliverer {
	return &Deliverer{Longitude: 0, Latitude: 0}
}
