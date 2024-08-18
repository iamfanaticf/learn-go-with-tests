package iteration

const repeatCount = 5
func Repeat(character string) string {
  var repeated string
  for _ = range repeatCount {
    repeated += character
  }
  return repeated
}
