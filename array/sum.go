package sum

func Sum(arr []int) int {
  var sum int
  for _,number := range arr {
    sum += number
  }
  return sum
} 

func SumAll(slices ...[]int) int {
  var sum int
  for _,slice := range slices {
    for _,number := range slice {
      sum += number
    }
  }
  return sum
}

func Flatten(slices ...[]int) []int {
  var flattened []int
  for _,slice := range slices {
    for _,number := range slice {
      flattened = append(flattened, number)
    }
  }

  return flattened
}
