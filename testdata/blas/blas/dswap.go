package main

import "github.com/Konstantin8105/f4go/intrinsic"

//*> \brief \b DSWAP
//*
//*  =========== DOCUMENTATION ===========
//*
//* Online html documentation available at
//*            http://www.netlib.org/lapack/explore-html/
//*
//*  Definition:
//*  ===========
//*
//*       SUBROUTINE DSWAP(N,DX,INCX,DY,INCY)
//*
//*       .. Scalar Arguments ..
//*       INTEGER INCX,INCY,N
//*       ..
//*       .. Array Arguments ..
//*       DOUBLE PRECISION DX(*),DY(*)
//*       ..
//*
//*
//*> \par Purpose:
//*  =============
//*>
//*> \verbatim
//*>
//*>    DSWAP interchanges two vectors.
//*>    uses unrolled loops for increments equal to 1.
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
//*> \param[in,out] DX
//*> \verbatim
//*>          DX is DOUBLE PRECISION array, dimension ( 1 + ( N - 1 )*abs( INCX ) )
//*> \endverbatim
//*>
//*> \param[in] INCX
//*> \verbatim
//*>          INCX is INTEGER
//*>         storage spacing between elements of DX
//*> \endverbatim
//*>
//*> \param[in,out] DY
//*> \verbatim
//*>          DY is DOUBLE PRECISION array, dimension ( 1 + ( N - 1 )*abs( INCY ) )
//*> \endverbatim
//*>
//*> \param[in] INCY
//*> \verbatim
//*>          INCY is INTEGER
//*>         storage spacing between elements of DY
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
//*> \ingroup double_blas_level1
//*
//*> \par Further Details:
//*  =====================
//*>
//*> \verbatim
//*>
//*>     jack dongarra, linpack, 3/11/78.
//*>     modified 12/3/93, array(1) declarations changed to array(*)
//*> \endverbatim
//*>
//*  =====================================================================
func DSWAP(N *int, DX *[]float64, INCX *int, DY *[]float64, INCY *int) {
	var DTEMP float64
	var I int
	var IX int
	var IY int
	var M int
	var MP1 int
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
	//*     .. Local Scalars ..
	//*     ..
	//*     .. Intrinsic Functions ..
	//*     ..
	if (*N) <= 0 {
		return
	}
	if (*INCX) == 1 && (*INCY) == 1 {
		//*
		//*       code for both increments equal to 1
		//*
		//*
		//*       clean-up loop
		//*
		M = intrinsic.MOD((*N), int(3))
		if M != 0 {
			for I = 1; I <= M; I++ {
				DTEMP = (*DX)[I-(1)]
				(*DX)[I-(1)] = (*DY)[I-(1)]
				(*DY)[I-(1)] = DTEMP
			}
			if (*N) < 3 {
				return
			}
		}
		MP1 = M + 1
		for I = MP1; I <= (*N); I += 3 {
			DTEMP = (*DX)[I-(1)]
			(*DX)[I-(1)] = (*DY)[I-(1)]
			(*DY)[I-(1)] = DTEMP
			DTEMP = (*DX)[I+1-(1)]
			(*DX)[I+1-(1)] = (*DY)[I+1-(1)]
			(*DY)[I+1-(1)] = DTEMP
			DTEMP = (*DX)[I+2-(1)]
			(*DX)[I+2-(1)] = (*DY)[I+2-(1)]
			(*DY)[I+2-(1)] = DTEMP
		}
	} else {
		//*
		//*       code for unequal increments or equal increments not equal
		//*         to 1
		//*
		IX = 1
		IY = 1
		if (*INCX) < 0 {
			IX = (-(*N)+1)*(*INCX) + 1
		}
		if (*INCY) < 0 {
			IY = (-(*N)+1)*(*INCY) + 1
		}
		for I = 1; I <= (*N); I++ {
			DTEMP = (*DX)[IX-(1)]
			(*DX)[IX-(1)] = (*DY)[IY-(1)]
			(*DY)[IY-(1)] = DTEMP
			IX = IX + (*INCX)
			IY = IY + (*INCY)
		}
	}
	return
}
