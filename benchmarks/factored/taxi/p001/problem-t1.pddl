(define (problem taxi-001) (:domain taxi)
(:objects
	a - location
	b - location
    c - location
)
(:init
	(directly-connected a b)
	(directly-connected b a)
	(directly-connected b c)
	(directly-connected c b)
	(at p1 a)
    (at p2 a)
	(at t1 a)
	(empty t1)
	(free b)
    (free c)
)
(:goal
	(and
		(at t1 c)
        (at p2 c)
		(at p1 c)
	)
)
)
