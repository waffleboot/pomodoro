package main

type typ int

const (
	WORK typ = iota
	SMALL
	LARGE
)

type item struct {
	typ     typ
	elapsed int
	full    bool
}

func makeRemainder(cfg *config) func(int, int) int {
	if cfg.mode {
		return func(work int, total int) int {
			return cfg.timelimit - work
		}
	}
	return func(work int, total int) int {
		return cfg.timelimit - total
	}
}

type supplier struct {
	i     interval
	limit int
	last  int
	n     int
}

func newSupplier(i interval, limit int) supplier {
	return supplier{i: i, limit: limit, n: 0}
}

func (s *supplier) request() (ans int) {
	if s.n == 0 {
		ans = s.i.start
	} else if s.n < s.limit-1 {
		delta := float32(s.i.end-s.i.start) / float32(s.limit-1)
		ans = int(float32(s.i.start) + delta*float32(s.n))
	} else {
		ans = s.i.end
	}
	s.last = ans
	s.n++
	return
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func find(cfg *config) (int, int, int) {

	limit := cfg.timelimit
	workStart := cfg.work.start
	workEnd := cfg.work.end
	smallStart := cfg.small.start
	smallEnd := cfg.small.end
	largeStart := cfg.large.start
	largeEnd := cfg.large.end

	workMin, _ := minmax(workStart, workEnd)
	smallMin, _ := minmax(smallStart, smallEnd)
	largeMin, _ := minmax(largeStart, largeEnd)

	for w := limit / workMin; w >= 1; w-- {
		for s := limit / smallMin; s >= 1; s-- {
			for l := limit / largeMin; l >= 1; l-- {
				workSupplier := newSupplier(cfg.work, w)
				smallSupplier := newSupplier(cfg.small, s)
				largeSupplier := newSupplier(cfg.large, l)
				calcImpl(cfg, func(typ, int, bool) {}, &workSupplier, &smallSupplier, &largeSupplier)
				if workSupplier.last == workEnd && smallSupplier.last == smallEnd && largeSupplier.last == largeEnd {
					return w, s, l
				}
			}
		}
	}
	return 0, 0, 0
}

func calc(cfg *config) []item {
	w, s, l := find(cfg)
	items := make([]item, 0, 10)
	workSupplier := newSupplier(cfg.work, w)
	smallSupplier := newSupplier(cfg.small, s)
	largeSupplier := newSupplier(cfg.large, l)
	calcImpl(cfg, func(t typ, e int, f bool) {
		items = append(items, item{t, e, f})
	}, &workSupplier, &smallSupplier, &largeSupplier)
	return items
}

func calcImpl(cfg *config, add func(typ, int, bool), workSupplier, smallSupplier, largeSupplier *supplier) {
	var work, total, workCount int

	remainder := makeRemainder(cfg)

	for {
		workCount++
		w := workSupplier.request()
		if remainder := remainder(work, total); remainder <= w {
			add(WORK, remainder, false)
			return
		}
		work += w
		total += w
		add(WORK, w, true)
		relaxtype := SMALL
		var relaxperiod int
		if workCount == cfg.worklimit {
			relaxtype = LARGE
			relaxperiod = largeSupplier.request()
			workCount = 0
		} else {
			relaxperiod = smallSupplier.request()
		}
		total += relaxperiod
		if remainder(work, total) <= 0 {
			return
		}
		add(relaxtype, relaxperiod, false)
	}
}
