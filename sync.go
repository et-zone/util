package util

type worker struct {
	chs []chan chStruct
}

type chStruct struct {
	Err error
}

func NewWorker() worker {
	return worker{make([]chan chStruct, 0)}
}

func do(ch chan chStruct, f func() error) {
	ch <- chStruct{Err: f()}
}

func (this *worker) Add(f func() error) {
	tmch := make(chan chStruct)
	this.chs = append(this.chs, tmch)
	go do(tmch, f)
}

func (this *worker) Wait() error {
	var err error
	for _, ch := range this.chs {
		select {
		case tmp, noclose := <-ch:
			if tmp.Err != nil {
				err = tmp.Err
			}
			if noclose {
				close(ch)
			}
		}
	}
	this.chs = nil
	return err
}
