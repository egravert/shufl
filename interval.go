package main

import (
	"fmt"
	"strconv"
	"strings"
)

type interval struct {
	min   int
	max   int
	isSet bool
}

func (i *interval) String() string {
	return fmt.Sprintf("%d-%d", i.min, i.max)
}

func (i *interval) Type() string {
	return "LO-HI"
}

func (i *interval) Set(str string) error {
	var err error
	vals := strings.Split(str, "-")
	switch len(vals) {
	case 1:
		i.min = 1
		if i.max, err = strconv.Atoi(vals[0]); err != nil {
			return err
		}
	case 2:
		if i.min, err = strconv.Atoi(vals[0]); err != nil {
			return err
		}
		if i.max, err = strconv.Atoi(vals[1]); err != nil {
			return err
		}
	}
	i.isSet = true
	return nil
}

func (i *interval) Sequence() []string {
	s := make([]string, i.max-i.min+1)

	for idx := range s {
		s[idx] = strconv.Itoa(i.min + idx)
	}
	return s
}
