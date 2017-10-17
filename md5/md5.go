// DO NOT EDIT.
// Generate with: go run gen.go -output main.go

package md5

const Size = 16

type Digest struct {
	s   [4]uint32
	len uint64
}

func New() Digest {
	return Digest{s: [4]uint32{0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476}, len: 0}
}

func (dig Digest) Block(X [16]uint32) Digest {
	s := dig.s
	ta0 := s[0]
	tb0 := s[1]
	tc0 := s[2]
	td0 := s[3]

	tmp_ta0 := ((ta0 + ((tb0 & tc0) | (td0 &^ tb0))) + (X[0] + 3614090360))
	ta1 := tmp_ta0<<uint8(7) | tmp_ta0>>uint8(32-7) + tb0

	tmp_td0 := ((td0 + ((ta1 & tb0) | (tc0 &^ ta1))) + (X[1] + 3905402710))
	td1 := tmp_td0<<uint8(12) | tmp_td0>>uint8(32-12) + ta1

	tmp_tc0 := ((tc0 + ((td1 & ta1) | (tb0 &^ td1))) + (X[2] + 606105819))
	tc1 := tmp_tc0<<uint8(17) | tmp_tc0>>uint8(32-17) + td1

	tmp_tb0 := ((tb0 + ((tc1 & td1) | (ta1 &^ tc1))) + (X[3] + 3250441966))
	tb1 := tmp_tb0<<uint8(22) | tmp_tb0>>uint8(32-22) + tc1

	tmp_ta1 := ((ta1 + ((tb1 & tc1) | (td1 &^ tb1))) + (X[4] + 4118548399))
	ta2 := tmp_ta1<<uint8(7) | tmp_ta1>>uint8(32-7) + tb1

	tmp_td1 := ((td1 + ((ta2 & tb1) | (tc1 &^ ta2))) + (X[5] + 1200080426))
	td2 := tmp_td1<<uint8(12) | tmp_td1>>uint8(32-12) + ta2

	tmp_tc1 := ((tc1 + ((td2 & ta2) | (tb1 &^ td2))) + (X[6] + 2821735955))
	tc2 := tmp_tc1<<uint8(17) | tmp_tc1>>uint8(32-17) + td2

	tmp_tb1 := ((tb1 + ((tc2 & td2) | (ta2 &^ tc2))) + (X[7] + 4249261313))
	tb2 := tmp_tb1<<uint8(22) | tmp_tb1>>uint8(32-22) + tc2

	tmp_ta2 := ((ta2 + ((tb2 & tc2) | (td2 &^ tb2))) + (X[8] + 1770035416))
	ta3 := tmp_ta2<<uint8(7) | tmp_ta2>>uint8(32-7) + tb2

	tmp_td2 := ((td2 + ((ta3 & tb2) | (tc2 &^ ta3))) + (X[9] + 2336552879))
	td3 := tmp_td2<<uint8(12) | tmp_td2>>uint8(32-12) + ta3

	tmp_tc2 := ((tc2 + ((td3 & ta3) | (tb2 &^ td3))) + (X[10] + 4294925233))
	tc3 := tmp_tc2<<uint8(17) | tmp_tc2>>uint8(32-17) + td3

	tmp_tb2 := ((tb2 + ((tc3 & td3) | (ta3 &^ tc3))) + (X[11] + 2304563134))
	tb3 := tmp_tb2<<uint8(22) | tmp_tb2>>uint8(32-22) + tc3

	tmp_ta3 := ((ta3 + ((tb3 & tc3) | (td3 &^ tb3))) + (X[12] + 1804603682))
	ta4 := tmp_ta3<<uint8(7) | tmp_ta3>>uint8(32-7) + tb3

	tmp_td3 := ((td3 + ((ta4 & tb3) | (tc3 &^ ta4))) + (X[13] + 4254626195))
	td4 := tmp_td3<<uint8(12) | tmp_td3>>uint8(32-12) + ta4

	tmp_tc3 := ((tc3 + ((td4 & ta4) | (tb3 &^ td4))) + (X[14] + 2792965006))
	tc4 := tmp_tc3<<uint8(17) | tmp_tc3>>uint8(32-17) + td4

	tmp_tb3 := ((tb3 + ((tc4 & td4) | (ta4 &^ tc4))) + (X[15] + 1236535329))
	tb4 := tmp_tb3<<uint8(22) | tmp_tb3>>uint8(32-22) + tc4

	// Round 2.

	tmp_ta4 := ((ta4 + ((tb4 & td4) | (tc4 &^ td4))) + (X[uint32(1+5*0)&15] + 4129170786))
	ta5 := tmp_ta4<<uint8(5) | tmp_ta4>>uint8(32-5) + tb4

	tmp_td4 := ((td4 + ((ta5 & tc4) | (tb4 &^ tc4))) + (X[uint32(1+5*1)&15] + 3225465664))
	td5 := tmp_td4<<uint8(9) | tmp_td4>>uint8(32-9) + ta5

	tmp_tc4 := ((tc4 + ((td5 & tb4) | (ta5 &^ tb4))) + (X[uint32(1+5*2)&15] + 643717713))
	tc5 := tmp_tc4<<uint8(14) | tmp_tc4>>uint8(32-14) + td5

	tmp_tb4 := ((tb4 + ((tc5 & ta5) | (td5 &^ ta5))) + (X[uint32(1+5*3)&15] + 3921069994))
	tb5 := tmp_tb4<<uint8(20) | tmp_tb4>>uint8(32-20) + tc5

	tmp_ta5 := ((ta5 + ((tb5 & td5) | (tc5 &^ td5))) + (X[uint32(1+5*4)&15] + 3593408605))
	ta6 := tmp_ta5<<uint8(5) | tmp_ta5>>uint8(32-5) + tb5

	tmp_td5 := ((td5 + ((ta6 & tc5) | (tb5 &^ tc5))) + (X[uint32(1+5*5)&15] + 38016083))
	td6 := tmp_td5<<uint8(9) | tmp_td5>>uint8(32-9) + ta6

	tmp_tc5 := ((tc5 + ((td6 & tb5) | (ta6 &^ tb5))) + (X[uint32(1+5*6)&15] + 3634488961))
	tc6 := tmp_tc5<<uint8(14) | tmp_tc5>>uint8(32-14) + td6

	tmp_tb5 := ((tb5 + ((tc6 & ta6) | (td6 &^ ta6))) + (X[uint32(1+5*7)&15] + 3889429448))
	tb6 := tmp_tb5<<uint8(20) | tmp_tb5>>uint8(32-20) + tc6

	tmp_ta6 := ((ta6 + ((tb6 & td6) | (tc6 &^ td6))) + (X[uint32(1+5*8)&15] + 568446438))
	ta7 := tmp_ta6<<uint8(5) | tmp_ta6>>uint8(32-5) + tb6

	tmp_td6 := ((td6 + ((ta7 & tc6) | (tb6 &^ tc6))) + (X[uint32(1+5*9)&15] + 3275163606))
	td7 := tmp_td6<<uint8(9) | tmp_td6>>uint8(32-9) + ta7

	tmp_tc6 := ((tc6 + ((td7 & tb6) | (ta7 &^ tb6))) + (X[uint32(1+5*10)&15] + 4107603335))
	tc7 := tmp_tc6<<uint8(14) | tmp_tc6>>uint8(32-14) + td7

	tmp_tb6 := ((tb6 + ((tc7 & ta7) | (td7 &^ ta7))) + (X[uint32(1+5*11)&15] + 1163531501))
	tb7 := tmp_tb6<<uint8(20) | tmp_tb6>>uint8(32-20) + tc7

	tmp_ta7 := ((ta7 + ((tb7 & td7) | (tc7 &^ td7))) + (X[uint32(1+5*12)&15] + 2850285829))
	ta8 := tmp_ta7<<uint8(5) | tmp_ta7>>uint8(32-5) + tb7

	tmp_td7 := ((td7 + ((ta8 & tc7) | (tb7 &^ tc7))) + (X[uint32(1+5*13)&15] + 4243563512))
	td8 := tmp_td7<<uint8(9) | tmp_td7>>uint8(32-9) + ta8

	tmp_tc7 := ((tc7 + ((td8 & tb7) | (ta8 &^ tb7))) + (X[uint32(1+5*14)&15] + 1735328473))
	tc8 := tmp_tc7<<uint8(14) | tmp_tc7>>uint8(32-14) + td8

	tmp_tb7 := ((tb7 + ((tc8 & ta8) | (td8 &^ ta8))) + (X[uint32(1+5*15)&15] + 2368359562))
	tb8 := tmp_tb7<<uint8(20) | tmp_tb7>>uint8(32-20) + tc8

	// Round 3.

	tmp_ta8 := ((ta8 + (tb8 ^ tc8 ^ td8)) + (X[uint32(5+3*0)&15] + 4294588738))
	ta9 := tmp_ta8<<uint8(4) | tmp_ta8>>uint8(32-4) + tb8

	tmp_td8 := ((td8 + (ta9 ^ tb8 ^ tc8)) + (X[uint32(5+3*1)&15] + 2272392833))
	td9 := tmp_td8<<uint8(11) | tmp_td8>>uint8(32-11) + ta9

	tmp_tc8 := ((tc8 + (td9 ^ ta9 ^ tb8)) + (X[uint32(5+3*2)&15] + 1839030562))
	tc9 := tmp_tc8<<uint8(16) | tmp_tc8>>uint8(32-16) + td9

	tmp_tb8 := ((tb8 + (tc9 ^ td9 ^ ta9)) + (X[uint32(5+3*3)&15] + 4259657740))
	tb9 := tmp_tb8<<uint8(23) | tmp_tb8>>uint8(32-23) + tc9

	tmp_ta9 := ((ta9 + (tb9 ^ tc9 ^ td9)) + (X[uint32(5+3*4)&15] + 2763975236))
	ta10 := tmp_ta9<<uint8(4) | tmp_ta9>>uint8(32-4) + tb9

	tmp_td9 := ((td9 + (ta10 ^ tb9 ^ tc9)) + (X[uint32(5+3*5)&15] + 1272893353))
	td10 := tmp_td9<<uint8(11) | tmp_td9>>uint8(32-11) + ta10

	tmp_tc9 := ((tc9 + (td10 ^ ta10 ^ tb9)) + (X[uint32(5+3*6)&15] + 4139469664))
	tc10 := tmp_tc9<<uint8(16) | tmp_tc9>>uint8(32-16) + td10

	tmp_tb9 := ((tb9 + (tc10 ^ td10 ^ ta10)) + (X[uint32(5+3*7)&15] + 3200236656))
	tb10 := tmp_tb9<<uint8(23) | tmp_tb9>>uint8(32-23) + tc10

	tmp_ta10 := ((ta10 + (tb10 ^ tc10 ^ td10)) + (X[uint32(5+3*8)&15] + 681279174))
	ta11 := tmp_ta10<<uint8(4) | tmp_ta10>>uint8(32-4) + tb10

	tmp_td10 := ((td10 + (ta11 ^ tb10 ^ tc10)) + (X[uint32(5+3*9)&15] + 3936430074))
	td11 := tmp_td10<<uint8(11) | tmp_td10>>uint8(32-11) + ta11

	tmp_tc10 := ((tc10 + (td11 ^ ta11 ^ tb10)) + (X[uint32(5+3*10)&15] + 3572445317))
	tc11 := tmp_tc10<<uint8(16) | tmp_tc10>>uint8(32-16) + td11

	tmp_tb10 := ((tb10 + (tc11 ^ td11 ^ ta11)) + (X[uint32(5+3*11)&15] + 76029189))
	tb11 := tmp_tb10<<uint8(23) | tmp_tb10>>uint8(32-23) + tc11

	tmp_ta11 := ((ta11 + (tb11 ^ tc11 ^ td11)) + (X[uint32(5+3*12)&15] + 3654602809))
	ta12 := tmp_ta11<<uint8(4) | tmp_ta11>>uint8(32-4) + tb11

	tmp_td11 := ((td11 + (ta12 ^ tb11 ^ tc11)) + (X[uint32(5+3*13)&15] + 3873151461))
	td12 := tmp_td11<<uint8(11) | tmp_td11>>uint8(32-11) + ta12

	tmp_tc11 := ((tc11 + (td12 ^ ta12 ^ tb11)) + (X[uint32(5+3*14)&15] + 530742520))
	tc12 := tmp_tc11<<uint8(16) | tmp_tc11>>uint8(32-16) + td12

	tmp_tb11 := ((tb11 + (tc12 ^ td12 ^ ta12)) + (X[uint32(5+3*15)&15] + 3299628645))
	tb12 := tmp_tb11<<uint8(23) | tmp_tb11>>uint8(32-23) + tc12

	// Round 4.

	tmp_ta12 := ((ta12 + (tc12 ^ (tb12 | ^td12))) + (X[uint32(7*0)&15] + 4096336452))
	ta13 := tmp_ta12<<uint8(6) | tmp_ta12>>uint8(32-6) + tb12

	tmp_td12 := ((td12 + (tb12 ^ (ta13 | ^tc12))) + (X[uint32(7*1)&15] + 1126891415))
	td13 := tmp_td12<<uint8(10) | tmp_td12>>uint8(32-10) + ta13

	tmp_tc12 := ((tc12 + (ta13 ^ (td13 | ^tb12))) + (X[uint32(7*2)&15] + 2878612391))
	tc13 := tmp_tc12<<uint8(15) | tmp_tc12>>uint8(32-15) + td13

	tmp_tb12 := ((tb12 + (td13 ^ (tc13 | ^ta13))) + (X[uint32(7*3)&15] + 4237533241))
	tb13 := tmp_tb12<<uint8(21) | tmp_tb12>>uint8(32-21) + tc13

	tmp_ta13 := ((ta13 + (tc13 ^ (tb13 | ^td13))) + (X[uint32(7*4)&15] + 1700485571))
	ta14 := tmp_ta13<<uint8(6) | tmp_ta13>>uint8(32-6) + tb13

	tmp_td13 := ((td13 + (tb13 ^ (ta14 | ^tc13))) + (X[uint32(7*5)&15] + 2399980690))
	td14 := tmp_td13<<uint8(10) | tmp_td13>>uint8(32-10) + ta14

	tmp_tc13 := ((tc13 + (ta14 ^ (td14 | ^tb13))) + (X[uint32(7*6)&15] + 4293915773))
	tc14 := tmp_tc13<<uint8(15) | tmp_tc13>>uint8(32-15) + td14

	tmp_tb13 := ((tb13 + (td14 ^ (tc14 | ^ta14))) + (X[uint32(7*7)&15] + 2240044497))
	tb14 := tmp_tb13<<uint8(21) | tmp_tb13>>uint8(32-21) + tc14

	tmp_ta14 := ((ta14 + (tc14 ^ (tb14 | ^td14))) + (X[uint32(7*8)&15] + 1873313359))
	ta15 := tmp_ta14<<uint8(6) | tmp_ta14>>uint8(32-6) + tb14

	tmp_td14 := ((td14 + (tb14 ^ (ta15 | ^tc14))) + (X[uint32(7*9)&15] + 4264355552))
	td15 := tmp_td14<<uint8(10) | tmp_td14>>uint8(32-10) + ta15

	tmp_tc14 := ((tc14 + (ta15 ^ (td15 | ^tb14))) + (X[uint32(7*10)&15] + 2734768916))
	tc15 := tmp_tc14<<uint8(15) | tmp_tc14>>uint8(32-15) + td15

	tmp_tb14 := ((tb14 + (td15 ^ (tc15 | ^ta15))) + (X[uint32(7*11)&15] + 1309151649))
	tb15 := tmp_tb14<<uint8(21) | tmp_tb14>>uint8(32-21) + tc15

	tmp_ta15 := ((ta15 + (tc15 ^ (tb15 | ^td15))) + (X[uint32(7*12)&15] + 4149444226))
	ta16 := tmp_ta15<<uint8(6) | tmp_ta15>>uint8(32-6) + tb15

	tmp_td15 := ((td15 + (tb15 ^ (ta16 | ^tc15))) + (X[uint32(7*13)&15] + 3174756917))
	td16 := tmp_td15<<uint8(10) | tmp_td15>>uint8(32-10) + ta16

	tmp_tc15 := ((tc15 + (ta16 ^ (td16 | ^tb15))) + (X[uint32(7*14)&15] + 718787259))
	tc16 := tmp_tc15<<uint8(15) | tmp_tc15>>uint8(32-15) + td16

	tmp_tb15 := ((tb15 + (td16 ^ (tc16 | ^ta16))) + (X[uint32(7*15)&15] + 3951481745))
	tb16 := tmp_tb15<<uint8(21) | tmp_tb15>>uint8(32-21) + tc16

	return Digest{
		s: [4]uint32{
			s[0] + ta16,
			s[1] + tb16,
			s[2] + tc16,
			s[3] + td16,
		},
		len: dig.len + (16 * 4),
	}
}

func (dig Digest) Digest() [4]uint32 {
	return dig.s
}

func (dig Digest) Sum() [Size]byte {
	var tmp [16]byte

	for i := uint(0); i < 4; i++ {
		a := dig.s[i]
		tmp[i*4] = byte(a)
		tmp[i*4+1] = byte(a >> 8)
		tmp[i*4+2] = byte(a >> 16)
		tmp[i*4+3] = byte(a >> 24)
	}
	return tmp
}
