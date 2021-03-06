(define (problem BLOCKS-4-0) (:domain blocks)
(:objects
	a - block
	c - block
	b - block
	e - block
	d - block
	g - block
	f - block
	i - block
	h - block
	k - block
	j - block
	m - block
	l - block
	n - block

	(:private
		a3 - agent
	)
)
(:init
	(handempty a3)
	(clear a)
	(clear g)
	(clear f)
	(ontable e)
	(ontable n)
	(ontable f)
	(on a j)
	(on j h)
	(on h m)
	(on m k)
	(on k c)
	(on c l)
	(on l b)
	(on b e)
	(on g d)
	(on d i)
	(on i n)
)
(:goal
	(and
		(on e l)
		(on l f)
		(on f b)
		(on b j)
		(on j i)
		(on i n)
		(on n c)
		(on c k)
		(on k g)
		(on g d)
		(on d m)
		(on m a)
		(on a h)
	)
)
)