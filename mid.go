package util

import (
"context"
)
//middleware

type Context struct {
	context.Context
	arg interface{}
}

type Worker interface {
	Before(ctx Context)error
	After(ctx Context)error
	Dao(ctx Context)error
}

func Run(ctx Context,worker Worker)error{
	var err error
	err=worker.Before(ctx)
	if err!=nil{
		return err
	}
	err=worker.Dao(ctx)
	if err!=nil{
		return err
	}
	err=worker.After(ctx)
	if err!=nil{
		return err
	}
	return nil
}

//arg can be nil
func NewContext(ctx context.Context)*Context{
	return &Context{Context:ctx}
}

func (c *Context)SetArg(arg interface{})*Context{
	c.arg=arg
	return c
}

func (c *Context)GetArg()interface{}{
	return 	c.arg
}