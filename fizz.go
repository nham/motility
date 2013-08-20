package main
import "fmt"

type Particle struct {
    pos []float64
    vel []float64
}

func main() {
    p := Particle{
        pos: []float64{0, 100},
        vel: []float64{0, 0},
    }

    dt := 0.1
    a := []float64{0, -10.0}


    fmt.Println(p)

    for p.pos[1] >= 0 {
        for i := range p.pos {
            p.pos[i] += p.vel[i] * dt
        }

        for i := range p.vel {
            p.vel[i] += a[i] * dt
        }
        fmt.Println(p.pos)
    }

}
