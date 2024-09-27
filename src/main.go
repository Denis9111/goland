package main

import (
	"fmt"
	"log"
	"net/http"
)

//import (
//	"GoProject/src/magazine"
//	"fmt"
//)
//
//func printInfo(s *magazine.Subscriber) {
//	fmt.Println("Name:", s.Name)
//	fmt.Println("Monthly rate:", s.Rate)
//	fmt.Println("Active?", s.Active)
//}
//
//func defaultSubscriber(name string) *magazine.Subscriber {
//	var s magazine.Subscriber
//	s.Name = name
//	s.Rate = 5.99
//	s.Active = true
//	return &s
//}
//
//func applyDiscount(s *magazine.Subscriber) {
//	s.Rate = 4.99
//}
//
//func main() {
//	subscriber1 := defaultSubscriber("Aman Singh")
//	applyDiscount(subscriber1)
//	printInfo(subscriber1)
//	subscriber2 := defaultSubscriber("Beth Ryan")
//	printInfo(subscriber2)
//}

//type Gallons float64
//type Liters float64
//type Milliliters float64
//
//func ToGallons(l Liters) Gallons {
//	return Gallons(l * 0.264)
//}
//func ToLiters(g Gallons) Liters {
//	return Liters(g * 3.785)
//}
//
//func main() {
//	gall := Gallons(1.2)
//	litt := Liters(4.5)
//
//	gall += ToGallons(Liters(40.0))
//	litt += ToLiters(Gallons(30.0))
//	fmt.Printf("Gallon %0.1f\n Liters %0.1f\n", gall, litt)
//
//}

//func main() {
//	date := calendar.Date{}
//	errY := date.SetYear(2019)
//	if errY != nil {
//		log.Fatal(errY)
//	}
//	errM := date.SetMonth(10)
//	if errM != nil {
//		log.Fatal(errM)
//	}
//	errD := date.SetDay(27)
//	if errD != nil {
//		log.Fatal(errD)
//	}
//	fmt.Println(date.Year())
//	fmt.Println(date.Month())
//	fmt.Println(date.Day())
//
//	event := calendar.Event{}
//	errT := event.SetTitle("Moms day")
//	if errT != nil {
//		log.Fatal(errT)
//	}
//	fmt.Println(event.Title())
//
//}

//func playList(device gadged.Player, songs []string) {
//	for _, song := range songs {
//		device.Play(song)
//	}
//	device.Stop()
//}
//
//func TryOut(play gadged.Player) {
//	play.Play("Play song")
//	play.Stop()
//	record, ok := play.(gadged.TapeRecorder)
//	if ok {
//		record.Record("record new song")
//	}
//}
//
//func main() {
//	TryOut(gadged.TapeRecorder{})
//	TryOut(gadged.TapePlayer{})
//}

//type Truck string
//
//func (t Truck) Accelerate() {
//	fmt.Println("Speeding up")
//}
//func (t Truck) Brake() {
//	fmt.Println("Stopping")
//}
//func (t Truck) Steer(direction string) {
//	fmt.Println("Turning", direction)
//}
//func (t Truck) LoadCargo(cargo string) {
//	fmt.Println("Loading", cargo)
//}
//
//type Vehicle interface {
//	Accelerate()
//	Brake()
//	Steer(string)
//}
//
//func TryVehicle(vehicle Vehicle) {
//	vehicle.Accelerate()
//	vehicle.Brake()
//	vehicle.Steer("left")
//	vehicle.Steer("right")
//	truck, ok := vehicle.(Truck)
//	if ok {
//		truck.LoadCargo("test cargo")
//	}
//}
//
//func main() {
//	var truck Truck = "Man"
//	TryVehicle(truck)
//
//}

type OverheatError float64

func (o OverheatError) Error() string {

	return fmt.Sprintf("OverheatError: %f", o)
}

func chekTemp(current float64, safe float64) error {
	if current > safe {
		return OverheatError(current)
	}
	return nil
}

func myH(w http.ResponseWriter) {

}

func init() {
	fmt.Println("hello world")
}

func main() {
	var err error = chekTemp(3.0, 2.0)
	if err != nil {
		log.Fatal(err)
	}
}

/*добавил коммент теперь надо запушить на удаленку 555*/

/*еще коммент*/
