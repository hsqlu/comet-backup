package traffic

type ThroughputLevel uint

const (
	HighThroughput ThroughputLevel = iota
	MediumThroughput
	LowThroughput
)

type Controller interface {
	Name() string
	PerformanceByLevel(intersection *Intersection) int32
	Performance(intersection *Intersection) int32
	CPMPerDollar(intersection *Intersection) int32
}

var controllers []Controller

func init() {
	controllers = []Controller{
		new(Roundabout),
		new(StopSign),
		new(TrafficLight),
	}
}

type Roundabout struct{}
type StopSign struct{}
type TrafficLight struct{}

func (c Roundabout) Name() string {
	return "Roundabout"
}

func (c Roundabout) CPMPerDollar(i *Intersection) int32 {
	return 0
}

func (c Roundabout) PerformanceByLevel(i *Intersection) int32 {
	switch i.ThroughputLevel() {
	case HighThroughput:
		return 50
	case MediumThroughput:
		return 75
	case LowThroughput:
		return 90
	}
	return 0
}

func (c Roundabout) Performance(i *Intersection) int32 {
	ans := c.PerformanceByLevel(i)

	if i.northCPM+i.southCPM > i.westCPM+i.eastCPM {
		ans = ans + 10
	} else {
		ans = ans - 10
	}
	if ans > 100 {
		ans = 100
	} else if ans < 0 {
		ans = 0
	}
	return ans
}

func (c StopSign) Name() string {
	return "StopSign"
}

func (c StopSign) CPMPerDollar(i *Intersection) int32 {
	return 0
}

func (c StopSign) PerformanceByLevel(i *Intersection) int32 {
	switch i.ThroughputLevel() {
	case HighThroughput:
		return 20
	case MediumThroughput:
		return 30
	case LowThroughput:
		return 40
	}
	return 0
}

func (c StopSign) Performance(i *Intersection) int32 {
	return c.PerformanceByLevel(i)
}

func (c TrafficLight) Name() string {
	return "TrafficLight"
}

func (c TrafficLight) CPMPerDollar(i *Intersection) int32 {
	return 0
}

func (c TrafficLight) PerformanceByLevel(i *Intersection) int32 {
	switch i.ThroughputLevel() {
	case HighThroughput:
		return 90
	case MediumThroughput:
		return 75
	case LowThroughput:
		return 30
	}
	return 0
}

func (c TrafficLight) Performance(i *Intersection) int32 {
	return c.PerformanceByLevel(i)
}

type Intersection struct {
	northCPM int32
	eastCPM  int32
	southCPM int32
	westCPM  int32
}

func New(n, e, s, w int32) Intersection {
	return Intersection{
		northCPM: n,
		eastCPM:  e,
		southCPM: s,
		westCPM:  w,
	}
}

func (i *Intersection) TotalCPM() int32 {
	return i.eastCPM + i.northCPM + i.southCPM + i.westCPM
}

func (i *Intersection) ThroughputLevel() ThroughputLevel {
	if i.TotalCPM() >= 20 {
		return HighThroughput
	}
	if i.TotalCPM() < 10 {
		return LowThroughput
	}
	return MediumThroughput
}

func (i *Intersection) EfficientController() (int32, Controller) {
	var controller Controller
	var max int32
	for _, c := range controllers {
		efficient := c.Performance(i)
		if efficient > max {
			max = efficient
			controller = c
		}
	}
	return max, controller
}
