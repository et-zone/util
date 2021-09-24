package util

import (
	"context"
	"errors"
	"time"
)

//说明，支持设置超时时间，具体可以设置多少超时routine，看单个routine消耗的资源情况

const no_timeout = 0

type routineWorker struct {
	chs     []chan chStruct
	timeout time.Duration
	cancel  context.CancelFunc
	ctx     context.Context
}

type chStruct struct {
	Err error
}


func NewRoutineWorker() routineWorker {
	w := routineWorker{chs: make([]chan chStruct, 0),timeout: 0}
	w.ctx = context.Background()
	return w
}

func NewTimeoutWorker(timeout time.Duration) routineWorker {
	w := routineWorker{chs: make([]chan chStruct, 0), timeout: timeout}
	if w.timeout == no_timeout {
		w.ctx = context.Background()
	} else {
		w.ctx, w.cancel = context.WithTimeout(context.Background(), timeout)
	}
	return w
}

func do(ctx context.Context, ch chan chStruct, f func() error) {
	ch <- chStruct{Err: f()}
}

func (this *routineWorker) Add(f func() error) {
	tmch := make(chan chStruct)
	this.chs = append(this.chs, tmch)
	go do(this.ctx,tmch, f)
}

func (this *routineWorker) Wait() error {
	var err error

	if this.timeout!=no_timeout{
		select {
		case <-this.ctx.Done():
			_,ok:=this.ctx.Deadline()
			if ok{
				this.cancel()
				return this.timeoutClose()
			}
		}
	}

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

func (this *routineWorker) timeoutClose() error {
	for _, ch := range this.chs {
		go func(){
			defer func() {
				recover()
			}()
			ch<-chStruct{Err: errors.New("the func timeout")}
		}()

		_, noclose := <-ch
		if noclose {
			close(ch)
		}
	}
	this.chs = nil
	return errors.New("routine_timeout:"+time.Now().Format("2006-01-02 15:04:05"))
}

