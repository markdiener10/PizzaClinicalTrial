package main

import (
	"testing"
)

func TestHouseFactory(t *testing.T) {

	house := NewHouse(0, 0)
	if house == nil {
		t.Errorf("House Factory should generate valid house instancest")
	}
	if !house.HitTest(0, 0) {
		t.Errorf("House Factory assigned to Equator/GMT should be located located there")
	}
	if house.HitTest(-1, 0) {
		t.Errorf("New House should not be located west of the equator")
	}
	if house.HitTest(0, 1) {
		t.Errorf("New House should be located north of equator")
	}

}

func TestHouseDelivery(t *testing.T) {

	house := NewHouse(0, 0)
	if house.DeliveryCount() != 0 {
		t.Errorf("House Factory should produce new houses with no deliveries")
	}
	house.DeliverPizza()
	if house.DeliveryCount() != 1 {
		t.Errorf("House should be located at the equator and greenwich mean time")
	}

}

func TestDelivererFactory(t *testing.T) {

	deliverer := NewDeliverer()
	if deliverer == nil {
		t.Errorf("Initialized Deliverer should be valid instance")
	}
	if !deliverer.HitTest(0, 0) {
		t.Errorf("Initialized Deliverer should be located at the equator/GMT")
	}

}

func TestDelivererMovement(t *testing.T) {

	deliverer := NewDeliverer()

	deliverer.GoEast()
	if !deliverer.HitTest(1, 0) {
		t.Errorf("Deliverer should be located east of GMT")
	}

	deliverer.GoNorth()
	if !deliverer.HitTest(1, 1) {
		t.Errorf("Deliverer should be located north and east of equator/GMT")
	}

	deliverer.GoWest()
	if !deliverer.HitTest(0, 1) {
		t.Errorf("Deliverer should be located north of equator and at GMT")
	}

	deliverer.GoSouth()
	if !deliverer.HitTest(0, 0) {
		t.Errorf("Initialized Deliverer should be located at the equator/GMT")
	}

}

func TestDeliveriesFactory(t *testing.T) {

	deliveries := NewDeliveries()
	if deliveries == nil {
		t.Errorf("Initialized Deliveries should be valid instance")
	}
	if deliveries.NumHouses() != 0 {
		t.Errorf("Delivery factory creating delivery history without actual events")
	}

	house := deliveries.FindHouse(0, 0)
	if house != nil {
		t.Errorf("Delivery factory creating delivery to houses without actual events")
	}

	house = deliveries.AddHouse(1, 1)
	if house == nil {
		t.Errorf("Delivery to houses not registering with delivery history cache")
	}
	if !house.HitTest(1, 1) {
		t.Errorf("Delivery to houses not registering with correct longitude and latitude")
	}

}

func TestPart1(t *testing.T) {

	//Part 1 = Given a set of dispatch commands and a single delivery agent, determine how many houses receive at least one pizza
	var house *House = nil
	deliveries := NewDeliveries()
	person := NewDeliverer()

	for idx := 0; idx < len(dispatchCommands); idx++ {

		switch dispatchCommands[idx] {
		case '^':
			person.GoNorth()
		case 'V':
			person.GoSouth()
		case '>':
			person.GoEast()
		case '<':
			person.GoWest()
		}

		house = deliveries.FindHouse(person.Longitude, person.Latitude)
		if house == nil {
			house = deliveries.AddHouse(person.Longitude, person.Latitude)
		}

		house.DeliverPizza()
	}

	if len(dispatchCommands) != 8192 {
		t.Errorf("Known dispatch command Set Not matching original known version:%d", len(dispatchCommands))
	}

	if deliveries.NumHouses() != 4498 {
		t.Errorf("Known dispatch command Set Not Correctly Executing :%d", deliveries.NumHouses())
	}

}

func TestPart2(t *testing.T) {

	//Part 2 = Given a set of dispatch commands and two delivery agents (Delivery Person + Rented Goat), determine how many houses receive at least one pizza
	var house *House = nil
	var goatNow bool = false
	var current *Deliverer = nil

	deliveries := NewDeliveries()

	person := NewDeliverer()
	goat := NewDeliverer()

	for idx := 0; idx < len(dispatchCommands); idx++ {

		if goatNow {
			current = goat
		} else {
			current = person
		}

		goatNow = !goatNow //Flip polarity each command

		switch dispatchCommands[idx] {
		case '^':
			current.GoNorth()
		case 'V':
			current.GoSouth()
		case '>':
			current.GoEast()
		case '<':
			current.GoWest()
		}

		house = deliveries.FindHouse(current.Longitude, current.Latitude)
		if house == nil {
			house = deliveries.AddHouse(current.Longitude, current.Latitude)
		}
		house.DeliverPizza()
	}

	if len(dispatchCommands) != 8192 {
		t.Errorf("Known dispatch command Set Not matching original known version:%d", len(dispatchCommands))
	}

	if deliveries.NumHouses() != 4438 {
		t.Errorf("Known dispatch command Set Not Correctly Executing :%d", deliveries.NumHouses())
	}

}
