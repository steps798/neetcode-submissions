func carFleet(target int, position []int, speed []int) int {
//  1  2  3  4  5  6  7  8  9  10
//  a        a        a        a
//           b     b     b     b

//  1  2  3  4  5  6  7  8  9  10
//  a        b
//           a     b
//                    a  b
//                             ab

//  0  1  2  3  4  5  6  7  8  9  10
//              a     a     a     a
//     b     b     b     b     b
//  c  c  c  c  c  c  c  c  c  c  c
//                       d  d  d  d

//  0  1  2  3  4  5  6  7  8  9  10
//  c  b        a        d
//     c     b        a     d
//        c        b        a  d
//           c           b        ad
//              c              b  ad
//                 c              bad
//                    c           bad
//                       c        bad
//                          c     bad
//                             c  bad
//                                cbad

//  0  1  2  3  4  5  6  7  8  9  10
//           a     b     c
//                    a  b  c
//                             abc

// target int, position []int, speed []int
n := len(position)
cars := make([][2]int, n)
for i := 0; i < n; i++ {
	cars[i] = [2]int{position[i], speed[i]}
}

sort.Slice(cars, func(i, j int) bool {
	return cars[i][0] > cars[j][0] // sort car position desc
})

// when the nearest car to the target reach target?
time := float64(target - cars[0][0]) / float64(cars[0][1])
// the nearest car reach the target = 1 fleet
fleet := 1
// loop the rest of cars
for i := 1; i < n; i++ {
	// when the next car reach target?
	curr := float64(target - cars[i][0]) / float64(cars[i][1])
	// if the next car overlap the previous time
	if curr > time {
		// create another fleet
		fleet++
		time = curr 
	}
}

return fleet
}
