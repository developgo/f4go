(string) (len=7947) "COMMENT\t> \\brief \\b SROTM\nCOMMENT\t*\nCOMMENT\t*  =========== DOCUMENTATION ===========\nCOMMENT\t*\nCOMMENT\t* Online html documentation available at\nCOMMENT\t*            http://www.netlib.org/lapack/explore-html/\nCOMMENT\t*\nCOMMENT\t*  Definition:\nCOMMENT\t*  ===========\nCOMMENT\t*\nCOMMENT\t*       SUBROUTINE SROTM(N,SX,INCX,SY,INCY,SPARAM)\nCOMMENT\t*\nCOMMENT\t*       .. Scalar Arguments ..\nCOMMENT\t*       INTEGER INCX,INCY,N\nCOMMENT\t*       ..\nCOMMENT\t*       .. Array Arguments ..\nCOMMENT\t*       REAL SPARAM(5),SX(*),SY(*)\nCOMMENT\t*       ..\nCOMMENT\t*\nCOMMENT\t*\nCOMMENT\t*> \\par Purpose:\nCOMMENT\t*  =============\nCOMMENT\t*>\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>\nCOMMENT\t*>    APPLY THE MODIFIED GIVENS TRANSFORMATION, H, TO THE 2 BY N MATRIX\nCOMMENT\t*>\nCOMMENT\t*>    (SX**T) , WHERE **T INDICATES TRANSPOSE. THE ELEMENTS OF SX ARE IN\nCOMMENT\t*>    (SX**T)\nCOMMENT\t*>\nCOMMENT\t*>    SX(LX+I*INCX), I = 0 TO N-1, WHERE LX = 1 IF INCX .GE. 0, ELSE\nCOMMENT\t*>    LX = (-INCX)*N, AND SIMILARLY FOR SY USING USING LY AND INCY.\nCOMMENT\t*>    WITH SPARAM(1)=SFLAG, H HAS ONE OF THE FOLLOWING FORMS..\nCOMMENT\t*>\nCOMMENT\t*>    SFLAG=-1.E0     SFLAG=0.E0        SFLAG=1.E0     SFLAG=-2.E0\nCOMMENT\t*>\nCOMMENT\t*>      (SH11  SH12)    (1.E0  SH12)    (SH11  1.E0)    (1.E0  0.E0)\nCOMMENT\t*>    H=(          )    (          )    (          )    (          )\nCOMMENT\t*>      (SH21  SH22),   (SH21  1.E0),   (-1.E0 SH22),   (0.E0  1.E0).\nCOMMENT\t*>    SEE  SROTMG FOR A DESCRIPTION OF DATA STORAGE IN SPARAM.\nCOMMENT\t*>\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*\nCOMMENT\t*  Arguments:\nCOMMENT\t*  ==========\nCOMMENT\t*\nCOMMENT\t*> \\param[in] N\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          N is INTEGER\nCOMMENT\t*>         number of elements in input vector(s)\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in,out] SX\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          SX is REAL array, dimension ( 1 + ( N - 1 )*abs( INCX ) )\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in] INCX\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          INCX is INTEGER\nCOMMENT\t*>         storage spacing between elements of SX\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in,out] SY\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          SY is REAL array, dimension ( 1 + ( N - 1 )*abs( INCY ) )\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in] INCY\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          INCY is INTEGER\nCOMMENT\t*>         storage spacing between elements of SY\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in] SPARAM\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          SPARAM is REAL array, dimension (5)\nCOMMENT\t*>     SPARAM(1)=SFLAG\nCOMMENT\t*>     SPARAM(2)=SH11\nCOMMENT\t*>     SPARAM(3)=SH21\nCOMMENT\t*>     SPARAM(4)=SH12\nCOMMENT\t*>     SPARAM(5)=SH22\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*\nCOMMENT\t*  Authors:\nCOMMENT\t*  ========\nCOMMENT\t*\nCOMMENT\t*> \\author Univ. of Tennessee\nCOMMENT\t*> \\author Univ. of California Berkeley\nCOMMENT\t*> \\author Univ. of Colorado Denver\nCOMMENT\t*> \\author NAG Ltd.\nCOMMENT\t*\nCOMMENT\t*> \\date November 2017\nCOMMENT\t*\nCOMMENT\t*> \\ingroup single_blas_level1\nCOMMENT\t*\nCOMMENT\t*  =====================================================================\ntoken(96)\tSUBROUTINE\nIDENT\tSROTM\n(\t(\nIDENT\tN\n,\t,\nIDENT\tSX\n,\t,\nIDENT\tINCX\n,\t,\nIDENT\tSY\n,\t,\nIDENT\tINCY\n,\t,\nIDENT\tSPARAM\n)\t)\nCOMMENT\t*\nCOMMENT\t*  -- Reference BLAS level1 routine (version 3.8.0) --\nCOMMENT\t*  -- Reference BLAS is a software package provided by Univ. of Tennessee,    --\nCOMMENT\t*  -- Univ. of California Berkeley, Univ. of Colorado Denver and NAG Ltd..--\nCOMMENT\t*     November 2017\nCOMMENT\t*\nCOMMENT\t*     .. Scalar Arguments ..\nIDENT\tINTEGER\nIDENT\tINCX\n,\t,\nIDENT\tINCY\n,\t,\nIDENT\tN\nCOMMENT\t*     ..\nCOMMENT\t*     .. Array Arguments ..\nIDENT\tREAL\nIDENT\tSPARAM\n(\t(\nIDENT\t5\n)\t)\n,\t,\nIDENT\tSX\n(\t(\n*\t*)\n,\t,\nIDENT\tSY\n(\t(\n*\t*)\nCOMMENT\t*     ..\nCOMMENT\t*\nCOMMENT\t*  =====================================================================\nCOMMENT\t*\nCOMMENT\t*     .. Local Scalars ..\nIDENT\tREAL\nIDENT\tSFLAG\n,\t,\nIDENT\tSH11\n,\t,\nIDENT\tSH12\n,\t,\nIDENT\tSH21\n,\t,\nIDENT\tSH22\n,\t,\nIDENT\tTWO\n,\t,\nIDENT\tW\n,\t,\nIDENT\tZ\n,\t,\nIDENT\tZERO\nIDENT\tINTEGER\nIDENT\tI\n,\t,\nIDENT\tKX\n,\t,\nIDENT\tKY\n,\t,\nIDENT\tNSTEPS\nCOMMENT\t*     ..\nCOMMENT\t*     .. Data statements ..\nIDENT\tDATA\nIDENT\tZERO\n,\t,\nIDENT\tTWO\n/\t/\nIDENT\t0\n.\t.\nIDENT\tE0\n,\t,\nIDENT\t2\n.\t.\nIDENT\tE0\n/\t/\nCOMMENT\t*     ..\nCOMMENT\t*\nIDENT\tSFLAG\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t1\n)\t)\nIDENT\tIF\n(\t(\nIDENT\tN\n.\t.\nIDENT\tLE\n.\t.\nIDENT\t0\n.\t.\nIDENT\tOR\n.\t.\n(\t(\nIDENT\tSFLAG\n+\t+\nIDENT\tTWO\n.\t.\nIDENT\tEQ\n.\t.\nIDENT\tZERO\n)\t)\n)\t)\nIDENT\tRETURN\nIDENT\tIF\n(\t(\nIDENT\tINCX\n.\t.\nIDENT\tEQ\n.\t.\nIDENT\tINCY\n.\t.\nIDENT\tAND\n.\t.\nIDENT\tINCX\n.\t.\nIDENT\tGT\n.\t.\nIDENT\t0\n)\t)\nIDENT\tTHEN\nCOMMENT\t*\nIDENT\tNSTEPS\n=\t=\nIDENT\tN\n*\t*I\nIDENT\tNCX\nIDENT\tIF\n(\t(\nIDENT\tSFLAG\n.\t.\nIDENT\tLT\n.\t.\nIDENT\tZERO\n)\t)\nIDENT\tTHEN\nIDENT\tSH11\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t2\n)\t)\nIDENT\tSH12\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t4\n)\t)\nIDENT\tSH21\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t3\n)\t)\nIDENT\tSH22\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t5\n)\t)\nIDENT\tDO\nIDENT\tI\n=\t=\nIDENT\t1\n,\t,\nIDENT\tNSTEPS\n,\t,\nIDENT\tINCX\nIDENT\tW\n=\t=\nIDENT\tSX\n(\t(\nIDENT\tI\n)\t)\nIDENT\tZ\n=\t=\nIDENT\tSY\n(\t(\nIDENT\tI\n)\t)\nIDENT\tSX\n(\t(\nIDENT\tI\n)\t)\n=\t=\nIDENT\tW\n*\t*S\nIDENT\tH11\n+\t+\nIDENT\tZ\n*\t*S\nIDENT\tH12\nIDENT\tSY\n(\t(\nIDENT\tI\n)\t)\n=\t=\nIDENT\tW\n*\t*S\nIDENT\tH21\n+\t+\nIDENT\tZ\n*\t*S\nIDENT\tH22\nIDENT\tEND\nIDENT\tDO\nIDENT\tELSE\nIDENT\tIF\n(\t(\nIDENT\tSFLAG\n.\t.\nIDENT\tEQ\n.\t.\nIDENT\tZERO\n)\t)\nIDENT\tTHEN\nIDENT\tSH12\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t4\n)\t)\nIDENT\tSH21\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t3\n)\t)\nIDENT\tDO\nIDENT\tI\n=\t=\nIDENT\t1\n,\t,\nIDENT\tNSTEPS\n,\t,\nIDENT\tINCX\nIDENT\tW\n=\t=\nIDENT\tSX\n(\t(\nIDENT\tI\n)\t)\nIDENT\tZ\n=\t=\nIDENT\tSY\n(\t(\nIDENT\tI\n)\t)\nIDENT\tSX\n(\t(\nIDENT\tI\n)\t)\n=\t=\nIDENT\tW\n+\t+\nIDENT\tZ\n*\t*S\nIDENT\tH12\nIDENT\tSY\n(\t(\nIDENT\tI\n)\t)\n=\t=\nIDENT\tW\n*\t*S\nIDENT\tH21\n+\t+\nIDENT\tZ\nIDENT\tEND\nIDENT\tDO\nIDENT\tELSE\nIDENT\tSH11\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t2\n)\t)\nIDENT\tSH22\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t5\n)\t)\nIDENT\tDO\nIDENT\tI\n=\t=\nIDENT\t1\n,\t,\nIDENT\tNSTEPS\n,\t,\nIDENT\tINCX\nIDENT\tW\n=\t=\nIDENT\tSX\n(\t(\nIDENT\tI\n)\t)\nIDENT\tZ\n=\t=\nIDENT\tSY\n(\t(\nIDENT\tI\n)\t)\nIDENT\tSX\n(\t(\nIDENT\tI\n)\t)\n=\t=\nIDENT\tW\n*\t*S\nIDENT\tH11\n+\t+\nIDENT\tZ\nIDENT\tSY\n(\t(\nIDENT\tI\n)\t)\n=\t=\n-\t-\nIDENT\tW\n+\t+\nIDENT\tSH22\n*\t*Z\nIDENT\tEND\nIDENT\tDO\nIDENT\tEND\nIDENT\tIF\nIDENT\tELSE\nIDENT\tKX\n=\t=\nIDENT\t1\nIDENT\tKY\n=\t=\nIDENT\t1\nIDENT\tIF\n(\t(\nIDENT\tINCX\n.\t.\nIDENT\tLT\n.\t.\nIDENT\t0\n)\t)\nIDENT\tKX\n=\t=\nIDENT\t1\n+\t+\n(\t(\nIDENT\t1\n-\t-\nIDENT\tN\n)\t)\n*\t*I\nIDENT\tNCX\nIDENT\tIF\n(\t(\nIDENT\tINCY\n.\t.\nIDENT\tLT\n.\t.\nIDENT\t0\n)\t)\nIDENT\tKY\n=\t=\nIDENT\t1\n+\t+\n(\t(\nIDENT\t1\n-\t-\nIDENT\tN\n)\t)\n*\t*I\nIDENT\tNCY\nCOMMENT\t*\nIDENT\tIF\n(\t(\nIDENT\tSFLAG\n.\t.\nIDENT\tLT\n.\t.\nIDENT\tZERO\n)\t)\nIDENT\tTHEN\nIDENT\tSH11\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t2\n)\t)\nIDENT\tSH12\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t4\n)\t)\nIDENT\tSH21\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t3\n)\t)\nIDENT\tSH22\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t5\n)\t)\nIDENT\tDO\nIDENT\tI\n=\t=\nIDENT\t1\n,\t,\nIDENT\tN\nIDENT\tW\n=\t=\nIDENT\tSX\n(\t(\nIDENT\tKX\n)\t)\nIDENT\tZ\n=\t=\nIDENT\tSY\n(\t(\nIDENT\tKY\n)\t)\nIDENT\tSX\n(\t(\nIDENT\tKX\n)\t)\n=\t=\nIDENT\tW\n*\t*S\nIDENT\tH11\n+\t+\nIDENT\tZ\n*\t*S\nIDENT\tH12\nIDENT\tSY\n(\t(\nIDENT\tKY\n)\t)\n=\t=\nIDENT\tW\n*\t*S\nIDENT\tH21\n+\t+\nIDENT\tZ\n*\t*S\nIDENT\tH22\nIDENT\tKX\n=\t=\nIDENT\tKX\n+\t+\nIDENT\tINCX\nIDENT\tKY\n=\t=\nIDENT\tKY\n+\t+\nIDENT\tINCY\nIDENT\tEND\nIDENT\tDO\nIDENT\tELSE\nIDENT\tIF\n(\t(\nIDENT\tSFLAG\n.\t.\nIDENT\tEQ\n.\t.\nIDENT\tZERO\n)\t)\nIDENT\tTHEN\nIDENT\tSH12\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t4\n)\t)\nIDENT\tSH21\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t3\n)\t)\nIDENT\tDO\nIDENT\tI\n=\t=\nIDENT\t1\n,\t,\nIDENT\tN\nIDENT\tW\n=\t=\nIDENT\tSX\n(\t(\nIDENT\tKX\n)\t)\nIDENT\tZ\n=\t=\nIDENT\tSY\n(\t(\nIDENT\tKY\n)\t)\nIDENT\tSX\n(\t(\nIDENT\tKX\n)\t)\n=\t=\nIDENT\tW\n+\t+\nIDENT\tZ\n*\t*S\nIDENT\tH12\nIDENT\tSY\n(\t(\nIDENT\tKY\n)\t)\n=\t=\nIDENT\tW\n*\t*S\nIDENT\tH21\n+\t+\nIDENT\tZ\nIDENT\tKX\n=\t=\nIDENT\tKX\n+\t+\nIDENT\tINCX\nIDENT\tKY\n=\t=\nIDENT\tKY\n+\t+\nIDENT\tINCY\nIDENT\tEND\nIDENT\tDO\nIDENT\tELSE\nIDENT\tSH11\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t2\n)\t)\nIDENT\tSH22\n=\t=\nIDENT\tSPARAM\n(\t(\nIDENT\t5\n)\t)\nIDENT\tDO\nIDENT\tI\n=\t=\nIDENT\t1\n,\t,\nIDENT\tN\nIDENT\tW\n=\t=\nIDENT\tSX\n(\t(\nIDENT\tKX\n)\t)\nIDENT\tZ\n=\t=\nIDENT\tSY\n(\t(\nIDENT\tKY\n)\t)\nIDENT\tSX\n(\t(\nIDENT\tKX\n)\t)\n=\t=\nIDENT\tW\n*\t*S\nIDENT\tH11\n+\t+\nIDENT\tZ\nIDENT\tSY\n(\t(\nIDENT\tKY\n)\t)\n=\t=\n-\t-\nIDENT\tW\n+\t+\nIDENT\tSH22\n*\t*Z\nIDENT\tKX\n=\t=\nIDENT\tKX\n+\t+\nIDENT\tINCX\nIDENT\tKY\n=\t=\nIDENT\tKY\n+\t+\nIDENT\tINCY\nIDENT\tEND\nIDENT\tDO\nIDENT\tEND\nIDENT\tIF\nIDENT\tEND\nIDENT\tIF\nIDENT\tRETURN\nIDENT\tEND\n"