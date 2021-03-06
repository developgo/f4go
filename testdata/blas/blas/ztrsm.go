package main

import "github.com/Konstantin8105/f4go/intrinsic"

//*> \brief \b ZTRSM
//*
//*  =========== DOCUMENTATION ===========
//*
//* Online html documentation available at
//*            http://www.netlib.org/lapack/explore-html/
//*
//*  Definition:
//*  ===========
//*
//*       SUBROUTINE ZTRSM(SIDE,UPLO,TRANSA,DIAG,M,N,ALPHA,A,LDA,B,LDB)
//*
//*       .. Scalar Arguments ..
//*       COMPLEX*16 ALPHA
//*       INTEGER LDA,LDB,M,N
//*       CHARACTER DIAG,SIDE,TRANSA,UPLO
//*       ..
//*       .. Array Arguments ..
//*       COMPLEX*16 A(LDA,*),B(LDB,*)
//*       ..
//*
//*
//*> \par Purpose:
//*  =============
//*>
//*> \verbatim
//*>
//*> ZTRSM  solves one of the matrix equations
//*>
//*>    op( A )*X = alpha*B,   or   X*op( A ) = alpha*B,
//*>
//*> where alpha is a scalar, X and B are m by n matrices, A is a unit, or
//*> non-unit,  upper or lower triangular matrix  and  op( A )  is one  of
//*>
//*>    op( A ) = A   or   op( A ) = A**T   or   op( A ) = A**H.
//*>
//*> The matrix X is overwritten on B.
//*> \endverbatim
//*
//*  Arguments:
//*  ==========
//*
//*> \param[in] SIDE
//*> \verbatim
//*>          SIDE is CHARACTER*1
//*>           On entry, SIDE specifies whether op( A ) appears on the left
//*>           or right of X as follows:
//*>
//*>              SIDE = 'L' or 'l'   op( A )*X = alpha*B.
//*>
//*>              SIDE = 'R' or 'r'   X*op( A ) = alpha*B.
//*> \endverbatim
//*>
//*> \param[in] UPLO
//*> \verbatim
//*>          UPLO is CHARACTER*1
//*>           On entry, UPLO specifies whether the matrix A is an upper or
//*>           lower triangular matrix as follows:
//*>
//*>              UPLO = 'U' or 'u'   A is an upper triangular matrix.
//*>
//*>              UPLO = 'L' or 'l'   A is a lower triangular matrix.
//*> \endverbatim
//*>
//*> \param[in] TRANSA
//*> \verbatim
//*>          TRANSA is CHARACTER*1
//*>           On entry, TRANSA specifies the form of op( A ) to be used in
//*>           the matrix multiplication as follows:
//*>
//*>              TRANSA = 'N' or 'n'   op( A ) = A.
//*>
//*>              TRANSA = 'T' or 't'   op( A ) = A**T.
//*>
//*>              TRANSA = 'C' or 'c'   op( A ) = A**H.
//*> \endverbatim
//*>
//*> \param[in] DIAG
//*> \verbatim
//*>          DIAG is CHARACTER*1
//*>           On entry, DIAG specifies whether or not A is unit triangular
//*>           as follows:
//*>
//*>              DIAG = 'U' or 'u'   A is assumed to be unit triangular.
//*>
//*>              DIAG = 'N' or 'n'   A is not assumed to be unit
//*>                                  triangular.
//*> \endverbatim
//*>
//*> \param[in] M
//*> \verbatim
//*>          M is INTEGER
//*>           On entry, M specifies the number of rows of B. M must be at
//*>           least zero.
//*> \endverbatim
//*>
//*> \param[in] N
//*> \verbatim
//*>          N is INTEGER
//*>           On entry, N specifies the number of columns of B.  N must be
//*>           at least zero.
//*> \endverbatim
//*>
//*> \param[in] ALPHA
//*> \verbatim
//*>          ALPHA is COMPLEX*16
//*>           On entry,  ALPHA specifies the scalar  alpha. When  alpha is
//*>           zero then  A is not referenced and  B need not be set before
//*>           entry.
//*> \endverbatim
//*>
//*> \param[in] A
//*> \verbatim
//*>          A is COMPLEX*16 array, dimension ( LDA, k ),
//*>           where k is m when SIDE = 'L' or 'l'
//*>             and k is n when SIDE = 'R' or 'r'.
//*>           Before entry  with  UPLO = 'U' or 'u',  the  leading  k by k
//*>           upper triangular part of the array  A must contain the upper
//*>           triangular matrix  and the strictly lower triangular part of
//*>           A is not referenced.
//*>           Before entry  with  UPLO = 'L' or 'l',  the  leading  k by k
//*>           lower triangular part of the array  A must contain the lower
//*>           triangular matrix  and the strictly upper triangular part of
//*>           A is not referenced.
//*>           Note that when  DIAG = 'U' or 'u',  the diagonal elements of
//*>           A  are not referenced either,  but are assumed to be  unity.
//*> \endverbatim
//*>
//*> \param[in] LDA
//*> \verbatim
//*>          LDA is INTEGER
//*>           On entry, LDA specifies the first dimension of A as declared
//*>           in the calling (sub) program.  When  SIDE = 'L' or 'l'  then
//*>           LDA  must be at least  max( 1, m ),  when  SIDE = 'R' or 'r'
//*>           then LDA must be at least max( 1, n ).
//*> \endverbatim
//*>
//*> \param[in,out] B
//*> \verbatim
//*>          B is COMPLEX*16 array, dimension ( LDB, N )
//*>           Before entry,  the leading  m by n part of the array  B must
//*>           contain  the  right-hand  side  matrix  B,  and  on exit  is
//*>           overwritten by the solution matrix  X.
//*> \endverbatim
//*>
//*> \param[in] LDB
//*> \verbatim
//*>          LDB is INTEGER
//*>           On entry, LDB specifies the first dimension of B as declared
//*>           in  the  calling  (sub)  program.   LDB  must  be  at  least
//*>           max( 1, m ).
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
//*> \ingroup complex16_blas_level3
//*
//*> \par Further Details:
//*  =====================
//*>
//*> \verbatim
//*>
//*>  Level 3 Blas routine.
//*>
//*>  -- Written on 8-February-1989.
//*>     Jack Dongarra, Argonne National Laboratory.
//*>     Iain Duff, AERE Harwell.
//*>     Jeremy Du Croz, Numerical Algorithms Group Ltd.
//*>     Sven Hammarling, Numerical Algorithms Group Ltd.
//*> \endverbatim
//*>
//*  =====================================================================
func ZTRSM(SIDE *byte, UPLO *byte, TRANSA *byte, DIAG *byte, M *int, N *int, ALPHA *complex128, A *[][]complex128, LDA *int, B *[][]complex128, LDB *int) {
	var TEMP complex128
	var I int
	var INFO int
	var J int
	var K int
	var NROWA int
	var LSIDE bool
	var NOCONJ bool
	var NOUNIT bool
	var UPPER bool
	var ONE complex128 = (1.0e+0 + (0.0e+0)*1i)
	var ZERO complex128 = (0.0e+0 + (0.0e+0)*1i)
	//*
	//*  -- Reference BLAS level3 routine (version 3.7.0) --
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
	//*     .. External Functions ..
	//*     ..
	//*     .. External Subroutines ..
	//*     ..
	//*     .. Intrinsic Functions ..
	//*     ..
	//*     .. Local Scalars ..
	//*     ..
	//*     .. Parameters ..
	//*     ..
	//*
	//*     Test the input parameters.
	//*
	LSIDE = LSAME(SIDE, func() *byte { y := byte('L'); return &y }())
	if LSIDE {
		NROWA = (*M)
	} else {
		NROWA = (*N)
	}
	NOCONJ = LSAME(TRANSA, func() *byte { y := byte('T'); return &y }())
	NOUNIT = LSAME(DIAG, func() *byte { y := byte('N'); return &y }())
	UPPER = LSAME(UPLO, func() *byte { y := byte('U'); return &y }())
	//*
	INFO = 0
	if (!LSIDE) && (!LSAME(SIDE, func() *byte { y := byte('R'); return &y }())) {
		INFO = 1
	} else if (!UPPER) && (!LSAME(UPLO, func() *byte { y := byte('L'); return &y }())) {
		INFO = 2
	} else if (!LSAME(TRANSA, func() *byte { y := byte('N'); return &y }())) && (!LSAME(TRANSA, func() *byte { y := byte('T'); return &y }())) && (!LSAME(TRANSA, func() *byte { y := byte('C'); return &y }())) {
		INFO = 3
	} else if (!LSAME(DIAG, func() *byte { y := byte('U'); return &y }())) && (!LSAME(DIAG, func() *byte { y := byte('N'); return &y }())) {
		INFO = 4
	} else if (*M) < 0 {
		INFO = 5
	} else if (*N) < 0 {
		INFO = 6
	} else if (*LDA) < intrinsic.MAX(int(1), NROWA) {
		INFO = 9
	} else if (*LDB) < intrinsic.MAX(int(1), (*M)) {
		INFO = 11
	}
	if INFO != 0 {
		XERBLA(func() *[]byte { y := []byte("ZTRSM "); return &y }(), &(INFO))
		return
	}
	//*
	//*     Quick return if possible.
	//*
	if (*M) == 0 || (*N) == 0 {
		return
	}
	//*
	//*     And when  alpha.eq.zero.
	//*
	if (*ALPHA) == ZERO {
		for J = 1; J <= (*N); J++ {
			for I = 1; I <= (*M); I++ {
				(*B)[I-(1)][J-(1)] = ZERO
			}
		}
		return
	}
	//*
	//*     Start the operations.
	//*
	if LSIDE {
		if LSAME(TRANSA, func() *byte { y := byte('N'); return &y }()) {
			//*
			//*           Form  B := alpha*inv( A )*B.
			//*
			if UPPER {
				for J = 1; J <= (*N); J++ {
					if (*ALPHA) != ONE {
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][J-(1)] = (*ALPHA) * (*B)[I-(1)][J-(1)]
						}
					}
					for K = (*M); K <= 1; K += -1 {
						if (*B)[K-(1)][J-(1)] != ZERO {
							if NOUNIT {
								(*B)[K-(1)][J-(1)] = (*B)[K-(1)][J-(1)] / (*A)[K-(1)][K-(1)]
							}
							for I = 1; I <= K-1; I++ {
								(*B)[I-(1)][J-(1)] = (*B)[I-(1)][J-(1)] - (*B)[K-(1)][J-(1)]*(*A)[I-(1)][K-(1)]
							}
						}
					}
				}
			} else {
				for J = 1; J <= (*N); J++ {
					if (*ALPHA) != ONE {
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][J-(1)] = (*ALPHA) * (*B)[I-(1)][J-(1)]
						}
					}
					for K = 1; K <= (*M); K++ {
						if (*B)[K-(1)][J-(1)] != ZERO {
							if NOUNIT {
								(*B)[K-(1)][J-(1)] = (*B)[K-(1)][J-(1)] / (*A)[K-(1)][K-(1)]
							}
							for I = K + 1; I <= (*M); I++ {
								(*B)[I-(1)][J-(1)] = (*B)[I-(1)][J-(1)] - (*B)[K-(1)][J-(1)]*(*A)[I-(1)][K-(1)]
							}
						}
					}
				}
			}
		} else {
			//*
			//*           Form  B := alpha*inv( A**T )*B
			//*           or    B := alpha*inv( A**H )*B.
			//*
			if UPPER {
				for J = 1; J <= (*N); J++ {
					for I = 1; I <= (*M); I++ {
						TEMP = (*ALPHA) * (*B)[I-(1)][J-(1)]
						if NOCONJ {
							for K = 1; K <= I-1; K++ {
								TEMP = TEMP - (*A)[K-(1)][I-(1)]*(*B)[K-(1)][J-(1)]
							}
							if NOUNIT {
								TEMP = TEMP / (*A)[I-(1)][I-(1)]
							}
						} else {
							for K = 1; K <= I-1; K++ {
								TEMP = TEMP - intrinsic.DCONJG((*A)[K-(1)][I-(1)])*(*B)[K-(1)][J-(1)]
							}
							if NOUNIT {
								TEMP = TEMP / intrinsic.DCONJG((*A)[I-(1)][I-(1)])
							}
						}
						(*B)[I-(1)][J-(1)] = TEMP
					}
				}
			} else {
				for J = 1; J <= (*N); J++ {
					for I = (*M); I <= 1; I += -1 {
						TEMP = (*ALPHA) * (*B)[I-(1)][J-(1)]
						if NOCONJ {
							for K = I + 1; K <= (*M); K++ {
								TEMP = TEMP - (*A)[K-(1)][I-(1)]*(*B)[K-(1)][J-(1)]
							}
							if NOUNIT {
								TEMP = TEMP / (*A)[I-(1)][I-(1)]
							}
						} else {
							for K = I + 1; K <= (*M); K++ {
								TEMP = TEMP - intrinsic.DCONJG((*A)[K-(1)][I-(1)])*(*B)[K-(1)][J-(1)]
							}
							if NOUNIT {
								TEMP = TEMP / intrinsic.DCONJG((*A)[I-(1)][I-(1)])
							}
						}
						(*B)[I-(1)][J-(1)] = TEMP
					}
				}
			}
		}
	} else {
		if LSAME(TRANSA, func() *byte { y := byte('N'); return &y }()) {
			//*
			//*           Form  B := alpha*B*inv( A ).
			//*
			if UPPER {
				for J = 1; J <= (*N); J++ {
					if (*ALPHA) != ONE {
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][J-(1)] = (*ALPHA) * (*B)[I-(1)][J-(1)]
						}
					}
					for K = 1; K <= J-1; K++ {
						if (*A)[K-(1)][J-(1)] != ZERO {
							for I = 1; I <= (*M); I++ {
								(*B)[I-(1)][J-(1)] = (*B)[I-(1)][J-(1)] - (*A)[K-(1)][J-(1)]*(*B)[I-(1)][K-(1)]
							}
						}
					}
					if NOUNIT {
						TEMP = ONE / (*A)[J-(1)][J-(1)]
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][J-(1)] = TEMP * (*B)[I-(1)][J-(1)]
						}
					}
				}
			} else {
				for J = (*N); J <= 1; J += -1 {
					if (*ALPHA) != ONE {
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][J-(1)] = (*ALPHA) * (*B)[I-(1)][J-(1)]
						}
					}
					for K = J + 1; K <= (*N); K++ {
						if (*A)[K-(1)][J-(1)] != ZERO {
							for I = 1; I <= (*M); I++ {
								(*B)[I-(1)][J-(1)] = (*B)[I-(1)][J-(1)] - (*A)[K-(1)][J-(1)]*(*B)[I-(1)][K-(1)]
							}
						}
					}
					if NOUNIT {
						TEMP = ONE / (*A)[J-(1)][J-(1)]
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][J-(1)] = TEMP * (*B)[I-(1)][J-(1)]
						}
					}
				}
			}
		} else {
			//*
			//*           Form  B := alpha*B*inv( A**T )
			//*           or    B := alpha*B*inv( A**H ).
			//*
			if UPPER {
				for K = (*N); K <= 1; K += -1 {
					if NOUNIT {
						if NOCONJ {
							TEMP = ONE / (*A)[K-(1)][K-(1)]
						} else {
							TEMP = ONE / intrinsic.DCONJG((*A)[K-(1)][K-(1)])
						}
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][K-(1)] = TEMP * (*B)[I-(1)][K-(1)]
						}
					}
					for J = 1; J <= K-1; J++ {
						if (*A)[J-(1)][K-(1)] != ZERO {
							if NOCONJ {
								TEMP = (*A)[J-(1)][K-(1)]
							} else {
								TEMP = intrinsic.DCONJG((*A)[J-(1)][K-(1)])
							}
							for I = 1; I <= (*M); I++ {
								(*B)[I-(1)][J-(1)] = (*B)[I-(1)][J-(1)] - TEMP*(*B)[I-(1)][K-(1)]
							}
						}
					}
					if (*ALPHA) != ONE {
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][K-(1)] = (*ALPHA) * (*B)[I-(1)][K-(1)]
						}
					}
				}
			} else {
				for K = 1; K <= (*N); K++ {
					if NOUNIT {
						if NOCONJ {
							TEMP = ONE / (*A)[K-(1)][K-(1)]
						} else {
							TEMP = ONE / intrinsic.DCONJG((*A)[K-(1)][K-(1)])
						}
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][K-(1)] = TEMP * (*B)[I-(1)][K-(1)]
						}
					}
					for J = K + 1; J <= (*N); J++ {
						if (*A)[J-(1)][K-(1)] != ZERO {
							if NOCONJ {
								TEMP = (*A)[J-(1)][K-(1)]
							} else {
								TEMP = intrinsic.DCONJG((*A)[J-(1)][K-(1)])
							}
							for I = 1; I <= (*M); I++ {
								(*B)[I-(1)][J-(1)] = (*B)[I-(1)][J-(1)] - TEMP*(*B)[I-(1)][K-(1)]
							}
						}
					}
					if (*ALPHA) != ONE {
						for I = 1; I <= (*M); I++ {
							(*B)[I-(1)][K-(1)] = (*ALPHA) * (*B)[I-(1)][K-(1)]
						}
					}
				}
			}
		}
	}
	//*
	return
	//*
	//*     End of ZTRSM .
	//*
}
