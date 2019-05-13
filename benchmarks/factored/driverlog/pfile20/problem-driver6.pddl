(define (problem DLOG-8-6-25) (:domain driverlog)
(:objects
	package25 - package
	package24 - package
	package23 - package
	package22 - package
	package21 - package
	package20 - package
	package12 - package
	p5-17 - location
	truck5 - truck
	truck4 - truck
	truck6 - truck
	truck1 - truck
	truck3 - truck
	truck2 - truck
	p6-11 - location
	p13-18 - location
	p8-0 - location
	p8-4 - location
	p0-2 - location
	p7-18 - location
	p7-19 - location
	p7-15 - location
	p2-9 - location
	p7-10 - location
	package18 - package
	p16-12 - location
	package11 - package
	p4-7 - location
	p4-3 - location
	p5-8 - location
	p4-9 - location
	p3-12 - location
	p3-11 - location
	p15-13 - location
	package16 - package
	p14-0 - location
	package8 - package
	package9 - package
	package19 - package
	p14-13 - location
	package17 - package
	package1 - package
	package2 - package
	package3 - package
	package4 - package
	package5 - package
	package6 - package
	package7 - package
	s9 - location
	s8 - location
	s3 - location
	s2 - location
	s1 - location
	s0 - location
	s7 - location
	s6 - location
	s5 - location
	s4 - location
	p0-16 - location
	package13 - package
	p7-6 - location
	package10 - package
	p2-15 - location
	p2-14 - location
	s19 - location
	s18 - location
	p18-6 - location
	s13 - location
	s12 - location
	s11 - location
	s10 - location
	s17 - location
	s16 - location
	s15 - location
	s14 - location
	p10-13 - location
	package14 - package
	package15 - package
	p1-14 - location
	p1-15 - location
	p4-11 - location
	p4-16 - location
	p13-5 - location
	p14-15 - location
	p13-8 - location
	p13-9 - location
	p12-1 - location
	p19-18 - location
	p15-5 - location
	p19-11 - location

	(:private
		driver6 - driver
	)
)
(:init
	(at driver6 s0)
	(at truck1 s13)
	(empty truck1)
	(at truck2 s8)
	(empty truck2)
	(at truck3 s5)
	(empty truck3)
	(at truck4 s17)
	(empty truck4)
	(at truck5 s16)
	(empty truck5)
	(at truck6 s16)
	(empty truck6)
	(at package1 s18)
	(at package2 s15)
	(at package3 s8)
	(at package4 s2)
	(at package5 s8)
	(at package6 s9)
	(at package7 s15)
	(at package8 s15)
	(at package9 s1)
	(at package10 s3)
	(at package11 s1)
	(at package12 s6)
	(at package13 s16)
	(at package14 s14)
	(at package15 s11)
	(at package16 s17)
	(at package17 s14)
	(at package18 s18)
	(at package19 s0)
	(at package20 s0)
	(at package21 s12)
	(at package22 s11)
	(at package23 s17)
	(at package24 s6)
	(at package25 s8)
	(path s0 p0-2)
	(path p0-2 s0)
	(path s2 p0-2)
	(path p0-2 s2)
	(path s0 p0-16)
	(path p0-16 s0)
	(path s16 p0-16)
	(path p0-16 s16)
	(path s1 p1-14)
	(path p1-14 s1)
	(path s14 p1-14)
	(path p1-14 s14)
	(path s1 p1-15)
	(path p1-15 s1)
	(path s15 p1-15)
	(path p1-15 s15)
	(path s2 p2-9)
	(path p2-9 s2)
	(path s9 p2-9)
	(path p2-9 s9)
	(path s2 p2-14)
	(path p2-14 s2)
	(path s14 p2-14)
	(path p2-14 s14)
	(path s2 p2-15)
	(path p2-15 s2)
	(path s15 p2-15)
	(path p2-15 s15)
	(path s3 p3-11)
	(path p3-11 s3)
	(path s11 p3-11)
	(path p3-11 s11)
	(path s3 p3-12)
	(path p3-12 s3)
	(path s12 p3-12)
	(path p3-12 s12)
	(path s4 p4-3)
	(path p4-3 s4)
	(path s3 p4-3)
	(path p4-3 s3)
	(path s4 p4-7)
	(path p4-7 s4)
	(path s7 p4-7)
	(path p4-7 s7)
	(path s4 p4-9)
	(path p4-9 s4)
	(path s9 p4-9)
	(path p4-9 s9)
	(path s4 p4-11)
	(path p4-11 s4)
	(path s11 p4-11)
	(path p4-11 s11)
	(path s4 p4-16)
	(path p4-16 s4)
	(path s16 p4-16)
	(path p4-16 s16)
	(path s5 p5-8)
	(path p5-8 s5)
	(path s8 p5-8)
	(path p5-8 s8)
	(path s5 p5-17)
	(path p5-17 s5)
	(path s17 p5-17)
	(path p5-17 s17)
	(path s6 p6-11)
	(path p6-11 s6)
	(path s11 p6-11)
	(path p6-11 s11)
	(path s7 p7-6)
	(path p7-6 s7)
	(path s6 p7-6)
	(path p7-6 s6)
	(path s7 p7-10)
	(path p7-10 s7)
	(path s10 p7-10)
	(path p7-10 s10)
	(path s7 p7-15)
	(path p7-15 s7)
	(path s15 p7-15)
	(path p7-15 s15)
	(path s7 p7-18)
	(path p7-18 s7)
	(path s18 p7-18)
	(path p7-18 s18)
	(path s7 p7-19)
	(path p7-19 s7)
	(path s19 p7-19)
	(path p7-19 s19)
	(path s8 p8-0)
	(path p8-0 s8)
	(path s0 p8-0)
	(path p8-0 s0)
	(path s8 p8-4)
	(path p8-4 s8)
	(path s4 p8-4)
	(path p8-4 s4)
	(path s10 p10-13)
	(path p10-13 s10)
	(path s13 p10-13)
	(path p10-13 s13)
	(path s12 p12-1)
	(path p12-1 s12)
	(path s1 p12-1)
	(path p12-1 s1)
	(path s13 p13-5)
	(path p13-5 s13)
	(path s5 p13-5)
	(path p13-5 s5)
	(path s13 p13-8)
	(path p13-8 s13)
	(path s8 p13-8)
	(path p13-8 s8)
	(path s13 p13-9)
	(path p13-9 s13)
	(path s9 p13-9)
	(path p13-9 s9)
	(path s13 p13-18)
	(path p13-18 s13)
	(path s18 p13-18)
	(path p13-18 s18)
	(path s14 p14-0)
	(path p14-0 s14)
	(path s0 p14-0)
	(path p14-0 s0)
	(path s14 p14-13)
	(path p14-13 s14)
	(path s13 p14-13)
	(path p14-13 s13)
	(path s14 p14-15)
	(path p14-15 s14)
	(path s15 p14-15)
	(path p14-15 s15)
	(path s15 p15-5)
	(path p15-5 s15)
	(path s5 p15-5)
	(path p15-5 s5)
	(path s15 p15-13)
	(path p15-13 s15)
	(path s13 p15-13)
	(path p15-13 s13)
	(path s16 p16-12)
	(path p16-12 s16)
	(path s12 p16-12)
	(path p16-12 s12)
	(path s18 p18-6)
	(path p18-6 s18)
	(path s6 p18-6)
	(path p18-6 s6)
	(path s19 p19-11)
	(path p19-11 s19)
	(path s11 p19-11)
	(path p19-11 s11)
	(path s19 p19-18)
	(path p19-18 s19)
	(path s18 p19-18)
	(path p19-18 s18)
	(link s1 s0)
	(link s0 s1)
	(link s1 s6)
	(link s6 s1)
	(link s1 s9)
	(link s9 s1)
	(link s2 s6)
	(link s6 s2)
	(link s2 s18)
	(link s18 s2)
	(link s2 s19)
	(link s19 s2)
	(link s3 s1)
	(link s1 s3)
	(link s3 s5)
	(link s5 s3)
	(link s3 s7)
	(link s7 s3)
	(link s3 s11)
	(link s11 s3)
	(link s3 s18)
	(link s18 s3)
	(link s4 s3)
	(link s3 s4)
	(link s4 s5)
	(link s5 s4)
	(link s4 s12)
	(link s12 s4)
	(link s5 s0)
	(link s0 s5)
	(link s5 s9)
	(link s9 s5)
	(link s5 s16)
	(link s16 s5)
	(link s5 s17)
	(link s17 s5)
	(link s6 s3)
	(link s3 s6)
	(link s6 s7)
	(link s7 s6)
	(link s6 s11)
	(link s11 s6)
	(link s6 s12)
	(link s12 s6)
	(link s6 s16)
	(link s16 s6)
	(link s6 s18)
	(link s18 s6)
	(link s7 s4)
	(link s4 s7)
	(link s7 s9)
	(link s9 s7)
	(link s7 s18)
	(link s18 s7)
	(link s8 s0)
	(link s0 s8)
	(link s8 s6)
	(link s6 s8)
	(link s8 s9)
	(link s9 s8)
	(link s8 s10)
	(link s10 s8)
	(link s8 s11)
	(link s11 s8)
	(link s8 s12)
	(link s12 s8)
	(link s8 s15)
	(link s15 s8)
	(link s8 s19)
	(link s19 s8)
	(link s9 s15)
	(link s15 s9)
	(link s9 s18)
	(link s18 s9)
	(link s10 s0)
	(link s0 s10)
	(link s10 s13)
	(link s13 s10)
	(link s10 s15)
	(link s15 s10)
	(link s10 s16)
	(link s16 s10)
	(link s11 s1)
	(link s1 s11)
	(link s11 s5)
	(link s5 s11)
	(link s11 s10)
	(link s10 s11)
	(link s11 s16)
	(link s16 s11)
	(link s12 s1)
	(link s1 s12)
	(link s12 s15)
	(link s15 s12)
	(link s12 s16)
	(link s16 s12)
	(link s13 s2)
	(link s2 s13)
	(link s13 s9)
	(link s9 s13)
	(link s13 s17)
	(link s17 s13)
	(link s14 s2)
	(link s2 s14)
	(link s14 s4)
	(link s4 s14)
	(link s14 s8)
	(link s8 s14)
	(link s15 s4)
	(link s4 s15)
	(link s15 s5)
	(link s5 s15)
	(link s15 s16)
	(link s16 s15)
	(link s15 s17)
	(link s17 s15)
	(link s16 s0)
	(link s0 s16)
	(link s16 s14)
	(link s14 s16)
	(link s16 s17)
	(link s17 s16)
	(link s17 s8)
	(link s8 s17)
	(link s17 s11)
	(link s11 s17)
	(link s18 s12)
	(link s12 s18)
	(link s18 s14)
	(link s14 s18)
	(link s18 s19)
	(link s19 s18)
	(link s19 s6)
	(link s6 s19)
	(link s19 s10)
	(link s10 s19)
)
(:goal
	(and
		(at truck3 s5)
		(at truck4 s18)
		(at truck6 s12)
		(at package1 s6)
		(at package2 s7)
		(at package3 s3)
		(at package4 s11)
		(at package5 s5)
		(at package6 s14)
		(at package7 s19)
		(at package8 s16)
		(at package9 s13)
		(at package10 s9)
		(at package11 s7)
		(at package12 s3)
		(at package13 s11)
		(at package14 s14)
		(at package15 s2)
		(at package16 s12)
		(at package18 s2)
		(at package19 s4)
		(at package20 s7)
		(at package21 s8)
		(at package22 s14)
		(at package23 s10)
		(at package24 s4)
		(at package25 s16)
	)
)
)