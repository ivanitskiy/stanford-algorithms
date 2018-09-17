"""
In this programming assignment you will implement one or more of the integer multiplication algorithms described in lecture.

To get the most out of this assignment, your program should restrict itself to multiplying only pairs of single-digit numbers. You can implement the grade-school algorithm if you want, but to get the most out of the assignment you'll want to implement recursive integer multiplication and/or Karatsuba's algorithm.

So: what's the product of the following two 64-digit numbers?

3141592653589793238462643383279502884197169399375105820974944592

2718281828459045235360287471352662497757247093699959574966967627

[TIP: before submitting, first test the correctness of your program on some small test cases of your own devising. Then post your best test cases to the discussion forums to help your fellow students!]

[Food for thought: the number of digits in each input number is a power of 2. Does this make your life easier? Does it depend on which algorithm you're implementing?]

TODO: add unit-tests
Clean up code
"""

import sys

def get_parts(x,y):
    l = len(x)
    k = len(y)
    retval = (x[:l//2], x[l//2 + l % 2:], y[:k//2], y[k//2 + k% 2:])
    print   "|%s|%s|\n|%s|%s|" % retval
    return retval

def reverse(string): 
    string = "".join(reversed(string)) 
    return string 

def a_plus_b(a,b):
    """bonus implementation.
    instead of using:
        int(ac)*( 10 ** lenx) + (int(bc) + int(ad))*( 10**(lenx//2)) + int(bd) 
    we can implement using summ of string with shifting 0 in the end:
    a_plus_b(
        a_plus_b(
            ac + "0"*(n + m),
            ad + "0"*n
        ),
        a_plus_b(
                bc + "0"*m,
                bd
        )
    )
    """
    maxlen = []
    minlen = []
    if len(a) >= len(b):
        maxlen = list(a)
        minlen = list(b)
    else:
        maxlen = list(b)
        minlen = list(a)
    maxlen.reverse()
    minlen.reverse()
    index = 0
    shift = 0
    retval = []
    while (index < len(maxlen)):
        x = int(maxlen[index])
        y = 0
        if index < len(minlen):
            y = int(minlen[index])
        sumxy = x + y + shift
        shift = sumxy // 10
        retval.insert(0, sumxy % 10)
        index += 1

    if (shift):
        sumxy = retval[0] + shift
        shift = sumxy // 10
        if shift:
            retval[0] = sumxy % 10
            retval.insert(0, sumxy % 10)
        else:
            retval[0] = sumxy
    retval = "".join([str(x) for x in retval])
    # strretval =  "{:} + {:} = {:}".format(a,b,retval)
    # print "{:>80}".format(strretval)
    return retval

def multiply(x, y):
    """
    @param  x: string
    @param  y: string
    """
    
    lenx = len(str(x))
    leny = len(str(y))
    if (lenx <= 2) or (leny <= 2):
        retval = int(x) * int(y)
        # print "%s*%s=%s" % (x,y, retval),
        # print 
        return retval
    a, b,c,d = get_parts(x,y)
    n = len(a)
    m = len(c)

    print ""
    ac = str(multiply(a,c))
    bd = str(multiply(b,d))
    bc = str(multiply(b,c))
    ad = str(multiply(a,d))

    retval = int(ac)*( 10 ** lenx) + (int(bc) + int(ad))*( 10**(lenx//2)) + int(bd) 

    print "all together:", a_plus_b(
            a_plus_b(
                ac + "0"*(n + m),
                ad + "0"*n),
            a_plus_b(
                bc + "0"*m,
                bd
            )
    )
    bc + "0"*m
    
    # print "*" * (n+n+m+m)
    # strFormatter = "{:>%s}" % (n+n+m+m)
    # print strFormatter.format(ac + "0"*(n + m))
    # print strFormatter.format(ad + "0"*n)
    # print strFormatter.format(bc + "0"*m)
    # print strFormatter.format(bd)
    # print "=" * (n+n+m+m)
    # print strFormatter.format(retval)
    print "retval:", retval
    return retval


if (__name__ == "__main__"):
    s1 = "3141592653589793238462643383279502884197169399375105820974944592"
    s2 = "2718281828459045235360287471352662497757247093699959574966967627"
    x = "5678"
    y = "1234"
    print x,"*",y
    multiply(x,y)
    
    x = "56785678"
    y = "12341234"
    print x,"*",y
    multiply(x,y)
    s1 = "3141592653589793238462643383279502884197169399375105820974944592"
    s2 = "2718281828459045235360287471352662497757247093699959574966967627"
    multiply(s1,s2)
    # a_plus_b("6720000", "2652")
    # a_plus_b("6722652", "284000")

    s1 = "3141592653589793238462643383279502884197169399375105820974944592314159265358979323846264338327950288419716939937510582097494459231415926535897932384626433832795028841971693993751058209749445923141592653589793238462643383279502884197169399375105820974944592"
    s2 = "31415926535897932384626433832795028841971693993751058209749445923141592653589793238462643383279502884197169399375105820974944592314159265358979323846264338327950288419716939937510582097494459231415926535897932384626433832795028841971693993751058209749445923141592653589793238462643383279502884197169399375105820974944592314159265358979323846264338327950288419716939937510582097494459231415926535897932384626433832795028841971693993751058209749445923141592653589793238462643383279502884197169399375105820974944592"
    multiply(s1,s2)
    
    64500313688838271489248293066764401740653362011787585375758514016111089716230965470280743257222857375869844910108853492895412404442187812448709336903988808039473509785679021632917170834329391086701679528811186242577323114266536226265990002244217096
