(define (problem taxi-00) (:domain taxi)
(:objects
	a - location
	b - location
)
(:init
	(directly-connected a b)
	(directly-connected b a)
	(at t1 a)
	(empty t1)
	(at p1 a)
	(free b)
	(goal-of p1 b)
	
)
(:goal
	(and
		(at t1 b)
		(at p1 b)
	)
)
)
