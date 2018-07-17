(string) (len=5637) "COMMENT\t> \\brief \\b ZDROT\nCOMMENT\t*\nCOMMENT\t*  =========== DOCUMENTATION ===========\nCOMMENT\t*\nCOMMENT\t* Online html documentation available at\nCOMMENT\t*            http://www.netlib.org/lapack/explore-html/\nCOMMENT\t*\nCOMMENT\t*  Definition:\nCOMMENT\t*  ===========\nCOMMENT\t*\nCOMMENT\t*       SUBROUTINE ZDROT( N, CX, INCX, CY, INCY, C, S )\nCOMMENT\t*\nCOMMENT\t*       .. Scalar Arguments ..\nCOMMENT\t*       INTEGER            INCX, INCY, N\nCOMMENT\t*       DOUBLE PRECISION   C, S\nCOMMENT\t*       ..\nCOMMENT\t*       .. Array Arguments ..\nCOMMENT\t*       COMPLEX*16         CX( * ), CY( * )\nCOMMENT\t*       ..\nCOMMENT\t*\nCOMMENT\t*\nCOMMENT\t*> \\par Purpose:\nCOMMENT\t*  =============\nCOMMENT\t*>\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>\nCOMMENT\t*> Applies a plane rotation, where the cos and sin (c and s) are real\nCOMMENT\t*> and the vectors cx and cy are complex.\nCOMMENT\t*> jack dongarra, linpack, 3/11/78.\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*\nCOMMENT\t*  Arguments:\nCOMMENT\t*  ==========\nCOMMENT\t*\nCOMMENT\t*> \\param[in] N\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          N is INTEGER\nCOMMENT\t*>           On entry, N specifies the order of the vectors cx and cy.\nCOMMENT\t*>           N must be at least zero.\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in,out] CX\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          CX is COMPLEX*16 array, dimension at least\nCOMMENT\t*>           ( 1 + ( N - 1 )*abs( INCX ) ).\nCOMMENT\t*>           Before entry, the incremented array CX must contain the n\nCOMMENT\t*>           element vector cx. On exit, CX is overwritten by the updated\nCOMMENT\t*>           vector cx.\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in] INCX\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          INCX is INTEGER\nCOMMENT\t*>           On entry, INCX specifies the increment for the elements of\nCOMMENT\t*>           CX. INCX must not be zero.\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in,out] CY\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          CY is COMPLEX*16 array, dimension at least\nCOMMENT\t*>           ( 1 + ( N - 1 )*abs( INCY ) ).\nCOMMENT\t*>           Before entry, the incremented array CY must contain the n\nCOMMENT\t*>           element vector cy. On exit, CY is overwritten by the updated\nCOMMENT\t*>           vector cy.\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in] INCY\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          INCY is INTEGER\nCOMMENT\t*>           On entry, INCY specifies the increment for the elements of\nCOMMENT\t*>           CY. INCY must not be zero.\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in] C\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          C is DOUBLE PRECISION\nCOMMENT\t*>           On entry, C specifies the cosine, cos.\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*>\nCOMMENT\t*> \\param[in] S\nCOMMENT\t*> \\verbatim\nCOMMENT\t*>          S is DOUBLE PRECISION\nCOMMENT\t*>           On entry, S specifies the sine, sin.\nCOMMENT\t*> \\endverbatim\nCOMMENT\t*\nCOMMENT\t*  Authors:\nCOMMENT\t*  ========\nCOMMENT\t*\nCOMMENT\t*> \\author Univ. of Tennessee\nCOMMENT\t*> \\author Univ. of California Berkeley\nCOMMENT\t*> \\author Univ. of Colorado Denver\nCOMMENT\t*> \\author NAG Ltd.\nCOMMENT\t*\nCOMMENT\t*> \\date December 2016\nCOMMENT\t*\nCOMMENT\t*> \\ingroup complex16_blas_level1\nCOMMENT\t*\nCOMMENT\t*  =====================================================================\ntoken(96)\tSUBROUTINE\nIDENT\tZDROT\n(\t(\nIDENT\tN\n,\t,\nIDENT\tCX\n,\t,\nIDENT\tINCX\n,\t,\nIDENT\tCY\n,\t,\nIDENT\tINCY\n,\t,\nIDENT\tC\n,\t,\nIDENT\tS\n)\t)\nCOMMENT\t*\nCOMMENT\t*  -- Reference BLAS level1 routine (version 3.7.0) --\nCOMMENT\t*  -- Reference BLAS is a software package provided by Univ. of Tennessee,    --\nCOMMENT\t*  -- Univ. of California Berkeley, Univ. of Colorado Denver and NAG Ltd..--\nCOMMENT\t*     December 2016\nCOMMENT\t*\nCOMMENT\t*     .. Scalar Arguments ..\nIDENT\tINTEGER\nIDENT\tINCX\n,\t,\nIDENT\tINCY\n,\t,\nIDENT\tN\nIDENT\tDOUBLE\nIDENT\tPRECISION\nIDENT\tC\n,\t,\nIDENT\tS\nCOMMENT\t*     ..\nCOMMENT\t*     .. Array Arguments ..\nIDENT\tCOMPLEX\n*\t*1\nIDENT\t6\nIDENT\tCX\n(\t(\n*\t* \n)\t)\n,\t,\nIDENT\tCY\n(\t(\n*\t* \n)\t)\nCOMMENT\t*     ..\nCOMMENT\t*\nCOMMENT\t* =====================================================================\nCOMMENT\t*\nCOMMENT\t*     .. Local Scalars ..\nIDENT\tINTEGER\nIDENT\tI\n,\t,\nIDENT\tIX\n,\t,\nIDENT\tIY\nIDENT\tCOMPLEX\n*\t*1\nIDENT\t6\nIDENT\tCTEMP\nCOMMENT\t*     ..\nCOMMENT\t*     .. Executable Statements ..\nCOMMENT\t*\nIDENT\tIF\n(\t(\nIDENT\tN\n.\t.\nIDENT\tLE\n.\t.\nIDENT\t0\n)\t)\nIDENT\tRETURN\nIDENT\tIF\n(\t(\nIDENT\tINCX\n.\t.\nIDENT\tEQ\n.\t.\nIDENT\t1\n.\t.\nIDENT\tAND\n.\t.\nIDENT\tINCY\n.\t.\nIDENT\tEQ\n.\t.\nIDENT\t1\n)\t)\nIDENT\tTHEN\nCOMMENT\t*\nCOMMENT\t*        code for both increments equal to 1\nCOMMENT\t*\nIDENT\tDO\nIDENT\tI\n=\t=\nIDENT\t1\n,\t,\nIDENT\tN\nIDENT\tCTEMP\n=\t=\nIDENT\tC\n*\t*C\nIDENT\tX\n(\t(\nIDENT\tI\n)\t)\n+\t+\nIDENT\tS\n*\t*C\nIDENT\tY\n(\t(\nIDENT\tI\n)\t)\nIDENT\tCY\n(\t(\nIDENT\tI\n)\t)\n=\t=\nIDENT\tC\n*\t*C\nIDENT\tY\n(\t(\nIDENT\tI\n)\t)\n-\t-\nIDENT\tS\n*\t*C\nIDENT\tX\n(\t(\nIDENT\tI\n)\t)\nIDENT\tCX\n(\t(\nIDENT\tI\n)\t)\n=\t=\nIDENT\tCTEMP\nIDENT\tEND\nIDENT\tDO\nIDENT\tELSE\nCOMMENT\t*\nCOMMENT\t*        code for unequal increments or equal increments not equal\nCOMMENT\t*          to 1\nCOMMENT\t*\nIDENT\tIX\n=\t=\nIDENT\t1\nIDENT\tIY\n=\t=\nIDENT\t1\nIDENT\tIF\n(\t(\nIDENT\tINCX\n.\t.\nIDENT\tLT\n.\t.\nIDENT\t0\n)\t)\nIDENT\tIX\n=\t=\n(\t(\n-\t-\nIDENT\tN\n+\t+\nIDENT\t1\n)\t)\n*\t*I\nIDENT\tNCX\n+\t+\nIDENT\t1\nIDENT\tIF\n(\t(\nIDENT\tINCY\n.\t.\nIDENT\tLT\n.\t.\nIDENT\t0\n)\t)\nIDENT\tIY\n=\t=\n(\t(\n-\t-\nIDENT\tN\n+\t+\nIDENT\t1\n)\t)\n*\t*I\nIDENT\tNCY\n+\t+\nIDENT\t1\nIDENT\tDO\nIDENT\tI\n=\t=\nIDENT\t1\n,\t,\nIDENT\tN\nIDENT\tCTEMP\n=\t=\nIDENT\tC\n*\t*C\nIDENT\tX\n(\t(\nIDENT\tIX\n)\t)\n+\t+\nIDENT\tS\n*\t*C\nIDENT\tY\n(\t(\nIDENT\tIY\n)\t)\nIDENT\tCY\n(\t(\nIDENT\tIY\n)\t)\n=\t=\nIDENT\tC\n*\t*C\nIDENT\tY\n(\t(\nIDENT\tIY\n)\t)\n-\t-\nIDENT\tS\n*\t*C\nIDENT\tX\n(\t(\nIDENT\tIX\n)\t)\nIDENT\tCX\n(\t(\nIDENT\tIX\n)\t)\n=\t=\nIDENT\tCTEMP\nIDENT\tIX\n=\t=\nIDENT\tIX\n+\t+\nIDENT\tINCX\nIDENT\tIY\n=\t=\nIDENT\tIY\n+\t+\nIDENT\tINCY\nIDENT\tEND\nIDENT\tDO\nIDENT\tEND\nIDENT\tIF\nIDENT\tRETURN\nIDENT\tEND\n"