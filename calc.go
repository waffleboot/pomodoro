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
	last  int
	start int
	end   int
	limit int
	n     int
}

func newSupplier(start, end, limit int) supplier {
	return supplier{start: start, end: end, limit: limit, n: 0}
}

func (s *supplier) request() (ans int) {
	if s.n == 0 {
		ans = s.start
	} else if s.n < s.limit-1 {
		delta := float32(s.end-s.start) / float32(s.limit-1)
		ans = int(float32(s.start) + delta*float32(s.n))
	} else {
		ans = s.end
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
	workStart := cfg.work
	workEnd := cfg.work
	smallStart := cfg.small
	smallEnd := cfg.small
	largeStart := cfg.large
	largeEnd := cfg.large

	workMin, _ := minmax(workStart, workEnd)
	smallMin, _ := minmax(smallStart, smallEnd)
	largeMin, _ := minmax(largeStart, largeEnd)

	for w := limit / workMin; w >= 1; w-- {
		for s := limit / smallMin; s >= 1; s-- {
			for l := limit / largeMin; l >= 1; l-- {
				workSupplier := newSupplier(workStart, workEnd, w)
				smallSupplier := newSupplier(smallStart, smallEnd, s)
				largeSupplier := newSupplier(largeStart, largeEnd, l)
				calcImpl(cfg, func(typ, int) {}, &workSupplier, &smallSupplier, &largeSupplier)
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
	workSupplier := newSupplier(cfg.work, cfg.work, w)
	smallSupplier := newSupplier(cfg.small, cfg.small, s)
	largeSupplier := newSupplier(cfg.large, cfg.large, l)
	calcImpl(cfg, func(t typ, e int) {
		items = append(items, item{t, e})
	}, &workSupplier, &smallSupplier, &largeSupplier)
	return items
}

func calcImpl(cfg *config, add func(typ, int), workSupplier, smallSupplier, largeSupplier *supplier) {
	var work, total, workCount int

	remainder := makeRemainder(cfg)

	for {
		workCount++
		w := workSupplier.request()
		if remainder := remainder(work, total); remainder <= w {
			add(WORK, remainder)
			return
		}
		work += w
		total += w
		add(WORK, w)
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
		add(relaxtype, relaxperiod)
	}
}
