(define (problem wood-prob) (:domain woodworking)
(:objects
	blue - acolour
	black - acolour
	red - acolour
	white - acolour
	mauve - acolour
	green - acolour
	pine - awood
	beech - awood
	walnut - awood
	cherry - awood
	teak - awood
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
	p18 - part
	p19 - part
	p20 - part
	b0 - board
	b1 - board
	b2 - board
	b3 - board
	b4 - board
	b5 - board
	b6 - board
	b7 - board
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
		immersion-varnisher0 - immersion-varnisher
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
	(has-colour immersion-varnisher0 blue)
	(has-colour immersion-varnisher0 natural)
	(has-colour immersion-varnisher0 mauve)
	(has-colour immersion-varnisher0 black)
	(has-colour immersion-varnisher0 green)
	(has-colour immersion-varnisher0 red)
	(unused p0)
	(goalsize p0 small)
	(unused p1)
	(goalsize p1 small)
	(unused p2)
	(goalsize p2 medium)
	(unused p3)
	(goalsize p3 small)
	(unused p4)
	(goalsize p4 medium)
	(available p5)
	(colour p5 mauve)
	(wood p5 teak)
	(surface-condition p5 rough)
	(treatment p5 glazed)
	(goalsize p5 small)
	(unused p6)
	(goalsize p6 large)
	(unused p7)
	(goalsize p7 large)
	(unused p8)
	(goalsize p8 large)
	(unused p9)
	(goalsize p9 medium)
	(unused p10)
	(goalsize p10 large)
	(available p11)
	(colour p11 white)
	(wood p11 walnut)
	(surface-condition p11 rough)
	(treatment p11 colourfragments)
	(goalsize p11 small)
	(available p12)
	(colour p12 black)
	(wood p12 teak)
	(surface-condition p12 rough)
	(treatment p12 glazed)
	(goalsize p12 small)
	(available p13)
	(colour p13 green)
	(wood p13 beech)
	(surface-condition p13 smooth)
	(treatment p13 varnished)
	(goalsize p13 medium)
	(unused p14)
	(goalsize p14 large)
	(unused p15)
	(goalsize p15 large)
	(unused p16)
	(goalsize p16 medium)
	(unused p17)
	(goalsize p17 small)
	(unused p18)
	(goalsize p18 medium)
	(unused p19)
	(goalsize p19 small)
	(available p20)
	(colour p20 white)
	(wood p20 beech)
	(surface-condition p20 smooth)
	(treatment p20 colourfragments)
	(goalsize p20 medium)
	(boardsize b0 s8)
	(wood b0 beech)
	(surface-condition b0 rough)
	(available b0)
	(boardsize b1 s3)
	(wood b1 beech)
	(surface-condition b1 rough)
	(available b1)
	(boardsize b2 s5)
	(wood b2 teak)
	(surface-condition b2 rough)
	(available b2)
	(boardsize b3 s5)
	(wood b3 cherry)
	(surface-condition b3 rough)
	(available b3)
	(boardsize b4 s7)
	(wood b4 pine)
	(surface-condition b4 rough)
	(available b4)
	(boardsize b5 s1)
	(wood b5 pine)
	(surface-condition b5 rough)
	(available b5)
	(boardsize b6 s9)
	(wood b6 walnut)
	(surface-condition b6 rough)
	(available b6)
	(boardsize b7 s3)
	(wood b7 walnut)
	(surface-condition b7 rough)
	(available b7)
	(= (total-cost) 0) 
	(= (spray-varnish-cost p0) 5) 
	(= (glaze-cost p0) 10) 
	(= (grind-cost p0) 15) 
	(= (plane-cost p0) 10) 
	(= (spray-varnish-cost p1) 5) 
	(= (glaze-cost p1) 10) 
	(= (grind-cost p1) 15) 
	(= (plane-cost p1) 10) 
	(= (spray-varnish-cost p2) 10) 
	(= (glaze-cost p2) 15) 
	(= (grind-cost p2) 30) 
	(= (plane-cost p2) 20) 
	(= (spray-varnish-cost p3) 5) 
	(= (glaze-cost p3) 10) 
	(= (grind-cost p3) 15) 
	(= (plane-cost p3) 10) 
	(= (spray-varnish-cost p4) 10) 
	(= (glaze-cost p4) 15) 
	(= (grind-cost p4) 30) 
	(= (plane-cost p4) 20) 
	(= (spray-varnish-cost p5) 5) 
	(= (glaze-cost p5) 10) 
	(= (grind-cost p5) 15) 
	(= (plane-cost p5) 10) 
	(= (spray-varnish-cost p6) 15) 
	(= (glaze-cost p6) 20) 
	(= (grind-cost p6) 45) 
	(= (plane-cost p6) 30) 
	(= (spray-varnish-cost p7) 15) 
	(= (glaze-cost p7) 20) 
	(= (grind-cost p7) 45) 
	(= (plane-cost p7) 30) 
	(= (spray-varnish-cost p8) 15) 
	(= (glaze-cost p8) 20) 
	(= (grind-cost p8) 45) 
	(= (plane-cost p8) 30) 
	(= (spray-varnish-cost p9) 10) 
	(= (glaze-cost p9) 15) 
	(= (grind-cost p9) 30) 
	(= (plane-cost p9) 20) 
	(= (spray-varnish-cost p10) 15) 
	(= (glaze-cost p10) 20) 
	(= (grind-cost p10) 45) 
	(= (plane-cost p10) 30) 
	(= (spray-varnish-cost p11) 5) 
	(= (glaze-cost p11) 10) 
	(= (grind-cost p11) 15) 
	(= (plane-cost p11) 10) 
	(= (spray-varnish-cost p12) 5) 
	(= (glaze-cost p12) 10) 
	(= (grind-cost p12) 15) 
	(= (plane-cost p12) 10) 
	(= (spray-varnish-cost p13) 10) 
	(= (glaze-cost p13) 15) 
	(= (grind-cost p13) 30) 
	(= (plane-cost p13) 20) 
	(= (spray-varnish-cost p14) 15) 
	(= (glaze-cost p14) 20) 
	(= (grind-cost p14) 45) 
	(= (plane-cost p14) 30) 
	(= (spray-varnish-cost p15) 15) 
	(= (glaze-cost p15) 20) 
	(= (grind-cost p15) 45) 
	(= (plane-cost p15) 30) 
	(= (spray-varnish-cost p16) 10) 
	(= (glaze-cost p16) 15) 
	(= (grind-cost p16) 30) 
	(= (plane-cost p16) 20) 
	(= (spray-varnish-cost p17) 5) 
	(= (glaze-cost p17) 10) 
	(= (grind-cost p17) 15) 
	(= (plane-cost p17) 10) 
	(= (spray-varnish-cost p18) 10) 
	(= (glaze-cost p18) 15) 
	(= (grind-cost p18) 30) 
	(= (plane-cost p18) 20) 
	(= (spray-varnish-cost p19) 5) 
	(= (glaze-cost p19) 10) 
	(= (grind-cost p19) 15) 
	(= (plane-cost p19) 10) 
	(= (spray-varnish-cost p20) 10) 
	(= (glaze-cost p20) 15) 
	(= (grind-cost p20) 30) 
	(= (plane-cost p20) 20) 
)
(:goal
	(and
		(available p0)
		(colour p0 green)
		(wood p0 pine)
		(available p1)
		(colour p1 natural)
		(treatment p1 glazed)
		(available p2)
		(colour p2 natural)
		(surface-condition p2 smooth)
		(available p3)
		(colour p3 natural)
		(treatment p3 glazed)
		(available p4)
		(colour p4 mauve)
		(wood p4 teak)
		(surface-condition p4 verysmooth)
		(available p5)
		(colour p5 black)
		(wood p5 teak)
		(surface-condition p5 smooth)
		(treatment p5 varnished)
		(available p6)
		(colour p6 natural)
		(treatment p6 varnished)
		(available p7)
		(colour p7 red)
		(wood p7 cherry)
		(surface-condition p7 smooth)
		(treatment p7 glazed)
		(available p8)
		(colour p8 green)
		(wood p8 beech)
		(surface-condition p8 verysmooth)
		(treatment p8 varnished)
		(available p9)
		(colour p9 red)
		(wood p9 pine)
		(surface-condition p9 smooth)
		(treatment p9 glazed)
		(available p10)
		(colour p10 mauve)
		(surface-condition p10 smooth)
		(treatment p10 varnished)
		(available p11)
		(colour p11 mauve)
		(wood p11 walnut)
		(surface-condition p11 smooth)
		(treatment p11 varnished)
		(available p12)
		(colour p12 red)
		(wood p12 teak)
		(available p13)
		(colour p13 red)
		(surface-condition p13 verysmooth)
		(treatment p13 varnished)
		(available p14)
		(colour p14 black)
		(wood p14 beech)
		(available p15)
		(colour p15 mauve)
		(wood p15 walnut)
		(available p16)
		(wood p16 beech)
		(treatment p16 glazed)
		(available p17)
		(colour p17 mauve)
		(treatment p17 varnished)
		(available p18)
		(colour p18 red)
		(treatment p18 glazed)
		(available p19)
		(colour p19 blue)
		(surface-condition p19 verysmooth)
		(available p20)
		(surface-condition p20 verysmooth)
		(treatment p20 glazed)
	)
)
(:metric minimize (total-cost))
)