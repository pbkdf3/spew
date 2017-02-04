create random lines of any length, printable ascii
useful for testing things
decently fast, hundreds of MBytes/sec on modern hardware

$ ./spew
fr9.&%dEQhm=ubf3Rt|"K?@PH\6\URKa$ ./spew -h

Usage: spew [LENGTH [LINES]]

generate random strings, one per line

Arguments:
  LENGTH=32    length of generated string
  LINES=0      number of lines of output (0 will output a string without a newline)

$ ./spew 32 1
H 4kvSFY)\;aYc?>:-`DbdZL{>S:|p<Q

$ ./spew 64 5
^UxGDX$+9OXwJE-EAX,]~0@aj*R;om7MUG7)xrriO!~_3$R)XnY~T%I>h:s@n~vS
d_y2Sxb<{N%B-mb'|^u<b\'y;lmRm+c1%n4bmlQ^bo'w]WPO-1i3u]cs<7^#s(\5
-\b$sW\%B%1!TSo>re_{Bst!s:nd\}~JwN,?e\\!K\AK70l5EvU~#$.n?6SI.tY{
J<)g{m,jj<cw>MzkE]G@H-K)Y}#VqF;oA0by>ZOQy64]\qF+4w7\$8tcZ0=l"F31
E@\A\in\0QA_c8jc\.6qZmERuKs;&VYya=k*{QWx=B v}#P"56(\/<+W,He/b\53

