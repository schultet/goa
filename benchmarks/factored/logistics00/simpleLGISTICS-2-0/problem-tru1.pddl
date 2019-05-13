(define (problem logistics-0) (:domain logistics)
(:objects
	p - package
	x - location
	y - location
	z - location

	(:private
		t1 - truck
	)
)
(:init
    (p-at p z)
    (at t1 z)
)
(:goal
    (p-at p x)
)
)
