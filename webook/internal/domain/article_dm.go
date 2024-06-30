package domain

import "time"

type Article struct {
	Id      int64
	Title   string
	Content string
	Ctime   time.Time
	Utime   time.Time
	Status  int
	Author  Author
}

func (a *Article) Abstract() string{
	abs := []rune(a.Content)
	if len(abs) < 100 {
		return string(abs)
	}
	return string(abs[:100])
}

type Author struct {
	Id   int64
	Name string
}
