package notice

import (
	"time"

	"github.com/golang/glog"
	"github.com/samuel/go-zookeeper/zk"
)

// ConfZK implements Notice interface.
type ConfZK struct {
	Addrs []string

	*zk.Conn
}

func (k *ConfZK) connectZk(addrs []string, timeout time.Duration) error {
	k.Addrs = addrs
	thisConn, ch, err := zk.Connect(zk.FormatServers(addrs), timeout)
	if err != nil {
		return err
	}

	connectOk := make(chan struct{})
	go k.checkConnectEvent(ch, connectOk)
	<-connectOk
	k.Conn = thisConn
	return nil
}

// CloseZK closes the zookeeper.
func (k *ConfZK) Close() {
	if k != nil {
		k.Conn.Close()
		k = nil
	}
}

func (k *ConfZK) checkConnectEvent(ch <-chan zk.Event, okChan chan<- struct{}) {
	for ev := range ch {
		switch ev.Type {
		case zk.EventSession:
			switch ev.State {
			case zk.StateConnecting:
			case zk.StateConnected:
				glog.Infof("Succeeded to connect to zk[%v].", k.Addrs)
			case zk.StateHasSession:
				glog.Infof("Succeeded to get session from zk[%v].", k.Addrs)
				okChan <- struct{}{}
			}
		default:
		}
	}
}

// CheckChildren sets a watcher on given path,
// the returned chan will be noticed when children changed.
func (k *ConfZK) CheckChildren(path string) (<-chan []string, <-chan error) {
	snapshots := make(chan []string)
	errors := make(chan error)

	go func() {
		for {
			snapshot, _, events, err := k.ChildrenW(path)
			if err != nil {
				errors <- err
				return
			}
			snapshots <- snapshot

			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}
			glog.V(2).Infof("event type: %v", evt.Type)
		}
	}()

	return snapshots, errors
}

// CheckDataChange sets a watcher on given path,
// the returned chan will be noticed when data changed.
func (k *ConfZK) CheckDataChange(path string) (<-chan []byte, <-chan error) {
	datas := make(chan []byte)
	errors := make(chan error)

	go func() {
		for {
			dataBytes, _, events, err := k.GetW(path)
			if err != nil {
				errors <- err
				return
			}
			datas <- dataBytes
			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}

		}
	}()

	return datas, errors
}

// NewConfZK creates a new DfsZk.
func NewConfZK(addrs []string, timeout time.Duration) *ConfZK {
	zk := new(ConfZK)
	if err := zk.connectZk(addrs, timeout); err != nil {
		return nil
	}
	return zk
}
