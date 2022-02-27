// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const (
	Big int = iota + 1
	Medium
	Small
)

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big, medium, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		{
			if this.big <= 0 {
				return false
			}
			this.big--
			return true
		}
	case 2:
		{
			if this.medium <= 0 {
				return false
			}
			this.medium--
			return true
		}
	case 3:
		{
			if this.small <= 0 {
				return false
			}
			this.small--
			return true
		}
	default:
		return false
	}

}

func main() {
	obj := Constructor(1, 2, 4)
	isBigCarPrkingAvailable := obj.AddCar(Big)
	fmt.Println("Big Car parking availability ", isBigCarPrkingAvailable)

	isBigCarPrkingAvailable = obj.AddCar(Big)
	fmt.Println("Big Car parking availability ", isBigCarPrkingAvailable)

	isMediumCarPrkingAvailable := obj.AddCar(Medium)
	fmt.Println("Medium Car parking availability ", isMediumCarPrkingAvailable)

	isSmallCarPrkingAvailable := obj.AddCar(Small)
	fmt.Println("Small Car parking availability ", isSmallCarPrkingAvailable)
}
