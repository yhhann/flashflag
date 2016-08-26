// Package notice processes the event fired by infrastructure.
package notice

// Notice is a interface process the notice operator.
type Notice interface {
	// CheckChildren checks the path, returned chan will be noticed
	// when its children changed.
	CheckChildren(path string) (<-chan []string, <-chan error)

	// CheckDataChange checks the path, returned chan will be noticed
	// when its data changed.
	CheckDataChange(path string) (<-chan []byte, <-chan error)

	// Close release resource hold by Notice.
	Close()
}
