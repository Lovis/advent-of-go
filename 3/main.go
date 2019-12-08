package main

import (
	"fmt"
	"strconv"
	"math"
)

// var left = []string{"R8","U5","L5","D3"}
// var right = []string{"U7","R6","D4","L4"}

// var left = []string{"R75","D30","R83","U83","L12","D49","R71","U7","L72"}
// var right = []string{"U62","R66","U55","R34","D71","R55","D58","R83"}

var left = []string{"R1002","U407","R530","D268","R516","U937","L74","U838","R784","D684","L912","U746","R189","U192","R868","D345","L972","D492","R942","U631","L559","U634","L80","U513","L746","D997","L348","D160","L655","U949","R717","U396","R549","D167","R591","U469","L22","U977","L167","D856","L320","D920","L396","U490","L895","U180","R661","D828","R864","U189","R307","U402","R409","U445","L101","D418","R812","U419","R319","U75","L813","D46","L491","U39","R737","U11","R177","U311","L278","U254","R475","U166","L515","D105","L694","D437","L298","U169","L613","D234","L999","U380","L711","D758","R932","D27","L951","D529","L935","D189","R816","D176","R98","D320","R965","D333","L367","U622","R18","U83","R275","D205","L960","U177","R750","D466","R442","U797","R355","D717","L569","D578","R384","U863","R541","U405","R527","D658","L514","U168","L64","D918","R947","D11","L189","D875","R599","U201","L165","U772","L679","U566","L195","U660","R896","D622","R678","U390","L984","D900","R889","D714","R557","U848","L176","U541","R518","D699","L904","D23","L55","D886","L206","D621","L48","D197","R502","D259","L779","D72","L183","U747","L424","U452","L603","U561","L430","D942","R515","D378","R962","U508","R230","D650","R804","D453","R899","D813","R484","U798","R456","D231","L316","U117","R630","D436","L985","D283","L393","D370","R158","U957","L914","D455","L875","U536","R889","U400","R347","U712","R487","D455","R428","U590","R127","D132","L202","U377","R138","U654","L760","D46","R213","D225","L817","U455","L612","U543","L525","U979","R591","D940","R446","U786","R750","U244","R325","U928","L44","U551","L955","U221","L986","U516","L916","D242","L280","D71","R80","U849","L271","U626","R737","D646","R82","U120","R646","U569","R463","D94","R570","U456","L548","D687","R221","D759","L606","D818","L859","U218","R682","U299","R818","D966","R407","U605","L859","D378","L53","D722","L216","D221","R639","U485","L865","D620","R99","D988","R944","D323","R540","U372","L470","U106","L804","D92","L177","U518","R277","U670","R451","D194","L695","D502","L601","U596","R374","U682","L19","D54","L156"}
var right = []string{"L1003","U22","R594","D241","L215","D906","R733","D831","L556","U421","L780","D242","R183","U311","R46","D52","R124","D950","L18","U985","R999","D528","R850","U575","L138","D62","L603","U467","R641","U155","L165","D63","L489","U4","R635","D460","L446","D938","R983","U494","L491","D433","L722","U427","L469","U832","L712","U798","R906","U804","R646","U721","R578","D194","L726","U803","L985","D934","R943","U198","R726","U341","R583","U998","L992","D401","L132","D681","L363","U949","L814","D977","R840","D145","L177","D291","L652","D396","R330","D951","L363","U813","R847","D374","R961","D912","R516","D178","R495","U49","R340","D395","R632","D991","R487","D263","R320","D948","R456","D142","L783","D888","R589","D999","L159","U686","R402","D586","L425","U946","R56","D979","L534","U313","R657","D401","L666","D504","R712","D232","L557","D620","R193","D670","L134","D975","R837","D901","R813","U459","L499","U450","L87","U84","L582","U310","R795","D280","L730","D458","L727","D196","R95","U210","R498","U760","R778","U325","R715","U479","R275","U557","L450","D196","L60","U115","R475","D265","L611","D372","R60","U935","L717","U809","L344","D854","L386","U473","R72","U968","L816","U900","R866","U172","R965","U383","R576","D774","R753","U788","L781","D237","L401","U786","R873","U331","R609","D232","L603","U685","L494","U177","L982","D173","L673","U772","L7","U7","R517","U573","R212","D413","L124","D810","L223","U137","L576","U95","R128","U896","L91","U932","L875","U917","R106","U911","L208","D507","L778","D59","L71","D352","R988","U708","L58","D429","L122","U771","L713","D801","R188","U661","R752","D374","R312","D848","L504","D540","R334","U517","R343","D739","L727","D552","L555","U580","L857","U474","R145","U188","L789","U698","R735","U131","L494","U162","L27","D849","L140","D849","R615","U798","R160","U492","R884","U521","L542","D729","R498","D853","R918","U565","R65","U32","L607","U552","L38","D822","L77","U490","L190","D93","L104","U268","R702","D112","L917","D876","L631","D139","L989","U810","R329","U253","L498","D767","L550","U666","L549","U616","R376"}

var coordsLeft []Coords
var coordsRight []Coords

type Coords struct {
	x, y, step int
}

func calcDistance(coords []Coords) (smallest int) {
	// traverse the list, find the smallest
	smallest = int(math.Abs(float64(coords[0].x)) + math.Abs(float64(coords[0].y)))
	for _, v := range coords {
		manhattan := int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)))
		// fmt.Println("manhattan", manhattan, v.x, v.y)
		if (manhattan < smallest) {
			smallest = manhattan
		}
	}
	return smallest
}

func calcDistanceWithSteps(coords []Coords) (smallest int) {
	// traverse the list, find the smallest
	smallest = coords[0].step
	for _, v := range coords {
		if (v.step < smallest) {
			smallest = v.step
		}
	}
	return smallest
}

func keep(point Coords, coordsRight []Coords) (bool){
	for _, v := range coordsRight {
		if (point.x == v.x && point.y == v.y) {
			return true
		} 
	}
	return false
}

func keepAndSumSteps(point *Coords, coordsRight []Coords) (bool){
	for _, v := range coordsRight {
		if (point.x == v.x && point.y == v.y) {
			point.step += v.step
			fmt.Println("point with added", point)
			return true
		} 
	}
	return false
}

func findIntersections(coordsLeft []Coords, coordsRight []Coords) []Coords {
	var intersections []Coords
	for _, v := range coordsLeft {
		// fmt.Println(i, v, coordsRight[i])
		if (keepAndSumSteps(&v, coordsRight)) {
			fmt.Println("match:", v)
			intersections = append(intersections, v)
		}
	}
	
	return intersections
}

func parse(input []string, coords []Coords) ([]Coords) {
	// parse and create array of coords
	// fmt.Println("len input", len(input))
	point := Coords{0, 0, 0}
	for i := 0; i < len(input); i++ {
		var instruction = input[i]
		var direction = string(instruction[0])
		steps, err := strconv.Atoi(string(instruction[1:]))
		if err != nil {
			break
		}
		for j := 0; j < steps; j++ {
			// x = int(coords[len(coords)-1].x)
			// y = int(coords[len(coords)-1].y)
			switch direction {
			case "R":
				point.x += 1
			case "U":
				point.y += 1
			case "L":
				point.x -= 1
			case "D":
				point.y -= 1
			}

			newCoord := Coords{point.x,point.y, len(coords)+1}
			fmt.Println("new coord", newCoord)
			coords = append(coords, newCoord)
		}
	}
	return coords
}
func main() {
	fmt.Println("here we go")
	coordsLeft = parse(left, coordsLeft)
	// fmt.Println("coordsLeft: ", coordsLeft)

	coordsRight = parse(right, coordsRight)
	//fmt.Println("coordsRight: ", coordsRight)

	intersections := findIntersections(coordsLeft, coordsRight)
	fmt.Println("intersections", intersections)

	smallest := calcDistance(intersections)
	fmt.Println("smallest manhattan: ", smallest)

	smallestPt2 := calcDistanceWithSteps(intersections)
	fmt.Println("smallest by steps: ", smallestPt2)
	// extend to two full arrays of cells (x,y)
	// find the common cells
	// measure how far they are from origo

}