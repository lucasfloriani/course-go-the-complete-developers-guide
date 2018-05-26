# Interfaces

É utilizado para reutilização de código.
No exemplo abaixo é demonstrado como um código pode ser repetido varias e varias vezes pelo falta de utilização de interface.

```go
func (d deck) shuffle() {
  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)

  for i := range d {
    newPosition := r.Intn(len(d) - 1)
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}

func (d []strings]) shuffle() {
  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)

  for i := range d {
    newPosition := r.Intn(len(d) - 1)
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}
func (d []float64]) shuffle() {
  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)

  for i := range d {
    newPosition := r.Intn(len(d) - 1)
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}
```