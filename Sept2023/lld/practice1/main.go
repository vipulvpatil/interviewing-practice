package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type VType string

const (
	CAR   VType = "CAR"
	BIKE  VType = "BIKE"
	TRUCK VType = "TRUCK"
)

type ErrorHandler struct{}

func (eh *ErrorHandler) handleError(e error) {
	// Other mechanism to handle errors.
	// Send to external service or log
	fmt.Printf("logging error %s\n", e)
}

func main() {
	var parkingLot *ParkingLot
	errorHandler := &ErrorHandler{}
	if len(os.Args) > 1 {
		// Try to read file and provide output
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			command := scanner.Text()
			if command == "exit" {
				break
			}
			parkingLot = executeCommand(parkingLot, command, errorHandler)
		}
	}
}

func executeCommand(parkingLot *ParkingLot, command string, errorHandler *ErrorHandler) *ParkingLot {
	commandName, commandArgs, err := parseCommand(command)
	if err != nil {
		errorHandler.handleError(err)
		return parkingLot
	}
	switch commandName {
	case "create_parking_lot": // PR1234 2 6
		if len(commandArgs) != 3 {
			errorHandler.handleError(errors.New("invalid command"))
		}
		floorCount, err := strconv.Atoi(commandArgs[1])
		if err != nil {
			errorHandler.handleError(err)
			return parkingLot
		}
		slotsPerFloor, err := strconv.Atoi(commandArgs[2])
		if err != nil {
			errorHandler.handleError(err)
			return parkingLot
		}
		parkingLot = CreateParkingLot(commandArgs[0], floorCount, slotsPerFloor)
	case "display free_count": // CAR
		if parkingLot == nil {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}
		if len(commandArgs) != 1 {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}
		freeCount := parkingLot.FreeCount(VType(commandArgs[0]))
		result := ""
		for i := range freeCount {
			result = fmt.Sprintf("%sNo. of free slots for %s on Floor %d: %d", result, commandArgs[0], i+1, freeCount[i])
			if i < len(freeCount)-1 {
				result = fmt.Sprintf("%s\n", result)
			}
		}
		fmt.Println(result)
	case "display free_slots": //CAR
		if parkingLot == nil {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}
		if len(commandArgs) != 1 {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}
		freeSlots := parkingLot.FreeSlots(VType(commandArgs[0]))
		result := ""
		for i := range freeSlots {
			row := ""
			for j := range freeSlots[i] {
				row = fmt.Sprintf("%s%d", row, freeSlots[i][j])
				if j < len(freeSlots[i])-1 {
					row = fmt.Sprintf("%s, ", row)
				}
			}
			result = fmt.Sprintf("%sFree slots for %s on Floor %d: %s", result, commandArgs[0], i+1, row)
			if i < len(freeSlots)-1 {
				result = fmt.Sprintf("%s\n", result)
			}
		}
		fmt.Println(result)
	case "display occupied_slots": // CAR
		if parkingLot == nil {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}
		if len(commandArgs) != 1 {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}

		occupiedSlots := parkingLot.OccupiedSlots(VType(commandArgs[0]))
		result := ""
		for i := range occupiedSlots {
			row := ""
			for j := range occupiedSlots[i] {
				row = fmt.Sprintf("%s%d", row, occupiedSlots[i][j])
				if j < len(occupiedSlots[i])-1 {
					row = fmt.Sprintf("%s, ", row)
				}
			}
			result = fmt.Sprintf("%sOccupied slots for %s on Floor %d: %s", result, commandArgs[0], i+1, row)
			if i < len(occupiedSlots)-1 {
				result = fmt.Sprintf("%s\n", result)
			}
		}
		fmt.Println(result)
	case "park_vehicle": // CAR KA-01-DB-1234 black
		if parkingLot == nil {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}
		if len(commandArgs) != 3 {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}
		result := parkingLot.ParkVehicle(VType(commandArgs[0]), commandArgs[1], commandArgs[2])
		fmt.Println(result)
	case "unpark_vehicle": // PR1234_2_5
		fmt.Println(commandArgs)
		if parkingLot == nil {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}
		if len(commandArgs) != 1 {
			errorHandler.handleError(errors.New("invalid command"))
			return parkingLot
		}
		result := parkingLot.UnparkVehicle(commandArgs[0])
		fmt.Println(result)
	}
	return parkingLot
}

func parseCommand(command string) (string, []string, error) {
	parts := strings.Split(command, " ")
	if len(parts) <= 1 {
		return "", nil, errors.New("invalid command length")
	}
	commandName := parts[0]
	commandArgs := parts[1:]
	if parts[0] == "display" {
		commandName = fmt.Sprintf("%s %s", commandName, parts[1])
		commandArgs = parts[2:]
	}
	return commandName, commandArgs, nil
}

type ParkingLot struct {
	plId    string
	slots   [][]*Slot
	carPQ   *SlotsPQ
	bikePQ  *SlotsPQ
	truckPQ *SlotsPQ
}

type Vehicle struct {
	vType VType
	regNo string
	color string
}

type Slot struct {
	vehicle *Vehicle
	vType   VType
	floorNo int
	slotNo  int
}

func CreateParkingLot(plId string, floorCount int, slotsPerFloor int) *ParkingLot {
	// Create floors and slots per floor
	// Maybe use PQ for slots per vehicle
	slots := make([][]*Slot, floorCount)
	for i := 0; i < floorCount; i++ {
		slots[i] = make([]*Slot, slotsPerFloor)
		for j := 0; j < slotsPerFloor; j++ {
			slots[i][j] = CreateSlot(i, j)
		}
	}

	lot := &ParkingLot{
		plId:    plId,
		slots:   slots,
		carPQ:   &SlotsPQ{},
		bikePQ:  &SlotsPQ{},
		truckPQ: &SlotsPQ{},
	}

	for i := 0; i < floorCount; i++ {
		for j := 0; j < slotsPerFloor; j++ {
			lot.MakeSlotAvailable(slots[i][j])
		}
	}

	fmt.Printf("Created parking lot with %d floors and %d slots per floor\n", floorCount, slotsPerFloor)

	return lot
}

func CreateSlot(i, j int) *Slot {
	vType := CAR
	if j == 0 {
		vType = TRUCK
	} else if j <= 2 {
		vType = BIKE
	}

	slot := Slot{
		vType:   vType,
		floorNo: i + 1,
		slotNo:  j + 1,
	}
	return &slot
}

type ParkingLotEntrance interface {
	ParkVehicle(vType int, regNo string, color string) string
	UnparkVehicle(ticketId string) string
}

func (lot *ParkingLot) MakeSlotAvailable(slot *Slot) {
	h := lot.getSlotsHeap(slot.vType)
	if h != nil {
		heap.Push(h, slot)
	}
}

func (lot *ParkingLot) getSlotsHeap(vType VType) *SlotsPQ {
	switch vType {
	case CAR:
		return lot.carPQ
	case BIKE:
		return lot.bikePQ
	case TRUCK:
		return lot.truckPQ
	}
	return nil
}

func (lot *ParkingLot) getTicketId(slot *Slot) string {
	return fmt.Sprintf("%s_%d_%d", lot.plId, slot.floorNo, slot.slotNo)
}

func (lot *ParkingLot) getSlotFromTicketId(id string) *Slot {
	strs := strings.Split(id, "_")
	if len(strs) != 3 {
		return nil
	}
	plId := strs[0]
	floorNo, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil
	}
	slotNo, err := strconv.Atoi(strs[2])
	if err != nil {
		return nil
	}
	if plId != lot.plId {
		return nil
	}
	if floorNo > len(lot.slots) || floorNo < 1 {
		return nil
	}
	floorIndex := floorNo - 1
	if slotNo > len(lot.slots[floorIndex]) || slotNo < 1 {
		return nil
	}
	slotIndex := slotNo - 1
	return lot.slots[floorIndex][slotIndex]
}

func (lot *ParkingLot) ParkVehicle(vType VType, regNo string, color string) string {
	h := lot.getSlotsHeap(vType)
	if h.Len() == 0 {
		return "Parking Lot Full"
	}
	top := heap.Pop(h).(*Slot)
	top.vehicle = &Vehicle{
		vType: vType,
		regNo: regNo,
		color: color,
	}
	result := fmt.Sprintf("Parked vehicle. Ticket ID: %s", lot.getTicketId(top))
	return result
}

func (lot *ParkingLot) UnparkVehicle(ticketId string) string {
	slot := lot.getSlotFromTicketId(ticketId)
	if slot == nil {
		return "Invalid Ticket"

	}
	if slot.vehicle == nil {
		return "Invalid Ticket"
	}
	result := fmt.Sprintf("Unparked vehicle with Registration Number: %s and Color: %s", slot.vehicle.regNo, slot.vehicle.color)
	slot.vehicle = nil
	lot.MakeSlotAvailable(slot)
	return result
}

type ParkingLotQueryer interface {
	FreeCount(vType int) int
	FreeSlots(vType int) []int
	OccupiedSlots(vType int) []int
}

func (lot *ParkingLot) FreeCount(vType VType) []int {
	fmt.Println(vType)
	freeSlots := lot.FreeSlots(vType)
	freeCount := make([]int, len(lot.slots))
	for i := range freeCount {
		freeCount[i] = len(freeSlots[i])
	}
	return freeCount
}

func (lot *ParkingLot) FreeSlots(vType VType) [][]int {
	freeSlots := make([][]int, len(lot.slots))
	for i := range freeSlots {
		freeSlots[i] = []int{}
		for _, s := range lot.slots[i] {
			if s.vType == vType && s.vehicle == nil {
				freeSlots[i] = append(freeSlots[i], s.slotNo)
			}
		}
	}
	return freeSlots
}
func (lot *ParkingLot) OccupiedSlots(vType VType) [][]int {
	freeSlots := make([][]int, len(lot.slots))
	for i := range freeSlots {
		freeSlots[i] = []int{}
		for _, s := range lot.slots[i] {
			if s.vType == vType && s.vehicle != nil {
				freeSlots[i] = append(freeSlots[i], s.slotNo)
			}
		}
	}
	return freeSlots
}

// minPQ for slots
type SlotsPQ []*Slot

func (s SlotsPQ) Len() int {
	return len(s)
}
func (s SlotsPQ) Less(i int, j int) bool {
	ith := s[i]
	jth := s[j]
	if ith.floorNo < jth.floorNo {
		return true
	} else if ith.floorNo > jth.floorNo {
		return false
	}
	return ith.slotNo < jth.slotNo
}
func (s SlotsPQ) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *SlotsPQ) Push(k any) {
	t := *s
	*s = append(t, k.(*Slot))
}

func (s *SlotsPQ) Pop() any {
	t := *s
	result := t[len(t)-1]
	*s = t[0 : len(t)-1]
	return result
}
