package main

import "math"
import "github.com/Konstantin8105/f4go/intrinsic"

//*> \brief \b SNRM2
//*
//*  =========== DOCUMENTATION ===========
//*
//* Online html documentation available at
//*            http://www.netlib.org/lapack/explore-html/
//*
//*  Definition:
//*  ===========
//*
//*       REAL FUNCTION SNRM2(N,X,INCX)
//*
//*       .. Scalar Arguments ..
//*       INTEGER INCX,N
//*       ..
//*       .. Array Arguments ..
//*       REAL X(*)
//*       ..
//*
//*
//*> \par Purpose:
//*  =============
//*>
//*> \verbatim
//*>
//*> SNRM2 returns the euclidean norm of a vector via the function
//*> name, so that
//*>
//*>    SNRM2 := sqrt( x'*x ).
//*> \endverbatim
//*
//*  Arguments:
//*  ==========
//*
//*> \param[in] N
//*> \verbatim
//*>          N is INTEGER
//*>         number of elements in input vector(s)
//*> \endverbatim
//*>
//*> \param[in] X
//*> \verbatim
//*>          X is REAL array, dimension ( 1 + ( N - 1 )*abs( INCX ) )
//*> \endverbatim
//*>
//*> \param[in] INCX
//*> \verbatim
//*>          INCX is INTEGER
//*>         storage spacing between elements of SX
//*> \endverbatim
//*
//*  Authors:
//*  ========
//*
//*> \author Univ. of Tennessee
//*> \author Univ. of California Berkeley
//*> \author Univ. of Colorado Denver
//*> \author NAG Ltd.
//*
//*> \date November 2017
//*
//*> \ingroup single_blas_level1
//*
//*> \par Further Details:
//*  =====================
//*>
//*> \verbatim
//*>
//*>  -- This version written on 25-October-1982.
//*>     Modified on 14-October-1993 to inline the call to SLASSQ.
//*>     Sven Hammarling, Nag Ltd.
//*> \endverbatim
//*>
//*  =====================================================================
func SNRM2(N *int, X *[]float64, INCX *int) (SNRM2_RES float64) {
	var ONE float64 = 1.0e+0
	var ZERO float64 = 0.0e+0
	var ABSXI float64
	var NORM float64
	var SCALE float64
	var SSQ float64
	var IX int
	//*
	//*  -- Reference BLAS level1 routine (version 3.8.0) --
	//*  -- Reference BLAS is a software package provided by Univ. of Tennessee,    --
	//*  -- Univ. of California Berkeley, Univ. of Colorado Denver and NAG Ltd..--
	//*     November 2017
	//*
	//*     .. Scalar Arguments ..
	//*     ..
	//*     .. Array Arguments ..
	//*     ..
	//*
	//*  =====================================================================
	//*
	//*     .. Parameters ..
	//*     ..
	//*     .. Local Scalars ..
	//*     ..
	//*     .. Intrinsic Functions ..
	//*     ..
	if (*N) < 1 || (*INCX) < 1 {
		NORM = ZERO
	} else if (*N) == 1 {
		NORM = intrinsic.ABS((*X)[1-(1)])
	} else {
		SCALE = ZERO
		SSQ = ONE
		//*        The following loop is equivalent to this call to the LAPACK
		//*        auxiliary routine:
		//*        CALL SLASSQ( N, X, INCX, SCALE, SSQ )
		//*
		for IX = 1; IX <= 1+((*N)-1)*(*INCX); IX += (*INCX) {
			if (*X)[IX-(1)] != ZERO {
				ABSXI = intrinsic.ABS((*X)[IX-(1)])
				if SCALE < ABSXI {
					SSQ = ONE + SSQ*math.Pow((SCALE/ABSXI), 2)
					SCALE = ABSXI
				} else {
					SSQ = SSQ + math.Pow((ABSXI/SCALE), 2)
				}
			}
		}
		NORM = SCALE * SQRT(&(SSQ))
	}
	//*
	(SNRM2_RES) = NORM
	return
	//*
	//*     End of SNRM2.
	//*
}
