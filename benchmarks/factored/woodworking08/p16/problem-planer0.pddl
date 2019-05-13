(define (problem wood-prob) (:domain woodworking)
(:objects
	red - acolour
	green - acolour
	mauve - acolour
	black - acolour
	blue - acolour
	white - acolour
	walnut - awood
	teak - awood
	cherry - awood
	beech - awood
	oak - awood
	p0 - part
	p1 - part
	p2 - part
	p3 - part
	p4 - part
	p5 - part
	p6 - part
	p7 - part
	p8 - part
	p9 - part
	p10 - part
	p11 - part
	p12 - part
	p13 - part
	p14 - part
	p15 - part
	p16 - part
	p17 - part
	b0 - board
	b1 - board
	b2 - board
	b3 - board
	b4 - board
	b5 - board
	b6 - board
	s0 - aboardsize
	s1 - aboardsize
	s2 - aboardsize
	s3 - aboardsize
	s4 - aboardsize
	s5 - aboardsize
	s6 - aboardsize
	s7 - aboardsize
	s8 - aboardsize
	s9 - aboardsize

	(:private
		planer0 - planer
	)
)
(:init
	(is-smooth smooth)
	(is-smooth verysmooth)
	(boardsize-successor s0 s1)
	(boardsize-successor s1 s2)
	(boardsize-successor s2 s3)
	(boardsize-successor s3 s4)
	(boardsize-successor s4 s5)
	(boardsize-successor s5 s6)
	(boardsize-successor s6 s7)
	(boardsize-successor s7 s8)
	(boardsize-successor s8 s9)
	(unused p0)
	(goalsize p0 large)
	(unused p1)
	(goalsize p1 small)
	(available p2)
	(colour p2 mauve)
	(wood p2 oak)
	(surface-condition p2 rough)
	(treatment p2 glazed)
	(goalsize p2 small)
	(unused p3)
	(goalsize p3 large)
	(unused p4)
	(goalsize p4 medium)
	(unused p5)
	(goalsize p5 large)
	(unused p6)
	(goalsize p6 small)
	(unused p7)
	(goalsize p7 large)
	(available p8)
	(colour p8 black)
	(wood p8 walnut)
	(surface-condition p8 verysmooth)
	(treatment p8 colourfragments)
	(goalsize p8 small)
	(unused p9)
	(goalsize p9 medium)
	(unused p10)
	(goalsize p10 small)
	(unused p11)
	(goalsize p11 large)
	(available p12)
	(colour p12 red)
	(wood p12 walnut)
	(surface-condition p12 verysmooth)
	(treatment p12 glazed)
	(goalsize p12 small)
	(unused p13)
	(goalsize p13 small)
	(unused p14)
	(goalsize p14 small)
	(unused p15)
	(goalsize p15 large)
	(unused p16)
	(goalsize p16 large)
	(available p17)
	(colour p17 natural)
	(wood p17 teak)
	(surface-condition p17 verysmooth)
	(treatment p17 varnished)
	(goalsize p17 large)
	(boardsize b0 s9)
	(wood b0 beech)
	(surface-condition b0 smooth)
	(available b0)
	(boardsize b1 s3)
	(wood b1 beech)
	(surface-condition b1 rough)
	(available b1)
	(boardsize b2 s4)
	(wood b2 teak)
	(surface-condition b2 rough)
	(available b2)
	(boardsize b3 s9)
	(wood b3 oak)
	(surface-condition b3 rough)
	(available b3)
	(boardsize b4 s2)
	(wood b4 oak)
	(surface-condition b4 rough)
	(available b4)
	(boardsize b5 s4)
	(wood b5 cherry)
	(surface-condition b5 rough)
	(available b5)
	(boardsize b6 s6)
	(wood b6 walnut)
	(surface-condition b6 rough)
	(available b6)
	(= (total-cost) 0) 
	(= (spray-varnish-cost p0) 15) 
	(= (glaze-cost p0) 20) 
	(= (grind-cost p0) 45) 
	(= (plane-cost p0) 30) 
	(= (spray-varnish-cost p1) 5) 
	(= (glaze-cost p1) 10) 
	(= (grind-cost p1) 15) 
	(= (plane-cost p1) 10) 
	(= (spray-varnish-cost p2) 5) 
	(= (glaze-cost p2) 10) 
	(= (grind-cost p2) 15) 
	(= (plane-cost p2) 10) 
	(= (spray-varnish-cost p3) 15) 
	(= (glaze-cost p3) 20) 
	(= (grind-cost p3) 45) 
	(= (plane-cost p3) 30) 
	(= (spray-varnish-cost p4) 10) 
	(= (glaze-cost p4) 15) 
	(= (grind-cost p4) 30) 
	(= (plane-cost p4) 20) 
	(= (spray-varnish-cost p5) 15) 
	(= (glaze-cost p5) 20) 
	(= (grind-cost p5) 45) 
	(= (plane-cost p5) 30) 
	(= (spray-varnish-cost p6) 5) 
	(= (glaze-cost p6) 10) 
	(= (grind-cost p6) 15) 
	(= (plane-cost p6) 10) 
	(= (spray-varnish-cost p7) 15) 
	(= (glaze-cost p7) 20) 
	(= (grind-cost p7) 45) 
	(= (plane-cost p7) 30) 
	(= (spray-varnish-cost p8) 5) 
	(= (glaze-cost p8) 10) 
	(= (grind-cost p8) 15) 
	(= (plane-cost p8) 10) 
	(= (spray-varnish-cost p9) 10) 
	(= (glaze-cost p9) 15) 
	(= (grind-cost p9) 30) 
	(= (plane-cost p9) 20) 
	(= (spray-varnish-cost p10) 5) 
	(= (glaze-cost p10) 10) 
	(= (grind-cost p10) 15) 
	(= (plane-cost p10) 10) 
	(= (spray-varnish-cost p11) 15) 
	(= (glaze-cost p11) 20) 
	(= (grind-cost p11) 45) 
	(= (plane-cost p11) 30) 
	(= (spray-varnish-cost p12) 5) 
	(= (glaze-cost p12) 10) 
	(= (grind-cost p12) 15) 
	(= (plane-cost p12) 10) 
	(= (spray-varnish-cost p13) 5) 
	(= (glaze-cost p13) 10) 
	(= (grind-cost p13) 15) 
	(= (plane-cost p13) 10) 
	(= (spray-varnish-cost p14) 5) 
	(= (glaze-cost p14) 10) 
	(= (grind-cost p14) 15) 
	(= (plane-cost p14) 10) 
	(= (spray-varnish-cost p15) 15) 
	(= (glaze-cost p15) 20) 
	(= (grind-cost p15) 45) 
	(= (plane-cost p15) 30) 
	(= (spray-varnish-cost p16) 15) 
	(= (glaze-cost p16) 20) 
	(= (grind-cost p16) 45) 
	(= (plane-cost p16) 30) 
	(= (spray-varnish-cost p17) 15) 
	(= (glaze-cost p17) 20) 
	(= (grind-cost p17) 45) 
	(= (plane-cost p17) 30) 
)
(:goal
	(and
		(available p0)
		(surface-condition p0 verysmooth)
		(treatment p0 varnished)
		(available p1)
		(surface-condition p1 smooth)
		(treatment p1 varnished)
		(available p2)
		(colour p2 green)
		(surface-condition p2 smooth)
		(treatment p2 glazed)
		(available p3)
		(colour p3 natural)
		(wood p3 walnut)
		(surface-condition p3 smooth)
		(treatment p3 glazed)
		(available p4)
		(colour p4 white)
		(wood p4 teak)
		(available p5)
		(colour p5 mauve)
		(wood p5 beech)
		(available p6)
		(colour p6 green)
		(wood p6 oak)
		(available p7)
		(surface-condition p7 verysmooth)
		(treatment p7 glazed)
		(available p8)
		(colour p8 natural)
		(treatment p8 varnished)
		(available p9)
		(wood p9 oak)
		(surface-condition p9 verysmooth)
		(available p10)
		(colour p10 blue)
		(surface-condition p10 verysmooth)
		(available p11)
		(colour p11 blue)
		(wood p11 beech)
		(available p12)
		(colour p12 green)
		(wood p12 walnut)
		(surface-condition p12 verysmooth)
		(treatment p12 glazed)
		(available p13)
		(colour p13 black)
		(surface-condition p13 verysmooth)
		(treatment p13 varnished)
		(available p14)
		(colour p14 blue)
		(wood p14 beech)
		(available p15)
		(colour p15 mauve)
		(wood p15 beech)
		(available p16)
		(surface-condition p16 smooth)
		(treatment p16 glazed)
		(available p17)
		(colour p17 blue)
		(wood p17 teak)
	)
)
(:metric minimize (total-cost))
)