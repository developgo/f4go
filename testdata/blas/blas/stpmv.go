package main

//*> \brief \b STPMV
//*
//*  =========== DOCUMENTATION ===========
//*
//* Online html documentation available at
//*            http://www.netlib.org/lapack/explore-html/
//*
//*  Definition:
//*  ===========
//*
//*       SUBROUTINE STPMV(UPLO,TRANS,DIAG,N,AP,X,INCX)
//*
//*       .. Scalar Arguments ..
//*       INTEGER INCX,N
//*       CHARACTER DIAG,TRANS,UPLO
//*       ..
//*       .. Array Arguments ..
//*       REAL AP(*),X(*)
//*       ..
//*
//*
//*> \par Purpose:
//*  =============
//*>
//*> \verbatim
//*>
//*> STPMV  performs one of the matrix-vector operations
//*>
//*>    x := A*x,   or   x := A**T*x,
//*>
//*> where x is an n element vector and  A is an n by n unit, or non-unit,
//*> upper or lower triangular matrix, supplied in packed form.
//*> \endverbatim
//*
//*  Arguments:
//*  ==========
//*
//*> \param[in] UPLO
//*> \verbatim
//*>          UPLO is CHARACTER*1
//*>           On entry, UPLO specifies whether the matrix is an upper or
//*>           lower triangular matrix as follows:
//*>
//*>              UPLO = 'U' or 'u'   A is an upper triangular matrix.
//*>
//*>              UPLO = 'L' or 'l'   A is a lower triangular matrix.
//*> \endverbatim
//*>
//*> \param[in] TRANS
//*> \verbatim
//*>          TRANS is CHARACTER*1
//*>           On entry, TRANS specifies the operation to be performed as
//*>           follows:
//*>
//*>              TRANS = 'N' or 'n'   x := A*x.
//*>
//*>              TRANS = 'T' or 't'   x := A**T*x.
//*>
//*>              TRANS = 'C' or 'c'   x := A**T*x.
//*> \endverbatim
//*>
//*> \param[in] DIAG
//*> \verbatim
//*>          DIAG is CHARACTER*1
//*>           On entry, DIAG specifies whether or not A is unit
//*>           triangular as follows:
//*>
//*>              DIAG = 'U' or 'u'   A is assumed to be unit triangular.
//*>
//*>              DIAG = 'N' or 'n'   A is not assumed to be unit
//*>                                  triangular.
//*> \endverbatim
//*>
//*> \param[in] N
//*> \verbatim
//*>          N is INTEGER
//*>           On entry, N specifies the order of the matrix A.
//*>           N must be at least zero.
//*> \endverbatim
//*>
//*> \param[in] AP
//*> \verbatim
//*>          AP is REAL array, dimension at least
//*>           ( ( n*( n + 1 ) )/2 ).
//*>           Before entry with  UPLO = 'U' or 'u', the array AP must
//*>           contain the upper triangular matrix packed sequentially,
//*>           column by column, so that AP( 1 ) contains a( 1, 1 ),
//*>           AP( 2 ) and AP( 3 ) contain a( 1, 2 ) and a( 2, 2 )
//*>           respectively, and so on.
//*>           Before entry with UPLO = 'L' or 'l', the array AP must
//*>           contain the lower triangular matrix packed sequentially,
//*>           column by column, so that AP( 1 ) contains a( 1, 1 ),
//*>           AP( 2 ) and AP( 3 ) contain a( 2, 1 ) and a( 3, 1 )
//*>           respectively, and so on.
//*>           Note that when  DIAG = 'U' or 'u', the diagonal elements of
//*>           A are not referenced, but are assumed to be unity.
//*> \endverbatim
//*>
//*> \param[in,out] X
//*> \verbatim
//*>          X is REAL array, dimension at least
//*>           ( 1 + ( n - 1 )*abs( INCX ) ).
//*>           Before entry, the incremented array X must contain the n
//*>           element vector x. On exit, X is overwritten with the
//*>           transformed vector x.
//*> \endverbatim
//*>
//*> \param[in] INCX
//*> \verbatim
//*>          INCX is INTEGER
//*>           On entry, INCX specifies the increment for the elements of
//*>           X. INCX must not be zero.
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
//*> \date December 2016
//*
//*> \ingroup single_blas_level2
//*
//*> \par Further Details:
//*  =====================
//*>
//*> \verbatim
//*>
//*>  Level 2 Blas routine.
//*>  The vector and matrix arguments are not referenced when N = 0, or M = 0
//*>
//*>  -- Written on 22-October-1986.
//*>     Jack Dongarra, Argonne National Lab.
//*>     Jeremy Du Croz, Nag Central Office.
//*>     Sven Hammarling, Nag Central Office.
//*>     Richard Hanson, Sandia National Labs.
//*> \endverbatim
//*>
//*  =====================================================================
func STPMV(UPLO *byte, TRANS *byte, DIAG *byte, N *int, AP *[]float64, X *[]float64, INCX *int) {
	var ZERO float64 = 0.0e+0
	var TEMP float64
	var I int
	var INFO int
	var IX int
	var J int
	var JX int
	var K int
	var KK int
	var KX int
	var NOUNIT bool
	//*
	//*  -- Reference BLAS level2 routine (version 3.7.0) --
	//*  -- Reference BLAS is a software package provided by Univ. of Tennessee,    --
	//*  -- Univ. of California Berkeley, Univ. of Colorado Denver and NAG Ltd..--
	//*     December 2016
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
	//*     .. External Functions ..
	//*     ..
	//*     .. External Subroutines ..
	//*     ..
	//*
	//*     Test the input parameters.
	//*
	INFO = 0
	if !LSAME(UPLO, func() *byte { y := byte('U'); return &y }()) && !LSAME(UPLO, func() *byte { y := byte('L'); return &y }()) {
		INFO = 1
	} else if !LSAME(TRANS, func() *byte { y := byte('N'); return &y }()) && !LSAME(TRANS, func() *byte { y := byte('T'); return &y }()) && !LSAME(TRANS, func() *byte { y := byte('C'); return &y }()) {
		INFO = 2
	} else if !LSAME(DIAG, func() *byte { y := byte('U'); return &y }()) && !LSAME(DIAG, func() *byte { y := byte('N'); return &y }()) {
		INFO = 3
	} else if (*N) < 0 {
		INFO = 4
	} else if (*INCX) == 0 {
		INFO = 7
	}
	if INFO != 0 {
		XERBLA(func() *[]byte { y := []byte("STPMV "); return &y }(), &(INFO))
		return
	}
	//*
	//*     Quick return if possible.
	//*
	if (*N) == 0 {
		return
	}
	//*
	NOUNIT = LSAME(DIAG, func() *byte { y := byte('N'); return &y }())
	//*
	//*     Set up the start point in X if the increment is not unity. This
	//*     will be  ( N - 1 )*INCX  too small for descending loops.
	//*
	if (*INCX) <= 0 {
		KX = 1 - ((*N)-1)*(*INCX)
	} else if (*INCX) != 1 {
		KX = 1
	}
	//*
	//*     Start the operations. In this version the elements of AP are
	//*     accessed sequentially with one pass through AP.
	//*
	if LSAME(TRANS, func() *byte { y := byte('N'); return &y }()) {
		//*
		//*        Form  x:= A*x.
		//*
		if LSAME(UPLO, func() *byte { y := byte('U'); return &y }()) {
			KK = 1
			if (*INCX) == 1 {
				for J = 1; J <= (*N); J++ {
					if (*X)[J-(1)] != ZERO {
						TEMP = (*X)[J-(1)]
						K = KK
						for I = 1; I <= J-1; I++ {
							(*X)[I-(1)] = (*X)[I-(1)] + TEMP*(*AP)[K-(1)]
							K = K + 1
						}
						if NOUNIT {
							(*X)[J-(1)] = (*X)[J-(1)] * (*AP)[KK+J-1-(1)]
						}
					}
					KK = KK + J
				}
			} else {
				JX = KX
				for J = 1; J <= (*N); J++ {
					if (*X)[JX-(1)] != ZERO {
						TEMP = (*X)[JX-(1)]
						IX = KX
						for K = KK; K <= KK+J-2; K++ {
							(*X)[IX-(1)] = (*X)[IX-(1)] + TEMP*(*AP)[K-(1)]
							IX = IX + (*INCX)
						}
						if NOUNIT {
							(*X)[JX-(1)] = (*X)[JX-(1)] * (*AP)[KK+J-1-(1)]
						}
					}
					JX = JX + (*INCX)
					KK = KK + J
				}
			}
		} else {
			KK = ((*N) * ((*N) + 1)) / 2
			if (*INCX) == 1 {
				for J = (*N); J <= 1; J += -1 {
					if (*X)[J-(1)] != ZERO {
						TEMP = (*X)[J-(1)]
						K = KK
						for I = (*N); I <= J+1; I += -1 {
							(*X)[I-(1)] = (*X)[I-(1)] + TEMP*(*AP)[K-(1)]
							K = K - 1
						}
						if NOUNIT {
							(*X)[J-(1)] = (*X)[J-(1)] * (*AP)[KK-(*N)+J-(1)]
						}
					}
					KK = KK - ((*N) - J + 1)
				}
			} else {
				KX = KX + ((*N)-1)*(*INCX)
				JX = KX
				for J = (*N); J <= 1; J += -1 {
					if (*X)[JX-(1)] != ZERO {
						TEMP = (*X)[JX-(1)]
						IX = KX
						for K = KK; K <= KK-((*N)-(J+1)); K += -1 {
							(*X)[IX-(1)] = (*X)[IX-(1)] + TEMP*(*AP)[K-(1)]
							IX = IX - (*INCX)
						}
						if NOUNIT {
							(*X)[JX-(1)] = (*X)[JX-(1)] * (*AP)[KK-(*N)+J-(1)]
						}
					}
					JX = JX - (*INCX)
					KK = KK - ((*N) - J + 1)
				}
			}
		}
	} else {
		//*
		//*        Form  x := A**T*x.
		//*
		if LSAME(UPLO, func() *byte { y := byte('U'); return &y }()) {
			KK = ((*N) * ((*N) + 1)) / 2
			if (*INCX) == 1 {
				for J = (*N); J <= 1; J += -1 {
					TEMP = (*X)[J-(1)]
					if NOUNIT {
						TEMP = TEMP * (*AP)[KK-(1)]
					}
					K = KK - 1
					for I = J - 1; I <= 1; I += -1 {
						TEMP = TEMP + (*AP)[K-(1)]*(*X)[I-(1)]
						K = K - 1
					}
					(*X)[J-(1)] = TEMP
					KK = KK - J
				}
			} else {
				JX = KX + ((*N)-1)*(*INCX)
				for J = (*N); J <= 1; J += -1 {
					TEMP = (*X)[JX-(1)]
					IX = JX
					if NOUNIT {
						TEMP = TEMP * (*AP)[KK-(1)]
					}
					for K = KK - 1; K <= KK-J+1; K += -1 {
						IX = IX - (*INCX)
						TEMP = TEMP + (*AP)[K-(1)]*(*X)[IX-(1)]
					}
					(*X)[JX-(1)] = TEMP
					JX = JX - (*INCX)
					KK = KK - J
				}
			}
		} else {
			KK = 1
			if (*INCX) == 1 {
				for J = 1; J <= (*N); J++ {
					TEMP = (*X)[J-(1)]
					if NOUNIT {
						TEMP = TEMP * (*AP)[KK-(1)]
					}
					K = KK + 1
					for I = J + 1; I <= (*N); I++ {
						TEMP = TEMP + (*AP)[K-(1)]*(*X)[I-(1)]
						K = K + 1
					}
					(*X)[J-(1)] = TEMP
					KK = KK + ((*N) - J + 1)
				}
			} else {
				JX = KX
				for J = 1; J <= (*N); J++ {
					TEMP = (*X)[JX-(1)]
					IX = JX
					if NOUNIT {
						TEMP = TEMP * (*AP)[KK-(1)]
					}
					for K = KK + 1; K <= KK+(*N)-J; K++ {
						IX = IX + (*INCX)
						TEMP = TEMP + (*AP)[K-(1)]*(*X)[IX-(1)]
					}
					(*X)[JX-(1)] = TEMP
					JX = JX + (*INCX)
					KK = KK + ((*N) - J + 1)
				}
			}
		}
	}
	//*
	return
	//*
	//*     End of STPMV .
	//*
}
