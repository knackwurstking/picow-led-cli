package picow

type RequestARGS interface {
	string | int
}

type Request[T RequestARGS] struct {
	ID      int          `json:"id"` // defaults to zero
	Group   CommandGroup `json:"group"`
	Type    CommandType  `json:"type"`
	Command Command      `json:"command"`
	Args    []T          `json:"args"`
}
